
data "aws_iam_role" "example" {
  name = "aws-elasticbeanstalk-ec2-role"
}

data "aws_iam_policy" "ecr" {
  name = "AmazonEC2ContainerRegistryReadOnly"
}

data "aws_iam_policy" "eb_web" {
  name = "AWSElasticBeanstalkWebTier"
}

data "aws_iam_policy" "eb_worker" {
  name = "AWSElasticBeanstalkWorkerTier"
}

data "aws_iam_policy" "eb_docker" {
  name = "AWSElasticBeanstalkMulticontainerDocker"
}

output "role_policy" {
  value = data.aws_iam_role.example.assume_role_policy
}

output "role_id" {
  value = data.aws_iam_role.example.id
}

output "role_arn" {
  value = data.aws_iam_role.example.arn
}

output "role_path" {
  value = data.aws_iam_role.example.path
}

output "role_permissions" {
  value = data.aws_iam_role.example.permissions_boundary
}

output "ecr_policy" {
  value = data.aws_iam_policy.ecr.policy
}


resource "aws_iam_role" "newrole" {
  name               = local.prefix
  assume_role_policy = data.aws_iam_role.example.assume_role_policy

}

resource "aws_iam_role_policy_attachment" "policy_attach" {
  for_each = toset([
    data.aws_iam_policy.ecr.arn,
    data.aws_iam_policy.eb_web.arn,
    data.aws_iam_policy.eb_worker.arn,
    data.aws_iam_policy.eb_docker.arn,
  ])

  role       = aws_iam_role.newrole.name
  policy_arn = each.value
}
