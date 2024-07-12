#!/bin/bash

rm -f sqlite.db
go run migrate/migrate.go
