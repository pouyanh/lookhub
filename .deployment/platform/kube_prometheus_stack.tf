resource "helm_release" "kube_prometheus_stack" {
	name       = "kube-prometheus-stack"
	repository = "https://prometheus-community.github.io/helm-charts"
	chart      = "kube-prometheus-stack"
	version    = var.kube_prometheus_stack_helm_version

	namespace        = var.kube_prometheus_stack_namespace
	create_namespace = true

	values = [
		file("${path.module}/kube_prometheus_values.yaml")
	]

	depends_on = [kind_cluster.default]
}
