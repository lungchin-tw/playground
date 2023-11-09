terraform {
  source = "../../../module/network"

  before_hook "copy_tf_files" {
    commands = get_terraform_commands_that_need_locking()
    execute  = [
      "cp", 
      "${get_terragrunt_dir()}/../dbsetup/db-setup.tf", 
      "${get_terragrunt_dir()}/../dbsetup/sql-script-db-setup.tf", 
      "${get_terragrunt_dir()}/../dbsetup/sql-script-fdw-setup.tf", 
      "./"
    ]
  }

  after_hook "delete_data_files" {
    commands = get_terraform_commands_that_need_locking()
    execute  = ["rm", "-fr", "./sample"]
  }
  
  after_hook "copy_data_files" {
    commands = get_terraform_commands_that_need_locking()
    execute  = ["cp", "-R", "${get_terragrunt_dir()}/../../../res/sample", "."]
  }
}

# include "root" {
#   path = find_in_parent_folders()
# }

include "main" {
  path = find_in_parent_folders("main.hcl")
}


