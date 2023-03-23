data-source "default" {
  backend          = "postgresql"
  postgresqlParams = {
    username = "dh_test"
    password = "dh_test"
    host     = "127.0.0.1"
    port     = 5432
    db_name  = "dh_test"
  }
}

resource "country" {
  name = "country"
  properties = [
    specialProperties(),
    {
      name   = "name"
      type   = "string"
      length = 124
    },
    {
      name   = "description"
      type   = "string"
      length = 124
    }
  ]
}

record "country" {
  description = "Land of fire"
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
