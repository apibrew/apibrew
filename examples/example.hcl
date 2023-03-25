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
  id        = "sample-id"
  name      = "country"
  namespace = "default"

  source_config {
    dataSource = "default"
    catalog    = "public"
    entity     = "country"
  }

  annotations {
    test  = "test2"
    test2 = "test2"
  }

  property "name" {
    type   = "string"
    length = 124
    unique = true
  }

  property "description" {
    type   = "string"
    length = 124
  }
}

record "default" "country" {
  name        = "Azerbaijan"
  description = "Land of fire"
}

record "default" "country" {
  name        = "Georgia"
  description = "sample-description"
}

#
#country "Georgia" {
#  description = "Georgia"
#}
#
#resource "news-letter" {
#  virtual    = false
#  abstract   = false
#  properties = [
#    specialProperties(),
#    {
#      name    = "email"
#      type    = string
#      length  = 124
#      primary = false
#    }
#  ]
#}
#
#extension "NewsLetterRegistration" {
#  resource = "news-letter"
#  after {
#    sync = true
#    exec = {
#      http = {
#        uri : "http://my-service-backend:1234/abcx"
#        method : "POST"
#        body : {}
#      }
#    }
#  }
#}
#
