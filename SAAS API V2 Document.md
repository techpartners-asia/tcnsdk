# SAAS 交易相关接口文档V2

# []{#anchor}SAAS transaction related interface documentation V2

# []{#anchor-1}API interface documentation

*Table of contents*

*API interface documentation*

*1. Authentication (Post)*

*2. Open the door (Post)*

*3. Payment callback interface (Post)*

*4. Common return value description*

## []{#anchor-2}Test environment address

*https://openapi1.ourvend.com*

## []{#anchor-3}Formal environment address

*https://openapi.aivendortech.com*

## []{#anchor-4}1. Authentication

***URL **: https://openapi1.ourvend.com/OpenApi/Login​​*

***Method **: POST*

***parameter **:*

  ------------ ------------ ----------------------------------
  ***name***   ***type***   ***describe***
  *appID*      *string*     **
  *key*        *string*     **
  *secret*     *string*     **
  ------------ ------------ ----------------------------------

***Return value **:*

  ------------- ------------ ------------------------
  ***name***    ***type***   ***describe***
  *Token*       *string*     *Access Token*
  *ExpiresIn*   *int*        *Valid time (seconds)*
  ------------- ------------ ------------------------

### []{#anchor-5}Request Body (JSON)

{\
\"appID\": \"xxxxxxxxxxxxxxxxxx\",\
\"key\": \"xxxxxxxxxxxxxxxxxxxxxxxx\",\
\"secret\": \"xxxxxxxxxxxxxxxxxxxxxxxxxxx\"\
}\

## []{#anchor-6}2. Get device information

***URL: **https://openapi1.ourvend.com/ OpenApi/Machine/Info/{mid}*

***Method: Get***

https://openapi1.ourvend.com/OpenApi/Machine/Info/2412090001\

{\
\"statusCode\": 200,\
\"data\": {\
\"machineId\": \"2412090001\",\
\"machineName\": \"2412090001\",\
\"signal\": 3, // Signal strength 0\~5\
\"temperature\": \"24\", // Temperature\
\"runningState\": 1 // Whether to lock the machine\
},\
\"succeeded\": true,\
\"errors\": null,\
\"extras\": null,\
\"timestamp\": 1743496289599\
}\

## []{#anchor-7}3. Get equipment product information

***URL: **https://openapi1.ourvend.com/ OpenApi/Machine /
Commoditys/{mid}*

***Method: Get***

https://openapi1.ourvend.com/OpenApi/Machine/Commoditys/2412090001\

{\
\"statusCode\": 200,\
\"data\": \[\
{\
\"commodityId\": \"2164367863923397\",\
\"commodityName\": \"Wahaha Mineral Water 596ml\",\
\"commoditySKU\": \"69353535\",\
\"weight\": 596,\
\"price\": 3.00,\
\"doorNo\": 0,\
\"pictureUrl\":
\"https://ourvendv3.oss-cn-qingdao.aliyuncs.com/ProductImage/9080b157a4b54c63a300907f138a613a.jpg\"\
},\
{\
\"commodityId\": \"2180392082047685\",\
\"commodityName\": \"Oriental Leaves 300ml\",\
\"commoditySKU\": \"69442310\",\
\"weight\": 300,\
\"price\": 3.00,\
\"doorNo\": 0,\
\"pictureUrl\":
\"https://ourvendv3.oss-cn-qingdao.aliyuncs.com/ProductImage/5b5e36a2de4f4ba38d2182982096373d.jpg\"\
}\
\],\
\"succeeded\": true,\
\"errors\": null,\
\"extras\": null,\
\"timestamp\": 1743496372493\
}\

## []{#anchor-8}4. Get device payment configuration information

***URL: **https://openapi1.ourvend.com/ OpenApi/Machine /
PayConfig/{mid}*

***Method: Get***

{\
\"statusCode\": 200,\
\"data\": {\
\"machineId\":\"2412090001\",\
\"preAuthorLimit\":\"50\", // Pre-authorization limit\
\"currency\":\"USD\", // Currency\
\"qrCodeUrl\":\"2412090001\", // Shopping QR code address\
},\
\"succeeded\": true,\
\"errors\": null,\
\"extras\": null,\
\"timestamp\": 1743496372493\
}\

## []{#anchor-9}5. Trading Opening

***URL **: https://openapi1.ourvend.com/ OpenApi/Order/OpenDoor*

***Method **: POST*

***parameter **:*

  ------------------- ------------ ----------------------------------------------------------
  ***name***          ***type***   ***describe***
  *OrderId*           *string*     *Order number (required)*
  *MachineId*         *string*     *Machine number (required)*
  *DoorNo*            *int*        *Door Number (required)*
  *TranseType*        *int*        *Door opening method (required) 0 Buy 2 Restock*
  *CustomerDetails*   *string*     *Custom message (consumer data)*
  *TimeSp*            *long*       *Timestamp*
  *NotifyUrl*         *string*     *Operation result callback URL (CallbackURL) (required)*
  *Remark*            *string*     *The maximum length of the comment is 128 characters.*
  ------------------- ------------ ----------------------------------------------------------

***Return value **:*

  -------------- ------------ ---------------------------------------------------------------------
  ***name***     ***type***   ***describe***
  *OrderId*      *string*     *Order Number*
  *TranseType*   *int*        *Opening method 0 Purchase 2 Restock*
  *Action*       *string*     *Action flags (PreOpenDoor, OpenedDoor, ClosedDoor, OrderDetected)*
  *Status*       *bool*       *Operational Status*
  *Msg*          *string*     *Operation Description*
  *Data*         *string*     *Results Data*
  -------------- ------------ ---------------------------------------------------------------------

### []{#anchor-10}Request Body (JSON)

{\
\"orderId\": \"xxxxxxxxxxxxxxxx\",\
\"machineId\": \"xxxxxxxxxxxxxx\",\
\"doorNo\": 1,\
\"transeType\":1,\
\"customerDetails\": \"1\",\
\"timeSp\": 1741758180,\
\"NotifyUrl\":\"xxxxxxxxxxxxxxx\"\
}\

## []{#anchor-11}6. Restock and open the door ( Testing )

***URL **: https://openapi1.ourvend.com/OpenApi /Repli/OpenDoorMethod
**: **POST*

## []{#anchor-12}7. Confirm replenishment data ( Testing )

***URL **: https://openapi1.ourvend.com/OpenApi /Repli/Confirm **Method
**: POST **Parameters **:*

  -------------- ---------------------------- ------------ -------------------------------------
  *name*         *type*                       *Required*   *describe*
  *Mid*          *string*                     *yes*        *Merchant ID*
  *TransId*      *long*                       *yes*        *Transaction ID*
  *TrackingId*   *string*                     *yes*        *Tracking ID*
  *DoorNo*       *int*                        *yes*        *Door Number*
  *Data*         *List\<ConfirmRepliItem\>*   *yes*        *Confirmed replenishment data list*
  -------------- ---------------------------- ------------ -------------------------------------

#### []{#anchor-13}ConfirmRepliItem structure:

  --------------- ---------- ------------ --------------------------
  *name*          *type*     *Required*   *describe*
  *CommodityId*   *string*   *yes*        *Product ID*
  *LayerNo*       *int*      *yes*        *Product layer*
  *SpotCount*     *int*      *yes*        *Replenishment quantity*
  *StockCount*    *int*      *yes*        *Calibrated Inventory*
  *Sno*           *int*      *yes*        *Arrival Number*
  --------------- ---------- ------------ --------------------------

#### []{#anchor-14}Return parameter:

  -------------- ---------- --------------------
  *name*         *type*     *describe*
  *TrackingId*   *string*   *Tracking ID*
  *TransId*      *long*     *Transaction ID*
  *Success*      *bool*     *Success*
  *Message*      *string*   *Returned message*
  -------------- ---------- --------------------

#### []{#anchor-15}Example

##### []{#anchor-16}Request Example

{\
\"Mid\": \"xxxxxxxxxxx\",\
\"TransId\": 1234567890123,\
\"TrackingId\": \"xxxxxxxxxxxxxxx\",\
\"DoorNo\": 1,\
\"Data\": \[\
{\
\"CommodityId\": \"xxxxxxxxxxxx1\",\
\"LayerNo\": 2,\
\"SpotCount\": 1,\
\"StockCount\": 500,\
\"Sno\": 123\
},\
{\
\"CommodityId\": \"xxxxxxxxxxxxx\",\
\"LayerNo\": 3,\
\"SpotCount\": 2,\
\"StockCount\": 1000,\
\"Sno\": 124\
}\
\]\
}\

##### []{#anchor-17}Example Response

{\
\"TrackingId\": \"tracking-id-001\",\
\"TransId\": 1234567890123,\
\"Success\": true,\
\"Message\": \"Replenishment data confirmed successfully\"\
}\

## []{#anchor-18}10. Order transaction result reporting interface

***illustrate **:*

***URL **: https://openapi1.ourvend.com/
OpenApi/Order/ReportPaymentResult*

***Method **: POST*

***parameter **:*

  ---------------------- -------------- -------------------------------------------------------------------------------------------------------------
  ***name***             ***type***     ***describe***
  *OrderId*              *string*       *Order number (required), when PayType is refund, it is the third-party refund number*
  *OrderNo*              *string*       *SaasAPI order number (required), when PayType is Refund, it is SaasAPI refund order number Refund OrderNo*
  *PayType*              *int*          *Payment type 0: normal transaction, 1: deduction, 2: refund*
  *PayStatus*            *int*          *Payment status (required) 1 Success 2 Failed*
  *ErrorMessage*         *string*       *Payment failed message*
  *~~ProductDetails~~*   *~~string~~*   *~~Product Details~~*
  ---------------------- -------------- -------------------------------------------------------------------------------------------------------------

***Return value **:*

  ------------ ------------ -----------------------------
  ***name***   ***type***   ***describe***
  *data*       *bool*       *Is the report successful?*
  ------------ ------------ -----------------------------

### []{#anchor-19}Request Body (JSON)

{\
\"OrderId\": \"your OrderId\",\
\"OrderNo\": \"233355444\",\
\"PayType\":0,\
\"PayStatus\": 1,\
\"ErrorMessage\": \"\"\
}\

### []{#anchor-20}Response Example

{\
\"OrderId\": \"your OrderId\",\
\"Status\": true,\
\"Message\": \"Order processed successfully\"\
}\

## []{#anchor-21}11. CallBack Interface Sample Response

### []{#anchor-22}11.1 OpenDoor door opening event response

  ------------------- ------------ -----------------------------------------------------------------------
  ***name***          ***type***   ***describe***
  *OrderId*           *string*     *Order Number*
  *OrderNo*           *string*     *SaasAPI Order Number*
  *TranseType*        *int*        *Opening method 0 Purchase 2 Restock*
  *Action*            *string*     *OpenedDoor*
  *Status*            *bool*       *Operational Status*
  *Msg*               *string*     *Operation Description*
  *CustomerDetails*   *string*     *User-defined information, parameters passed before opening the door*
  *Data*              *string*     *Results Data*
  ------------------- ------------ -----------------------------------------------------------------------

### []{#anchor-23}11.2 CloseDoor event response

  ------------------- ------------ -----------------------------------------------------------------------
  ***name***          ***type***   ***describe***
  *OrderId*           *string*     *Order Number*
  *OrderNo*           *string*     *SaasAPI Order Number*
  *TranseType*        *int*        *Opening method 0 Purchase 2 Restock*
  *Action*            *string*     *CloseDoor*
  *Status*            *bool*       *Operational Status*
  *Msg*               *string*     *Operation Description*
  *CustomerDetails*   *string*     *User-defined information, parameters passed before opening the door*
  *Data*              *string*     *Results Data*
  ------------------- ------------ -----------------------------------------------------------------------

### []{#anchor-24}11.3 Cancel Order Cancellation Event Response

*Description: **This event is triggered when the machine fails to open
the door or reports an exception when closing the door.***

  ------------------- ------------ -----------------------------------------------------------------------
  ***name***          ***type***   ***describe***
  *OrderId*           *string*     *Order Number*
  *OrderNo*           *string*     *SaasAPI Order Number*
  *TranseType*        *int*        *Opening method 0 Purchase 2 Restock*
  *Action*            *string*     *Cancel*
  *Status*            *bool*       *Operational Status*
  *Msg*               *string*     *Operation Description*
  *CustomerDetails*   *string*     *User-defined information, parameters passed before opening the door*
  *Data*              *string*     *Results Data*
  ------------------- ------------ -----------------------------------------------------------------------

### []{#anchor-25}11.4 OrderDetected

*Description: **This event will be triggered when the order has entered
the AI recognition stage.***

[]{.image .placeholder original-image-src="image1.png"
original-image-title="" width="6.26875in" height="1.57971in"}

  -------------- ----------------------------- -----------------------------------------------
  ***name***     ***type***                    ***describe***
  *OrderId*      *string*                      *Order Number*
  *OrderNo*      *string*                      *SaasAPI Order Number*
  *TranseType*   *int*                         *Transaction type (0-order, 2-replenishment)*
  *OrgId*        *string*                      *Merchant Id*
  *Action*       *string*                      *Action Type*
  *Status*       *bool*                        *Transaction Status*
  *Msg*          *string*                      *Identify prompt content*
  *Data*         ***DetectOrderDetailJson***   *Identification results*
  -------------- ----------------------------- -----------------------------------------------

***DetectOrderDetailJson***

+--------------------------+------------+---------------------------+
| ***name***               | ***type*** | ***describe***            |
+--------------------------+------------+---------------------------+
| *TradeProcessMode*       | *int*      | *Transaction processing   |
|                          |            | method: (0=normal         |
|                          |            | transaction, 1=cancel     |
|                          |            | transaction, 2=interrupt  |
|                          |            | transaction)*             |
|                          |            |                           |
|                          |            | *when TradeProcessMode=0* |
|                          |            |                           |
|                          |            | *TradeProcessMode=1       |
|                          |            | cancels the transaction*  |
|                          |            |                           |
|                          |            | *TradeProcessMode=2       |
|                          |            | interrupts the            |
|                          |            | transaction and leaves it |
|                          |            | to the merchant to        |
|                          |            | identify it*              |
+--------------------------+------------+---------------------------+
| ***TradeProductModels*** |            | *Product Contents*        |
+--------------------------+------------+---------------------------+
| *Id*                     | *string*   | *Product ID*              |
+--------------------------+------------+---------------------------+
| *AlisName*               | *string*   | *Product Name*            |
+--------------------------+------------+---------------------------+
| *BuyCount*               | *int*      | *Purchase quantity*       |
+--------------------------+------------+---------------------------+
| *MeterType*              | *int*      | *Pricing method: (1=piece |
|                          |            | (number), 2=weight)*      |
+--------------------------+------------+---------------------------+
| *Price*                  | *money*    | *Unit price*              |
+--------------------------+------------+---------------------------+
| *PictureUrl*             | *string*   | *Product image URL*       |
+--------------------------+------------+---------------------------+

*Return status description:*

1.  ***Normal recognition ( ****Status=True ****)***

```{=html}
<!-- -->
```
1.  a.  ***Data.TradeProcessMode = 0 Normal transactions will return
        identification results***
    b.  ***TradeProcessMode=1 The recognition result is \[Cancel
        Transaction\], common reasons are:***

```{=html}
<!-- -->
```
1.  1.  i.  *Open but no shopping*

```{=html}
<!-- -->
```
1.  c.  ***TradeProcessMode=2 The recognition result is
        \[Merchant****'****s own judgment\], the common reasons are:***

  ------------------------------------------------------ ----------------------- ----------------
  ***Exception name***                                   ***type***              ***describe***
  *Block the camera*                                     *Unfriendly shopping*   
  *Replace the goods in the cabinet*                     *Unfriendly shopping*   
  *Insert foreign matter*                                *Unfriendly shopping*   
  *Malicious destruction of goods*                       *Unfriendly shopping*   
  *Cover the product*                                    *Unfriendly shopping*   
  *Long shopping time*                                   *Unfriendly shopping*   
  *Suspected merchant replenishment*                     *Unfriendly shopping*   
  *Other abnormal behaviors*                             *Unfriendly shopping*   
                                                                                 
  *The video is incomplete or the screen is distorted*   *Alerts*                
  *The product in the device is not replenished*         *Alerts*                
  *Products are too similar*                             *Alerts*                
  *The video is too long after closing*                  *Alerts*                
  *Power outage during purchase*                         *Alerts*                
                                                                                 
  ------------------------------------------------------ ----------------------- ----------------

2.  ***Identify exceptions when the following occurs ( ****Status=False,
    Msg=reason of exception ****)***

  ------------------------------------------------------------------------------------ ---------------------- --
  ***Identify abnormal prompts***                                                      ***Container Type***   
  *The gravity cabinet recognition is abnormal and the gravity data is not uploaded*   *Gravity Cabinet*      
  *Unable to obtain sensor data*                                                       *Gravity Cabinet*      
  *Abnormal overweight*                                                                *Gravity Cabinet*      
  *Open door product snapshot query failed*                                            *Gravity Cabinet*      
  *The number of mirrored items is less than the number of identified items.*          *Gravity Cabinet*      
  *Failed to construct the identification request object*                              *Visual Cabinet*       
  *Failed to obtain the identification request object*                                 *Visual Cabinet*       
  *Failed to obtain video*                                                             *Visual Cabinet*       
  *No video*                                                                           *Visual Cabinet*       
  *Order item mirror table query failed*                                               *Visual Cabinet*       
  *Failed to identify the platform*                                                    *Visual Cabinet*       
  ------------------------------------------------------------------------------------ ---------------------- --

### []{#anchor-26}11.5 Manual Settlement Order Event Response

*Description: **This event is triggered when the merchant or platform
settles the order due to abnormal order recognition.***

  ------------------- ----------------------------- -----------------------------------------------------------------------
  ***name***          ***type***                    ***describe***
  *OrderId*           *string*                      *Order Number*
  *OrderNo*           *string*                      *SaasAPI Order Number*
  *TranseType*        *int*                         *Opening method 0 Buy*
  *Action*            *string*                      *OrderSettlement*
  *Status*            *bool*                        *Operational Status*
  *Msg*               *string*                      *Operation Description*
  *CustomerDetails*   *string*                      *User-defined information, parameters passed before opening the door*
  *Data*              ***DetectOrderDetailJson***   *Results Data*
  ------------------- ----------------------------- -----------------------------------------------------------------------

***DetectOrderDetailJson***

+--------------------------+------------+---------------------------+
| ***name***               | ***type*** | ***describe***            |
+--------------------------+------------+---------------------------+
| *TradeProcessMode*       | *int*      | *Transaction processing   |
|                          |            | method: (0=normal         |
|                          |            | transaction, 1=cancel     |
|                          |            | transaction, 2=interrupt  |
|                          |            | transaction)*             |
|                          |            |                           |
|                          |            | *when TradeProcessMode=0* |
|                          |            |                           |
|                          |            | *TradeProcessMode=1       |
|                          |            | cancels the transaction*  |
|                          |            |                           |
|                          |            | *TradeProcessMode=2       |
|                          |            | interrupts the            |
|                          |            | transaction and leaves it |
|                          |            | to the merchant to        |
|                          |            | identify it*              |
+--------------------------+------------+---------------------------+
| ***TradeProductModels*** |            | *Product Contents*        |
+--------------------------+------------+---------------------------+
| *Id*                     | *string*   | *Product ID*              |
+--------------------------+------------+---------------------------+
| *AlisName*               | *string*   | *Product Name*            |
+--------------------------+------------+---------------------------+
| *BuyCount*               | *int*      | *Purchase quantity*       |
+--------------------------+------------+---------------------------+
| *MeterType*              | *int*      | *Pricing method: (1=piece |
|                          |            | (number), 2=weight)*      |
+--------------------------+------------+---------------------------+
| *price*                  | *money*    | *Unit price*              |
+--------------------------+------------+---------------------------+

### []{#anchor-27}11.6 Response to the incident of deducting an order

*Description: **This event is triggered when the merchant clicks the
supplementary deduction button.***

  ------------------- ----------------------- -----------------------------------------------------------------------
  ***name***          ***type***              ***describe***
  *OrderId*           *string*                *Order Number*
  *OrderNo*           *string*                *SaasAPI Order Number*
  *TranseType*        *int*                   *Opening method 0 Buy*
  *Action*            *string*                *OrderAdjustment*
  *Status*            *bool*                  *Operational Status*
  *Msg*               *string*                *Operation Description*
  *CustomerDetails*   *string*                *User-defined information, parameters passed before opening the door*
  *Data*              ***OrderAdjustment***   *Results Data*
  ------------------- ----------------------- -----------------------------------------------------------------------

***OrderAdjustment***

  --------------------- ------------ ------------------------------------------------
  ***name***            ***type***   ***describe***
  *OrderNo*             *string*     *Deduction order number*
  ***OrderProducts***                *Product Contents*
  *Id*                  *string*     *Product ID*
  *AlisName*            *string*     *Product Name*
  *BuyCount*            *int*        *Purchase quantity*
  *MeterType*           *int*        *Pricing method: (1=piece (number), 2=weight)*
  *Price*               *money*      *Unit price*
  --------------------- ------------ ------------------------------------------------

### []{#anchor-28}11.7 Refund Order Event Response

*Description: **This event is triggered when the merchant clicks the
refund button.***

  ------------------- ------------------- -----------------------------------------------------------------------
  ***name***          ***type***          ***describe***
  *OrderId*           *string*            *Order Number*
  *OrderNo*           *string*            *SaasAPI Order Number*
  *TranseType*        *int*               *Opening method 0 Buy*
  *Action*            *string*            *OrderRefund*
  *Status*            *bool*              *Operational Status*
  *Msg*               *string*            *Operation Description*
  *CustomerDetails*   *string*            *User-defined information, parameters passed before opening the door*
  *Data*              ***OrderRefund***   *Results Data*
  ------------------- ------------------- -----------------------------------------------------------------------

***OrderRefund***

  --------------------------- ------------ ------------------------------------------------
  ***name***                  ***type***   ***describe***
  *RefundRemark*              *string*     *Refund Instructions*
  *Refund Order No*           *string*     *Refund order number*
  ***OrderRefundProducts***                *Product Contents*
  *Id*                        *string*     *Product ID*
  *AlisName*                  *string*     *Product Name*
  *BuyCount*                  *int*        *Purchase quantity*
  *MeterType*                 *int*        *Pricing method: (1=piece (number), 2=weight)*
  *Price*                     *money*      *Unit price*
  --------------------------- ------------ ------------------------------------------------

## []{#anchor-29}Common Return Value Description Return value:

***Return value **:*

  ------------------ ------------ ---------------------------------------
  ***name***         ***type***   ***describe***
  ***statusCode***   *int*        *Request status code*
  ***succeeded***    *boolean*    *Whether the operation is successful*
  ***errors***       *string*     *Error message*
  ***extras***       *string*     *Other Tips*
  ***timestamp***    *Long*       *Timestamp*
  ------------------ ------------ ---------------------------------------

## []{#anchor-30}12. Order exception description:

### []{#anchor-31}12.1 Visual Cabinet

  ------------------------------------------------------ ----------------------- ----------------
  ***Exception name***                                   ***type***              ***describe***
  *Block the camera*                                     *Unfriendly shopping*   
  *Replace the goods in the cabinet*                     *Unfriendly shopping*   
  *Insert foreign matter*                                *Unfriendly shopping*   
  *Malicious destruction of goods*                       *Unfriendly shopping*   
  *Cover the product*                                    *Unfriendly shopping*   
  *Long shopping time*                                   *Unfriendly shopping*   
  *Suspected merchant replenishment*                     *Unfriendly shopping*   
  *Other abnormal behaviors*                             *Unfriendly shopping*   
                                                                                 
  *The video is incomplete or the screen is distorted*   *Alerts*                
  *The product in the device is not replenished*         *Alerts*                
  *Products are too similar*                             *Alerts*                
  *The video is too long after closing*                  *Alerts*                
  *Power outage during purchase*                         *Alerts*                
                                                                                 
  ------------------------------------------------------ ----------------------- ----------------

### []{#anchor-32}12.2 Gravity Cabinet
