# apiserver

Dummy repo for serving a grpc api. After creating an EKS cluster using
https://github.com/xh3b4sd/kia this api server app can be deployed used.

```
helm -n infra install apiserver ./helm/apiserver
```

```
grpcurl apiserver.aws.example.com:443 post.API/Search
```
