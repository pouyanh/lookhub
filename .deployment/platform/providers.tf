terraform {
	required_providers {
		kind = {
			source  = "tehcyx/kind"
			version = "0.7.0"
		}

		kubernetes = {
			source  = "hashicorp/kubernetes"
			version = "~> 2.20"
		}
	}
}

provider "kind" {}

provider "kubernetes" {
	config_path = kind_cluster.default.kubeconfig_path
}
