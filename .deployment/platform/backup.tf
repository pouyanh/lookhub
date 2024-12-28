resource "kubernetes_manifest" "etcd_backup_cronjob" {
	manifest = yamldecode(file("${path.module}/backup_etcd.yaml"))

	depends_on = [kind_cluster.default]
}
