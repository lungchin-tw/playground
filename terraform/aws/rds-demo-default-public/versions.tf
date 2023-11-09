terraform {
  required_version = "~> 1.5"
  
  required_providers {
    aws = {
      #   source  = "hashicorp/aws" # Default, can be omitted
      version = "~> 5.16"
    }

    local = {
      # source  = "hashicorp/local" # Default, can be omitted
      version = "~> 2.4"
    }

    null = {
      # source  = "hashicorp/null" # Default, can be omitted
      version = "~> 3.2"
    }
  }
}

