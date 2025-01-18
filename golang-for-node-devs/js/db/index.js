const { MongoClient } = require("mongodb");
require("dotenv").config();

const URI = process.env.MONGO_URI;
const databaseStr = process.env.DATABASE_NAME;

const collections = {
  products: "products",
};

const client = new MongoClient(URI, {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

async function connectToDB() {
  try {
    await client.connect();
    console.log("Database connected successfully");
  } catch (error) {
    console.error("Error connecting to database:", error.message);
    await client.close();
    process.exit(1);
  }
}

const database = client.db(databaseStr);

module.exports = {
  database,
  collections,
  connectToDB,
};
