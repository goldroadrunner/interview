#!/bin/sh

curl \
    --request POST \
    --data "{\"PowerLevel\": 100, \"Exhaustion\": 50, \"Name\": \"${@}\"}" \
    http://127.0.0.1:8081/hero
