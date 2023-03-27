schema {
  user "tax_calculator_extension_simple" {
    password = "tax_calculator_extension_simple"

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

  resource "income" {
    virtual = true

    property "gross_income" {
      type = "int32"
    }

    property "tax" {
      type = "int32"
    }

    property "net_income" {
      type = "int32"
    }

  }

  resource "tax_rate" {
    annotations {
      HclBlock = "tax_rate"
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

    property "city" {
      type     = "reference"
      required = false

      reference {
        referenced_resource = "city"
        cascade             = true
      }

      annotations {
        HclBlock = "city"
      }
    }

    property "order" {
      type = "int32"
    }

    property "until" {
      type     = "int32"
      required = true
    }

    property "rate" {
      type     = "float32"
      required = true
    }
  }

  extension "income_calculator" {
    namespace = "default"
    resource  = "income"

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