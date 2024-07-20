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

function list_user() {
  curl -s -L http://localhost:8000/api/user/${*} | jq '.'
  echo
}

function get_token() {
  curl -s -L \
    -X POST -H "Content-Type: application/json" \
    -d "${*}" \
    http://localhost:8000/api/token | jq '.token' |
    sed -e "s/\"//g"
}

function update_user() {
  curl -s -L \
    -X POST -H "Content-Type: application/json" \
    -d "${*}" \
    http://localhost:8000/api/user/update | jq '.'
  echo
}

function delete_user() {
  username="${1}"
  shift
  password="${1}"
  shift

  curl -s -L \
    -X POST -H "Content-Type: application/json" \
    -d "{\"username\": \"${username}\", \"password\": \"${password}\"}" \
    http://localhost:8000/api/user/delete | jq '.'
  echo
}

function generate_document_key {
  token="${1}"
  shift

  curl -s -L -X POST -H "Content-Type: application/json" \
    -H "Authorization: ${token}" \
    -d "${1}" \
    http://localhost:8000/api/document/upload | jq '.key' |
    sed -e "s/\"//g"
}

function generate_document_key_error() {
  token="${1}"
  shift

  curl -s -L -X POST -H "Content-Type: application/json" \
    -H "Authorization: ${token}" \
    -d "${1}" \
    http://localhost:8000/api/document/upload | jq '.'
  echo
}

function upload_document() {
  key="${1}"
  shift
  token="${1}"
  shift
  filename="${1}"
  shift

  curl -s -L -X POST -H "Content-Type: multipart/form-data" \
    -H "Authorization: ${token}" \
    -F "file=@${filename}" \
    http://localhost:8000/api/document/upload/${key} | jq '.'
  echo
}

function list_documents() {
  curl -s -L http://localhost:8000/api/documents | jq '.'
  echo
}
