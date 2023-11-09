output "arn" {
  value = aws_elastic_beanstalk_application.eb_app.arn
}

output "id" {
  value = aws_elastic_beanstalk_environment.eb_app_env.id
}

output "name" {
  value = aws_elastic_beanstalk_environment.eb_app_env.name
}

output "description" {
  value = aws_elastic_beanstalk_environment.eb_app_env.description
}

output "tier" {
  value = aws_elastic_beanstalk_environment.eb_app_env.tier
}

output "application" {
  value = aws_elastic_beanstalk_environment.eb_app_env.application
}

output "setting" {
  value = aws_elastic_beanstalk_environment.eb_app_env.setting
}

output "all_setting" {
  value = aws_elastic_beanstalk_environment.eb_app_env.setting
}

output "cname" {
  value = aws_elastic_beanstalk_environment.eb_app_env.cname
}

output "autoscaling_groups" {
  value = aws_elastic_beanstalk_environment.eb_app_env.autoscaling_groups
}

output "instances" {
  value = aws_elastic_beanstalk_environment.eb_app_env.instances
}

output "launch_configurations" {
  value = aws_elastic_beanstalk_environment.eb_app_env.launch_configurations
}

output "load_balancers" {
  value = aws_elastic_beanstalk_environment.eb_app_env.load_balancers
}

output "queues" {
  value = aws_elastic_beanstalk_environment.eb_app_env.queues
}

output "triggers" {
  value = aws_elastic_beanstalk_environment.eb_app_env.triggers
}

output "endpoint_url" {
  value = aws_elastic_beanstalk_environment.eb_app_env.endpoint_url
}