output "postgres_url" {
  value = "postgresql://${local.username}:@${aws_db_instance.postgres.endpoint}/${local.db_name}"
}

output "fdw_url" {
  value = "postgresql://${local.fdw_username}:@${aws_db_instance.postgres_fdw.endpoint}/${local.fdw_db_name}"
}