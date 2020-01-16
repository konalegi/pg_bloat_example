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
	workersCount    = 8
	cleanerInterval = time.Second * 30
	vacuumInterval  = time.Minute * 2
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
	go manualVacuumer(db, logger.WithField("system", "vacuumer"))
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

func manualVacuumer(db *reform.DB, logger *logrus.Entry) {
	for {
		_, err := db.Exec("vacuum")
		if err != nil {
			log.Fatal(err)
			continue
		}

		logger.Info("Manual vacuum completed")
		time.Sleep(vacuumInterval)
	}
}

func cleaner(db *reform.DB, logger *logrus.Entry) {
	var counter int64
	for {
		res, err := db.Exec("DELETE FROM specs WHERE created_at < $1", time.Now().Add(-20*time.Second))
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
			CreatedAt:  time.Now(),
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
