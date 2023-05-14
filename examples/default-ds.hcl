schema {
  data_source "default" {
    backend = "postgresql"

    postgresql_params {
      username = "dh_data"
      password = "dh_data"
      host     = "127.0.0.1"
      port     = 5432
      db_name  = "dh_data"
    }
  }
}
