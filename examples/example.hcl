var1 = "test"


data-source "dhTest" {
  backend : "mysql"
  mysqlParams : {

  }
}

resource "country" {
  data_source = "dhTest ${var1}"
  entity      = "country"
  properties  = [
    specialProperties(),
    {
      name    = "name"
      type    = string
      length  = 124
      primary = false
    },
    {
      name    = "description"
      type    = string
      length  = 124
      primary = false
    }
  ]
  annotations = {
    KeepHistory : true
  }
}

country "Azerbaijan" {
  description = "Land of fire"
}

country "Georgia" {
  description = "Georgia"
}

resource "news-letter" {
  data_source = "dhTest"
  entity      = "news-letter"
  virtual     = false
  abstract    = false
  properties  = [
    specialProperties(),
    {
      name    = "email"
      type    = string
      length  = 124
      primary = false
    }
  ]
  annotations = {
    KeepHistory : true
  }
  securityContext = {
    constraints = {}
  }
}

extension "NewsLetterRegistration" {
  resource = "news-letter"
  after {
    sync = true
    exec = {
      http = {
        uri : "http://my-service-backend:1234/abcx"
        method : "POST"
        body : {

        }
      }
    }
  }
}

