resource "helm_release" "metallb" {
	name       = "metallb"
	repository = "https://metallb.github.io/metallb"
	chart      = "metallb"
	version    = var.metallb_helm_version

	namespace        = var.metallb_namespace
	create_namespace = true

	values = [
		file("${path.module}/metallb_values.yaml")
	]

	depends_on = [kind_cluster.default]
}
