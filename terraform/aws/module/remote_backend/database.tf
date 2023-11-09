resource "aws_dynamodb_table" "tf_state_locks" {
  name         = var.table_tf_state_locks
  billing_mode = var.billing_mode
  hash_key     = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}