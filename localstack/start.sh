#!/bin/bash

awslocal s3api create-bucket --bucket ${AWS_BUCKET_NAME}

awslocal sqs create-queue --queue-name ${QUEUE_PROCESS_VIDEO}
awslocal sqs create-queue --queue-name ${QUEUE_RESULT_VIDEO}
awslocal sqs create-queue --queue-name ${DEAD_LETTER_QUEUE_PROCESS_VIDEO}
awslocal sqs create-queue --queue-name ${DEAD_LETTER_QUEUE_RESULT_VIDEO}

awslocal sqs set-queue-attributes \
--queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/${AWS_SQS_ACCESS_KEY_ID}/${QUEUE_PROCESS_VIDEO} \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:us-east-1:${AWS_SQS_ACCESS_KEY_ID}:${DEAD_LETTER_QUEUE_PROCESS_VIDEO}\",\"maxReceiveCount\":\"1\"}"
}'

awslocal sqs set-queue-attributes \
--queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/${AWS_SQS_ACCESS_KEY_ID}/${QUEUE_RESULT_VIDEO} \
--attributes '{
    "RedrivePolicy": "{\"deadLetterTargetArn\":\"arn:aws:sqs:us-east-1:${AWS_SQS_ACCESS_KEY_ID}:${DEAD_LETTER_QUEUE_RESULT_VIDEO}\",\"maxReceiveCount\":\"1\"}"
}'
