resource "kubernetes_namespace" "monitoring" {
	metadata {
		name = "monitoring"
	}
}

resource "helm_release" "metrics_server" {
	name       = "metrics-server"
	repository = "https://kubernetes-sigs.github.io/metrics-server"
	chart      = "metrics-server"
	version    = var.metrics_server_helm_version

	namespace = kubernetes_namespace.monitoring.metadata[0].name

	values = [
		file("${path.module}/metrics_server_values.yaml")
	]

	depends_on = [kind_cluster.default]
}

resource "helm_release" "kube_prometheus_stack" {
	name       = "kube-prometheus-stack"
	repository = "https://prometheus-community.github.io/helm-charts"
	chart      = "kube-prometheus-stack"
	version    = var.kube_prometheus_stack_helm_version

	namespace = kubernetes_namespace.monitoring.metadata[0].name

	values = [
		file("${path.module}/kube_prometheus_values.yaml")
	]

	depends_on = [kind_cluster.default]
}

data "kubernetes_service" "grafana" {
	metadata {
		name      = "kube-prometheus-stack-grafana"
		namespace = kubernetes_namespace.monitoring.metadata[0].name
	}
}

data "kubernetes_service" "prometheus" {
	metadata {
		name      = "kube-prometheus-stack-prometheus"
		namespace = kubernetes_namespace.monitoring.metadata[0].name
	}
}
