package http

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-hex-arch/internal/core/domain"
	"go-fiber-hex-arch/internal/core/service"
	"go-fiber-hex-arch/internal/dto"
	"strconv"
	"strings"
)

type ProductHandler struct {
	Service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{
		Service: service,
	}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	err := h.Service.InsertProduct(product)
	if err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			var errors []string
			for _, fieldErr := range validationErr {
				errors = append(errors, fmt.Sprintf("%s is %s", fieldErr.Field(), fieldErr.Tag()))
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": strings.Join(errors, ", ")})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToProductResponse(product))
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var product domain.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id, _ := strconv.ParseUint(c.Params("id"), 10, 0)
	product.ProductID = uint(id)

	err := h.Service.UpdateProduct(product)
	if err != nil {
		if validationErr, ok := err.(validator.ValidationErrors); ok {
			var errors []string
			for _, fieldErr := range validationErr {
				errors = append(errors, fmt.Sprintf("%s is %s", fieldErr.Field(), fieldErr.Tag()))
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": strings.Join(errors, ", ")})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(dto.ToProductResponse(product))
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.Service.DeleteProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	product, err := h.Service.GetProduct(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ToProductResponse(product))
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	products, err := h.Service.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(dto.ToProductResponses(products))
}
