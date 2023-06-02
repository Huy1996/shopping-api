#!/bin/sh

# Exit immediately if received a non 0 state
set -e

# Run DB Migration
echo "DB Migration"
/app/migrate -path /app/src/db/migration -database "$DB_SOURCE" -verbose up

# Run App
echo "App Starting"
exec "$@"