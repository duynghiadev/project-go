**Nginx Web server**

**MySql Running on Host machine**

---

# Compile go app

- `CGO_ENABLED=0  GOOS=linux go build -o build/server`
- Back-end run at port 8000

# Build react app

- npm run build
- Front-end run at port 8080

# Build docker image

`docker build -t blog-app .`

# Run container

`docker run --name blog_container -p 8080:80 -p 8000:8000 -it --rm blog-app`

# Run container in detach mode

`docker run --name blog_container -p 8080:80 -p 8000:8000 -it -d --rm blog-app`
