locals {
  project          = "jacky"
  environment_name = "rds-demo-default-public"
  prefix           = "${local.project}-${local.environment_name}"
  region           = "eu-central-1"
}

terraform {
  backend "s3" {
    bucket         = "jacky-demo-tf-state"
    key            = "jacky/rds-demo-default-public/terraform.tfstate"
    region         = "eu-central-1"
    dynamodb_table = "jacky-demo-tf-state-locking"
    encrypt        = true
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


