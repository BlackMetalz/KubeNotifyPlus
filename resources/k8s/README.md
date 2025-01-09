### Get API Server info
- Just for quick solution
```
kubectl config view --minify -o jsonpath='{.clusters[0].cluster.server}'
```
We have to map those apiserver address to a domain with roundrobin or something like that

### Permission requireds
- List,Get pod
- Read logs of pod

### Token
- Get secret
```
kubectl get secret pod-inspector-token -n kube-notify-plus -o yaml
```
- Extract token xD
```
kubectl get secret pod-inspector-token -n kube-notify-plus -o jsonpath='{.data.token}' | base64 --decode
```
- Verify token expires if it does
```
TOKEN=$(kubectl get secret pod-inspector-token -n kube-notify-plus -o jsonpath='{.data.token}' | base64 --decode)
echo $TOKEN | cut -d "." -f2 | base64 --decode | jq
```

If the token has no exp field, it does not expire!

- Rotate the token: Simply delete and create again xD
```
