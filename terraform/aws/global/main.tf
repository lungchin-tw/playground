locals {
  project          = "jacky-demo"
  environment_name = "dev"
}


terraform {
  required_version = "~> 1.5"
  # Put this in the project

  # backend "s3" {
  #   bucket = "${local.project}-tf-state"
  #   key = "${local.environment_name}/terraform.tfstate"
  #   region = "eu-central-1"
  #   dynamodb_table = "${local.project}-tf-state-locking"
  #   encrypt = true
  # }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.12"
    }
  }
}

provider "aws" {
  region = "eu-central-1"
}

module "remote_backend" {
  source               = "../module/remote_backend"
  table_tf_state_locks = "${local.project}-tf-state-locking"
  bucket_tf_state      = "${local.project}-tf-state"
}

