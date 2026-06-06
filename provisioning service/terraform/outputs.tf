output "namespace" {
  value = module.namespace.namespace_name
}

output "postgres_release" {
  value = module.postgres.release_name
}

output "redis_release" {
  value = module.redis.release_name
}
