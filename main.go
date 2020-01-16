package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	reform "gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

const (
	insertBatchSize = 10_000
	workersCount    = 5
	deleteCount     = 300_000
	cleanerInterval = time.Second * 5
)

func main() {
	rand.Seed(time.Now().UnixNano())
	logger := buildLogger()
	logger.Info("starting......")

	conn, err := sql.Open("postgres", "postgres://postgres:5432@postgres/pg_bloat_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		logger.Error(err)
		return
	}

	db := reform.NewDB(conn, postgresql.Dialect, nil)

	go cleaner(db, logger.WithField("system", "cleaner"))
	runWorkers(db, logger.WithField("system", "worker"))
}

func runWorkers(db *reform.DB, logger *logrus.Entry) {
	results := make(chan int64, 100)

	for w := 1; w <= workersCount; w++ {
		go worker(db, logger, results)
	}

	var counter int64
	for amount := range results {
		counter += amount
		if counter%200_000 == 0 {
			logger.Infof("Amount inserted is: %+v", counter)
		}
	}
}

func cleaner(db *reform.DB, logger *logrus.Entry) {
	var counter int64
	for {

		res, err := db.Exec("DELETE FROM specs WHERE ctid IN (SELECT ctid FROM specs LIMIT $1)", deleteCount)
		if err != nil {
			log.Fatal(err)
			continue
		}
		affected, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
			continue
		}

		counter += affected
		logger.Infof("Amount droppedis: %+v", counter)
		time.Sleep(cleanerInterval)
	}
}

func batchOfSpecs() []reform.Struct {
	batch := make([]reform.Struct, insertBatchSize)

	for i := 0; i < insertBatchSize; i++ {
		batch[i] = &Spec{
			Status:     0,
			LineNumber: 10,
			Filename:   "/spec/data",
			CommitID:   23233,
		}
	}

	return batch
}

func worker(db *reform.DB, logger *logrus.Entry, results chan<- int64) {
	for {
		specs := batchOfSpecs()
		err := db.InsertMulti(specs...)
		if err != nil {
			logger.Fatal(err)
			continue
		}
		results <- int64(len(specs))
	}
}

func buildLogger() *logrus.Logger {
	log := logrus.New()
	log.Level = logrus.InfoLevel

	return log
}
