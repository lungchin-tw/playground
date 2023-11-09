
resource "aws_default_vpc" "default" {
}

data "aws_subnets" "default" {
  filter {
    name   = "vpc-id"
    values = [aws_default_vpc.default.id]
  }
}

resource "aws_default_subnet" "default_a" {
  availability_zone = "${var.region}a"
}

resource "aws_default_subnet" "default_b" {
  availability_zone = "${var.region}b"
}

resource "aws_default_subnet" "default_c" {
  availability_zone = "${var.region}c"
}