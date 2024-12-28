resource "rancher2_etcd_backup" "etcd_backup" {
	cluster_id = kind_cluster.default.id
	name       = "etcd_auto_backup"

	backup_config {
		enabled = true
	}
}
