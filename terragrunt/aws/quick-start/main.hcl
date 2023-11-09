locals {
  project = "quick-start"
  region_vars  = read_terragrunt_config("region.hcl")
  profile_vars = read_terragrunt_config(find_in_parent_folders("profile.hcl"))
  prefix       = "${local.project}-${local.profile_vars.locals.profile_name}-${local.profile_vars.locals.env_name}-${local.region_vars.locals.region}"
}

generate "provider" {
  path      = "provider.tf"
  if_exists = "overwrite"
  contents  = <<EOF
provider "aws" {
    profile = "${local.profile_vars.locals.profile_name}"
    region  = "${local.region_vars.locals.region}"
}
EOF
}

remote_state {
  backend = "s3"
  generate = {
    path      = "backend.tf"
    if_exists = "overwrite"
  }
  config = {
    bucket         = "jacky-demo-tf-state"
    key            = "jacky/${local.project}/${path_relative_to_include()}/terraform.tfstate"
    region         = "eu-central-1"
    dynamodb_table = "jacky-demo-tf-state-locking"
    encrypt        = true
  }
}


inputs = {
  profile        = local.profile_vars.locals.profile_name
  env_name    = local.profile_vars.locals.env_name
  region         = local.region_vars.locals.region
  prefix         = local.prefix
  upload_folder  = "sample"
  db_password    = get_env("TF_VAR_db_password")
  db_username    = get_env("TF_VAR_db_username")
  fdwdb_username = get_env("TF_VAR_fdwdb_username")
}
