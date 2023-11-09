
locals {
  username = "mock"
  db_name = "mockdb"
}

resource "aws_db_instance" "postgres" {
  depends_on = [module.network]

  identifier_prefix          = "${local.prefix}-"
  instance_class             = "db.t4g.micro"
  multi_az                   = false
  storage_type               = "standard"
  allocated_storage          = 5
  storage_encrypted          = true
  engine                     = "postgres"
  engine_version             = "15.3"
  auto_minor_version_upgrade = false
  username                   = local.username
  password                   = var.db_password
  db_name                    = local.db_name
  port                       = module.network.postgres_port
  db_subnet_group_name       = module.network.postgres_default_subnet_group_name
  vpc_security_group_ids     = module.network.postgres_default_security_group_id
  parameter_group_name       = "default.postgres15"
  publicly_accessible        = true
  skip_final_snapshot        = true
  apply_immediately          = true
  # backup_retention_period    = 1
}

resource "null_resource" "db_setup" {
  triggers = {
    sql_script_content = local_file.sql_script_db_setup.content
  }

  provisioner "local-exec" {
    command = "psql postgresql://${local.username}:${var.db_password}@${aws_db_instance.postgres.endpoint}/${local.db_name} -f ${local_file.sql_script_db_setup.filename}"
  }
}



# resource "aws_db_instance" "postgresql_replica" {
#   identifier_prefix    = "${local.prefix}-replica-"
#   replicate_source_db  = aws_db_instance.postgresql.identifier
#   instance_class       = "db.t4g.micro"
#   publicly_accessible  = false
#   skip_final_snapshot  = true
#   apply_immediately    = true
#   parameter_group_name = "default.postgres15"

#   # disable backups to create DB faster
#   backup_retention_period = 0
# }