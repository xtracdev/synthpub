#!/bin/bash
aws cloudformation create-stack --stack-name synth-stack-queue \
--template-body file://create-and-sub-queue.yml \
--parameters ParameterKey=TopicARN,ParameterValue=$1 \
ParameterKey=QueueName,ParameterValue=SynthEventQ
