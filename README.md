# manytimesjob

Kubernetes Controller for applying job of same name more than once.

## Install

```
# controller
kubectl apply -f manifest.yaml -n default

# rbac
kubectl kustomize config/rbac | kubectl apply -f -
```

## Demo

1. apply ManyTimesJob Resource * 4

```
kubectl apply -f config/samples/batch_v1alpha1_manytimesjob.yaml
kubectl apply -f config/samples/batch_v1alpha1_manytimesjob.yaml
kubectl apply -f config/samples/batch_v1alpha1_manytimesjob.yaml
kubectl apply -f config/samples/batch_v1alpha1_manytimesjob.yaml
```

2. get Job Resource

```
$ kubectl get job
NAME              COMPLETIONS   DURATION   AGE
test-1582530561   1/1           15s        8m20s
test-1582531041   1/1           15s        30s
test-1582531047   1/1           15s        24s
test-1582531048   1/1           15s        23s
```

3. apply ManyTimesJob Resource

```
kubectl apply -f config/samples/batch_v1alpha1_manytimesjob.yaml
```

4. get Job Resource
    * deleted oldest Job resource!

```
$ kubectl get job
NAME              COMPLETIONS   DURATION   AGE
test-1582531041   1/1           15s        32s
test-1582531047   1/1           15s        26s
test-1582531048   1/1           15s        25s
test-1582531072   0/1            1s         1s
```
