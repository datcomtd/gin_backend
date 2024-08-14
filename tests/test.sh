#!/bin/bash

source ./load.sh

#Print "should return 'required fields are not filled'"
#add_user "{\"username\": \"patrick\", \"password\": \"patrick123\", \"course\": 1}"

# get admin credentials
ADMIN_PWD=$(cat ../initializers/env.go | grep ADMIN_PWD | awk {'print $NF'} | sed -e "s,\",,g")

Print "add the computer engineer president"
add_user "{\"admin-username\": \"admin\", \"admin-password\": \"${ADMIN_PWD}\", \"username\": \"patrick\", \"password\": \"patrick123\", \"role\": 1, \"course\": 1}"

Print "list user: patrick"
list_user "patrick"

Print "get patrick's token"
token=$(get_token "{\"username\": \"patrick\", \"password\": \"patrick123\"}")
echo "Token: ${token}"

Print "update patrick record"
update_user "patrick" "{\"username\": \"patrick\", \"password\": \"patrick123\", \"email\": \"newemail@gmail.com\"}"

Print "list user: patrick"
list_user "patrick"

Print "generate a key for document upload"
#generate_document_key_error "${token}" "{\"title\": \"Simple Text File\", \"source\": \"Tester\", \"category\": \"edital\"}"
key=$(generate_document_key "${token}" "{\"title\": \"Simple Text File\", \"source\": \"Tester\", \"category\": \"edital\"}")
echo "Key: ${key}"

Print "uploading the document text.txt"
upload_document "${key}" "${token}" "text.txt"

Print "renaming the document"
update_document "${token}" "{\"id\": 1, \"filename\": \"newfilename.txt\"}"

Print "list all documents"
list_documents

Print "deleting the document"
delete_document "${token}" "{\"id\": 1}"

Print "creating a product"
create_product "${token}" "{\"title\": \"camiseta DATCOM\"}"

Print "list all products"
list_products
