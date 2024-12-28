resource "helm_release" "argo_cd" {
	name       = "argo-cd"
	repository = "https://argoproj.github.io/argo-helm"
	chart      = "argo-cd"
	version    = var.argo_cd_helm_version

	namespace        = var.argo_cd_namespace
	create_namespace = true

	values = [
		file("${path.module}/argo_cd_values.yaml")
	]

	depends_on = [kind_cluster.default]
}
