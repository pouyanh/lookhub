# LookHub
LookHub is domain name lookup service over http api.
It collects a domain name resource records (A, NS, MX, ...) which includes ip address, name servers and other details.

# Run local development environment
There is a simple development environment built using docker-compose which runs LookHub from source code.
And whenever source code changes locally it reloads the service (hot-reload) powered by [PolyWatch][polywatch].

```shell
docker-compose up -d --remove-orphans
```

By using [autodns][autodns] you can reach services locally using their virtual fqdn under _cloud.snp_ domain.
* LookHub API Swagger UI: http://lookhub.cloud.snp/docs
* LookHub API: http://lookhub.cloud.snp
* Postgres UI: http://pgadmin.cloud.snp:8080

# Deployment
There are two deployment mechanisms which are described below.

## Local Kubernetes Cluster
To simulate production environment install [Terraform][terraform] and
create the local k8s cluster provided in [.deployment/platform](.deployment/platform):
```shell
cd .deployment/platform
terraform init
```

In order to let scheduler pull docker images from private docker registry
you should set credentials in _terraform.tfvars.json_ file. It's ignored by git vcs:
```shell
copy terraform.tfvars.sample.json terraform.tfvars.json
```
And fill in private docker registry credentials
```json
{
	"private_docker_registry_server": "registry.snapp.tech",
	"private_docker_registry_username": "your-gitlab-username",
	"private_docker_registry_password": "your-gitlab-access-token-with-read-registry-scope"
}
```
Finally, bring the local cluster up:
```shell
terraform apply
```
Ensure successful setup using [kubectl][kubectl]:
```shell
KUBECONFIG=$(terraform output -raw kubeconfig_path) kubectl get pods --all-namespaces -o wide
```
Install [helm][helm] and deploy LookHub [helm chart](.deployment/lookhub/Chart.yaml) on local cluster using [local values file](.deployment/lookhub/values-local.yaml)
```shell
KUBECONFIG=$(terraform output -raw kubeconfig_path) helm upgrade --install \
  --namespace="lookhub" \
  --kube-token="$(terraform output -raw lookhub_ci_token)" \
  lookhub ../lookhub -f ../lookhub/values-local.yaml
```

## Remote Kubernetes Cluster using Gitlab CI/CD
By default, ci/cd builds the docker image and pushes it to GitLab's container registry.
Whenever a user run the pipeline using **New pipeline** in the GitLab UI,
from the project’s **Build > Pipelines** section a **deploy** job runs on successful release step
which should be triggered manually.

# Source code

[autodns]: https://github.com/pouyanh/autodns
[polywatch]: https://pouyanh.github.io/polywatch
[terraform]: https://www.terraform.io/
[kubectl]: https://kubernetes.io/docs/reference/kubectl
[helm]: https://helm.sh/
