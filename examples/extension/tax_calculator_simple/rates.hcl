data {
  country "Azerbaijan" {}

  city "Baku" {
    country {
      name = "Azerbaijan"
    }
  }

  city "Sumgayit" {
    country {
      name = "Azerbaijan"
    }
  }

  tax_rate "simple_rate" {
    country {
      name = "Azerbaijan"
    }

    order = 1
    until = 120000
    rate  = 0.05
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