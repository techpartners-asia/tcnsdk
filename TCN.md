
**Visual Recognition SDK**  
**Product Description**  
**Product Introduction**  
Yunshu provides an external product recognition algorithm service.

**Instructions**  
Before integration, please contact customer service to apply for API access. Customer service will generate the AppID, Secret, and Key.  
Before actual integration, please read the following documentation carefully.

Each API has a call frequency limit: **100 QPS** (queries per second).

**Process Description**


**Start**

⬇️  
**Query Cloud Product Inventory** _(Sync Products)_  
  ↳ _No product returned_ → **Contact customer service to add new product**  
⬇️  
**Request Recognition**  
  ↳ → **Query Recognition Result**  
⬇️  
**Recognition Callback**  
⬇️  
**Return Result**  
⬇️  
**End**

## API Rules

### Request Headers

```
Connection: keep-alive
Content-Type: application/json;charset=utf-8
Referer: https://openapi1.ourvend.com
Authorization: Bearer + xxxxxx(token)
Sign: xxxxx
```

### Signature Rules

1. All parameters are sorted in ASCII order.
    
2. Join parameters as a string using the format: `key1=value1&key2=value2`
    
3. Append the secret key at the end: `key1=value1&key2=value2&key=xxxxxxxxxxx`. Array parameters are excluded.
    
4. Apply MD5 hash to the string and convert it to uppercase.
    

### Response Parameters

|Field|Type|Description|
|---|---|---|
|Code|int|Status code, 200 means success|
|Data|object|All data results will be in this field|
|Msg|string|Only returned when request fails|

### Status Codes

|Code|Description|
|---|---|
|200|Success|
|402|Parameter error|
|418|Signature error|
|420|Insufficient remaining service uses|

- **Test Environment:** `openapi1.ourvend.com`
    
- **Production Environment:** `openapi.ourvend.com`
    

---

## API Calls

### Get TOKEN

- **Description:** Retrieve API secret key, valid for 2 hours
    
- **Method:** POST
    
- **URL:** `https://openapi1.ourvend.com/OpenApi/Login`
    

#### Request Body (JSON):

|Parameter|Type|Required|Description|
|---|---|---|---|
|AppID|string|Yes|Generated AppID (unique merchant ID)|
|Key|string|Yes|Generated Key|
|Secret|string|Yes|Generated Secret (used for signing)|

#### Response:

```json
{
  "Code": 200,
  "Data": {
    "Token": "f7b82ca1ebc7f9838d0ac399149d6577",
    "ExpiresIn": 72000
  }
}
```

#### Error Example:

```json
{
  "Code": 402,
  "Msg": "Invalid AppID"
}
```

---

## Cloud Product Query

- **Description:** Actively query product info
    
- **Method:** POST
    
- **URL:** `https://openapi1.ourvend.com/OpenApi/CommodityConfirm`
    

#### Request Body (JSON):

|Parameter|Type|Required|Description|
|---|---|---|---|
|AppID|string|Yes|Merchant unique ID|
|CommoditySku|string|Yes|Product SKU code (EAN-13)|

#### Response:

```json
{
  "Code": 200,
  "Data": {
    "CommodityName": "Non-fried Cucumber Chips",
    "CommodityId": "6901028185905",
    "PictureUrl": "https://ourvendv3.oss-cn-qingdao.aliyuncs.com/ProductImage/example.jpg"
  }
}
```

---

## Product Registration

- **Method:** POST
    
- **URL:** `https://openapi1.ourvend.com/OpenApi/Commodity/Apply`
    

#### Request Parameters (JSON):

|Parameter|Type|Required|Description|
|---|---|---|---|
|AppID|string|Yes|Merchant ID|
|AskId|string|Yes|Application ID|
|SkuName|string|Yes|SKU Name|
|Sku|string|Yes*|SKU Code (mandatory if `IsStandard=true`)|
|IsStandard|bool|No|Whether standard SKU|
|ImgUrls|Array|Yes|1-9 white background image URLs|
|NotifyUrl|string|Yes|Callback URL (currently not supported)|

#### Response:

|Parameter|Type|Description|
|---|---|---|
|YsSkuId|string|Product ID used for recognition later|

---

## Product Review Query

- **Method:** POST
    
- **URL:** `https://openapi1.ourvend.com/OpenApi/Commodity/ApplyQuery`
    

#### Parameters:

|Parameter|Type|Required|Description|
|---|---|---|---|
|AppID|string|Yes|Merchant ID|
|YsSkuId|string|Yes|Product ID|

#### Response:

|Parameter|Type|Description|
|---|---|---|
|State|int|0: Pending, 1: Reviewing, 2: Approved, 3: Rejected|
|Desc|string|Reason (only for rejected submissions)|

---

## Product Review Notification

- **Method:** POST
    
- **URL:** Callback URL
    

#### Parameters:

|Parameter|Type|Required|Description|
|---|---|---|---|
|AppID|string|Yes|Merchant ID|
|YsSkuId|string|Yes|Product ID|
|State|int|Yes|0: Pending, 1: Reviewing, 2: Approved, 3: Rejected|
|Desc|string|No|Rejection reason (if any)|

#### Notification Result:

- Success: HTTP 200 or 204
    
- Failure: HTTP 4XX or 5XX with:
    

```json
{
  "code": "ERROR_CODE",
  "message": "Reason for error"
}
```

---

## Recognition Request

- **Description:** Submit video for recognition (must contain registered products)
    
- **Method:** POST
    
- **URL:** `https://openapi1.ourvend.com/OpenApi/PushVideoAsk`
    

#### Parameters:

|Name|Type|Required|Description|
|---|---|---|---|
|AppID|string|Yes|Merchant ID|
|TaskId|string|Yes|Unique task ID|
|ResourceType|int|No|1: Video URL, 2: Video resource ID|
|ResourceUrl|Array|Yes|List of accessible video URLs (MP4/AVI/WebM, H.264)|
|ProductRange|Array|Yes|List of YsSkuId (up to 40)|
|NotifyUrl|string|Yes|Callback URL|
|Weight|JSON Object|No|Tray weight difference (if using weight)|

#### Response:

```json
{
  "Code": 200,
  "Data": {
    "YsTaskId": "1273922496732089",
    "RemainingServiceNumber": 99
  }
}
```

---

## Recognition Callback

- **Method:** POST
    
- **URL:** NotifyUrl
    

#### Parameters:

```json
{
  "State": 1,
  "TaskId": "121212121",
  "YsTaskId": "...",
  "ResultStatus": {
    "Code": 2,
    "Desc": "Recognition success"
  },
  "ResultData": [
    {
      "CommoditySku": "6901028185905",
      "Qty": 1
    }
  ],
  "VideoUrl": "url1,url2"
}
```

---

## Recognition Result Query

- **Method:** POST
    
- **URL:** `https://openapi1.ourvend.com/OpenApi/VideoAskResult`
    

#### Parameters:

|Parameter|Type|Required|Description|
|---|---|---|---|
|AppID|string|Yes|Merchant ID|
|TaskId|string|Yes|Your task ID|

---

## Remaining Recognition Quota

- **Method:** POST
    
- **URL:** `https://openapi1.ourvend.com/OpenApi/GetVideoAskServiceNumber/{AppID}`
    

#### Response:

```json
{
  "Code": 200,
  "Data": {
    "Number": 99
  }
}
```