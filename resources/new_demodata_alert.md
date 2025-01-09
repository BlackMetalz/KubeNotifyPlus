### KubePodCrashLooping
```
 {
    "receiver": "notify-services-k8s",
    "status": "firing",
    "alerts": [
        {
            "status": "firing",
            "labels": {
                "alertname": "KubePodCrashLooping",
                "cluster": "kienlt-rke2",
                "cluster_id": "local",
                "cluster_name": "rke2-cluster",
                "container": "cpu-throttle-example",
                "endpoint": "http",
                "instance": "10.42.5.138:8080",
                "job": "kube-state-metrics",
                "namespace": "default",
                "pod": "cpu-throttle-example-7df487bbd-wfhhl",
                "prometheus": "cattle-monitoring-system/rancher-monitoring-prometheus",
                "prometheus_replica": "prometheus-rancher-monitoring-prometheus-0",
                "service": "rancher-monitoring-kube-state-metrics",
                "severity": "critical",
                "uid": "2a0ac152-5736-4a1f-9ad8-26585237b3e9"
            },
            "annotations": {
                "message": "Pod default/cpu-throttle-example-7df487bbd-wfhhl (cpu-throttle-example) is restarting 1.03 times / 5 minutes.",
                "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodcrashlooping"
            },
            "startsAt": "2025-01-05T07:44:37.42Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=rate%28kube_pod_container_status_restarts_total%7Bjob%3D%22kube-state-metrics%22%7D%5B15m%5D%29+%2A+60+%2A+5+%3E+0\u0026g0.tab=1",
            "fingerprint": "2253cda00cf19998"
        },
        {
            "status": "firing",
            "labels": {
                "alertname": "KubePodCrashLooping",
                "cluster": "kienlt-rke2",
                "cluster_id": "local",
                "cluster_name": "rke2-cluster",
                "container": "crash-loopbackoff-example",
                "endpoint": "http",
                "instance": "10.42.5.138:8080",
                "job": "kube-state-metrics",
                "namespace": "default",
                "pod": "crash-loopbackoff-example-c87d8fdbf-s4mnx",
                "prometheus": "cattle-monitoring-system/rancher-monitoring-prometheus",
                "prometheus_replica": "prometheus-rancher-monitoring-prometheus-0",
                "service": "rancher-monitoring-kube-state-metrics",
                "severity": "critical",
                "uid": "b76279b8-a034-4834-9de5-0d48f541505f"
            },
            "annotations": {
                "message": "Pod default/crash-loopbackoff-example-c87d8fdbf-s4mnx (crash-loopbackoff-example) is restarting 1.03 times / 5 minutes.",
                "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodcrashlooping"
            },
            "startsAt": "2025-01-05T07:44:37.42Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=rate%28kube_pod_container_status_restarts_total%7Bjob%3D%22kube-state-metrics%22%7D%5B15m%5D%29+%2A+60+%2A+5+%3E+0\u0026g0.tab=1",
            "fingerprint": "809e675d93c793ed"
        },
        {
            "status": "firing",
            "labels": {
                "alertname": "KubePodCrashLooping",
                "cluster": "kienlt-rke2",
                "cluster_id": "local",
                "cluster_name": "rke2-cluster",
                "container": "probe-failure-example",
                "endpoint": "http",
                "instance": "10.42.5.138:8080",
                "job": "kube-state-metrics",
                "namespace": "default",
                "pod": "probe-failure-example-744dbc6bf7-5lrh6",
                "prometheus": "cattle-monitoring-system/rancher-monitoring-prometheus",
                "prometheus_replica": "prometheus-rancher-monitoring-prometheus-0",
                "service": "rancher-monitoring-kube-state-metrics",
                "severity": "critical",
                "uid": "25db64d6-5e59-4357-a846-77a113a72415"
            },
            "annotations": {
                "message": "Pod default/probe-failure-example-744dbc6bf7-5lrh6 (probe-failure-example) is restarting 2.07 times / 5 minutes.",
                "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodcrashlooping"
            },
            "startsAt": "2025-01-05T07:44:37.42Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=rate%28kube_pod_container_status_restarts_total%7Bjob%3D%22kube-state-metrics%22%7D%5B15m%5D%29+%2A+60+%2A+5+%3E+0\u0026g0.tab=1",
            "fingerprint": "e366ca7d0ffb4a45"
        }
    ],
    "groupLabels": {
        "alertname": "KubePodCrashLooping",
        "cluster": "kienlt-rke2",
        "service": "rancher-monitoring-kube-state-metrics"
    },
    "commonLabels": {
        "alertname": "KubePodCrashLooping",
        "cluster": "kienlt-rke2",
        "cluster_id": "local",
        "cluster_name": "rke2-cluster",
        "endpoint": "http",
        "instance": "10.42.5.138:8080",
        "job": "kube-state-metrics",
        "namespace": "default",
        "prometheus": "cattle-monitoring-system/rancher-monitoring-prometheus",
        "prometheus_replica": "prometheus-rancher-monitoring-prometheus-0",
        "service": "rancher-monitoring-kube-state-metrics",
        "severity": "critical"
    },
    "commonAnnotations": {
        "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodcrashlooping"
    },
    "externalURL": "http://kienlt-salt-master-107-145:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"KubePod.*\"}:{alertname=\"KubePodCrashLooping\", cluster=\"kienlt-rke2\", service=\"rancher-monitoring-kube-state-metrics\"}",
    "truncatedAlerts": 0
}
 ```

 ### KubePodNotReady
 ```
{
    "receiver": "notify-services-k8s",
    "status": "firing",
    "alerts": [
        {
            "status": "firing",
            "labels": {
                "alertname": "KubePodNotReady",
                "cluster": "kienlt-rke2",
                "namespace": "default",
                "pod": "configmap-error-example-549c447bf7-trnv8",
                "severity": "critical"
            },
            "annotations": {
                "message": "Pod default/configmap-error-example-549c447bf7-trnv8 has been in a non-ready state for longer than 5 minutes.",
                "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodnotready"
            },
            "startsAt": "2025-01-05T07:43:37.42Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=sum+by+%28cluster%2C+namespace%2C+pod%29+%28kube_pod_status_phase%7Bjob%3D%22kube-state-metrics%22%2Cphase%3D~%22Pending%7CUnknown%22%7D%29+%3E+0\u0026g0.tab=1",
            "fingerprint": "fa8be574a8306c0a"
        },
        {
            "status": "firing",
            "labels": {
                "alertname": "KubePodNotReady",
                "cluster": "kienlt-rke2",
                "namespace": "default",
                "pod": "image-pullbackoff-example-76ff6c495c-whj69",
                "severity": "critical"
            },
            "annotations": {
                "message": "Pod default/image-pullbackoff-example-76ff6c495c-whj69 has been in a non-ready state for longer than 5 minutes.",
                "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodnotready"
            },
            "startsAt": "2025-01-05T07:43:37.42Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=sum+by+%28cluster%2C+namespace%2C+pod%29+%28kube_pod_status_phase%7Bjob%3D%22kube-state-metrics%22%2Cphase%3D~%22Pending%7CUnknown%22%7D%29+%3E+0\u0026g0.tab=1",
            "fingerprint": "a0132f23c6972abe"
        },
        {
            "status": "firing",
            "labels": {
                "alertname": "KubePodNotReady",
                "cluster": "kienlt-rke2",
                "namespace": "default",
                "pod": "secret-error-example-74c4f9f599-7bp64",
                "severity": "critical"
            },
            "annotations": {
                "message": "Pod default/secret-error-example-74c4f9f599-7bp64 has been in a non-ready state for longer than 5 minutes.",
                "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodnotready"
            },
            "startsAt": "2025-01-05T07:44:37.42Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=sum+by+%28cluster%2C+namespace%2C+pod%29+%28kube_pod_status_phase%7Bjob%3D%22kube-state-metrics%22%2Cphase%3D~%22Pending%7CUnknown%22%7D%29+%3E+0\u0026g0.tab=1",
            "fingerprint": "2be93c6da890339c"
        }
    ],
    "groupLabels": {
        "alertname": "KubePodNotReady",
        "cluster": "kienlt-rke2"
    },
    "commonLabels": {
        "alertname": "KubePodNotReady",
        "cluster": "kienlt-rke2",
        "namespace": "default",
        "severity": "critical"
    },
    "commonAnnotations": {
        "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodnotready"
    },
    "externalURL": "http://kienlt-salt-master-107-145:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"KubePod.*\"}:{alertname=\"KubePodNotReady\", cluster=\"kienlt-rke2\"}",
    "truncatedAlerts": 0
}

 ```