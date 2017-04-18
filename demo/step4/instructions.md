The web app exposes an endpoint that consumes close to 100% available cpu.
Usa that endpoint only once (haproxy will send it to one random pod)

http://spclvz8cv5.stackpoint.io/spoil

It's going to be 30 seconds or less until hpa notices that a new node is needed. You can use kubectl or the dashboard to show CPU consumption:

```
$ kubectl top pod
NAME                      CPU(cores)   MEMORY(bytes)
kwtest-3110424013-sq5n7   0m           1Mi
kwtest-3110424013-bvf0v   199m         1Mi
```

The pod is limited to 200m, so we are consuming most of the CPU.
After 1 minute, you will probably have one more pod.

You can use FIREFOX to check that a new pod has been added to the cluster.
