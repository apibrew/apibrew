package rest

//go:generate statik -include=*.yaml -src=. -ns=rest -dest=.
//go:generate npx openapi-typescript openapi.yaml --output schema.ts
