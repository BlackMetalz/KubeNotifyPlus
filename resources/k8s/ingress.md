### Tools for test
- https://github.com/tsenart/vegeta
- Ref: https://www.aviator.co/blog/how-to-monitor-and-alert-on-nginx-ingress-in-kubernetes/#

### Example query:
- List queries are error i guess.
```
nginx_ingress_controller_requests{status!~"1..|2..|3..|401|404"}
```

### Repo for testing:
- https://github.com/BlackMetalz/Golang-Webapp-Testing
- How to use it: just build and deploy to k8s then using endpoint for testing, tools for test located at: `resources/test/test_requests.sh`