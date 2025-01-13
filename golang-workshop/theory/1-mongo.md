# MongoDB Lesson: Using MongoDB for a TODO List Service

## Overview

MongoDB is a popular NoSQL database known for its flexibility and scalability. It stores data in JSON-like documents, making it easy to work with and highly adaptable to a variety of applications. In this lesson, we will explore how MongoDB can be used to create a TODO list service, covering key concepts and practical examples.

## Key Concepts

### 1. MongoDB Basics

- **Documents**: MongoDB stores data in documents, which are similar to JSON objects. Each document is a set of key-value pairs.
- **Collections**: Documents are grouped into collections. Collections are akin to tables in relational databases, but they don't enforce a fixed schema.
- **Database**: A database in MongoDB is a container for collections.

### 2. ObjectID

MongoDB uses a special type called `ObjectID` as the default unique identifier for documents. `ObjectID` is a 12-byte identifier that ensures uniqueness and includes a timestamp.

### 3. CRUD Operations

CRUD stands for Create, Read, Update, and Delete. These are the basic operations for managing data in a database.

## Practical Examples

### Adding a Task

To add a new task to our TODO list, we will insert a document into the `tasks` collection. MongoDB will automatically generate an `_id` field for the document if it is not provided.

#### Example Document

```json
{
  "title": "Buy groceries",
  "completed": false
}
```

#### Insert Operation

```sh
db.tasks.insertOne({
  "title": "Buy groceries",
  "completed": false
})
```

### Updating a Task

To update an existing task, we need to use the task's `_id` to identify it. We can then update specific fields using the `$set` operator.

#### Update Operation

```sh
db.tasks.updateOne(
  { "_id": ObjectId("60b8d295f9c5de3b2a7a564d") },
  { "$set": { "title": "Buy groceries and cook dinner" } }
)
```

### Marking a Task as Completed

To mark a task as completed, we update the `completed` field to `true`.

#### Update Operation

```sh
db.tasks.updateOne(
  { "_id": ObjectId("60b8d295f9c5de3b2a7a564d") },
  { "$set": { "completed": true } }
)
```

### Retrieving All Tasks

To retrieve all tasks from the `tasks` collection, we use the `find` method. This will return a cursor that we can iterate over to access each task document.

#### Find Operation

```sh
db.tasks.find()
```

### Deleting a Task

To delete a task, we use the task's `_id` to identify and remove it from the collection.

#### Delete Operation

```sh
db.tasks.deleteOne({ "_id": ObjectId("60b8d295f9c5de3b2a7a564d") })
```

## Advanced Concepts

### Indexes

Indexes improve the performance of search queries by creating a data structure that MongoDB can quickly traverse. By default, MongoDB creates an index on the `_id` field. You can create additional indexes to optimize queries on other fields.

#### Creating an Index

```sh
db.tasks.createIndex({ "title": 1 })
```

### Aggregation

Aggregation operations process data records and return computed results. This is useful for performing complex queries and data transformations.

#### Aggregation Example

Find all completed tasks and count them:

```sh
db.tasks.aggregate([
  { "$match": { "completed": true } },
  { "$count": "completedTasksCount" }
])
```

## Summary

MongoDB's flexible schema and powerful querying capabilities make it an excellent choice for building a TODO list service. By understanding the basics of documents, collections, and CRUD operations, as well as more advanced features like indexing and aggregation, you can effectively manage and query your data.

In this lesson, we've covered:
- The basic structure of MongoDB and its key components.
- How to perform CRUD operations on a TODO list.
- Advanced concepts like indexing and aggregation to optimize and analyze data.

These concepts provide a solid foundation for working with MongoDB in a variety of applications.