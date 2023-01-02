#!/bin/bash
# **************************
# Localstack SQS Delete Queue Tool
# Version: 1.0.0
# **************************
if [ -z "$1" ]; then
  echo 'Queue name must be informed'
  exit 1
else
  if [ $RUNNING_IN_CONTAINER ]; then
    HOST=localstack
  else
    HOST=0.0.0.0
  fi
  echo "aws --endpoint-url=http://$HOST:4566 sqs get-queue-url --queue-name $1"
  aws --endpoint-url=http://$HOST:4566 sqs get-queue-url --queue-name $1
  if [ $? -eq 0 ]; then
    echo "aws --endpoint-url=http://$HOST:4566 sqs delete-queue --queue-url http://$HOST:4566/000000000000/$1"
    aws --endpoint-url=http://$HOST:4566 sqs delete-queue --queue-url http://$HOST:4566/000000000000/$1
    if [ $? -eq 0 ]; then
      echo "Queue deleted"
    else
      echo "Queue not deleted"
      exit 1
    fi
  else
    echo "Queue doesn't exists"
    exit 1
  fi
fi
