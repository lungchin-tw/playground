resource "aws_dynamodb_table" "table1" {
  name           = local.prefix
  billing_mode   = "PROVISIONED"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "id"
  range_key      = "card"

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "card"
    type = "N"
  }

  attribute {
    name = "uid"
    type = "S"
  }

  global_secondary_index {
    name     = "GSI-01"
    hash_key = "uid"

    read_capacity  = 1
    write_capacity = 1

    projection_type = "ALL"
  }
}
