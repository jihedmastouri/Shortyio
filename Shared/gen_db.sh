#!/bin/bash

# Read PostgreSQL database details from environment variables
PG_HOST="$PG_HOST_ENV"
PG_PORT="$PG_PORT_ENV"
PG_USER="$PG_USER_ENV"
PG_PASSWORD="$PG_PASSWORD_ENV"
PG_DATABASE="$PG_DATABASE_ENV"

# Define the SQL script file path
SQL_FILE="./schema.sql"

# Connect to the PostgreSQL database and execute the script
psql -h $PG_HOST -p $PG_PORT -U $PG_USER -W $PG_PASSWORD -d $PG_DATABASE -f $SQL_FILE

