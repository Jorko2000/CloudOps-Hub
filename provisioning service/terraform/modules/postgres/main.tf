resource "helm_release" "postgres" {

  name = var.postgres_name

  repository = "https://charts.bitnami.com/bitnami"

  chart = "postgresql"

  namespace = var.namespace
}
