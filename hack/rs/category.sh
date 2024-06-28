#!/usr/bin/env bash

# URL của API
BASE_URL="http://localhost:8000"


curl --location "$BASE_URL/categories" \
--header "Content-Type: application/json" \
--header "Accept: application/json" \
--data '{
  "description": "iPhone",
  "name": "Phone"
}'
check_result $?
echo "categories ... ok"