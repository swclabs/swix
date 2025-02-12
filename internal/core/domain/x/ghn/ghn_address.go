package ghn

type Province struct {
	ProvinceID   int    `json:"ProvinceID"`
	ProvinceName string `json:"ProvinceName"`
	Code         string `json:"Code"`
}

type ProvinceDTO struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Province `json:"data"`
}

type DistrictDTO struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []District `json:"data"`
}

type District struct {
	DistrictID   int    `json:"DistrictID"`
	ProvinceID   int    `json:"ProvinceID"`
	DistrictName string `json:"DistrictName"`
	Code         string `json:"Code"`
	Types        int    `json:"Type"`
	SupportType  int    `json:"SupportType"`
}

type Ward struct {
	WardCode   string `json:"WardCode"`
	DistrictID int    `json:"DistrictID"`
	WardName   string `json:"WardName"`
}

type WardDTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Ward `json:"data"`
}
