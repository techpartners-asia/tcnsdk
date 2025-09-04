package tcnsdk

type ListSlotRequest struct {
	VendID string `json:"vendId"`
}
type ListSlotResponse struct {
	BaseResponse
	Data struct {
		Counter []Counter `json:"counter"`
	} `json:"data"`
}

type Counter struct {
	DoorNo            int     `json:"doorNo"`
	DoorName          string  `json:"doorName"`
	CommodityCategory int     `json:"commodityCategory"`
	Layer             []Layer `json:"layer"`
}

type Layer struct {
	LayerNo       int         `json:"layerNo"`
	LayerName     string      `json:"layerName"`
	ErrorCode     int         `json:"errorCode"`
	CommodityInfo []Commodity `json:"commodityInfo"`
}

type AddProductToMachineRequest struct {
	VendID          string  `json:"vendId"`
	CommodityID     string  `json:"commodityId"`
	LayerNo         int     `json:"layerNo"`
	DoorNo          int     `json:"doorNo"`
	Capacity        int     `json:"capacity"`
	EarlyRate       bool    `json:"earlyrate"`
	EarlyWarigCount int     `json:"earlywarigcount"`
	Price           float64 `json:"price"`
}

type AddProductToMachineResponse struct {
	BaseResponse
}

type DeleteProductFromMachineRequest struct {
	VendInventedSlotIds []int  `json:"vendInventedSlotIds"`
	VendID              string `json:"vendId"`
}
type DeleteProductFromMachineResponse struct {
	BaseResponse
}

type UpdateProductOnMachineRequest struct {
	InvSlot         []InvSlot `json:"invslot"`
	VendID          string    `json:"vendId"`
	DoorNo          int       `json:"doorNo"`
	Price           float64   `json:"price"`
	EarlyWaringRate bool      `json:"earlyrate"`
	EarlyWarigCount int       `json:"earlywarigcount"`
	Capacity        int       `json:"capacity"`
	LayerNo         int       `json:"layerNo"`
}

type InvSlot struct {
	VendInventedSlotId int `json:"vendInventedSlotId"`
}
type UpdateProductOnMachineResponse struct {
	BaseResponse
}

type ProductUpdateRequest struct {
	ID             string  `json:"id"`
	Name           string  `json:"name"`
	PictureUrl     string  `json:"pictureUrl"`
	Weight         float64 `json:"weight"`
	Price          float64 `json:"price"`
	CostPrice      float64 `json:"costPrice"`
	SyncAllMachine bool    `json:"syncAllMachine"`
}

type ProductUpdateResponse struct {
	BaseResponse
}

type ProductListRequest struct {
	PageIndex     uint   `json:"pageIndex"`
	PageSize      uint   `json:"pageSize"`
	Name          string `json:"name"`
	CommodityType string `json:"commodityType"`
	SupportType   string `json:"supportType"`
}

type ProductListResponse struct {
	BaseResponse
	Data struct {
		Total     uint      `json:"total"`
		PageSize  uint      `json:"pageSize"`
		PageIndex uint      `json:"pageIndex"`
		Message   string    `json:"message"`
		Items     []Product `json:"items"`
	} `json:"data"`
}

type Product struct {
	ID                string    `json:"id"`
	ExtCommodityID    string    `json:"extCommodityId"`
	Sku               string    `json:"sku"`
	PictureUrl        string    `json:"pictureUrl"`
	Name              string    `json:"name"`
	IsOnShelve        bool      `json:"isOnShelve"`
	Packing           string    `json:"packing"`
	MeterType         MeterType `json:"meterType"`
	MeterTypeName     string    `json:"meterTypeName"`
	MeterUnit         string    `json:"meterUnit"`
	Weight            float64   `json:"weight"`
	Price             float64   `json:"price"`
	CostPrice         float64   `json:"costPrice"`
	CustomerCostPrice float64   `json:"customerCostPrice"`
	Types             []string  `json:"types"`
	SupportType       []string  `json:"supportType"`
	WeightFloatValue  float64   `json:"weightFloatValue"`
	SupportTypeName   string    `json:"supportTypeName"`
}

type ListProductTrainRequest struct {
	PageIndex   uint   `json:"pageIndex"`
	PageSize    uint   `json:"pageSize"`
	Sku         string `json:"sku"`
	RecordID    string `json:"recordId"`
	CommodityID string `json:"commodityId"`
	State       State  `json:"state"`
}

type ListProductTrainResponse struct {
	BaseResponse
	Data struct {
		Total     uint           `json:"total"`
		PageSize  uint           `json:"pageSize"`
		PageIndex uint           `json:"pageIndex"`
		Message   string         `json:"message"`
		Items     []ProductTrain `json:"items"`
	} `json:"data"`
}

type ProductTrain struct {
	Product
	State           State    `json:"state"`
	RejectCause     string   `json:"rejectCause"`
	OrganizeID      string   `json:"organizeId"`
	QualityID       string   `json:"qualityId"`
	QualityPeroid   int      `json:"qualityPeroid"`
	ModifyUserName  string   `json:"modifyUserName"`
	IsPovertyRelief bool     `json:"isPovertyRelief"`
	Nature          string   `json:"nature"`
	Tags            []string `json:"tags"`
	ApplicationType int      `json:"applicationType"`
}

type State int

const (
	StatePending  State = 1
	StateApproved State = 2
	StateRejected State = 3
)

type ProductTrainRequest struct {
	Sku               string  `json:"sku"`
	Name              string  `json:"name"`
	Price             float64 `json:"price"`
	PictureURL        string  `json:"pictureUrl"`
	BackImage         string  `json:"backImage"`
	ProfileImage      string  `json:"profileImage"`
	TopOrBottomImage  string  `json:"topOrBottomImage"`
	PictureUrl_Bottom string  `json:"pictureUrl_bottom"`
	PictureUrl_Other  string  `json:"pictureUrl_other"`
	Weight            float64 `json:"weight"`
	Packing           string  `json:"packing"`
}

type ProductTrainResponse struct {
	BaseResponse
	Data string `json:"data"` //if its not empty, it means the product is undergoing review, otherwise it's already exists
}

type TrainResultListRequest struct {
	PageIndex   uint   `json:"pageIndex"`
	PageSize    uint   `json:"pageSize"`
	RecordID    string `json:"recordId"`
	CommodityID string `json:"commodityId"`
	Sku         string `json:"sku"`
	State       State  `json:"state"`
}

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
	Sku                string  `json:"sku"`
	CommodityID        string  `json:"commodityId"`
	CommodityName      string  `json:"commodityName"`
	CommoditySKU       string  `json:"commoditySKU"`
	Weight             int     `json:"weight"`
	Price              float64 `json:"price"`
	DoorNo             int     `json:"doorNo"`
	PictureURL         string  `json:"pictureUrl"`
	LayerNo            int     `json:"layerNo"`
	InventedSlotNo     int     `json:"inventedSlotNo"`
	VendInventedSlotId int     `json:"vendInventedSlotId"`
	SlotId             string  `json:"slotId"`
	IsLack             bool    `json:"isLack"`
	IsEmpty            bool    `json:"isEmpty"`
	Stock              int     `json:"stock"`
	Capacity           int     `json:"capacity"`
	EarlyWaringRate    bool    `json:"earlyWaringRate"`
	EarlyWarigCount    int     `json:"earlyWarigCount"`
	ReplenishedSales   int     `json:"replenishedSales"`
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

// TranseType represents transaction type
type TranseType int

// Constants for transaction types
const (
	TranseTypePurchase TranseType = 0
	TranseTypeRestock  TranseType = 2
)

// Constants for payment types
const (
	PayTypeNormal    PayType = 0
	PayTypeDeduction PayType = 1
	PayTypeRefund    PayType = 2
)

// PayType represents payment type
type PayType int

// Constants for payment status
const (
	PayStatusSuccess PayStatus = 1
	PayStatusFailed  PayStatus = 2
)

// PayStatus represents payment status
type PayStatus int

// Constants for trade process modes
const (
	TradeProcessModeNormal    TradeProcessMode = 0 // normal transaction will return identification results
	TradeProcessModeCancel    TradeProcessMode = 1 // cancels the transaction
	TradeProcessModeInterrupt TradeProcessMode = 2 // interrupts the transaction and leaves it to the merchant to identify it
)

// *when TradeProcessMode=2*
// exception types:
// 1. Block the camera															| Unfriendly shopping
// 2. Replace the goods in the cabinet											| Unfriendly shopping
// 3. Insert foreign object														| Unfriendly shopping
// 4. Malicious destruction of goods											| Unfriendly shopping
// 5. Covering the product														| Unfriendly shopping
// 6. Long shopping time														| Unfriendly shopping
// 7. Suspected merchant replenishment											| Unfriendly shopping
// 8. Other abnormal behaviors													| Unfriendly shopping
// 9. The video is incomplete or the screen is distorted						| Alarm
// 10. The product in the device is not replenished								| Alarm
// 11. Products are too similar													| Alarm
// 12. The video is too long after closing										| Alarm
// 13. Power outage during purchase												| Alarm

// identify exceptions when the following occurs ( ****Status=False, Msg=reason of exception ****)
// 1. Gravity cabinet recogntion error, gravity data is not uploaded 								| Gravity Cabinet
// 2. Unable to obtain sensor data 																	| Gravity Cabinet
// 3. Abnormal overweight 																			| Gravity Cabinet
// 4. Open door product snapshot query failed 														| Gravity Cabinet
// 5. The number of products in the mirror table is less than the number of identified products 	| Gravity Cabinet
// 6. Failed to construct the identification request object											| Visual Cabinet
// 7. Failed to obtain the identification request object											| Visual Cabinet
// 8. Failed to retrieve video																		| Visual Cabinet
// 9. No video																						| Visual Cabinet
// 10. Order item mirror table query failed															| Visual Cabinet
// 11. Failed to identify the platform																| Visual Cabinet

// TradeProcessMode represents trade process mode
type TradeProcessMode int

// Constants for meter types
const (
	MeterTypePiece  MeterType = 1
	MeterTypeWeight MeterType = 2
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
	RecognitionStatePending   RecognitionState = 0
	RecognitionStateReviewing RecognitionState = 1
	RecognitionStateApproved  RecognitionState = 2
	RecognitionStateRejected  RecognitionState = 3
)

// RecognitionState represents recognition state
type RecognitionState int

// Constants for recognition result status codes
const (
	RecognitionResultSuccess RecognitionResultCode = 2
)

// RecognitionResultCode represents recognition result code
type RecognitionResultCode int

// Constants for resource types
const (
	ResourceTypeURL ResourceType = 1
	ResourceTypeID  ResourceType = 2
)

// ResourceType represents resource type
type ResourceType int

// CallbackAction represents callback action
type CallbackAction string

// Constants for callback actions
const (
	CallbackActionPreOpenDoor     CallbackAction = "PreOpenDoor"
	CallbackActionOpenedDoor      CallbackAction = "OpenedDoor"
	CallbackActionCloseDoor       CallbackAction = "CloseDoor"
	CallbackActionCancel          CallbackAction = "Cancel"
	CallbackActionOrderDetected   CallbackAction = "OrderDetected"
	CallbackActionOrderSettlement CallbackAction = "OrderSettlement"
	CallbackActionOrderAdjustment CallbackAction = "OrderAdjustment"
	CallbackActionOrderRefund     CallbackAction = "OrderRefund"
	CallbackActionProductReview   CallbackAction = "ProductReview"
	CallbackActionRecognition     CallbackAction = "Recognition"
)

// CallbackEvent represents callback event structure
type CallbackEvent struct {
	OrderID         string         `json:"OrderId"`
	OrderNo         string         `json:"OrderNo"`
	TranseType      TranseType     `json:"TranseType"`
	Action          CallbackAction `json:"Action"`
	Status          bool           `json:"Status"`
	Msg             string         `json:"Msg"`
	CustomerDetails string         `json:"CustomerDetails"`
	Data            interface{}    `json:"Data"`
}

// OrderDetectedCallback represents OrderDetected callback event
type OrderDetectedCallback struct {
	OrderID    string            `json:"OrderId"`
	OrderNo    string            `json:"OrderNo"`
	TranseType TranseType        `json:"TranseType"`
	OrgID      string            `json:"OrgId"`
	Action     CallbackAction    `json:"Action"`
	Status     bool              `json:"Status"`
	Msg        string            `json:"Msg"`
	Data       DetectOrderDetail `json:"Data"`
}

// OpenedDoorCallback represents OpenedDoor callback event
type OpenedDoorCallback struct {
	OrderID         string         `json:"OrderId"`
	OrderNo         string         `json:"OrderNo"`
	TranseType      TranseType     `json:"TranseType"`
	Action          CallbackAction `json:"Action"`
	Status          bool           `json:"Status"`
	Msg             string         `json:"Msg"`
	CustomerDetails string         `json:"CustomerDetails"`
	Data            string         `json:"Data"`
}

// PreOpenDoorCallback represents PreOpenDoor callback event
type PreOpenDoorCallback struct {
	OrderID         string         `json:"OrderId"`
	OrderNo         string         `json:"OrderNo"`
	TranseType      TranseType     `json:"TranseType"`
	Action          CallbackAction `json:"Action"`
	Status          bool           `json:"Status"`
	Msg             string         `json:"Msg"`
	CustomerDetails string         `json:"CustomerDetails"`
	Data            string         `json:"Data"`
}

// CloseDoorCallback represents CloseDoor callback event
type CloseDoorCallback struct {
	OrderID         string         `json:"OrderId"`
	OrderNo         string         `json:"OrderNo"`
	TranseType      TranseType     `json:"TranseType"`
	Action          CallbackAction `json:"Action"`
	Status          bool           `json:"Status"`
	Msg             string         `json:"Msg"`
	CustomerDetails string         `json:"CustomerDetails"`
	Data            string         `json:"Data"`
}

// this event is triggered when the machine fails to open the door or reports an exception when closing the door.
type CancelCallback struct {
	OrderID         string         `json:"OrderId"`
	OrderNo         string         `json:"OrderNo"`
	TranseType      TranseType     `json:"TranseType"`
	Action          CallbackAction `json:"Action"`
	Status          bool           `json:"Status"`
	Msg             string         `json:"Msg"`
	CustomerDetails string         `json:"CustomerDetails"`
	Data            string         `json:"Data"`
}

// OrderSettlementCallback represents OrderSettlement callback event
type OrderSettlementCallback struct {
	OrderID         string            `json:"OrderId"`
	OrderNo         string            `json:"OrderNo"`
	TranseType      TranseType        `json:"TranseType"`
	Action          CallbackAction    `json:"Action"`
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
	Action          CallbackAction  `json:"Action"`
	Status          bool            `json:"Status"`
	Msg             string          `json:"Msg"`
	CustomerDetails string          `json:"CustomerDetails"`
	Data            OrderAdjustment `json:"Data"`
}

// OrderRefundCallback represents OrderRefund callback event
type OrderRefundCallback struct {
	OrderID         string         `json:"OrderId"`
	OrderNo         string         `json:"OrderNo"`
	TranseType      TranseType     `json:"TranseType"`
	Action          CallbackAction `json:"Action"`
	Status          bool           `json:"Status"`
	Msg             string         `json:"Msg"`
	CustomerDetails string         `json:"CustomerDetails"`
	Data            OrderRefund    `json:"Data"`
}

// ProductReviewNotificationCallback represents product review notification callback
type ProductReviewNotificationCallback struct {
	AppID   string           `json:"AppID"`
	YsSkuId string           `json:"YsSkuId"`
	State   RecognitionState `json:"State"`
	Action  CallbackAction   `json:"Action"`
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

// OpenDoorRequest represents door opening request
type OpenDoorRequest struct {
	OrderID         string     `json:"orderId"`
	MachineID       string     `json:"machineId"`
	DoorNo          int        `json:"doorNo"`
	TranseType      TranseType `json:"transeType"`
	CustomerDetails string     `json:"customerDetails,omitempty"`
	TimeSp          int64      `json:"timeSp"`
	NotifyURL       string     `json:"NotifyUrl"`
	Remark          string     `json:"remark,omitempty"`
}

// OpenDoorResponse represents door opening response
type OpenDoorResponse struct {
	OrderID    string         `json:"OrderId"`
	TranseType TranseType     `json:"TranseType"`
	Action     CallbackAction `json:"Action"`
	Status     bool           `json:"Status"`
	Msg        string         `json:"Msg"`
	Data       string         `json:"Data"`
}

// RestockOpenDoorRequest represents restock door opening request
type RestockOpenDoorRequest struct {
	OrderID         string     `json:"orderId"`
	MachineID       string     `json:"machineId"`
	DoorNo          int        `json:"doorNo"`
	TranseType      TranseType `json:"transeType"`
	CustomerDetails string     `json:"customerDetails,omitempty"`
	TimeSp          int64      `json:"timeSp"`
	NotifyURL       string     `json:"NotifyUrl"`
	Remark          string     `json:"remark,omitempty"`
}

// RestockOpenDoorResponse represents restock door opening response
type RestockOpenDoorResponse struct {
	OrderID    string     `json:"OrderId"`
	TranseType TranseType `json:"TranseType"`
	Action     string     `json:"Action"`
	Status     bool       `json:"Status"`
	Msg        string     `json:"Msg"`
	Data       string     `json:"Data"`
}
