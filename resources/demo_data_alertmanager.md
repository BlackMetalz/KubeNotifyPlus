### Payload for KubePodCrashLooping
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
                "container": "cpu-throttle-container",
                "endpoint": "http",
                "instance": "10.42.5.138:8080",
                "job": "kube-state-metrics",
                "namespace": "cpu-throttle",
                "pod": "cpu-throttle-example-65cb7677f9-ncs5t",
                "prometheus": "cattle-monitoring-system/rancher-monitoring-prometheus",
                "prometheus_replica": "prometheus-rancher-monitoring-prometheus-0",
                "service": "rancher-monitoring-kube-state-metrics",
                "severity": "critical",
                "uid": "0c59ceef-09d1-46a9-8e00-75513a666026"
            },
            "annotations": {
                "message": "Pod cpu-throttle/cpu-throttle-example-65cb7677f9-ncs5t (cpu-throttle-container) is restarting 0.69 times / 5 minutes.",
                "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodcrashlooping"
            },
            "startsAt": "2024-12-31T05:47:37.42Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=rate%28kube_pod_container_status_restarts_total%7Bjob%3D%22kube-state-metrics%22%7D%5B15m%5D%29+%2A+60+%2A+5+%3E+0\u0026g0.tab=1",
            "fingerprint": "1df482aa1c360857"
        },
        {
            "status": "firing",
            "labels": {
                "alertname": "KubePodCrashLooping",
                "cluster": "kienlt-rke2",
                "cluster_id": "local",
                "cluster_name": "rke2-cluster",
                "container": "crash-loop-container",
                "endpoint": "http",
                "instance": "10.42.5.138:8080",
                "job": "kube-state-metrics",
                "namespace": "crash-loopbackoff",
                "pod": "crash-loopbackoff-example-6768587b8-lgw5g",
                "prometheus": "cattle-monitoring-system/rancher-monitoring-prometheus",
                "prometheus_replica": "prometheus-rancher-monitoring-prometheus-0",
                "service": "rancher-monitoring-kube-state-metrics",
                "severity": "critical",
                "uid": "7211a2f0-b11d-4934-92c8-2dba4e8fd35e"
            },
            "annotations": {
                "message": "Pod crash-loopbackoff/crash-loopbackoff-example-6768587b8-lgw5g (crash-loop-container) is restarting 1.03 times / 5 minutes.",
                "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodcrashlooping"
            },
            "startsAt": "2024-12-31T04:53:37.42Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=rate%28kube_pod_container_status_restarts_total%7Bjob%3D%22kube-state-metrics%22%7D%5B15m%5D%29+%2A+60+%2A+5+%3E+0\u0026g0.tab=1",
            "fingerprint": "07b9bb37aaa20360"
        },
        {
            "status": "firing",
            "labels": {
                "alertname": "KubePodCrashLooping",
                "cluster": "kienlt-rke2",
                "cluster_id": "local",
                "cluster_name": "rke2-cluster",
                "container": "oom-container",
                "endpoint": "http",
                "instance": "10.42.5.138:8080",
                "job": "kube-state-metrics",
                "namespace": "oom",
                "pod": "oom-example-56d77685f5-dqv7z",
                "prometheus": "cattle-monitoring-system/rancher-monitoring-prometheus",
                "prometheus_replica": "prometheus-rancher-monitoring-prometheus-0",
                "service": "rancher-monitoring-kube-state-metrics",
                "severity": "critical",
                "uid": "a5d00a5c-5032-4288-9ed5-6b519106814a"
            },
            "annotations": {
                "message": "Pod oom/oom-example-56d77685f5-dqv7z (oom-container) is restarting 1.03 times / 5 minutes.",
                "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodcrashlooping"
            },
            "startsAt": "2024-12-31T04:58:37.42Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=rate%28kube_pod_container_status_restarts_total%7Bjob%3D%22kube-state-metrics%22%7D%5B15m%5D%29+%2A+60+%2A+5+%3E+0\u0026g0.tab=1",
            "fingerprint": "328c4b7fcec75544"
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

### Payload for KubePodNotReady
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
                "namespace": "non-exists-image",
                "pod": "image-pullbackoff-example-d8689bfb8-nlb8b",
                "severity": "critical"
            },
            "annotations": {
                "message": "Pod non-exists-image/image-pullbackoff-example-d8689bfb8-nlb8b has been in a non-ready state for longer than 5 minutes.",
                "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodnotready"
            },
            "startsAt": "2025-01-02T04:31:37.42Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=sum+by+%28cluster%2C+namespace%2C+pod%29+%28kube_pod_status_phase%7Bjob%3D%22kube-state-metrics%22%2Cphase%3D~%22Pending%7CUnknown%22%7D%29+%3E+0\u0026g0.tab=1",
            "fingerprint": "c9b70396653b4ad7"
        }
    ],
    "groupLabels": {
        "alertname": "KubePodNotReady",
        "cluster": "kienlt-rke2"
    },
    "commonLabels": {
        "alertname": "KubePodNotReady",
        "cluster": "kienlt-rke2",
        "namespace": "non-exists-image",
        "pod": "image-pullbackoff-example-d8689bfb8-nlb8b",
        "severity": "critical"
    },
    "commonAnnotations": {
        "message": "Pod non-exists-image/image-pullbackoff-example-d8689bfb8-nlb8b has been in a non-ready state for longer than 5 minutes.",
        "runbook_url": "https://github.com/kubernetes-monitoring/kubernetes-mixin/tree/master/runbook.md#alert-name-kubepodnotready"
    },
    "externalURL": "http://kienlt-salt-master-107-145:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"KubePod.*\"}:{alertname=\"KubePodNotReady\", cluster=\"kienlt-rke2\"}",
    "truncatedAlerts": 0
}
```


### Payload for notify-ingress-k8s (499)
```
{
    "receiver": "notify-ingress-k8s",
    "status": "firing",
    "alerts": [
        {
            "status": "firing",
            "labels": {
                "alertname": "ingress-k8s-http-error-499",
                "cluster": "kienlt-rke2",
                "ingress": "golang-webapp-testing-production",
                "severity": "critical",
                "status": "499"
            },
            "annotations": {
                "description": "Service golang-webapp-testing-production/kienlt-rke2 too many HTTP requests with status 499 (\u003e50%): 100%"
            },
            "startsAt": "2025-01-08T05:32:21.539Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=sum+by+%28ingress%2C+cluster%2C+status%29+%28rate%28nginx_ingress_controller_requests%7Bingress%21%3D%22%22%2Cjob%3D%22ingress-nginx%22%2Cnamespace%3D%22kube-system%22%2Cstatus%3D%22499%22%7D%5B2m%5D%29%29+%2F+sum+by+%28ingress%2C+cluster%2C+status%29+%28rate%28nginx_ingress_controller_requests%7Bingress%21%3D%22%22%2Cjob%3D%22ingress-nginx%22%2Cnamespace%3D%22kube-system%22%7D%5B2m%5D%29%29+%2A+100+%3E+50\u0026g0.tab=1",
            "fingerprint": "8cd759d5eac14f05"
        }
    ],
    "groupLabels": {
        "alertname": "ingress-k8s-http-error-499",
        "cluster": "kienlt-rke2"
    },
    "commonLabels": {
        "alertname": "ingress-k8s-http-error-499",
        "cluster": "kienlt-rke2",
        "ingress": "golang-webapp-testing-production",
        "severity": "critical",
        "status": "499"
    },
    "commonAnnotations": {
        "description": "Service golang-webapp-testing-production/kienlt-rke2 too many HTTP requests with status 499 (\u003e50%): 100%"
    },
    "externalURL": "http://kienlt-salt-master-107-145:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"ingress-k8s.*\"}:{alertname=\"ingress-k8s-http-error-499\", cluster=\"kienlt-rke2\"}",
    "truncatedAlerts": 0
}
```

### Payload for notify-ingress-k8s (500)
```
{
    "receiver": "notify-ingress-k8s",
    "status": "firing",
    "alerts": [
        {
            "status": "firing",
            "labels": {
                "alertname": "ingress-k8s-http-error-500",
                "cluster": "kienlt-rke2",
                "ingress": "golang-webapp-testing-production",
                "severity": "critical",
                "status": "500"
            },
            "annotations": {
                "description": "Service golang-webapp-testing-production/kienlt-rke2 too many HTTP requests with status 500 (\u003e5%): 100%"
            },
            "startsAt": "2025-01-08T05:32:21.539Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=sum+by+%28ingress%2C+cluster%2C+status%29+%28rate%28nginx_ingress_controller_requests%7Bingress%21%3D%22%22%2Cjob%3D%22ingress-nginx%22%2Cnamespace%3D%22kube-system%22%2Cstatus%3D%22500%22%7D%5B2m%5D%29%29+%2F+sum+by+%28ingress%2C+cluster%2C+status%29+%28rate%28nginx_ingress_controller_requests%7Bingress%21%3D%22%22%2Cjob%3D%22ingress-nginx%22%2Cnamespace%3D%22kube-system%22%7D%5B2m%5D%29%29+%2A+100+%3E+5\u0026g0.tab=1",
            "fingerprint": "0f26057cfa37a329"
        }
    ],
    "groupLabels": {
        "alertname": "ingress-k8s-http-error-500",
        "cluster": "kienlt-rke2"
    },
    "commonLabels": {
        "alertname": "ingress-k8s-http-error-500",
        "cluster": "kienlt-rke2",
        "ingress": "golang-webapp-testing-production",
        "severity": "critical",
        "status": "500"
    },
    "commonAnnotations": {
        "description": "Service golang-webapp-testing-production/kienlt-rke2 too many HTTP requests with status 500 (\u003e5%): 100%"
    },
    "externalURL": "http://kienlt-salt-master-107-145:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"ingress-k8s.*\"}:{alertname=\"ingress-k8s-http-error-500\", cluster=\"kienlt-rke2\"}",
    "truncatedAlerts": 0
}
```

### Payload for notify-ingress-k8s (502-503)
```
{
    "receiver": "notify-ingress-k8s",
    "status": "firing",
    "alerts": [
        {
            "status": "firing",
            "labels": {
                "alertname": "ingress-k8s-http-error-502-503",
                "cluster": "kienlt-rke2",
                "ingress": "golang-webapp-testing-production",
                "severity": "critical",
                "status": "502"
            },
            "annotations": {
                "description": "Service golang-webapp-testing-production/kienlt-rke2 too many HTTP requests with status 502|503 (\u003e5%): 100%"
            },
            "startsAt": "2025-01-08T05:32:21.539Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=sum+by+%28ingress%2C+cluster%2C+status%29+%28rate%28nginx_ingress_controller_requests%7Bingress%21%3D%22%22%2Cjob%3D%22ingress-nginx%22%2Cnamespace%3D%22kube-system%22%2Cstatus%3D~%2250%282%7C3%29%22%7D%5B2m%5D%29%29+%2F+sum+by+%28ingress%2C+cluster%2C+status%29+%28rate%28nginx_ingress_controller_requests%7Bingress%21%3D%22%22%2Cjob%3D%22ingress-nginx%22%2Cnamespace%3D%22kube-system%22%7D%5B2m%5D%29%29+%2A+100+%3E+5\u0026g0.tab=1",
            "fingerprint": "908d5920f0a54ec0"
        },
        {
            "status": "firing",
            "labels": {
                "alertname": "ingress-k8s-http-error-502-503",
                "cluster": "kienlt-rke2",
                "ingress": "golang-webapp-testing-production",
                "severity": "critical",
                "status": "503"
            },
            "annotations": {
                "description": "Service golang-webapp-testing-production/kienlt-rke2 too many HTTP requests with status 502|503 (\u003e5%): 100%"
            },
            "startsAt": "2025-01-08T05:32:21.539Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=sum+by+%28ingress%2C+cluster%2C+status%29+%28rate%28nginx_ingress_controller_requests%7Bingress%21%3D%22%22%2Cjob%3D%22ingress-nginx%22%2Cnamespace%3D%22kube-system%22%2Cstatus%3D~%2250%282%7C3%29%22%7D%5B2m%5D%29%29+%2F+sum+by+%28ingress%2C+cluster%2C+status%29+%28rate%28nginx_ingress_controller_requests%7Bingress%21%3D%22%22%2Cjob%3D%22ingress-nginx%22%2Cnamespace%3D%22kube-system%22%7D%5B2m%5D%29%29+%2A+100+%3E+5\u0026g0.tab=1",
            "fingerprint": "908b5920f0a4cbe9"
        }
    ],
    "groupLabels": {
        "alertname": "ingress-k8s-http-error-502-503",
        "cluster": "kienlt-rke2"
    },
    "commonLabels": {
        "alertname": "ingress-k8s-http-error-502-503",
        "cluster": "kienlt-rke2",
        "ingress": "golang-webapp-testing-production",
        "severity": "critical"
    },
    "commonAnnotations": {
        "description": "Service golang-webapp-testing-production/kienlt-rke2 too many HTTP requests with status 502|503 (\u003e5%): 100%"
    },
    "externalURL": "http://kienlt-salt-master-107-145:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"ingress-k8s.*\"}:{alertname=\"ingress-k8s-http-error-502-503\", cluster=\"kienlt-rke2\"}",
    "truncatedAlerts": 0
}
```

### Payload for notify-ingress-k8s (504)
```
{
    "receiver": "notify-ingress-k8s",
    "status": "firing",
    "alerts": [
        {
            "status": "firing",
            "labels": {
                "alertname": "ingress-k8s-http-error-504",
                "cluster": "kienlt-rke2",
                "ingress": "golang-webapp-testing-production",
                "severity": "critical",
                "status": "504"
            },
            "annotations": {
                "description": "Service golang-webapp-testing-production/kienlt-rke2 too many HTTP requests with status 504 (\u003e5%): 100%"
            },
            "startsAt": "2025-01-08T05:32:21.539Z",
            "endsAt": "0001-01-01T00:00:00Z",
            "generatorURL": "http://50cb6c5d4c40:9090/graph?g0.expr=sum+by+%28ingress%2C+cluster%2C+status%29+%28rate%28nginx_ingress_controller_requests%7Bingress%21%3D%22%22%2Cjob%3D%22ingress-nginx%22%2Cnamespace%3D%22kube-system%22%2Cstatus%3D%22504%22%7D%5B2m%5D%29%29+%2F+sum+by+%28ingress%2C+cluster%2C+status%29+%28rate%28nginx_ingress_controller_requests%7Bingress%21%3D%22%22%2Cjob%3D%22ingress-nginx%22%2Cnamespace%3D%22kube-system%22%7D%5B2m%5D%29%29+%2A+100+%3E+5\u0026g0.tab=1",
            "fingerprint": "2d4e9e34107bd119"
        }
    ],
    "groupLabels": {
        "alertname": "ingress-k8s-http-error-504",
        "cluster": "kienlt-rke2"
    },
    "commonLabels": {
        "alertname": "ingress-k8s-http-error-504",
        "cluster": "kienlt-rke2",
        "ingress": "golang-webapp-testing-production",
        "severity": "critical",
        "status": "504"
    },
    "commonAnnotations": {
        "description": "Service golang-webapp-testing-production/kienlt-rke2 too many HTTP requests with status 504 (\u003e5%): 100%"
    },
    "externalURL": "http://kienlt-salt-master-107-145:9093",
    "version": "4",
    "groupKey": "{}/{alertname=~\"ingress-k8s.*\"}:{alertname=\"ingress-k8s-http-error-504\", cluster=\"kienlt-rke2\"}",
    "truncatedAlerts": 0
}
```