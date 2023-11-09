resource "aws_db_subnet_group" "postgres_default" {
  name       = "postgres-default-${var.prefix}"
  subnet_ids = data.aws_subnets.default.ids
}

resource "aws_security_group" "postgres_default" {
  name   = "postgres-default-${var.prefix}"
  vpc_id = aws_default_vpc.default.id

  ingress {
    from_port   = var.postgres_port
    to_port     = var.postgres_port
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = var.postgres_port
    to_port     = var.postgres_port
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}