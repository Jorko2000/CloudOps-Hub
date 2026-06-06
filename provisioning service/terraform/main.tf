module "namespace" {

  source = "./modules/namespace"

  namespace = var.namespace
}

module "postgres" {

  source = "./modules/postgres"

  namespace = var.namespace

  postgres_name = var.postgres_name
}

module "redis" {

  source = "./modules/redis"

  namespace = var.namespace

  redis_name = var.redis_name
}
