# 🚀 Docker Commands for This Project

## ⚠ Issue: Running Multiple Commands Repeatedly

Each time I run Docker Compose, I have to execute three separate commands:

```sh
docker compose up -d db
docker compose up -d goapp
docker compose up -d nextapp
```

This is repetitive. **How can I run all containers with a single command?**

✅ **Solution:** Run all services at once:

```sh
docker compose up -d
```

This will start **all services** defined in `docker-compose.yml`.

---

# 🛠 Key Docker Commands in This Project

## 📌 1. Connect to PostgreSQL Database in Docker

Run this command to access the PostgreSQL database via the terminal:

```sh
docker exec -it db psql -U postgres
```

## 📌 2. Initialize the Database (`db`) in Docker

This command starts the database container based on `docker-compose.yml`:

```sh
docker compose up -d db
```

## 📌 3. Build the App Image in Docker

Navigate to the `go-fullstack-app` folder and build the Docker image:

```sh
docker compose build
```

## 📌 4. Start the Backend API (`goapp`)

This command creates and starts the `goapp` container:

```sh
docker compose up -d goapp
```

## 📌 5. Start the Frontend (`nextapp`)

This command creates and starts the `nextapp` container:

```sh
docker compose up -d nextapp
```

---

# 🔍 Check Running Containers

To verify that all containers are running, use:

```sh
docker ps -a
```

---

# 🎯 Final Steps: Run Everything with One Command

Instead of running separate commands for each service, use:

```sh
docker compose up -d
```

This starts **all services** defined in `docker-compose.yml`.

---

This document provides a clear and structured guide for managing Docker containers efficiently. 🚀
