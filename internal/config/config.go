package config

import (
    "context"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var (
    AppPort  string
    MongoURI string
    DBName   string
    MongoDB  *mongo.Database
)

func LoadEnv() {
    godotenv.Load()
    AppPort = getEnv("APP_PORT", "3030")
    MongoURI = getEnv("MONGO_URI", "mongodb://mongo:27017")
    DBName = getEnv("MONGO_DB", "crud_mongo_db")
    connectMongo()
}

func connectMongo() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoURI))
    if err != nil {
        log.Fatal(err)
    }

    MongoDB = client.Database(DBName)
    log.Println("✅ MongoDB connected")
}

func getEnv(key, fallback string) string {
    if v, ok := os.LookupEnv(key); ok {
        return v
    }
    return fallback
}
