#!/usr/bin/env bash

# URL của API
BASE_URL="http://localhost:8000"


check_result() {
  if [ $1 -ne 0 ]; then
    echo "Request failed with status $1"
    exit 1
  fi
}

curl --location "$BASE_URL/products" \
--header "Content-Type: application/json" \
--header "Accept: application/json" \
--data '{
  "name": "iPhone 16",
  "description": "iPhone 16",
  "price": "1666-111111",
  "status": "active",
  "category_id": "1",
  "supplier_id": "1"
}'
check_result $?
echo "products ... ok"