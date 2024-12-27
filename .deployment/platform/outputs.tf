output "kubeconfig_path" {
	value = kind_cluster.default.kubeconfig_path
}

output "endpoint" {
	value = kind_cluster.default.endpoint
}

output "lookhub_ci_user" {
	value = kubernetes_service_account.lookhub_ci.metadata[0].name
}

output "lookhub_ci_token" {
	sensitive = true
	value = lookup(kubernetes_secret.lookhub_ci_token.data, "token")
}
