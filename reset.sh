#!/bin/bash

rm -rf ./media/
rm -f sqlite.db
go run migrate/migrate.go
