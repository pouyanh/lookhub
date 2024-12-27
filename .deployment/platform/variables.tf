variable "cluster_name" {
	description = "Name of the Kubernetes cluster"
	type        = string
	default     = "cpe-cluster"
}

variable "ingress_nginx_helm_version" {
	description = "The Helm version for the nginx ingress controller."
	type        = string
	default     = "4.11.3"
}

variable "ingress_nginx_namespace" {
	description = "The nginx ingress namespace (it will be created if needed)."
	type        = string
	default     = "ingress-nginx"
}

variable "metrics_server_helm_version" {
	description = "The Helm version for the metrics-server."
	type        = string
	default     = "3.12.2"
}

variable "kube_prometheus_stack_helm_version" {
	description = "The Helm version for the kube-prometheus-stack."
	type        = string
	default     = "67.5.0"
}
