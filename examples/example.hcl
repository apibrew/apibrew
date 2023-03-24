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
  }

  property "description" {
    type   = "string"
    length = 124
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

record "default" "rich-test-3995" {
  date      = "2022-01-03"
  time      = "12:03"
  timestamp = "2022-01-03 12:03"
  bool      = false
  bytes     = ""
  int32     = 123
  int64     = 123
  float     = 231
  double    = 123
  string    = "asdasdsa"
  text      = "asdasdsa"
  uuid      = "1945d115-80d5-4cda-abd3-ac636ab60184"
  object {
    abc = 123
    cde = {
      asd = "asdasd"
    }
  }

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
