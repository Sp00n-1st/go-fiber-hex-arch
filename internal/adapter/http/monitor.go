package http

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type MonitoringHandler struct {
	MongoClient *mongo.Client
}

func NewMonitoringHandler(client *mongo.Client) *MonitoringHandler {
	return &MonitoringHandler{MongoClient: client}
}

func (h *MonitoringHandler) GetMonitoringData(c *fiber.Ctx) error {
	collection := h.MongoClient.Database("Monitoring").Collection("logs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, map[string]interface{}{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer cursor.Close(ctx)

	var results []map[string]interface{}
	if err = cursor.All(ctx, &results); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(results)
}
