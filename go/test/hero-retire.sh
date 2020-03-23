#!/bin/sh

curl \
    --request POST \
    --data "{\"PowerLevel\": 100, \"Exhaustion\": 50, \"Name\": \"fred\"}" \
    "http://127.0.0.1:8081/hero" &&
    curl \
	--request GET \
	"http://127.0.0.1:8081/hero/fred" &&
    curl \
	--request POST \
	"http://127.0.0.1:8081/hero/retire/fred" &&
    curl \
	--request GET \
	"http://127.0.0.1:8081/hero/fred"
