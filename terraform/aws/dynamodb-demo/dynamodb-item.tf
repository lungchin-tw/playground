resource "aws_dynamodb_table_item" "table1_item1" {
  depends_on = [aws_dynamodb_table.table1]

  table_name = aws_dynamodb_table.table1.name
  hash_key   = aws_dynamodb_table.table1.hash_key
  range_key  = aws_dynamodb_table.table1.range_key

  item = <<ITEM
    {
        "id": {"S": "card01"},
        "card": {"N": "1"},
        "uid": {"S": "1234"},
        "uname": {"S": "User1234"}
    }
    ITEM
}