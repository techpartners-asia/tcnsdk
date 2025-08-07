package tcnsdk

// AuthRequest represents authentication request
type AuthRequest struct {
	AppID  string `json:"appID"`
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// AuthResponse represents authentication response
type AuthResponse struct {
	Token     string `json:"Token"`
	ExpiresIn int    `json:"ExpiresIn"`
}

// BaseResponse represents common API response structure
type BaseResponse struct {
	StatusCode int         `json:"statusCode"`
	Succeeded  bool        `json:"succeeded"`
	Errors     interface{} `json:"errors"`
	Extras     interface{} `json:"extras"`
	Timestamp  int64       `json:"timestamp"`
}

// MachineInfo represents device information
type MachineInfo struct {
	MachineID    string `json:"machineId"`
	MachineName  string `json:"machineName"`
	Signal       int    `json:"signal"`
	Temperature  string `json:"temperature"`
	RunningState int    `json:"runningState"`
}

// MachineInfoResponse represents device information response
type MachineInfoResponse struct {
	BaseResponse
	Data MachineInfo `json:"data"`
}

// Commodity represents product information
type Commodity struct {
	CommodityID   string  `json:"commodityId"`
	CommodityName string  `json:"commodityName"`
	CommoditySKU  string  `json:"commoditySKU"`
	Weight        int     `json:"weight"`
	Price         float64 `json:"price"`
	DoorNo        int     `json:"doorNo"`
	PictureURL    string  `json:"pictureUrl"`
}

// CommodityResponse represents product information response
type CommodityResponse struct {
	BaseResponse
	Data []Commodity `json:"data"`
}

// PayConfig represents payment configuration
type PayConfig struct {
	MachineID      string `json:"machineId"`
	PreAuthorLimit string `json:"preAuthorLimit"`
	Currency       string `json:"currency"`
	QRCodeURL      string `json:"qrCodeUrl"`
}

// PayConfigResponse represents payment configuration response
type PayConfigResponse struct {
	BaseResponse
	Data PayConfig `json:"data"`
}

// ConfirmRepliItem represents replenishment confirmation item
type ConfirmRepliItem struct {
	CommodityID string `json:"CommodityId"`
	LayerNo     int    `json:"LayerNo"`
	SpotCount   int    `json:"SpotCount"`
	StockCount  int    `json:"StockCount"`
	Sno         int    `json:"Sno"`
}

// ConfirmRepliRequest represents replenishment confirmation request
type ConfirmRepliRequest struct {
	Mid        string             `json:"Mid"`
	TransID    int64              `json:"TransId"`
	TrackingID string             `json:"TrackingId"`
	DoorNo     int                `json:"DoorNo"`
	Data       []ConfirmRepliItem `json:"Data"`
}

// ConfirmRepliResponse represents replenishment confirmation response
type ConfirmRepliResponse struct {
	TrackingID string `json:"TrackingId"`
	TransID    int64  `json:"TransId"`
	Success    bool   `json:"Success"`
	Message    string `json:"Message"`
}

// ReportPaymentResultRequest represents payment result reporting request
type ReportPaymentResultRequest struct {
	OrderID      string `json:"OrderId"`
	OrderNo      string `json:"OrderNo"`
	PayType      int    `json:"PayType"`
	PayStatus    int    `json:"PayStatus"`
	ErrorMessage string `json:"ErrorMessage,omitempty"`
}

// ReportPaymentResultResponse represents payment result reporting response
type ReportPaymentResultResponse struct {
	OrderID string `json:"OrderId"`
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
}

// CallbackEvent represents callback event structure
type CallbackEvent struct {
	OrderID         string      `json:"OrderId"`
	OrderNo         string      `json:"OrderNo"`
	TranseType      int         `json:"TranseType"`
	Action          string      `json:"Action"`
	Status          bool        `json:"Status"`
	Msg             string      `json:"Msg"`
	CustomerDetails string      `json:"CustomerDetails"`
	Data            interface{} `json:"Data"`
}

// DetectOrderDetail represents order detection details
type DetectOrderDetail struct {
	TradeProcessMode   int                 `json:"TradeProcessMode"`
	TradeProductModels []TradeProductModel `json:"TradeProductModels"`
}

// TradeProductModel represents product model in detection
type TradeProductModel struct {
	ID         string  `json:"Id"`
	AlisName   string  `json:"AlisName"`
	BuyCount   int     `json:"BuyCount"`
	MeterType  int     `json:"MeterType"`
	Price      float64 `json:"Price"`
	PictureURL string  `json:"PictureUrl,omitempty"`
}

// OrderAdjustment represents order adjustment data
type OrderAdjustment struct {
	OrderNo       string              `json:"OrderNo"`
	OrderProducts []TradeProductModel `json:"OrderProducts"`
}

// OrderRefund represents order refund data
type OrderRefund struct {
	RefundRemark        string              `json:"RefundRemark"`
	RefundOrderNo       string              `json:"Refund Order No"`
	OrderRefundProducts []TradeProductModel `json:"OrderRefundProducts"`
}

// Constants for transaction types
const (
	TranseTypePurchase = 0
	TranseTypeRestock  = 2
)

// Constants for payment types
const (
	PayTypeNormal    = 0
	PayTypeDeduction = 1
	PayTypeRefund    = 2
)

// Constants for payment status
const (
	PayStatusSuccess = 1
	PayStatusFailed  = 2
)

// Constants for trade process modes
const (
	TradeProcessModeNormal    = 0
	TradeProcessModeCancel    = 1
	TradeProcessModeInterrupt = 2
)

// Constants for meter types
const (
	MeterTypePiece  = 1
	MeterTypeWeight = 2
)
