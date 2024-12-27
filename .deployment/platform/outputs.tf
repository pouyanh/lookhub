output "kubeconfig_path" {
	value = kind_cluster.default.kubeconfig_path
}

output "endpoint" {
	value = kind_cluster.default.endpoint
}

output "ci_user" {
	value = kubernetes_service_account.gitlab_ci.metadata[0].name
}

output "ci_token" {
	sensitive = true
	value = lookup(kubernetes_secret.gitlab_ci_token.data, "token")
}
