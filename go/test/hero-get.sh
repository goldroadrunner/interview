#!/bin/sh

curl \
    --request GET \
    "http://127.0.0.1:8081/hero/${@}"
