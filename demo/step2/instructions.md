Create the ingress controller at this folder that points to the kwtest service

```
kubectl create -f kwtest-ingress.yaml
```

This ingress creates a default backend at port 80.
You should be able to use the master DNS name to display the page.
USE FIREFOX.
