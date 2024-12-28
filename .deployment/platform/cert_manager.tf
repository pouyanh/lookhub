resource "helm_release" "cert_manager" {
	name       = "cert-manager"
	repository = "https://charts.jetstack.io"
	chart      = "cert-manager"
	version    = var.cert_manager_helm_version

	namespace        = var.cert_manager_namespace
	create_namespace = true

	values = [
		file("${path.module}/cert_manager_values.yaml")
	]

	depends_on = [kind_cluster.default]
}
