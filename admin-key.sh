#!/bin/bash

key=$(echo $RANDOM | md5sum | head -c 32)

echo -e "[\e[33;1m!\e[0m] admin key: ${key}"

sed -i "s,.*DATCOM_ADMIN_KEY.*,var DATCOM_ADMIN_KEY string = \"${key}\",g" initializers/env.go
