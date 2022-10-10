package main

import "data-handler/stub/model"

// main static data
func main() {
	systemDataSource := model.DataSource{
		Backend: model.DataSourceBackend_POSTGRESQL,
		Type:    model.DataType_SYSTEM,
		Options: model.,
	}
}
