package main

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readconcern"
    "log"
    "time"
)

// You will be using this Trainer type later in the program
type Trainer struct {
    Name string
    Age  int
    City string
}

func main() {
    // Set client options
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

    clientOptions.SetLocalThreshold(3 * time.Second)     //只使用与mongo操作耗时小于3秒的
    clientOptions.SetMaxConnIdleTime(5 * time.Second)    //指定连接可以保持空闲的最大毫秒数
    clientOptions.SetMaxPoolSize(200)                    //使用最大的连接数
    // clientOptions.SetReadPreference(want)                //表示只使用辅助节点
    clientOptions.SetReadConcern(readconcern.Majority()) //指定查询应返回实例的最新数据确认为，已写入副本集中的大多数成员
    // clientOptions.SetWriteConcern(wc)                    //请求确认写操作传播到大多数mongod实例

    // Connect to MongoDB
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")


    collection := client.Database("test").Collection("trainers")
    ash := Trainer{"Ash", 10, "Pallet Town"}
    misty := Trainer{"Misty", 10, "Cerulean City"}
    brock := Trainer{"Brock", 15, "Pewter City"}

    insertResult, err := collection.InsertOne(context.TODO(), ash)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted a single document: ", insertResult.InsertedID)
    trainers := []interface{}{misty, brock}

    insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)


    filter := bson.D{{"name", "Ash"}}

    update := bson.D{
        {"$inc", bson.D{
            {"age", 1},
        }},
    }
    updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)


    // create a value into which the result can be decoded
    var result Trainer
    err = collection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found a single document: %+v\n", result)



    // Pass these options to the Find method
    findOptions := options.Find()
    findOptions.SetLimit(2)
    // Here's an array in which you can store the decoded documents
    var results []*Trainer
    // Passing bson.D{{}} as the filter matches all documents in the collection
    cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
    if err != nil {
        log.Fatal(err)
    }
    // Finding multiple documents returns a cursor
    // Iterating through the cursor allows us to decode documents one at a time
    for cur.Next(context.TODO()) {
        // create a value into which the single document can be decoded
        var elem Trainer
        err := cur.Decode(&elem)
        if err != nil {
            log.Fatal(err)
        }
        results = append(results, &elem)
    }

    if err := cur.Err(); err != nil {
        log.Fatal(err)
    }
    // Close the cursor once finished
    cur.Close(context.TODO())
    fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
    for k,v:=range results {
        log.Println(k,v.Name,v.Age,v.City)
    }


    deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)


    err = client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connection to MongoDB closed.")


}