package handlers

import (
	"context"
	"fmt"
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

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var todo models.Todo
	if err := config.TodoCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&todo); err != nil {
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

	newTodo.ID = result.InsertedID.(primitive.ObjectID).String()
	return c.JSON(newTodo)

}
func UpdateTodo(c *fiber.Ctx) error {

	id := c.Params("id")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	var updatedTodo models.Todo

	if err := c.BodyParser(&updatedTodo); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	update := bson.M{"$set": bson.M{
		"title":  updatedTodo.Title,
		"isDone": updatedTodo.IsDone,
	}}
	if _, err := config.TodoCollection.UpdateOne(ctx, bson.M{"_id": objId}, update); err != nil {
		fmt.Print(err)
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(updatedTodo)

}
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := config.TodoCollection.DeleteOne(ctx, bson.M{"_id": objId}); err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
