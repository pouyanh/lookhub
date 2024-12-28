variable "cluster_name" {
	description = "Name of the Kubernetes cluster"
	type        = string
	default     = "cpe-cluster"
}

variable "cert_manager_helm_version" {
	description = "The Helm version for the cert manager."
	type        = string
	default     = "1.16.2"
}

variable "cert_manager_namespace" {
	description = "The cert manager namespace (it will be created if needed)."
	type        = string
	default     = "cert-manager"
}

variable "kubernetes_dashboard_helm_version" {
	description = "The Helm version for the kubernetes dashboard."
	type        = string
	default     = "7.10.0"
}

variable "kubernetes_dashboard_namespace" {
	description = "The kubernetes dashboard namespace (it will be created if needed)."
	type        = string
	default     = "kubernetes-dashboard"
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

variable "metallb_helm_version" {
	description = "The Helm version for the nginx ingress controller."
	type        = string
	default     = "0.14.9"
}

variable "metallb_namespace" {
	description = "The nginx ingress namespace (it will be created if needed)."
	type        = string
	default     = "metallb"
}

variable "metrics_server_helm_version" {
	description = "The Helm version for the metrics-server."
	type        = string
	default     = "3.12.2"
}

variable "metrics_server_namespace" {
	description = "The metrics namespace (it will be created if needed)."
	type        = string
	default     = "metrics"
}

variable "kube_prometheus_stack_helm_version" {
	description = "The Helm version for the kube-prometheus-stack."
	type        = string
	default     = "67.5.0"
}

variable "kube_prometheus_stack_namespace" {
	description = "The kube prometheus stack namespace (it will be created if needed)."
	type        = string
	default     = "monitoring"
}
