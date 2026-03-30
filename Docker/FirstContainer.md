# First Alpine Linux Containers

*Sep 19, 2017 • @jimcodified*

In this lab you will run a popular, free, lightweight container and explore the basics of how containers work and how the Docker Engine executes and isolates containers.

If you already have experience running containers and basic Docker commands you can probably skip this intro exercise.

---

## Concepts in this exercise

- Docker engine  
- Containers & images  
- Image registries and Docker Hub  
- Container isolation  

---

## Tips

- Inline snippets are examples:
  ```bash
  uname -a
  ```

- Commands you should type (replace placeholders like `<container ID>`):
  ```bash
  docker container start <container ID>
  ```

---

## 1.0 Running your first container

Run your first Docker container:

```bash
docker container run hello-world
```

### What happened?

- Docker looked for the `hello-world` image locally  
- It wasn’t found, so it pulled it from Docker Hub  
- The container ran and printed output  
- Then it exited  

### Containers vs VMs

- **VM** = hardware abstraction (CPU, RAM, OS)
- **Container** = application abstraction (app + dependencies)

---

## 1.1 Docker Images

Run an Alpine Linux container.

### Pull the image

```bash
docker image pull alpine
```

### List images

```bash
docker image ls
```

Example output:

```
REPOSITORY     TAG     IMAGE ID       CREATED        SIZE
alpine         latest  c51f86c28340   4 weeks ago    1.1 MB
hello-world    latest  690ed74de00f   5 months ago   960 B
```

---

## 1.1 Running Containers

### Run a command in a container

```bash
docker container run alpine ls -l
```

This:
1. Finds the image  
2. Creates a container  
3. Runs the command  
4. Stops the container  

---

### Another example

```bash
docker container run alpine echo "hello from alpine"
```

Output:
```
hello from alpine
```

---

### Run a shell (non-interactive)

```bash
docker container run alpine /bin/sh
```

This exits immediately because no interaction was provided.

---

### Run an interactive shell

```bash
docker container run -it alpine /bin/sh
```

Now inside the container, try:

```bash
ls -l
uname -a
```

Exit with:

```bash
exit
```

---

## List Containers

### Running containers

```bash
docker container ls
```

### All containers

```bash
docker container ls -a
```

Example:

```
CONTAINER ID   IMAGE        COMMAND        STATUS
36171a5da744   alpine       "/bin/sh"      Exited
a6a9d46d0b2f   alpine       "echo ..."     Exited
ff0a5c3750b9   alpine       "ls -l"        Exited
c317d0a9e3d2   hello-world  "/hello"       Exited
```

---

## 1.2 Container Isolation

Each container is isolated, even if based on the same image.

---

### Create a file in a container

```bash
docker container run -it alpine /bin/ash
```

Inside the container:

```bash
echo "hello world" > hello.txt
ls
exit
```

---

### Run a new container

```bash
docker container run alpine ls
```

Notice:

👉 `hello.txt` is **missing**

This proves containers are isolated.

---

## Why so many containers?

Each `docker run` creates a **new container instance** with:

- Its own filesystem  
- Its own namespace  
- No access to other containers by default  

---

## Find your previous container

```bash
docker container ls -a
```

Example:

```
CONTAINER ID   IMAGE   COMMAND     STATUS
3030c9c91e12   alpine  "/bin/ash"  Exited
```

---

## Restart a container

```bash
docker container start <container ID>
```

Tip: You can use a shortened ID:

```bash
docker container start 3030
```

---

## Verify it's running

```bash
docker container ls
```

---

## Execute a command inside it

```bash
docker container exec <container ID> ls
```

Now you should see:

```
hello.txt
```

---

## Key Takeaways

- Each `docker run` = new container  
- Containers are **isolated by default**  
- Use `exec` to interact with running containers  
- Containers are fast because they skip full OS boot  
