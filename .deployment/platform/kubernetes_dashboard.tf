resource "helm_release" "kubernetes_dashboard" {
	name       = "kubernetes-dashboard"
	repository = "https://kubernetes.github.io/dashboard"
	chart      = "kubernetes-dashboard"
	version    = var.kubernetes_dashboard_helm_version

	namespace        = var.kubernetes_dashboard_namespace
	create_namespace = true

	values = [
		file("${path.module}/kubernetes_dashboard_values.yaml")
	]

	depends_on = [kind_cluster.default]
}
