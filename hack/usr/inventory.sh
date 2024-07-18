#!/usr/bin/env bash

curl --location 'http://localhost:8000/inventories' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
--data '{
  "available": "1000",
  "model": "phone",
  "price": "12312313",
  "product_id": "1",
  "currency_code": "USD"
}'