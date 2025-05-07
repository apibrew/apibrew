package nano

//go:generate apbr generate --platform=golang --path=./model/ --package=model --source-file=schema/schema.yml
//go:generate apbr generate --platform=golang --path=./model/ --package=model --source-file=schema/Action.yml
//go:generate apbr generate --platform=golang --path=./model/ --package=model --source-file=schema/CronJob.yml
//go:generate apbr generate --platform=golang --path=./model/ --package=model --source-file=schema/Job.yml
