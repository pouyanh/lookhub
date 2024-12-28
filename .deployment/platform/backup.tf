resource "kubernetes_manifest" "etcd_backup_cronjob" {
	manifest = yamldecode(<<-EOT
    apiVersion: batch/v1
    kind: CronJob
    metadata:
      name: etcd-backup
      namespace: kube-system
    spec:
      schedule: "0 * * * *"
      jobTemplate:
        spec:
          template:
            spec:
              tolerations:
                - key: "node-role.kubernetes.io/control-plane"
                  operator: "Exists"
                  effect: "NoSchedule"
              containers:
              - name: etcd-backup
                image: bitnami/etcd:latest
                command:
                - /bin/sh
                - -c
                - |
                  export ETCDCTL_API=3
                  etcdctl snapshot save /etcd-backups/etcd-snapshot-$(date +%Y%m%d%H%M%S).db \
                    --endpoints=https://127.0.0.1:2379 \
                    --cacert=/certs/ca.crt \
                    --cert=/certs/server.crt \
                    --key=/certs/server.key
                volumeMounts:
                - name: etcd-certs
                  mountPath: /certs
                  readOnly: true
                - name: etcd-backups
                  mountPath: /etcd-backups
              restartPolicy: OnFailure
              volumes:
              - name: etcd-certs
                hostPath:
                  path: /etc/kubernetes/pki/etcd
              - name: etcd-backups
                hostPath:
                  path: /backups/etcd
    EOT
	)

	depends_on = [kind_cluster.default]
}
