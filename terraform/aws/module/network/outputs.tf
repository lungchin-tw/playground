output "aws_default_vpc_id" {
  value = aws_default_vpc.default.id
}

output "aws_default_subnet_ids" {
  value = data.aws_subnets.default.ids
}

output "postgres_port" {
  value = var.postgres_port
}

output "postgres_default_subnet_group_name" {
  value = aws_db_subnet_group.postgres_default.name
}

output "postgres_default_security_group_id" {
  value = [ aws_security_group.postgres_default.id ]
}