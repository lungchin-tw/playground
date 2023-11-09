
variable "table_tf_state_locks" {
  description = "The DynamoDB table for terrafrom state lock"
  type        = string
  default     = "tf-state-locking"
}

variable "bucket_tf_state" {
  description = "The S3 bucket for terrafrom state file"
  type        = string
  default     = "tf-state"
}

variable "billing_mode" {
  description = "The prefix for each resource's naming"
  type        = string
  default     = "PAY_PER_REQUEST"
}
