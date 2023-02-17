#!/bin/bash
# **************************
# Localstack SQS Send Message Tool
# Version: 1.0.0
# **************************
# -----------------------------------------------------------------------------
# Current file variables
# -----------------------------------------------------------------------------
debug=true
parent_folder="../"
current_path=$(pwd)/
current_path_basename=$(basename $(pwd))
current_file_full_path=$0
current_file_name=$(basename -- "$0")
if [ $current_file_full_path = $current_file_name ] || [ $current_file_full_path = "./$current_file_name" ]; then
  current_file_full_path="./${current_file_full_path}"
  current_file_path="./"
else
  current_file_path="${current_file_full_path/$current_file_name/''}"
fi

current_file_path_basename=$(basename -- "$current_file_path")

if [ -z "$current_file_path_basename" ] || [ $current_file_path = "./" ]; then
  current_parent_folder="../"
else
  current_file_path_basename=$current_file_path_basename/
  # TODO: está com problema (corrigir)
  current_parent_folder="${current_file_path/$current_file_path_basename/''}"
fi


if [ debug ]; then
  echo '----------------------------------------'
  echo "$0 - Script variables"
  echo '----------------------------------------'
  echo "current_path: $current_path"
  echo "current_path_basename: $current_path_basename"
  echo "current_file_full_path: $current_file_full_path"
  echo "current_file_name: $current_file_name"
  echo "current_file_path: $current_file_path"
  echo "current_parent_folder: $current_parent_folder"
  echo '----------------------------------------'
fi

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
MESSAGE=$2
if [ -z "$MESSAGE" ]
then
  MESSAGE=$(cat ${current_file_path}sample.json)
else
  if test -f $2 ; then
    MESSAGE=$(cat $2)
  fi
fi

# cat ${current_file_path}sample.json
#echo $MESSAGE
echo "aws --endpoint-url=http://$HOST:4566 sqs send-message --queue-url $QUEUE --message-body '$MESSAGE'"
aws --endpoint-url=http://$HOST:4566 sqs send-message --queue-url $QUEUE --message-body "$MESSAGE"
