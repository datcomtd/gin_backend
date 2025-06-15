#!/bin/bash

DATCOM_DF=${DATCOM_DF:-"./Dockerfile"}

if [ ! -f "${DATCOM_DF}" ] ; then
  echo " - Entre no diretório do gin_backend antes de executar o docker.sh"
  exit 1
fi

function cleandocker() {
  docker stop gin-postgres-db 1>/dev/null 2>&1
  docker rm gin-postgres-db 1>/dev/null 2>&1
  docker rmi postgres:latest 1>/dev/null 2>&1

  docker rm datcom-backend 1>/dev/null 2>&1
  docker rmi datcom-backend 1>/dev/null 2>&1
  docker rmi golang:1.23.5-alpine 1>/dev/null 2>&1
}

function setup_postgres() {
  SETUP_DB=1

  docker run --name gin-postgres-db \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=postgres \
    -e POSTGRES_DB=postgres \
    -p 4145:5432 \
    -d postgres
  if [ $? -ne 0 ] ; then
    echo "- ERROR while running the postgres database"
    exit 1
  fi
}

function setup_backend() {
  docker build -f "${DATCOM_DF}" -t datcom-backend .
  if [ $? -ne 0 ] ; then
    echo "- ERROR while building the backend image"
    exit 1
  fi

  docker run --name datcom-backend \
    --network host \
    -d datcom-backend
  if [ $? -ne 0 ] ; then
    echo "- ERROR while running the backend"
    exit 1
  fi
}

docker info 1>/dev/null 2>&1
if [ $? -ne 0 ] ; then
  echo " - ERROR O daemon Docker não está iniciado."
  exit 1
fi

cleandocker

while [ $# -gt 0 ] ; do
  case "${1}" in
  "-stop") ;&
  "stop")
    exit
    ;;
  "postgres") ;&
  "postgresql") ;&
  "psql") ;&
  "pg")
    setup_postgres
    ;;
  "reset") ;&
  "-r")
    RESET_DB=1
    ;;
  "backend") ;&
  "run")
    setup_backend
    ;;
  esac

  shift
done

if [ "${SETUP_DB}" ] && [ "${RESET_DB}" ] ; then ./reset.sh ; fi
