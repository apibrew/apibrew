schema {
  data_source "default" {
    backend = "postgresql"

    postgresql_params {
      username = "dh_test"
      password = "dh_test"
      host     = "127.0.0.1"
      port     = 5432
      db_name  = "dh_test"
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
    name      = "country"

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
    area  = 12332123
  }
}
