#  Create a container that run an iPython Docker lesson

### 1. Buid the Doker image
```
docker build -t ipython-notebook .
```
### 2. Run the container
```
docker run -p 8888:80 -v $(pwd):/app ipython-notebook
```
What this does:
-p 8888:80 → maps container port to your machine
-v $(pwd):/app → mounts current folder so notebooks persist
