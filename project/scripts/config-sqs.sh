#!/bin/bash

aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name $QUEUE_NAME