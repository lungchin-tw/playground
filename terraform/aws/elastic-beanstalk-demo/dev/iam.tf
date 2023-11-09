
data "aws_iam_role" "eb_app_service_role" {
  name = "aws-elasticbeanstalk-service-role"
}

data "aws_iam_role" "aws_eb_ec2_role" {
  name = "aws-elasticbeanstalk-ec2-role"
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

data "aws_iam_policy" "ecr" {
  name = "AmazonEC2ContainerRegistryReadOnly"
}

resource "aws_iam_role" "eb_ec2_role" {
  name               = local.prefix
  assume_role_policy = data.aws_iam_role.aws_eb_ec2_role.assume_role_policy
}

resource "aws_iam_role_policy_attachment" "eb_ec2_role_policy_attach" {
  for_each = toset([
    data.aws_iam_policy.eb_web.arn,
    data.aws_iam_policy.eb_worker.arn,
    data.aws_iam_policy.eb_docker.arn,
    data.aws_iam_policy.ecr.arn,
  ])

  role       = aws_iam_role.eb_ec2_role.name
  policy_arn = each.value
}
