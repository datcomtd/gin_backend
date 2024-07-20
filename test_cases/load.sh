#!/bin/bash

function Print() {
  echo -e " --- \e[33;1m${*}\e[0m"
}

function add_user() {
  curl -s -L \
    -X POST -H "Content-Type: application/json" \
    -d "${*}" \
    http://localhost:8000/api/register | jq '.'
  echo
}

function list_users() {
  curl -s -L http://localhost:8000/api/users | jq '.'
  echo
}

function get_token() {
  curl -s -L \
    -X POST -H "Content-Type: application/json" \
    -d "${*}" \
    http://localhost:8000/api/token | jq '.'
  echo
}
