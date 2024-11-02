package ghn

import "time"

type CreateOrderDTO struct {
	Token            string      `json:"token" validate:"required"`
	ShopID           int         `json:"shop_id" validate:"required"`
	ToName           string      `json:"to_name" validate:"required,max=1024"`
	FromName         string      `json:"from_name" validate:"required,max=1024"`
	FromPhone        string      `json:"from_phone" validate:"required"`
	FromAddress      string      `json:"from_address" validate:"required,max=1024"`
	FromWardName     string      `json:"from_ward_name" validate:"required"`
	FromDistrictName string      `json:"from_district_name" validate:"required"`
	FromProvinceName string      `json:"from_province_name" validate:"required"`
	ToPhone          string      `json:"to_phone" validate:"required"`
	ToAddress        string      `json:"to_address" validate:"required,max=1024"`
	ToWardCode       string      `json:"to_ward_code" validate:"required"`
	ToDistrictID     int         `json:"to_district_id" validate:"required"`
	ReturnPhone      string      `json:"return_phone"`
	ReturnAddress    string      `json:"return_address" validate:"max=1024"`
	ReturnDistrictID int         `json:"return_district_id"`
	ReturnWardCode   string      `json:"return_ward_code"`
	ClientOrderCode  string      `json:"client_order_code" validate:"max=50"`
	CodAmount        int         `json:"cod_amount" validate:"max=50000000"`
	Content          string      `json:"content" validate:"max=2000"`
	Weight           int         `json:"weight" validate:"required,max=50000"`
	Length           int         `json:"length" validate:"required,max=200"`
	Width            int         `json:"width" validate:"required,max=200"`
	Height           int         `json:"height" validate:"required,max=200"`
	InsuranceValue   int         `json:"insurance_value" validate:"max=5000000"`
	PickStationID    int         `json:"pick_station_id"`
	Coupon           string      `json:"coupon"`
	ServiceTypeID    int         `json:"service_type_id" validate:"required"`
	PaymentTypeID    int         `json:"payment_type_id" validate:"required"`
	Note             string      `json:"note" validate:"max=5000"`
	RequiredNote     string      `json:"required_note" validate:"required,oneof=CHOTHUHANG CHOXEMHANGKHONGTHU KHONGCHOXEMHANG"`
	PickShift        string      `json:"pick_shift"`
	Items            []OrderItem `json:"items" validate:"required"`
}

type OrderItem struct {
	Name     string `json:"name" validate:"required"`
	Code     string `json:"code"`
	Quantity int    `json:"quantity" validate:"required"`
	Price    int    `json:"price"`
	Length   int    `json:"length"`
	Weight   int    `json:"weight" validate:"required"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Category string `json:"category"`
}

type OrderDTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	DistrictEncode       string `json:"district_encode"`
	ExpectedDeliveryTime string `json:"expected_delivery_time"`
	Fee                  Fee    `json:"fee"`
	OrderCode            string `json:"order_code"`
	SortCode             string `json:"sort_code"`
	TotalFee             string `json:"total_fee"`
	TransType            string `json:"trans_type"`
	WardEncode           string `json:"ward_encode"`
}

type Fee struct {
	Coupon      int `json:"coupon"`
	Insurance   int `json:"insurance"`
	MainService int `json:"main_service"`
	R2S         int `json:"r2s"`
	Return      int `json:"return"`
	StationDo   int `json:"station_do"`
	StationPu   int `json:"station_pu"`
}

type Log struct {
	Status      string    `json:"status"`
	UpdatedDate time.Time `json:"updated_date"`
}

type OrderInfo struct {
	ShopID               int       `json:"shop_id"`
	ClientID             int       `json:"client_id"`
	ReturnName           string    `json:"return_name"`
	ReturnPhone          string    `json:"return_phone"`
	ReturnAddress        string    `json:"return_address"`
	ReturnWardCode       string    `json:"return_ward_code"`
	ReturnDistrictID     int       `json:"return_district_id"`
	FromName             string    `json:"from_name"`
	FromPhone            string    `json:"from_phone"`
	FromAddress          string    `json:"from_address"`
	FromWardCode         string    `json:"from_ward_code"`
	FromDistrictID       int       `json:"from_district_id"`
	DeliverStationID     int       `json:"deliver_station_id"`
	ToName               string    `json:"to_name"`
	ToPhone              string    `json:"to_phone"`
	ToAddress            string    `json:"to_address"`
	ToWardCode           string    `json:"to_ward_code"`
	ToDistrictID         int       `json:"to_district_id"`
	Weight               int       `json:"weight"`
	Length               int       `json:"length"`
	Width                int       `json:"width"`
	Height               int       `json:"height"`
	ConvertedWeight      int       `json:"converted_weight"`
	ServiceTypeID        int       `json:"service_type_id"`
	ServiceID            int       `json:"service_id"`
	PaymentTypeID        int       `json:"payment_type_id"`
	CustomServiceFee     int       `json:"custom_service_fee"`
	CODAmount            int       `json:"cod_amount"`
	CODCollectDate       string    `json:"cod_collect_date"`
	CODTransferDate      string    `json:"cod_transfer_date"`
	IsCODTransferred     bool      `json:"is_cod_transferred"`
	IsCODCollected       bool      `json:"is_cod_collected"`
	InsuranceValue       int       `json:"insurance_value"`
	OrderValue           int       `json:"order_value"`
	PickStationID        int       `json:"pick_station_id"`
	ClientOrderCode      string    `json:"client_order_code"`
	CODFailedAmount      int       `json:"cod_failed_amount"`
	CODFailedCollectDate string    `json:"cod_failed_collect_date"`
	RequiredNote         string    `json:"required_note"`
	Content              string    `json:"content"`
	Note                 string    `json:"note"`
	EmployeeNote         string    `json:"employee_note"`
	Coupon               string    `json:"coupon"`
	ID                   string    `json:"_id"`
	OrderCode            string    `json:"order_code"`
	VersionNo            string    `json:"version_no"`
	UpdatedIP            string    `json:"updated_ip"`
	UpdatedEmployee      int       `json:"updated_employee"`
	UpdatedClient        int       `json:"updated_client"`
	UpdatedSource        string    `json:"updated_source"`
	UpdatedDate          time.Time `json:"updated_date"`
	UpdatedWarehouse     int       `json:"updated_warehouse"`
	CreatedIP            string    `json:"created_ip"`
	CreatedEmployee      int       `json:"created_employee"`
	CreatedClient        int       `json:"created_client"`
	CreatedSource        string    `json:"created_source"`
	CreatedDate          time.Time `json:"created_date"`
	Status               string    `json:"status"`
	PickWarehouseID      int       `json:"pick_warehouse_id"`
	DeliverWarehouseID   int       `json:"deliver_warehouse_id"`
	CurrentWarehouseID   int       `json:"current_warehouse_id"`
	ReturnWarehouseID    int       `json:"return_warehouse_id"`
	NextWarehouseID      int       `json:"next_warehouse_id"`
	Leadtime             time.Time `json:"leadtime"`
	OrderDate            time.Time `json:"order_date"`
	SOCID                string    `json:"soc_id"`
	FinishDate           string    `json:"finish_date"`
	Tag                  []string  `json:"tag"`
	Logs                 []Log     `json:"log"`
}

type OrderInfoDTO struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    []OrderInfo `json:"data"`
}
