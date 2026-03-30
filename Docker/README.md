# Running containers in AWS

### Setting up Docker and Docker Compose
Ideally this would all go into a n EC2 template
```
#!/bin/sh
yum -y update
yum -y install docker
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

echo "Docker installed" >>/home/ec2-user/docker-status
usermod -a -G docker ec2-user
systemctl enable docker.service
systemctl start docker.service
echo "Docker configured" >>/home/ec2-user/docker-status

```

### Examples
Run ngnix
```
docker run --rm -p 88:80 nginx
```

Run the Juice Shop hacking exercise 
```
docker run --rm -p 3000:3000 bkimminich/juice-shop
```

#  Create a container that will run an iPython Docker lesson

### 1. Buid the Doker image
```
docker build -t ipython-notebook .
```
### 2. Run the container
```
docker run -p 8888:8888 -v $(pwd):/app ipython-notebook
```
What this does:
-p 8888:80 → maps container port to your machine
-v $(pwd):/app → mounts current folder so notebooks persist
