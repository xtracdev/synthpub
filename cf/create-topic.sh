aws cloudformation create-stack --stack-name synth-stack-topic \
--template-body file://create-topic.yml \
--parameters ParameterKey=TopicName,ParameterValue=SynthTopic
