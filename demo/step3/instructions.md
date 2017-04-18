We are going to scale the deployment pods using Horizontal Pod Autoscaler.
Let's set it to 2 min / 5 max, scaling when average pod CPU reaches 40%
(it can also be done with custom metrics, but that's too much for the demo)


```
kubectl autoscale deployment kwtest --min=2 --max=5 --cpu-percent=40
```

Wait 1 minute, and you should have 2 pods.
open FIREFOX (chrome caches the page and doesn't round robin) and refresh the page a number of times.
You should get to a different pod at each requeset.


You can check hpa status

```
kubectl describe hpa kwtest
```
