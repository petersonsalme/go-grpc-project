#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE DATABASE auth_svc;
    CREATE DATABASE order_svc;
    CREATE DATABASE product_svc;
EOSQL