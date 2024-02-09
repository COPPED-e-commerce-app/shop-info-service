package rest

import (
	"github.com/COPPED/shop-info-service/internal/shop-info-service/service"
	"github.com/COPPED/shop-info-service/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func FetchShopInfo(c *fiber.Ctx) error {
	id := c.Params("id")
	shopInfo := service.FetchShopInfo(id)
	return c.JSON(shopInfo)
}

func CreateShopInfo(c *fiber.Ctx) error {
	request := &models.ShopInfoRequest{}
	if err := c.BodyParser(request); err != nil {
		return err
	}
	service.CreateShopInfo(request)
	return c.Status(fiber.StatusCreated).SendString("Shop Info entity created!")
}

func UpdateShopInfo(c *fiber.Ctx) error {
	request := &models.ShopInfoRequest{}
	if err := c.BodyParser(request); err != nil {
		return err
	}
	service.UpdateShopInfo(request)
	return c.Status(fiber.StatusNoContent).SendString("Shop Info entity updated!")
}

func FetchShopInfoList(c *fiber.Ctx) error {
	request := &models.ShopInfoListRequest{}
	if err := c.BodyParser(request); err != nil {
		return err
	}
	shopInfo := service.FetchShopInfoList(request)
	return c.JSON(shopInfo)
}

func UpdateShopInfoList(c *fiber.Ctx) error {
	request := &models.ShopInfoListRequest{}
	if err := c.BodyParser(request); err != nil {
		return err
	}
	service.UpdateShopInfoList(request)
	return c.Status(fiber.StatusNoContent).SendString("Shop Info List entities updated!")
}
