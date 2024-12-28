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

resource "kubernetes_service_account" "dashboard_admin" {
	metadata {
		name      = "cpe"
		namespace = helm_release.kubernetes_dashboard.namespace
	}
}

resource "kubernetes_cluster_role_binding" "dashboard_admin" {
	metadata {
		name = "dashboard-admin-role-binding"
	}

	role_ref {
		api_group = "rbac.authorization.k8s.io"
		kind      = "ClusterRole"
		name      = kubernetes_cluster_role.admin.metadata[0].name
	}

	subject {
		kind      = "ServiceAccount"
		name      = kubernetes_service_account.dashboard_admin.metadata[0].name
		namespace = helm_release.kubernetes_dashboard.namespace
	}
}
