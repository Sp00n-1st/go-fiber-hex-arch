package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go-fiber-hex-arch/internal/adapter/config"
	"go-fiber-hex-arch/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"time"
)

func MonitoringFuncPerformance(client *mongo.Client, cfg *config.Container) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Path() == cfg.HTTP.Prefix+"/monitoring" {
			return c.Next()
		}

		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		mongo := client.Database(cfg.MONGO.MongoDB).Collection(cfg.MONGO.MongoCollection)

		go func() {
			result, err := mongo.InsertOne(context.Background(), map[string]interface{}{
				"path":     c.Path(),
				"method":   c.Method(),
				"duration": util.FormatDuration(duration),
				"time":     time.Now().In(util.Loc),
			})
			if err != nil {
				slog.Error("Failed to insert performance data", "error", err)
			} else {
				slog.Info("Inserted performance data", "result", result)
			}
		}()

		return err
	}
}
