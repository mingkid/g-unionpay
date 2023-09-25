package aggregation

type Prepay struct {
	StoreID           string          `json:"store_id"`                      // 聚合收单平台为商户分配的惟一 ID
	OutTradeNo        string          `json:"out_trade_no"`                  // 服务商订单号
	TotalAmount       float64         `json:"total_amount"`                  // 订单总金额，单位为元
	MerchantRate      float64         `json:"merchant_rate,omitempty"`       // 商户终端费率
	SubsidyFee        float64         `json:"subsidy_fee"`                   // 交易补贴金额
	PayPlatformType   PayPlatformType `json:"pay_platform_type"`             // 支付平台
	OrderType         OrderType       `json:"order_type"`                    // 订单类型
	SubMerchantNo     string          `json:"sub_merchant_no,omitempty"`     // 可使用指定的微信、支付宝子商户号进行交易
	SubAppID          string          `json:"sub_appid,omitempty"`           // 微信公众号ID
	UserID            string          `json:"user_id,omitempty"`             // 付款用户id
	OpenID            string          `json:"open_id,omitempty"`             // 微信openid
	Subject           string          `json:"subject"`                       // 订单标题
	Body              string          `json:"body"`                          // 商品描述
	NotifyURL         string          `json:"notify_url,omitempty"`          // 支付成功后回调地址
	SpbillCreateIP    string          `json:"spbill_create_ip,omitempty"`    // APP和网页支付提交用户端ip
	CallbackURL       string          `json:"callback_url,omitempty"`        // 小程序支付成功后的跳转地址
	CancelURL         string          `json:"cancel_url,omitempty"`          // 公众号、小程序支付过程中用户取消支付或支付失败跳转地址
	PullUpType        int             `json:"pull_up_type,omitempty"`        // 拉起支付方式
	ShareOrderType    int             `json:"share_order_type,omitempty"`    // 分账标识
	ShareType         ShareType       `json:"share_type,omitempty"`          // 分账类型
	DeviceInfo        string          `json:"device_info,omitempty"`         // 终端设备号
	TimeExpire        string          `json:"time_expire,omitempty"`         // 订单失效时间
	StoreNo           string          `json:"store_no,omitempty"`            // 门店编号
	TerminalID        string          `json:"terminal_id,omitempty"`         // 终端设备号
	NetworkLicense    string          `json:"network_license,omitempty"`     // 银行卡受理终端产品应用认证编号
	Location          string          `json:"location,omitempty"`            // （付款 APP）设备 GPS 位置
	DelayedSettleType int             `json:"delayed_settle_type,omitempty"` // 是否延迟结算订单
	SubsidyAccount    string          `json:"subsidy_account,omitempty"`     // 补贴账户
	FoodOrderType     FoodOrderType   `json:"food_order_type,omitempty"`     // 支付宝扫码点餐类型
}

// AliPayPrepay 支付宝统一下单
type AliPayPrepay struct {
	Prepay
	DisablePayChannels   string        `json:"disable_pay_channels,omitempty"`  // 该笔交易需禁用的支付渠道
	EnablePayChannels    string        `json:"enable_pay_channels,omitempty"`   // 该笔交易可用的支付渠道
	GoodsDetail          string        `json:"goods_detail,omitempty"`          // 商品详情
	DiscountableAmount   float64       `json:"discountable_amount,omitempty"`   // 参与优惠计算的金额
	UndiscountableAmount float64       `json:"undiscountable_amount,omitempty"` // 不参与优惠计算的金额
	BusinessParams       string        `json:"business_params,omitempty"`       // 支付宝商户传入业务信息
	AlipayExtend         *AliPayExtend `json:"alipay_extend_params,omitempty"`  // 支付宝扩展信息
}

type AliPayExtend struct {
	HbFqNum              HbFqNum `json:"hb_fq_num,omitempty"`               // 花呗分期期数
	HbFqSellerPercent    int     `json:"hb_fq_seller_percent,omitempty"`    // 花呗分期手续费承担模式
	SysServiceProviderID string  `json:"sys_service_provider_id,omitempty"` // 系统商编号
	IndustryRefluxInfo   string  `json:"industry_reflux_info,omitempty"`    // 行业数据回流信息
	CardType             string  `json:"card_type,omitempty"`               // 卡类型
}

// WeChatPayPrepay 微信支付统一下单
type WeChatPayPrepay struct {
	Prepay
	GoodsTag  string     `json:"goods_tag,omitempty"`  // 订单优惠标记
	LimitPay  string     `json:"limit_pay,omitempty"`  // 微信禁用信用卡
	SceneInfo *SceneInfo `json:"scene_info,omitempty"` // 场景信息
}

// SceneInfo 场景信息
type SceneInfo struct {
	ID       string `json:"id,omitempty"`        // 门店ID
	Name     string `json:"name,omitempty"`      // 门店名称
	AreaCode string `json:"area_code,omitempty"` // 门店行政区划码
	Address  string `json:"address,omitempty"`   // 门店详细地址
}
