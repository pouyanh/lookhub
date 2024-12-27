resource "kind_cluster" "default" {
	name = "cpe-cluster"
	wait_for_ready = true

	kind_config {
		api_version = "kind.x-k8s.io/v1alpha4"
		kind        = "Cluster"

		node {
			role = "control-plane"
		}

		node {
			role = "worker"
		}

		node {
			role = "worker"
		}
	}
}
