#!/bin/sh

set -e

psql -U postgres -h postgres -d postgres -c "DROP DATABASE IF EXISTS pg_bloat_db"
psql -U postgres -h postgres -d postgres -c "CREATE DATABASE pg_bloat_db"
psql -U postgres -h postgres -d pg_bloat_db -a -f db.sql