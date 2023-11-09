locals {
  project          = "jacky"
  environment_name = "s3-cloudfront-demov2"
  prefix           = "${local.project}-${local.environment_name}"
  region           = "eu-central-1"
}

terraform {
  required_version = "~> 1.5"

  backend "s3" {
    bucket         = "jacky-demo-tf-state"
    key            = "jacky/s3-cloudfront-demov2/terraform.tfstate"
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
  profile = local.project
  region  = local.region
}

module "network" {
  source = "../module/network"
  region = local.region
}

