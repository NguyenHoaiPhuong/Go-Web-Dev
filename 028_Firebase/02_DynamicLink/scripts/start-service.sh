#!/bin/bash

FILE=config/.env.sh
if [ -f "$FILE" ]; then
    echo "$FILE exist"
else 
    echo "$FILE does not exist"
    cp config/.env_sample.sh config/.env.sh
fi

source config/.env.sh && \
go run main.go