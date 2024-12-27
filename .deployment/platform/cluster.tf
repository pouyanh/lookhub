provider "kind" {}

provider "kubernetes" {
	config_path = kind_cluster.default.kubeconfig_path
}

resource "kind_cluster" "default" {
	name            = var.cluster_name
	kubeconfig_path = "${path.module}/kubeconfig.yaml"

	kind_config {
		api_version = "kind.x-k8s.io/v1alpha4"
		kind        = "Cluster"

		node {
			role = "control-plane"

			kubeadm_config_patches = [
				<<-EOT
				kind: KubeletConfiguration
				serverTLSBootstrap: true
				EOT
			,

				<<-EOT
				kind: ClusterConfiguration
				apiServer:
				  extraArgs:
				    authorization-mode: "Node,RBAC"
				EOT
			,

				<<-EOT
				kind: InitConfiguration
				nodeRegistration:
				  kubeletExtraArgs:
				    node-labels: "ingress-ready=true"
				EOT
			]
		}

		node {
			role = "worker"
		}

		node {
			role = "worker"
		}
	}
}

resource "kubernetes_cluster_role" "admin" {
	metadata {
		name = "cluster-admin-role"
	}

	rule {
		api_groups = [""]
		resources  = ["*"]
		verbs      = ["*"]
	}
}

resource "kubernetes_cluster_role_binding" "admin" {
	metadata {
		name = "cluster-admin-role-binding"
	}

	role_ref {
		api_group = "rbac.authorization.k8s.io"
		kind      = "ClusterRole"
		name      = kubernetes_cluster_role.admin.metadata[0].name
	}

	subject {
		kind      = "User"
		name      = "cpe"
		api_group = ""
	}
}
