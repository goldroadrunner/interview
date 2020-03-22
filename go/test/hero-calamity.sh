#!/bin/sh

curl \
    --request POST \
    --data "{\"Level\": 100, \"Heros\": [ \"a\", \"b\", \"c\" ] }" \
    http://127.0.0.1:8081/hero/calamity
