resource "helm_release" "grafana_operator" {
	name       = "grafana-operator"
	repository = "oci://ghcr.io/grafana/helm-charts"
	chart      = "grafana-operator"
	version    = var.grafana_operator_helm_version

	namespace        = var.grafana_operator_namespace
	create_namespace = true

	values = [
		file("${path.module}/grafana_operator_values.yaml")
	]

	depends_on = [kind_cluster.default]
}
