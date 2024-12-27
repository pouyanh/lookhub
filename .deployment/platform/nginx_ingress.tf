resource "helm_release" "ingress_nginx" {
	name       = "ingress-nginx"
	repository = "https://kubernetes.github.io/ingress-nginx"
	chart      = "ingress-nginx"
	version    = var.ingress_nginx_helm_version

	namespace        = var.ingress_nginx_namespace
	create_namespace = true

	values = [
		file("${path.module}/nginx_ingress_values.yaml")
	]

	depends_on = [kind_cluster.default]
}
