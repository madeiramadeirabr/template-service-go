#!/bin/bash
# **************************
# Localstack SQS Receive Messages Tool
# Version: 1.0.0
# **************************
if [ $RUNNING_IN_CONTAINER ]; then
  HOST=localstack
else
  HOST=0.0.0.0
fi

QUEUE=$1
if [ -z "$QUEUE" ]
then
  QUEUE='http://$HOST:4566/000000000000/test-queue'
else
  QUEUE=$(basename -- $QUEUE)
  QUEUE="http://$HOST:4566/000000000000/${QUEUE}"
fi
echo "aws --endpoint-url=http://$HOST:4566 sqs receive-message --queue-url $QUEUE"
aws --endpoint-url=http://$HOST:4566 sqs receive-message --queue-url $QUEUE

if [ ! $? -eq 0 ]; then
    QUEUE="http://$HOST:4566/000000000000/$QUEUE"
    echo "aws --endpoint-url=http://$HOST:4566 sqs receive-message --queue-url $QUEUE"
    aws --endpoint-url=http://$HOST:4566 sqs receive-message --queue-url $QUEUE
fi