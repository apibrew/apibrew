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

  user "admin2" {
    password = "admin123"

    securityContext {
      constraint {
        property = "idx"
      }
    }
  }

  resource "country" {
    name = "country"

    annotations {
      HclBlock = "country"
    }

    property "name" {
      type   = "string"
      length = 124
      unique = true

      annotations {
        IsHclLabel = "true"
      }
    }

    property "description" {
      type   = "string"
      length = 124
    }

    property "population" {
      type = "int64"
    }

    property "area" {
      type = "int64"
    }

    extension {
        step = "BEFORE"
        action = "CREATE_UPDATE"
        execution = {
            script = {
                language = "nodejs"
                code = """
                    record.description = "test"
                """
            }
        }
    }
  }

  resource "state" {
    name = "state"

    annotations {
      HclBlock = "state"
    }

    property "name" {
      type   = "string"
      length = 124
      unique = true

      annotations {
        IsHclLabel = "true"
      }
    }

    property "country" {
      type     = "reference"
      required = true

      reference {
        referenced_resource = "country"
        cascade             = true
      }
    }

    property "description" {
      type   = "string"
      length = 124
    }
  }

  resource "city" {
    name = "city"

    annotations {
      HclBlock = "city"
    }

    property "name" {
      type   = "string"
      length = 124
      unique = true

      annotations {
        IsHclLabel = "true"
      }
    }

    property "country" {
      type     = "reference"
      required = true

      reference {
        referenced_resource = "country"
        cascade             = true
      }

      annotations {
        HclBlock = "country"
      }
    }

    property "state" {
      type = "reference"

      reference {
        referenced_resource = "state"
        cascade             = true
      }
    }

    property "description" {
      type   = "string"
      length = 124
    }
  }

  extension "income_calculator" {
    namespace = "default"
    resource  = "tax_rate"

    instead {
      create {
        function_call {
          host          = "127.0.0.1:37612"
          function_name = "income_calculator_calculate"
        }
      }
    }
  }
}

data {
  record "default" "country" {
    name        = "Azerbaijan"
    description = "Land of fire2"
    population  = 10000000
  }

  record "default" "country" {
    name        = "Georgia"
    description = "sample-description"
  }


  country "Spain" {
    description = "Country Spain"
    population  = 40000002
  }

  country "USA" {
    description = "Country USA"
    population  = 400000001
    area        = 12332123
  }

  city "Baku" {
    country {
      name = "Azerbaijan"
    }

    description = "City of Wind"
  }

  tax_rate "vat_rate" {
    country {
      name = "Azerbaijan"
    }

    order = 2
    until = 1000000000
    rate  = 0.18
  }
}
