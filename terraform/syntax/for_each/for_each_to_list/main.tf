locals {
  result = [
    for key, value in fileset("res/", "*") : value
  ]
}

output "files" {
  value = local.result
}



