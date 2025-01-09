### Config for alertmanager receivers support basic auth
```
receivers:
- name: notify-services-k8s
  webhook_configs:
  - send_resolved: true
    url: http://example.com/alert
    http_config:
      basic_auth:
        username: "your_username"
        password: "your_password"
```

