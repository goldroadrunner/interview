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
	--data "{\"Level\": 299, \"Heros\": [ \"a\", \"b\", \"c\" ] }" \
	http://127.0.0.1:8081/hero/calamity &&
    curl \
	--request POST \
	--data "{\"PowerLevel\": 100, \"Exhaustion\": 50, \"Name\": \"d\"}" \
	http://127.0.0.1:8081/hero &&
    curl \
	--request POST \
	--data "{\"PowerLevel\": 100, \"Exhaustion\": 50, \"Name\": \"e\"}" \
	http://127.0.0.1:8081/hero &&
    curl \
	--request POST \
	--data "{\"PowerLevel\": 100, \"Exhaustion\": 50, \"Name\": \"f\"}" \
	http://127.0.0.1:8081/hero &&
    curl \
	--request GET \
	"http://127.0.0.1:8081/hero/d" &&
    curl \
	--request POST \
	--data "{\"Level\": 301, \"Heros\": [ \"d\", \"e\", \"f\" ] }" \
	http://127.0.0.1:8081/hero/calamity &&
    curl \
	--request GET \
	"http://127.0.0.1:8081/hero/d"
