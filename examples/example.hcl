data-source "default" {
  backend : "postgresql"
  params : {

  }
}

resource "country" {
  properties = [
    specialProperties(),
    {
      name   = "name"
      type   = string
      length = 124
    },
    {
      name   = "description"
      type   = string
      length = 124
    }
  ]
}

record "country" "Azerbaijan" {
  description = "Land of fire"
}

country "Georgia" {
  description = "Georgia"
}

resource "news-letter" {
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

