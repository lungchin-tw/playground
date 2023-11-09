

resource "aws_elasticache_subnet_group" "default" {
  name       = "vpc-${local.prefix}"
  subnet_ids = data.aws_subnets.default.ids
}


resource "aws_elasticache_cluster" "redis_cluster" {
  cluster_id                 = local.prefix
  engine                     = "redis"
  node_type                  = "cache.t4g.micro"
  num_cache_nodes            = 1
  parameter_group_name       = "default.redis7"
  apply_immediately          = true
  auto_minor_version_upgrade = false
  engine_version             = "7.0"
  port                       = 6379

  subnet_group_name = aws_elasticache_subnet_group.default.name
}

