resource "aws_db_instance" "postgresql" {
  identifier_prefix          = "${local.prefix}-"
  instance_class             = "db.t4g.micro"
  multi_az                   = false
  storage_type               = "standard"
  allocated_storage          = 5
  storage_encrypted          = true
  engine                     = "postgres"
  engine_version             = "15.3"
  auto_minor_version_upgrade = false
  username                   = "hero"
  password                   = var.db_password
  db_name                    = "rumbleheroes"
  parameter_group_name       = "default.postgres15"
  publicly_accessible        = false
  skip_final_snapshot        = true
  apply_immediately          = true
  backup_retention_period    = 1
}

resource "aws_db_instance" "postgresql_replica" {
  identifier_prefix    = "${local.prefix}-replica-"
  replicate_source_db  = aws_db_instance.postgresql.identifier
  instance_class       = "db.t4g.micro"
  publicly_accessible  = false
  skip_final_snapshot  = true
  apply_immediately    = true
  parameter_group_name = "default.postgres15"

  # disable backups to create DB faster
  backup_retention_period = 0
}