locals {
  project          = "dynamodb-demo"
  environment_name = "jacky"
  prefix           = "${local.project}-${local.environment_name}"
}

terraform {
  required_version = "~> 1.5"

  backend "s3" {
    bucket         = "jacky-demo-tf-state"
    key            = "jacky/terraform.tfstate"
    region         = "eu-central-1"
    dynamodb_table = "jacky-demo-tf-state-locking"
    encrypt        = true
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.12"
    }
  }
}

provider "aws" {
  profile = local.environment_name
  region  = "eu-central-1"
}

