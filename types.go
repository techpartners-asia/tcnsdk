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
	OrderID      string    `json:"OrderId"`
	OrderNo      string    `json:"OrderNo"`
	PayType      PayType   `json:"PayType"`
	PayStatus    PayStatus `json:"PayStatus"`
	ErrorMessage string    `json:"ErrorMessage,omitempty"`
}

// ReportPaymentResultResponse represents payment result reporting response
type ReportPaymentResultResponse struct {
	OrderID string `json:"OrderId"`
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
}

// Constants for transaction types
const (
	TranseTypePurchase = 0
	TranseTypeRestock  = 2
)

// TranseType represents transaction type
type TranseType int

// Constants for payment types
const (
	PayTypeNormal    = 0
	PayTypeDeduction = 1
	PayTypeRefund    = 2
)

// PayType represents payment type
type PayType int

// Constants for payment status
const (
	PayStatusSuccess = 1
	PayStatusFailed  = 2
)

// PayStatus represents payment status
type PayStatus int

// Constants for trade process modes
const (
	TradeProcessModeNormal    = 0
	TradeProcessModeCancel    = 1
	TradeProcessModeInterrupt = 2
)

// TradeProcessMode represents trade process mode
type TradeProcessMode int

// Constants for meter types
const (
	MeterTypePiece  = 1
	MeterTypeWeight = 2
)

// MeterType represents meter type
type MeterType int

// Visual Recognition SDK Types

// CommodityConfirmRequest represents product confirmation request
type CommodityConfirmRequest struct {
	AppID        string `json:"AppID"`
	CommoditySku string `json:"CommoditySku"`
}

// CommodityConfirmResponse represents product confirmation response
type CommodityConfirmResponse struct {
	Code int `json:"Code"`
	Data struct {
		CommodityName string `json:"CommodityName"`
		CommodityId   string `json:"CommodityId"`
		PictureUrl    string `json:"PictureUrl"`
	} `json:"Data"`
	Msg string `json:"Msg,omitempty"`
}

// ProductRegistrationRequest represents product registration request
type ProductRegistrationRequest struct {
	AppID      string   `json:"AppID"`
	AskId      string   `json:"AskId"`
	SkuName    string   `json:"SkuName"`
	Sku        string   `json:"Sku,omitempty"`
	IsStandard bool     `json:"IsStandard,omitempty"`
	ImgUrls    []string `json:"ImgUrls"`
	NotifyUrl  string   `json:"NotifyUrl"`
}

// ProductRegistrationResponse represents product registration response
type ProductRegistrationResponse struct {
	Code int `json:"Code"`
	Data struct {
		YsSkuId string `json:"YsSkuId"`
	} `json:"Data"`
	Msg string `json:"Msg,omitempty"`
}

// ProductReviewQueryRequest represents product review query request
type ProductReviewQueryRequest struct {
	AppID   string `json:"AppID"`
	YsSkuId string `json:"YsSkuId"`
}

// ProductReviewQueryResponse represents product review query response
type ProductReviewQueryResponse struct {
	Code int `json:"Code"`
	Data struct {
		State RecognitionState `json:"State"`
		Desc  string           `json:"Desc,omitempty"`
	} `json:"Data"`
	Msg string `json:"Msg,omitempty"`
}

// ProductReviewNotification represents product review notification
type ProductReviewNotification struct {
	AppID   string           `json:"AppID"`
	YsSkuId string           `json:"YsSkuId"`
	State   RecognitionState `json:"State"`
	Desc    string           `json:"Desc,omitempty"`
}

// RecognitionRequest represents video recognition request
type RecognitionRequest struct {
	AppID        string       `json:"AppID"`
	TaskId       string       `json:"TaskId"`
	ResourceType ResourceType `json:"ResourceType,omitempty"`
	ResourceUrl  []string     `json:"ResourceUrl"`
	ProductRange []string     `json:"ProductRange"`
	NotifyUrl    string       `json:"NotifyUrl"`
	Weight       *Weight      `json:"Weight,omitempty"`
}

// Weight represents tray weight difference
type Weight struct {
	// Add weight-related fields as needed
}

// RecognitionResponse represents video recognition response
type RecognitionResponse struct {
	Code int `json:"Code"`
	Data struct {
		YsTaskId               string `json:"YsTaskId"`
		RemainingServiceNumber int    `json:"RemainingServiceNumber"`
	} `json:"Data"`
	Msg string `json:"Msg,omitempty"`
}

// RecognitionCallback represents recognition callback data
type RecognitionCallback struct {
	State        RecognitionState        `json:"State"`
	TaskId       string                  `json:"TaskId"`
	YsTaskId     string                  `json:"YsTaskId"`
	ResultStatus RecognitionResultStatus `json:"ResultStatus"`
	ResultData   []RecognitionResultItem `json:"ResultData"`
	VideoUrl     string                  `json:"VideoUrl"`
}

// RecognitionResultStatus represents recognition result status
type RecognitionResultStatus struct {
	Code RecognitionResultCode `json:"Code"`
	Desc string                `json:"Desc"`
}

// RecognitionResultItem represents individual recognition result item
type RecognitionResultItem struct {
	CommoditySku string `json:"CommoditySku"`
	Qty          int    `json:"Qty"`
}

// RecognitionResultQueryRequest represents recognition result query request
type RecognitionResultQueryRequest struct {
	AppID  string `json:"AppID"`
	TaskId string `json:"TaskId"`
}

// RecognitionResultQueryResponse represents recognition result query response
type RecognitionResultQueryResponse struct {
	Code int                 `json:"Code"`
	Data RecognitionCallback `json:"Data"`
	Msg  string              `json:"Msg,omitempty"`
}

// RemainingQuotaResponse represents remaining recognition quota response
type RemainingQuotaResponse struct {
	Code int `json:"Code"`
	Data struct {
		Number int `json:"Number"`
	} `json:"Data"`
	Msg string `json:"Msg,omitempty"`
}

// Constants for recognition states
const (
	RecognitionStatePending   = 0
	RecognitionStateReviewing = 1
	RecognitionStateApproved  = 2
	RecognitionStateRejected  = 3
)

// RecognitionState represents recognition state
type RecognitionState int

// Constants for recognition result status codes
const (
	RecognitionResultSuccess = 2
)

// RecognitionResultCode represents recognition result code
type RecognitionResultCode int

// Constants for resource types
const (
	ResourceTypeURL = 1
	ResourceTypeID  = 2
)

// ResourceType represents resource type
type ResourceType int

// Constants for callback actions
const (
	CallbackActionPreOpenDoor     = "PreOpenDoor"
	CallbackActionOpenedDoor      = "OpenedDoor"
	CallbackActionCloseDoor       = "CloseDoor"
	CallbackActionCancel          = "Cancel"
	CallbackActionOrderDetected   = "OrderDetected"
	CallbackActionOrderSettlement = "OrderSettlement"
	CallbackActionOrderAdjustment = "OrderAdjustment"
	CallbackActionOrderRefund     = "OrderRefund"
	CallbackActionProductReview   = "ProductReview"
	CallbackActionRecognition     = "Recognition"
)

// CallbackEvent represents callback event structure
type CallbackEvent struct {
	OrderID         string      `json:"OrderId"`
	OrderNo         string      `json:"OrderNo"`
	TranseType      TranseType  `json:"TranseType"`
	Action          string      `json:"Action"`
	Status          bool        `json:"Status"`
	Msg             string      `json:"Msg"`
	CustomerDetails string      `json:"CustomerDetails"`
	Data            interface{} `json:"Data"`
}

// OrderDetectedCallback represents OrderDetected callback event
type OrderDetectedCallback struct {
	OrderID    string            `json:"OrderId"`
	OrderNo    string            `json:"OrderNo"`
	TranseType TranseType        `json:"TranseType"`
	OrgID      string            `json:"OrgId"`
	Action     string            `json:"Action"`
	Status     bool              `json:"Status"`
	Msg        string            `json:"Msg"`
	Data       DetectOrderDetail `json:"Data"`
}

// OpenedDoorCallback represents OpenedDoor callback event
type OpenedDoorCallback struct {
	OrderID         string     `json:"OrderId"`
	OrderNo         string     `json:"OrderNo"`
	TranseType      TranseType `json:"TranseType"`
	Action          string     `json:"Action"`
	Status          bool       `json:"Status"`
	Msg             string     `json:"Msg"`
	CustomerDetails string     `json:"CustomerDetails"`
	Data            string     `json:"Data"`
}

// PreOpenDoorCallback represents PreOpenDoor callback event
type PreOpenDoorCallback struct {
	OrderID         string     `json:"OrderId"`
	OrderNo         string     `json:"OrderNo"`
	TranseType      TranseType `json:"TranseType"`
	Action          string     `json:"Action"`
	Status          bool       `json:"Status"`
	Msg             string     `json:"Msg"`
	CustomerDetails string     `json:"CustomerDetails"`
	Data            string     `json:"Data"`
}

// CloseDoorCallback represents CloseDoor callback event
type CloseDoorCallback struct {
	OrderID         string     `json:"OrderId"`
	OrderNo         string     `json:"OrderNo"`
	TranseType      TranseType `json:"TranseType"`
	Action          string     `json:"Action"`
	Status          bool       `json:"Status"`
	Msg             string     `json:"Msg"`
	CustomerDetails string     `json:"CustomerDetails"`
	Data            string     `json:"Data"`
}

// CancelCallback represents Cancel callback event
type CancelCallback struct {
	OrderID         string     `json:"OrderId"`
	OrderNo         string     `json:"OrderNo"`
	TranseType      TranseType `json:"TranseType"`
	Action          string     `json:"Action"`
	Status          bool       `json:"Status"`
	Msg             string     `json:"Msg"`
	CustomerDetails string     `json:"CustomerDetails"`
	Data            string     `json:"Data"`
}

// OrderSettlementCallback represents OrderSettlement callback event
type OrderSettlementCallback struct {
	OrderID         string            `json:"OrderId"`
	OrderNo         string            `json:"OrderNo"`
	TranseType      TranseType        `json:"TranseType"`
	Action          string            `json:"Action"`
	Status          bool              `json:"Status"`
	Msg             string            `json:"Msg"`
	CustomerDetails string            `json:"CustomerDetails"`
	Data            DetectOrderDetail `json:"Data"`
}

// OrderAdjustmentCallback represents OrderAdjustment callback event
type OrderAdjustmentCallback struct {
	OrderID         string          `json:"OrderId"`
	OrderNo         string          `json:"OrderNo"`
	TranseType      TranseType      `json:"TranseType"`
	Action          string          `json:"Action"`
	Status          bool            `json:"Status"`
	Msg             string          `json:"Msg"`
	CustomerDetails string          `json:"CustomerDetails"`
	Data            OrderAdjustment `json:"Data"`
}

// OrderRefundCallback represents OrderRefund callback event
type OrderRefundCallback struct {
	OrderID         string      `json:"OrderId"`
	OrderNo         string      `json:"OrderNo"`
	TranseType      TranseType  `json:"TranseType"`
	Action          string      `json:"Action"`
	Status          bool        `json:"Status"`
	Msg             string      `json:"Msg"`
	CustomerDetails string      `json:"CustomerDetails"`
	Data            OrderRefund `json:"Data"`
}

// ProductReviewNotificationCallback represents product review notification callback
type ProductReviewNotificationCallback struct {
	AppID   string           `json:"AppID"`
	YsSkuId string           `json:"YsSkuId"`
	State   RecognitionState `json:"State"`
	Desc    string           `json:"Desc,omitempty"`
}

// DetectOrderDetail represents order detection details
type DetectOrderDetail struct {
	TradeProcessMode   TradeProcessMode    `json:"TradeProcessMode"`
	TradeProductModels []TradeProductModel `json:"TradeProductModels"`
}

// TradeProductModel represents product model in detection
type TradeProductModel struct {
	ID         string    `json:"Id"`
	AlisName   string    `json:"AlisName"`
	BuyCount   int       `json:"BuyCount"`
	MeterType  MeterType `json:"MeterType"`
	Price      float64   `json:"Price"`
	PictureURL string    `json:"PictureUrl,omitempty"`
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
