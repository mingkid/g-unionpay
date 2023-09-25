package aggregation

// OrderType  订单类型
type OrderType string

const (
	OrderTypeNative     OrderType = "2"  // 原生扫码支付
	OrderTypeJSAPI      OrderType = "3"  // 公众号支付
	OrderTypeMiniApp    OrderType = "4"  // 小程序支付
	OrderTypeCloudPay   OrderType = "6"  // 云闪付JS支付
	OrderTypePengPeng   OrderType = "7"  // 碰一碰
	OrderTypeApp        OrderType = "8"  // APP拉起支付
	OrderTypeSubWallet  OrderType = "9"  // 子钱包支付
	OrderTypeHardWallet OrderType = "10" // 硬钱包支付
)

// PayPlatformType 支付平台类型
type PayPlatformType string

const (
	PayPlatformAlipay      PayPlatformType = "1"  // 支付宝
	PayPlatformWeChat      PayPlatformType = "2"  // 微信
	PayPlatformUnionPay    PayPlatformType = "6"  // 云闪付
	PayPlatformCMBDigital  PayPlatformType = "9"  // 中行数字人民币
	PayPlatformUnionPayStg PayPlatformType = "13" // 银联聚分期
	PayPlatformCMBUnified  PayPlatformType = "16" // 中行数币统一结算
	PayPlatformPrepaidCard PayPlatformType = "17" // 预付卡
)

// ShareType 分账类型
type ShareType string

const (
	ShareTypeOrder   ShareType = "1" // 订单分账
	ShareTypeBalance ShareType = "2" // 余额分账
)

// FoodOrderType 支付宝扫码点餐类型
type FoodOrderType string

const (
	FoodOrderQR       FoodOrderType = "qr_order"       // 店内扫码点餐
	FoodOrderPreOrder FoodOrderType = "pre_order"      // 预点到店自提
	FoodOrderDelivery FoodOrderType = "home_delivery"  // 外送到家
	FoodOrderDirect   FoodOrderType = "direct_payment" // 直接付款
	FoodOrderOther    FoodOrderType = "other"          // 其它
)

// HbFqNum 花呗分期期数
type HbFqNum uint8

const (
	HbFqNum3  HbFqNum = 3  // 花呗分 3 期
	HbFqNum6  HbFqNum = 6  // 花呗分 6 期
	HbFqNum12 HbFqNum = 12 //  花呗分 12 期
)
