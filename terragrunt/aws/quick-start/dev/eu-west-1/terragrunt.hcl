terraform {
  source = "../../../module/network"
}

# include "root" {
#   path = find_in_parent_folders()
# }

include "main" {
  path = find_in_parent_folders("main.hcl")
}