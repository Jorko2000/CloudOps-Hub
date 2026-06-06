resource "helm_release" "redis" {

  name = var.redis_name

  repository = "https://charts.bitnami.com/bitnami"

  chart = "redis"

  namespace = var.namespace
}
