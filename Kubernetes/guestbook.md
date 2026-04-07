## Tutorial: Guestbook 
https://cloud.google.com/kubernetes-engine/docs/tutorials/guestbook

Use Google’s cloud shell: https://shell.cloud.google.com/ 

To deploy and run the guestbook application on GKE:
Set up a Redis master
Set up Redis workers
Set up the guestbook web frontend
Visit the professor's guestbook website,  leave a message about your guestbook.

Use us-central1 for your location

The tutorial overprovsions and may cause resource quota issues.
Use the following command instead of the tutorial's
```
gcloud container clusters create guestbook \
  --location=us-central1 \
  --disk-type=pd-standard \
  --disk-size=50GB \
  --num-nodes=2
```
