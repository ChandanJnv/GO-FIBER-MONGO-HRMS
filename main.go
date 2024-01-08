package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var mg MongoInstance

const (
	dbName   = "fiber-hrms"
	mongoURI = "mongodb://localhost:27017"
)

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int64   `json:"age"`
}

func Connect() error {
	connOpt := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), connOpt)
	if err != nil {
		return err
	}

	db := client.Database(dbName)
	mg = MongoInstance{
		Client: client,
		DB:     db,
	}
	return nil
}

func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {
		query := bson.D{{}}
		var employees []Employee = make([]Employee, 0)

		cur, err := mg.DB.Collection("employee").Find(context.Background(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		if err := cur.All(c.Context(), &employees); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(employees)
	})

	app.Post("/employee", func(c *fiber.Ctx) error {
		collection := mg.DB.Collection("employee")

		emp := new(Employee)

		if err := c.BodyParser(emp); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		emp.ID = ""

		result, err := collection.InsertOne(c.Context(), emp)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: result.InsertedID}}
		createdRecord := collection.FindOne(c.Context(), filter)

		createdEmployee := &Employee{}
		createdRecord.Decode(&createdEmployee)

		return c.Status(201).JSON(createdEmployee)
	})

	app.Put("/employee/:id", func(c *fiber.Ctx) error {
		idParams := c.Params("id")

		empId, err := primitive.ObjectIDFromHex(idParams)
		if err != nil {
			return c.SendStatus(400)
		}

		emp := new(Employee)
		if err := c.BodyParser(emp); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		query := bson.D{{Key: "_id", Value: empId}}
		update := bson.D{
			{Key: "$set",
				Value: bson.D{
					{Key: "name", Value: emp.Name},
					{Key: "age", Value: emp.Age},
					{Key: "salary", Value: emp.Salary},
				},
			},
		}

		if err := mg.DB.Collection("employee").FindOneAndUpdate(c.Context(), query, update).Err(); err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(400)
			}
			return c.SendStatus(500)
		}

		emp.ID = idParams

		return c.Status(200).JSON(emp)
	})

	app.Delete("/employee/:id", func(c *fiber.Ctx) error {
		empID, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.SendStatus(400)
		}

		query := bson.D{{Key: "_id", Value: empID}}
		result, err := mg.DB.Collection("employee").DeleteOne(c.Context(), &query)
		if err != nil {
			return c.SendStatus(404)
		}

		if result.DeletedCount < 1 {
			return c.SendStatus(404)
		}

		return c.Status(200).JSON("record-deleted")
	})

	log.Fatalln(app.Listen(":3000"))

}
