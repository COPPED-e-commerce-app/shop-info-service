package models

type ShopInfoListRequest struct {
	UserIdList     []string `json:"user_id_list"`
	ShopInfoIdList []string `json:"shop_info_id_list"`
}
