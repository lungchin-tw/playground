
variable "prefix" {
  description = "prefix for resources' name"
  type        = string
}

variable "profile" {
  type = string
}


variable "region" {
  description = "The region the vpc is located"
  type        = string
}


variable "postgres_port" {
  description = "The region the vpc is located"
  type        = number
  default     = 5432
}

variable "upload_folder" {
  description = "Files will be uploaded to the Storage(S3 for now)"
  type        = string
  default     = null
}

variable "db_password" {
  description = "Password for Rumble Heroes PostgreSQL DB"
  type        = string
  sensitive   = true
}

variable "db_username" {
  description = "Username for Rumble Heroes PostgreSQL DB"
  type        = string
  sensitive   = true
}

variable "fdwdb_username" {
  description = "Username for Rumble Heroes PostgreSQL FDW DB"
  type        = string
  sensitive   = true
}

variable "env_name" {
  type = string
  default = ""
}