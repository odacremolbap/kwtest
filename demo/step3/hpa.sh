#!/bin/bash

kubectl autoscale deployment kwtest --min=2 --max=5 --cpu-percent=40

# Now you can take a look at the dashboard deployment, and see how it has been updated to 2 pods
# There should also appear an HPA

# you can demo Round Robin with FIREFOX, DO NOT USE chrome, since it caches the page


# You can show the horizontal pod autoscaler events and configuration here
kubectl describe hpa kwtest

# Also, there should be 2 kwtest pods
