#!/bin/sh

curl \
    --request POST \
    --data "{\"Status\": \"${@}\"}" \
    http://127.0.0.1:8081/health/check
