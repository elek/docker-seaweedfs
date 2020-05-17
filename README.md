# [Seaweedfs](https://github.com/chrislusf/seaweedfs) Docker image and Kubernetes resource definitions

Docker image: `elek/seaweedfs`

Using in Kubernetes:

Install [flekszible](https://github.com/elek/flekszible)

```
cd /tmp/test
flekszible init
flekszible source add github.com/elek/docker-seaweedfs
flekszible generate
kubectl apply -f
```
