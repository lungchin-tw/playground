locals {
  project          = "jacky"
  environment_name = "elastic-beanstalk-demo-dev"
  prefix           = "${local.project}-${local.environment_name}"
  region           = "eu-central-1"
}


terraform {
  required_version = "~> 1.5"

  backend "s3" {
    bucket         = "jacky-demo-tf-state"
    key            = "jacky/elastic-beanstalk-demo-dev/terraform.tfstate"
    region         = "eu-central-1"
    dynamodb_table = "jacky-demo-tf-state-locking"
    encrypt        = true
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.12"
    }
  }
}

provider "aws" {
  profile = local.project
  region  = local.region
}


resource "aws_iam_instance_profile" "eb_ec2" {
  name = local.prefix
  role = aws_iam_role.eb_ec2_role.name
}


resource "aws_elastic_beanstalk_application" "eb_app" {
  name        = local.project
  description = "An elastic beanstalk application."

  appversion_lifecycle {
    service_role          = data.aws_iam_role.eb_app_service_role.arn
    max_count             = 128
    delete_source_from_s3 = true
  }
}

resource "aws_elastic_beanstalk_environment" "eb_app_env" {
  name                = local.environment_name
  application         = aws_elastic_beanstalk_application.eb_app.name
  cname_prefix        = local.prefix
  solution_stack_name = "64bit Amazon Linux 2023 v4.0.0 running Docker"
  tier                = "WebServer"

  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "IamInstanceProfile"
    value     = aws_iam_instance_profile.eb_ec2.name
  }

  setting {
    namespace = "aws:elasticbeanstalk:environment"
    name      = "ServiceRole"
    value     = data.aws_iam_role.eb_app_service_role.name
  }

  setting {
    namespace = "aws:autoscaling:launchconfiguration"
    name      = "InstanceType"
    value     = "t3.micro"
  }

  setting {
    namespace = "aws:elasticbeanstalk:environment"
    name      = "LoadBalancerType"
    value     = "application"
  }

  setting {
    namespace = "aws:autoscaling:asg"
    name      = "MinSize"
    value     = 1
  }
  setting {
    namespace = "aws:autoscaling:asg"
    name      = "MaxSize"
    value     = 1
  }
}

# resource "aws_s3_bucket" "eb_app_bucket" {
#   bucket = local.project
#   acl    = "private"
# }


# resource "aws_elastic_beanstalk_application_version" "eb_app_version" {
#   name        = "latest"
#   application = aws_elastic_beanstalk_application.eb_app.name
#   description = "The latest version of this application"
#   bucket      = aws_s3_bucket.eb_app_bucket.id
# }




