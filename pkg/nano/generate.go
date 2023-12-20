package nano

//go:generate apbr apply -f schema.yml
//go:generate apbr generate --filter namespace=nano --filter name=Code --platform=golang --package=nano .
