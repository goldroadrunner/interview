#!/bin/sh

curl \
    --request POST \
    --data "${@}" \
    http://127.0.0.1:8081/health/check
