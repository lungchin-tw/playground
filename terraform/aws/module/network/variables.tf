
variable "region" {
  description = "The region the vpc is located"
  type        = string
}

variable "postgres_port" {
  description = "The region the vpc is located"
  type        = number
  default     = 5432
}