### Kubernetes

There are several options to run data handler on kubernetes. You can use helm chart or kustomize.

## yaml deployment file

***Pre-requisites***

First you need to have a postgresql database on kubernetes.

You can use the following deployment file to deploy data handler on kubernetes.

Download https://raw.githubusercontent.com/tislib/data-handler/master/deploy/kubernetes/data-handler.yaml

you need to change the following values:

```
"username":  "root",
"password":  "root",
"host":  "db",
"port":  5432,
"dbName":  "dh_system",
"defaultSchema":  "public"
```
Keep in mind, you need to modify this in two places, for dh_system and dh_data

Now let's apply our deployment file

```bash
kubectl apply -f data-handler.yaml
```

It will deploy postgresql and data handler on kubernetes.

```
## See pods
kubectl get pods

## See services
kubectl get svc
```
