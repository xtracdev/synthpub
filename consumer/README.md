aws logs create-log-group --log-group-name SynthEvents

aws ecs register-task-definition --cli-input-json file://$PWD/taskdef.json

aws ecs run-task --cluster orange --task-definition consynthevent