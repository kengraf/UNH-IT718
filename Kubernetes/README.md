# UNH-IT718-k8s
Lab to support K8s discussion in class

[Kubernetes Concepts Comic](https://cloud.google.com/kubernetes-engine/kubernetes-comic/)

This is an example of using Kubernetes' horizontal pod autoscaling (HPA).  A container that provides moderate CPU utilization is needed.
One is provided in [Github UNH-IT718-docker](https://github.com/kengraf/UNH-IT718-docker).  The result being a repo of YOUR_NAME/hpa-example:v1 hosted on Docker Hub.  If not, this lesson without edits pulls from the default repo of billiardyoda/hpa-example.

## Lesson 
Now that we have a deploy focused Docker image, it is time to learn about scaling it with Kubernetes.  
If you are interested in a more complete deployment with back & front ends check out Google's Kubenetes examples: [GCP demo](https://cloud.google.com/kubernetes-engine/docs/tutorials/guestbook)  

Insure your Cloudshell settings are current.  This is needed if you have to reconnect after an inactivity timeout.
```
gcloud config set project YOUR_PROJECT_ID
gcloud config set compute/zone us-central1
```
You may need to enable service for your account
```
gcloud services enable compute.googleapis.com; gcloud services enable container.googleapis.com
```

Create a new cluster for the deployment
```
gcloud container clusters create hpa-example --num-nodes=2
gcloud container clusters list
gcloud container clusters describe hpa-example
```

Pull example yaml
```
git clone https://github.com/kengraf/UNH-IT718.git
cd  UNH-IT718/Kubernetes
```

Apply a "yaml" to deploy what is needed
```
kubectl apply -f hpa-example.yaml
```

Commands to check status
```
kubectl get -f hpa-example.yaml

# Review separate parts
kubectl get deployments
kubectl get services
kubectl rollout status deployment/hpa-example
kubectl get rs
kubectl get pods --show-labels
```
You can monitor the Kubernetes Engine Dashboard in the GCP console  
```
wget -q -O- http://<external-ip>/hello
wget -q -O- http://<external-ip>/dowork
while sleep 1; do wget -q -O- http://<external-ip>/dowork; done"

kubectl run -i --tty load-generator --rm --image=busybox --restart=Never -- /bin/sh -c "while sleep 1; do wget -q -O- http://hpa-example/dowork; done"

```

Roll an update to fix performance
```
kubectl get deployments
kubectl get pods -o wide
kubectl scale deployment hpa-example --replicas=1
kubectl set image deployments/hpa-example hpa-example=billiardyoda/hpa-example-fast:v1
kubectl get pods -o custom-columns="POD:.metadata.name,IMAGE:.spec.containers[*].image,STATUS:.status.phase"


kubectl scale deployment hpa-example --replicas=4
```

Clean up Kubernetes
```
kubectl delete service hpa-example
kubectl delete deployment hpa-example
kubectl delete hpa hpa-example
gcloud compute forwarding-rules list
gcloud container clusters delete hpa-example 
```



