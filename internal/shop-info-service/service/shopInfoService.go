package service

import (
	"log"

	"github.com/COPPED/shop-info-service/pkg/models"
)

func FetchShopInfo(id string) *models.ShopInfo {
	return &models.ShopInfo{}
}

func CreateShopInfo(request *models.ShopInfoRequest) {
	shopInfoShopInfo := &models.ShopInfo{
		Name: request.Name,
	}
	log.Printf("Created shopInfoShopInfo: %+v", shopInfoShopInfo)
}

func UpdateShopInfo(request *models.ShopInfoRequest) {
	shopInfoShopInfo := &models.ShopInfo{
		Name: request.Name,
	}
	log.Printf("Updated shopInfoShopInfo: %+v", shopInfoShopInfo)
}

func FetchShopInfoList(request *models.ShopInfoListRequest) *[]models.ShopInfo {
	shopInfoShopInfo := []models.ShopInfo{}
	for _, id := range request.ShopInfoIdList {
		shopInfoShopInfo = append(shopInfoShopInfo, models.ShopInfo{
			Id: id,
		})
	}
	return &shopInfoShopInfo
}

func UpdateShopInfoList(request *models.ShopInfoListRequest) {
	log.Printf("Shop Infos updated!")
}
