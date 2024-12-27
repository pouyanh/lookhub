provider "helm" {
	kubernetes {
		config_path = kind_cluster.default.kubeconfig_path
	}
}
