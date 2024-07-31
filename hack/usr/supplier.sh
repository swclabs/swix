#!/usr/bin/env bash

# URL của API
BASE_URL="http://localhost:8000"


curl --location "$BASE_URL/suppliers" \
--header "Content-Type: application/json" \
--header "Accept: application/json" \
--data-raw '{
  "email": "exam@example2.com",
  "name": "Apple",
  "city": "Ho Chi Minh City",
  "district": "D1",
  "street": "Ton Duc Thang",
  "ward": "14"
}'
echo "suppliers ... ok"