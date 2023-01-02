#!/bin/bash
# **************************
# Localstack SQS List Queues Tool
# Version: 1.0.0
# **************************
if [ $RUNNING_IN_CONTAINER ]; then
  HOST=localstack
else
  HOST=0.0.0.0
fi
echo "aws --endpoint-url=http://$HOST:4566 sqs list-queues"
aws --endpoint-url=http://$HOST:4566 sqs list-queues