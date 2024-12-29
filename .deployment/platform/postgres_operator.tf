resource "helm_release" "postgres_operator" {
	name       = "postgres-operator"
	repository = "https://opensource.zalando.com/postgres-operator/charts/postgres-operator"
	chart      = "postgres-operator"
	version    = var.postgres_operator_helm_version

	namespace        = var.postgres_operator_namespace
	create_namespace = true

	values = [
		file("${path.module}/postgres_operator_values.yaml")
	]

	depends_on = [kind_cluster.default]
}
