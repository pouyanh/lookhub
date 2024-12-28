resource "helm_release" "metrics_server" {
	name       = "metrics-server"
	repository = "https://kubernetes-sigs.github.io/metrics-server"
	chart      = "metrics-server"
	version    = var.metrics_server_helm_version

	namespace        = var.metrics_server_namespace
	create_namespace = true

	values = [
		file("${path.module}/metrics_server_values.yaml")
	]

	depends_on = [kind_cluster.default]
}
