#!/bin/bash
# **************************
# Localstack SQS Purged Queue Tool
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

  if [ -z "$2" ]; then
    REGION=us-east-1
  else
    REGION=$2
  fi

  echo "aws --endpoint-url=http://$HOST:4566 sqs get-queue-url --queue-name $1 --region $REGION"
  aws --endpoint-url=http://$HOST:4566 sqs get-queue-url --queue-name $1 --region $REGION
  if [ $? -eq 0 ]; then
    echo "aws --endpoint-url=http://$HOST:4566 sqs purge-queue --queue-url http://$HOST:4566/000000000000/$1 --region $REGION"
    aws --endpoint-url=http://$HOST:4566 sqs purge-queue --queue-url http://$HOST:4566/000000000000/$1 --region $REGION
    if [ $? -eq 0 ]; then
      echo "Queue purged"
    else
      echo "Queue not purged"
      exit 1
    fi
  else
    echo "Queue doesn't exists"
    exit 1
  fi
fi
