resource "kubernetes_namespace" "lookhub" {
	metadata {
		name = "lookhub"
	}

	depends_on = [kind_cluster.default]
}

resource "kubernetes_role" "lookhub_ns_owner" {
	metadata {
		name      = "namespace-owner-role"
		namespace = kubernetes_namespace.lookhub.metadata[0].name
	}

	rule {
		api_groups = [""]  # Core resources
		resources  = ["*"] # All resources
		verbs      = ["*"] # Full access
	}

	rule {
		api_groups = ["apps"] # Apps group
		resources  = ["*"]
		verbs      = ["*"]
	}

	rule {
		api_groups = ["batch"] # Batch resources (CronJobs, Jobs)
		resources  = ["*"]
		verbs      = ["*"]
	}

	rule {
		api_groups = ["rbac.authorization.k8s.io"] # RBAC resources
		resources  = ["roles", "rolebindings"]
		verbs      = ["*"]
	}
}

resource "kubernetes_role" "lookhub_ci" {
	metadata {
		name      = "ci-role"
		namespace = kubernetes_namespace.lookhub.metadata[0].name
	}

	rule {
		api_groups = [""] # Core resources
		resources  = ["pods", "services", "configmaps", "secrets"]
		verbs      = ["get", "list", "create", "update", "delete"]
	}

	rule {
		api_groups = ["apps"] # Apps group
		resources  = ["deployments", "statefulsets", "daemonsets", "replicasets"]
		verbs      = ["get", "list", "create", "update", "delete"]
	}

	rule {
		api_groups = ["batch"] # Batch resources
		resources  = ["jobs", "cronjobs"]
		verbs      = ["get", "list", "create", "update", "delete"]
	}
}

resource "kubernetes_role" "lookhub_pods_reader" {
	metadata {
		name      = "pods-reader-role"
		namespace = kubernetes_namespace.lookhub.metadata[0].name
	}

	rule {
		api_groups = [""] # Core resources
		resources  = ["pods"]
		verbs      = ["get", "list", "watch"]
	}
}

resource "kubernetes_service_account" "lookhub_ci" {
	metadata {
		name      = "ci"
		namespace = kubernetes_namespace.lookhub.metadata[0].name
	}
}

resource "kubernetes_secret" "lookhub_ci_token" {
	metadata {
		annotations = {
			"kubernetes.io/service-account.name" = kubernetes_service_account.lookhub_ci.metadata[0].name
		}

		generate_name = "ci-"
		namespace     = kubernetes_namespace.lookhub.metadata[0].name
	}

	type                           = "kubernetes.io/service-account-token"
	wait_for_service_account_token = true
}

resource "kubernetes_role_binding" "lookhub_ns_owner" {
	metadata {
		name      = "namespace-owner-role-binding"
		namespace = kubernetes_namespace.lookhub.metadata[0].name
	}

	role_ref {
		api_group = "rbac.authorization.k8s.io"
		kind      = "Role"
		name      = kubernetes_role.lookhub_ns_owner.metadata[0].name
	}

	subject {
		kind      = "User"
		name      = "cto"
		api_group = ""
	}
}

resource "kubernetes_role_binding" "lookhub_ci" {
	metadata {
		name      = "ci-role-binding"
		namespace = kubernetes_namespace.lookhub.metadata[0].name
	}

	role_ref {
		api_group = "rbac.authorization.k8s.io"
		kind      = "Role"
		name      = kubernetes_role.lookhub_ci.metadata[0].name
	}

	subject {
		kind      = "ServiceAccount"
		name      = kubernetes_service_account.lookhub_ci.metadata[0].name
		namespace = kubernetes_namespace.lookhub.metadata[0].name
	}
}

resource "kubernetes_role_binding" "lookhub_pods_reader" {
	metadata {
		name      = "pods-reader-role-binding"
		namespace = kubernetes_namespace.lookhub.metadata[0].name
	}

	role_ref {
		api_group = "rbac.authorization.k8s.io"
		kind      = "Role"
		name      = kubernetes_role.lookhub_pods_reader.metadata[0].name
	}

	subject {
		kind      = "User"
		name      = "qa"
		api_group = ""
	}

	subject {
		kind      = "User"
		name      = "pm"
		api_group = ""
	}
}
