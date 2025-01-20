#!/bin/bash

awslocal s3api create-bucket --bucket ${AWS_BUCKERT_NAME}

awslocal sqs create-queue --queue-name ${QUEUE_PROCESS_VIDEO}
awslocal sqs create-queue --queue-name ${DEAD_LETTER_QUEUE_PROCESS_VIDEO}

awslocal sqs set-queue-attributes \
--queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/${AWS_ACCESS_KEY_ID}/${QUEUE_PROCESS_VIDEO} \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:us-east-1:${AWS_ACCESS_KEY_ID}:${DEAD_LETTER_QUEUE_PROCESS_VIDEO}\",\"maxReceiveCount\":\"1\"}"
}'
