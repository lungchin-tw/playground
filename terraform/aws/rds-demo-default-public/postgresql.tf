
resource "aws_db_subnet_group" "postgres" {
  name       = "postgres-public"
  subnet_ids = module.network.aws_default_subnet_ids
}

resource "aws_security_group" "postgres" {
  name   = "postgres-public"
  vpc_id = module.network.aws_default_vpc_id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_db_instance" "postgres" {
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
  db_name                    = "rumbleheroes"
  port                       = 5432

  password               = var.db_password
  db_subnet_group_name   = aws_db_subnet_group.postgres.name
  vpc_security_group_ids = [aws_security_group.postgres.id]
  parameter_group_name   = "default.postgres15"
  publicly_accessible    = true
  # publicly_accessible    = false
  skip_final_snapshot = true
}

resource "null_resource" "db_setup" {
  triggers = {
    sql_script_content = local_file.sql_script.content
  }

  provisioner "local-exec" {
    command = "psql postgresql://hero:${var.db_password}@${aws_db_instance.postgres.endpoint}/rumbleheroes -f ${local_file.sql_script.filename}"
  }
}

