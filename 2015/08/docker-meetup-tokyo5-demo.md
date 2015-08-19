# Docker meetup Tokyo #5 (DEMO)

Let's create nginx application,

First, run nginx app with 2 replicas

```bash
$ kubectl run my-nginx --image=nginx --replicas=2 --port=80
```

Then, expost it as a service,


```bash
$ kubectl expose rc my-nginx --port=80 --type=LoadBalancer
```

Get external IP of nginx service,

```bash
$ kubectl get svc my-nginx -o json | grep IP
```

And request, you can see the nginx top page,


```bash
boot2docker ssh curl ...
```


