data "aws_availability_zones" "available" {}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.1.2"

  name                 = "postgres-public"
  cidr                 = "10.0.0.0/16"
  azs                  = data.aws_availability_zones.available.names
  public_subnets       = ["10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24"]
  enable_dns_hostnames = true
  enable_dns_support   = true
}

resource "aws_db_subnet_group" "postgres" {
  name       = "postgres-public"
  subnet_ids = module.vpc.public_subnets
}

resource "aws_security_group" "postgres" {
  name   = "postgres-public"
  vpc_id = module.vpc.vpc_id

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

  password               = var.db_password
  db_subnet_group_name   = aws_db_subnet_group.postgres.name
  vpc_security_group_ids = [aws_security_group.postgres.id]
  parameter_group_name   = "default.postgres15"
  publicly_accessible    = true
  skip_final_snapshot    = true
}

