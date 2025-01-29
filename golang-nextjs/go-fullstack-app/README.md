# Summary these command build and run docker in this project

---

this is command to see database postgres using command line, show in terminal

`docker exec -it db psql -U postgres`

---

this is command initial database `db` in docker, it based on file `compose.yaml`, at part `db: ..`

`docker compose up -d db`

---

this is command to build app into docker, stand up at folder `go-fullstack-app`, then run command:

`docker compose build`

---

this is command to create `./api` in folder backend

`docker compose up -d goapp`

---

then, i running command `docker ps -a` to see add container are running in docker

# Run these are command line below to open app

Running this command to create container `blogapp`

`docker compose up -d goapp`

Running this command to create container `nextapp`

`docker compose up -d nextapp`
