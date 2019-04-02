#!/usr/bin/env bash

echo "Start"

dropdb -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME    
echo "Drop DB"

createdb -h $DB_HOST -p $DB_PORT -U $DB_USER -O $DB_OWNER $DB_NAME
echo "Create DB"

psql -p $DB_PORT -d $DB_NAME -U $DB_USER -f schema.sql
echo "Schema upload"

echo "Finish"