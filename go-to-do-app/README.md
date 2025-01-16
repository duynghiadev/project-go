[source code](https://github.com/schadokar/go-to-do-app/)

---

# 📝 Go To Do App

**Server: Golang
Client: React, semantic-ui-react
Database: Local MongoDB**

The offline version of application `Get Shit Done` is hosted a

# Highlights

1. DB connection string, name and collection name moved to `.env` file as environment variable. Using `github.com/joho/godotenv` to read the environment variables.

## Application Requirement

### golang server requirement

1. golang https://golang.org/dl/
2. gorilla/mux package for router `go get -u github.com/gorilla/mux`
3. mongo-driver package to connect with mongoDB `go get go.mongodb.org/mongo-driver`
4. github.com/joho/godotenv to access the environment variable.

### react client

From the Application directory

`create-react-app client`

## 💻 Start the application

1. Make sure your mongoDB is started
2. Create a `.env` file inside the `go-server` and copy the keys from `.env.example` and update the DB connection string.
3. From go-server directory, open a terminal and run
   `go run main.go`
4. From client directory,
   a. install all the dependencies using `npm install`
   b. start client `npm start`

## 🐼 Walk through the application

Open application at http://localhost:3000

### Index page

![](https://github.com/duynghiadev/project-go/blob/main/go-to-do-app/images/index.PNG?raw=true)

### Create task

Enter a task and Enter

### Task Complete

On completion of a task, click "done" Icon of the respective task card.

![](https://github.com/duynghiadev/project-go/blob/main/go-to-do-app/images/taskComplete.PNG?raw=true)

You'll notice on completion of task, card's bottom line color changed from yellow to green.

### Undo a task

To undone a task, click on "undo" Icon,

![](https://github.com/duynghiadev/project-go/blob/main/go-to-do-app/images/createTask.PNG?raw=true)

You'll notice on completion of task, card's bottom line color changed from green to yellow.

### Delete a task

To delete a task, click on "delete" Icon.

![](https://github.com/duynghiadev/project-go/blob/main/go-to-do-app/images/deletetask.PNG?raw=true)

---

## Author

#### 🌞 Duy Nghia Dev

I am software engineer and love golang, reactjs, nextjs, and nodejs.

# References

https://godoc.org/go.mongodb.org/mongo-driver/mongo
https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
https://vkt.sh/go-mongodb-driver-cookbook/
