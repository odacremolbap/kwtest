- Import the compose file
- Review manifest
- Add resource limits
look for `resources: {}` and replace with
```
resources:
  limits:
    cpu: "200m"
```  


Result:

```
apiVersion: v1
items:
- apiVersion: extensions/v1beta1
  kind: Deployment
  metadata:
    creationTimestamp: null
    name: kwtest
  spec:
    replicas: 1
    strategy: {}
    template:
      metadata:
        creationTimestamp: null
        labels:
          service: kwtest
      spec:
        containers:
        - image: pmercado/kwtest:0.1.0
          name: kwtest
          ports:
          - containerPort: 80
            protocol: TCP
          resources:
            limits:
              cpu: "200m"
        restartPolicy: Always
  status: {}
- apiVersion: v1
  kind: Service
  metadata:
    creationTimestamp: null
    labels:
      service: kwtest
    name: kwtest
  spec:
    ports:
    - name: '80'
      port: 80
      protocol: TCP
      targetPort: 80
    selector:
      service: kwtest
  status:
    loadBalancer: {}
kind: List
metadata: {}
```
