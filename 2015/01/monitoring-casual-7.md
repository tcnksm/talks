# Monitoring CoreOS with datadog

## 前提

- DEV -\> DEVOPS
    - About me
- Monitoring Docker container on CoreOS cluster
    - 最近試していたことについて話す


## Difference of monitoring

- Physical (Machines) -\> VM -\> Service
- Physical (Machines) -\> VM -\> Container -\> Service

## Monitoring by your self

自分でモニタリングサービスを持っている場合

- Write you script
	- [http://www.labouisse.com/how-to/2014/11/17/simple-monitoring-for-docker-part-1/](http://www.labouisse.com/how-to/2014/11/17/simple-monitoring-for-docker-part-1/)
	- `/sys/fs/cgroup/memory/docker/`
- cAdvisor
	- [https://github.com/google/cadvisor](https://github.com/google/cadvisor)
	- :8080でアクセスできる
	- DEMO on my own local PC
- Heapster
	- [https://github.com/GoogleCloudPlatform/heapster](https://github.com/GoogleCloudPlatform/heapster)
	- CoreOS clusterの監視に利用（kubernetesの内部で使われているが，kubernetes以外でも使える）
	- 現在はInfluxDBのバックエンドをサポートしている
		- なのでGraphanaで可視化できる


```bash
docker run \
    --volume=/:/rootfs:ro \
	--volume=/var/run:/var/run:rw \
	--volume=/sys:/sys:ro \
	--volume=/var/lib/docker/:/var/lib/docker:ro \
	--publish=8080:8080 \
	--detach=true \
	--name=cadvisor \
	google/cadvisor:latest
```

## Monitoring as a Service

- Mackerel
	-　Not support Container specific feature ..? 	
		- If supported, I want to use..	　
	- [https://github.com/stanaka/mackerel-docker](https://github.com/stanaka/mackerel-docker) ?

- Newrelic 
	- Not support Contaienr specific feature ...?
	- Provide dockerized agent but only for host metrics
		- Not Containers, Need to install agent to its container
	- For CoreOS host
		- [https://github.com/johanneswuerbach/newrelic-sysmond-service](https://github.com/johanneswuerbach/newrelic-sysmond-service)
- Datadog
	- Support Container specific feature
		- Tagging by docker iamage, container name, Collecting docker event
	- [https://github.com/DataDog/docker-dd-agent](https://github.com/DataDog/docker-dd-agent)
		- Dockerized agent. Collect all containers metrics on a host (same as cAdvisor)


## DataDog (Docker-specific)

- All containers metrics on a host
	- CPU, memory, network I/O and disk I/O
- Tag
	- By container name, its images name 
	- -> easy to summarize and draw graph
- Lifecycle Event
	- create, start, stop, destroy Event 
	- -> We can see event on graph, so we can understand abnormal graph meanings

-  Datadog agent & DogStatsD
	- Datadog agent -\> like number of containers, load, memory, disk usage, and latency
	- DogStatsD -\> custom metrics you have instrumented in containerized application

## DEMO 

## Requirement

- Just run agent by docker container
	- Container **only** approach is Docker-way 
		- DataDog [https://www.datadoghq.com/wp-content/uploads/2014/06/DockerizeImage2.png](https://www.datadoghq.com/wp-content/uploads/2014/06/DockerizeImage2.png)
		- cAdivisor

## Health Check of Service in Docker 

- Enter container by `docker exec`
- Check file tests or pid tests … etc
	- 従来の方法が使えるはず… 

## etcdやconsulでのTokenの扱い（番外編）

分散Key-valueストアにTokenとかを保存するのをうまくやりたい

- [https://github.com/xordataexchange/crypt/tree/master/bin/crypt](https://github.com/xordataexchange/crypt/tree/master/bin/crypt)を使う

## References

- [https://speakerdeck.com/stanaka/monitoring-kubernetes](https://speakerdeck.com/stanaka/monitoring-kubernetes)
- [https://speakerdeck.com/stanaka/monitoring-docker-with-mackerel](https://speakerdeck.com/stanaka/monitoring-docker-with-mackerel)
- [http://www.slideshare.net/Docker/orchestrating-docker-with-terraform-and-consul-by-mitchell-hashimoto](http://www.slideshare.net/Docker/orchestrating-docker-with-terraform-and-consul-by-mitchell-hashimoto)
- [https://www.datadoghq.com/2014/08/monitor-coreos-scale-datadog/](https://www.datadoghq.com/2014/08/monitor-coreos-scale-datadog/)
- [http://docs.docker.com/articles/runmetrics/](http://docs.docker.com/articles/runmetrics/)
- [https://www.datadoghq.com/2014/06/monitor-docker-datadog/](https://www.datadoghq.com/2014/06/monitor-docker-datadog/)
- [https://www.datadoghq.com/2014/06/docker-ize-datadog/](https://www.datadoghq.com/2014/06/docker-ize-datadog/)
- [https://www.datadoghq.com/2014/08/monitor-coreos-scale-datadog/](https://www.datadoghq.com/2014/08/monitor-coreos-scale-datadog/)
- [http://qiita.com/jhotta/items/6b7ba9cd5602d1f9208d](http://qiita.com/jhotta/items/6b7ba9cd5602d1f9208d)

