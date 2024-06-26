#!/usr/bin/env bash

# URL của API
BASE_URL="http://localhost:8000"


check_result() {
  if [ $1 -ne 0 ]; then
    echo "Request failed with status $1"
    exit 1
  fi
}


curl --location "$BASE_URL/categories" \
--header "Content-Type: application/json" \
--header "Accept: application/json" \
--data '{
  "description": "iPhone",
  "name": "Phone"
}'
check_result $?
echo "categories ... ok"


curl --location "$BASE_URL/suppliers" \
--header "Content-Type: application/json" \
--header "Accept: application/json" \
--data-raw '{
  "email": "exam@example2.com",
  "name": "Apple2",
  "city": "Ho Chi Minh City",
  "district": "D1",
  "street": "Ton Duc Thang",
  "ward": "14"
}'
check_result $?
echo "suppliers ... ok"


curl --location "$BASE_URL/products" \
--header "Content-Type: application/json" \
--header "Accept: application/json" \
--data '{
  "category_id": "1",
  "description": "iphone",
  "name": "iphone",
  "price": "1666-111111",
  "status": "active",
  "supplier_id": "1"
}'
check_result $?
echo "products ... ok"