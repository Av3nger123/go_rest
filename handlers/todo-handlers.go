package handlers

import (
	"context"
	"time"
	"todo-app/config"
	"todo-app/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTodo(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := config.TodoCollection.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	todos := []models.Todo{}
	for cur.Next(ctx) {
		var todo models.Todo
		if err := cur.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)

}

func GetTodoById(c *fiber.Ctx) error {

	id := c.Params("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var todo models.Todo
	if err := config.TodoCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&todo); err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(todo)

}
func CreateTodo(c *fiber.Ctx) error {
	var newTodo models.Todo

	if err := c.BodyParser(&newTodo); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := config.TodoCollection.InsertOne(ctx, newTodo)
	if err != nil {
		return err
	}

	newTodo.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return c.JSON(newTodo)

}
func UpdateTodo(c *fiber.Ctx) error {

	id := c.Params("id")
	var updatedTodo models.Todo

	if err := c.BodyParser(&updatedTodo); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.M{"$set": updatedTodo}
	if _, err := config.TodoCollection.UpdateOne(ctx, bson.M{"_id": id}, update); err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(updatedTodo)

}
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := config.TodoCollection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
