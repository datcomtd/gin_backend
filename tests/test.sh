#!/bin/bash

source ./load.sh

#Print "should return 'required fields are not filled'"
#add_user "{\"username\": \"patrick\", \"password\": \"patrick123\", \"course\": 1}"

Print "add the computer engineer president"
add_user "{\"username\": \"patrick\", \"password\": \"patrick123\", \"role\": 1, \"course\": 1}"

Print "list user: patrick"
list_user "patrick"

Print "get patrick's token"
get_token "{\"username\": \"patrick\", \"password\": \"patrick123\"}"

Print "update patrick record"
update_user "patrick" "{\"username\": \"patrick\", \"password\": \"patrick123\", \"email\": \"newemail@gmail.com\"}"

Print "list user: patrick"
list_user "patrick"
