resource "aws_ecr_repository" "repo" {
  name                 = local.prefix
  image_tag_mutability = "MUTABLE"
  force_delete         = true
}
