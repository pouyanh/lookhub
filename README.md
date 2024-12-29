# LookHub
Domain name lookup center

# Usage
Visit: [http://lookhub.cloud.snp/docs][lookhub-swagger]
Visit: [http://lookhub.cloud.snp][lookhub]
Visit: [http://pgadmin.cloud.snp:8080][pgadmin]

# Run local development environment

```shell
docker-compose up -d --remove-orphans
```

# Deployment

## Local Kubernetes Cluster

Install Terraform

```shell
cd .deployment/platform
terraform init
```

Create _terraform.tfvars.json_ file
```shell
copy terraform.tfvars.sample.json terraform.tfvars.json
```

Fill in private docker registry credentials
```json
{
	"private_docker_registry_server": "registry.snapp.tech",
	"private_docker_registry_username": "your-gitlab-username",
	"private_docker_registry_password": "your-gitlab-access-token-with-read-registry-scope"
}
```

Setup local cluster
```shell
terraform apply
```

Ensure successful setup
```shell
KUBECONFIG=$(terraform output -raw kubeconfig_path) kubectl get pods --all-namespaces -o wide
```

Install helm and deploy lookhub helm chart on local cluster using local values file
```shell
KUBECONFIG=$(terraform output -raw kubeconfig_path) helm upgrade --install \
  --namespace="lookhub" \
  --kube-token="$(terraform output -raw lookhub_ci_token)" \
  lookhub ../lookhub -f ../lookhub/values-local.yaml
```

# Development

## Packages

[pgadmin]: http://pgadmin.cloud.snp:8080
[lookhub]: http://lookhub.cloud.snp
[lookhub-swagger]: http://lookhub.cloud.snp/docs
