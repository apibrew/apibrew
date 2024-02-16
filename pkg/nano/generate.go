package nano

//go:generate apbr generate --platform=golang --path=. --package=nano --source-file=schema.yml
//go:generate statik -src=./builtin -ns=nano-builtin -dest=.
