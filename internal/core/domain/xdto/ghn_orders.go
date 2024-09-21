package xdto

type CreateOrderDTO struct {
	Token            string `validate:"required"`
	ShopID           int    `validate:"required"`
	ToName           string `validate:"required,max=1024"`
	FromName         string `validate:"required,max=1024"`
	FromPhone        string `validate:"required"`
	FromAddress      string `validate:"required,max=1024"`
	FromWardName     string `validate:"required"`
	FromDistrictName string `validate:"required"`
	FromProvinceName string `validate:"required"`
	ToPhone          string `validate:"required"`
	ToAddress        string `validate:"required,max=1024"`
	ToWardCode       string `validate:"required"`
	ToDistrictID     int    `validate:"required"`
	ReturnPhone      string
	ReturnAddress    string `validate:"max=1024"`
	ReturnDistrictID int
	ReturnWardCode   string
	ClientOrderCode  string `validate:"max=50"`
	CodAmount        int    `validate:"max=50000000"`
	Content          string `validate:"max=2000"`
	Weight           int    `validate:"required,max=50000"`
	Length           int    `validate:"required,max=200"`
	Width            int    `validate:"required,max=200"`
	Height           int    `validate:"required,max=200"`
	InsuranceValue   int    `validate:"max=5000000"`
	PickStationID    int
	Coupon           string
	ServiceTypeID    int    `validate:"required"`
	PaymentTypeID    int    `validate:"required"`
	Note             string `validate:"max=5000"`
	RequiredNote     string `validate:"required,oneof=CHOTHUHANG CHOXEMHANGKHONGTHU KHONGCHOXEMHANG"`
	PickShift        string
	Items            []OrderItem `validate:"required"`
}

type OrderItem struct {
	Name     string `validate:"required"`
	Code     string
	Quantity int `validate:"required"`
	Price    int
	Length   int
	Weight   int `validate:"required"`
	Width    int
	Height   int
	Category string
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
