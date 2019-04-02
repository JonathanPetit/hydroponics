#!/usr/bin/env bash

psql -p $DB_PORT -c "TRUNCATE TABLE temperature RESTART IDENTITY CASCADE;"