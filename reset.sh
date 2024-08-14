#!/bin/bash

function Try() {
  color=${1}
  shift
  char=${1}
  shift

  psql -U postgres -d postgres -p 4145 -c "${*}" 1>/dev/null 2>&1
  code=$?
  if [ ${code} -ne 0 ]; then
    echo -e "\e[${color};1m[${char}] (code: ${code}) ${*}\e[0m"
  else
    echo -e "[+] ${*}"
  fi
}

curl -s -L "http://localhost:4145"
if [ $? -ne 52 ]; then
  echo -e "\e[31;1m[-] postgresql não iniciado\e[0m"
  exit
fi

db_pwd=$(echo $RANDOM | md5sum | head -c 10)
echo -e "[\e[33;1m!\e[0m] new password: ${db_pwd}"
sed -i "s,.*DATCOM_DB_PWD.*,var DATCOM_DB_PWD string = \"${db_pwd}\",g" initializers/env.go

pwd=$(echo $RANDOM | md5sum | head -c 32)
echo -e "[\e[33;1m!\e[0m] admin password: ${pwd}"
sed -i "s,.*DATCOM_ADMIN_PWD.*,var DATCOM_ADMIN_PWD string = \"${pwd}\",g" initializers/env.go

Try 33 "!" "drop database datcom_db;"
Try 33 "!" "drop user datcom_user;"

Try 31 "-" "create database datcom_db;"
Try 31 "-" "create user datcom_user with encrypted password '${pwd}';"
Try 31 "-" "grant all privileges on database datcom_db to datcom_user;"
Try 31 "-" "alter database datcom_db owner to datcom_user;"

rm -f sqlite.db

go run migrate/migrate.go
if [ $? -eq 0 ]; then
  echo -e "[+] go run migrate/migrate.go"
fi
