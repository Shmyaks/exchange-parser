schema "public" {}

table "arbitrage" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "asset"{
    null = false
    type = varchar(32)
  }
  column "fiat"{
    null = false
    type = varchar(8)
  }
  column "buy"{
    null = false
    type = float
  }
  column "first_pay_type"{
    null = false
    type = varchar(32)
  }
  column "sell"{
    null = false
    type = float
  }
  column "second_pay_type"{
    null = false
    type = varchar(32)
  }
  column "percent" {
    null = false
    type = float
  }
  column "market_id" {
    null = false
    type = int
  }
  primary_key {
    columns = [
      column.id
    ]
  }
}