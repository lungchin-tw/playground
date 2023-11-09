locals {
  result = {
    for k, v in fileset("res/", "*") : k => v
  }
}

output "files" {
  value = local.result
}



