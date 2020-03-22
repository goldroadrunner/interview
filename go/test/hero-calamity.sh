#!/bin/sh

curl \
    --request POST \
    --data "{\"PowerLevel\": 100, \"Exhaustion\": 50, \"Name\": \"a\"}" \
    http://127.0.0.1:8081/hero &&
    curl \
	--request POST \
	--data "{\"PowerLevel\": 100, \"Exhaustion\": 50, \"Name\": \"b\"}" \
	http://127.0.0.1:8081/hero &&
    curl \
	--request POST \
	--data "{\"PowerLevel\": 100, \"Exhaustion\": 50, \"Name\": \"c\"}" \
	http://127.0.0.1:8081/hero &&
    curl \
	--request POST \
	--data "{\"Level\": 100, \"Heros\": [ \"a\", \"b\", \"c\" ] }" \
	http://127.0.0.1:8081/hero/calamity
