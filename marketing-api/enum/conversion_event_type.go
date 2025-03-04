package enum

// ConversionEventType 转化回传事件
// https://event-manager.oceanengine.com/docs/8650/all_events/
type ConversionEventType = string

const (
	// 用户点击拨打电话
	Conversion_PHONE ConversionEventType = "phone"
	// 用户在页面内完成表单填写并提交
	Conversion_FORM ConversionEventType = "form"
	// 用户下载app完成激活行为
	Conversion_ACTIVE ConversionEventType = "active"
	// 用户在推广的应用/落地页场景下完成注册（手机号/微信等）行为，具体注册形式取决于广告主业务模式，如：招商加盟行业，用户下载app后完成注册行为；游戏行业，用户下载游戏后完成激活+注册；电商行业，完成注册
	Conversion_ACTIVE_REGISTER ConversionEventType = "active_register"
	// 用户在推广的应用/落地页场景下发生交易并完成至少一笔付款，具体支付形式取决于广告主业务模式
	Conversion_ACTIVE_PAY ConversionEventType = "active_pay"
	// 智能电话-确认拨打
	Conversion_PHONE_CONFIRM ConversionEventType = "phone_confirm"
	// 智能电话-确认接通
	Conversion_PHONE_CONNECT ConversionEventType = "phone_connect"
	// 智能电话-有效接通
	Conversion_PHONE_EFFECT ConversionEventType = "phone_effective"
	// 一般指用户点击“在线咨询”，通过IM向商户进行有效咨询，发言≥1条，广告主也可自定义有效的标准。
	Conversion_CONSULT_EFFECTIVE ConversionEventType = "consult_effective"
	// 用户进入电商平台APP，并完成提交订单行为
	Conversion_IN_APP_ORDER ConversionEventType = "in_app_order"
	// 用户成功打开访问应用打开并访问通常考核吊起打开，不考核活动/商品 uv/pv
	Conversion_IN_APP_UV ConversionEventType = "in_app_uv"
	// 用户将商品加入购物车
	Conversion_IN_APP_CART ConversionEventType = "in_app_cart"
	// 用户完成商品下单并完成付款
	Conversion_IN_APP_PAY ConversionEventType = "in_app_pay"
	// 用户在推广的应用/落地页场景下发生的关键行为/行为集合，具体行为取决于广告主业务模式。如：商务服务行业的用户下载激活app后提交简历，房地产行业的用户有看房意向；游戏行业的用户在游戏内发生关键行为或行为组合等
	Conversion_GAME_ADDICTION ConversionEventType = "game_addiction"
	// 提交线索的用户进一步完成了一次有价值的行为，具体行为取决于广告主业务模式，如预约到店，完成授权，完成支付等
	Conversion_CUSTOMER_EFFECTIVE ConversionEventType = "customer_effective"
	// 用户完成卡券领取行为
	Conversion_COUPON = "coupon"
	// 浏览重点页面，比如商品页、活动页、文章等考核活动/商品页面uv/pv
	Conversion_IN_APP_DETAIL_UV = "in_app_detail_uv"
	// 表单提交后，电话回访已回应
	Conversion_FROM_ANSWER = "form_answer"
	// 激活且在第二天内打开app
	Conversion_NEXT_DAY_OPEN = "next_day_open"
	// 用户在游戏内多次付费
	Conversion_DEEP_PURCHASE = "deep_purchase"
	// 用户完成指定目标页面加载并到访
	Conversion_PAGE_VIEW = "page_view"
	// 用户完成商品购买行为
	Conversion_SHOPPING = "shopping"
	// 用户完成微信复制行为，引导用户微信加粉
	Conversion_WECHAT = "wechat"
	// 其他
	Conversion_OTHER = "other"
	// 用户发生表单提交、电话拨打、咨询、预约等多种留咨行为
	Conversion_MULTIPLE = "multiple"
	// 贷款完件
	Conversion_LOAN_COMPLETION = "loan_completion"
	// 贷款预授信
	Conversion_PRE_LOAN_CREDIT = "pre_loan_credit"
	// 贷款授信
	Conversion_LOAN_CREDIT = "loan_credit"
	// 广告变现ROI
	Conversion_IT_ROI = "lt_roi"
	// 贷款放款
	Conversion_LOAN = "loan"
	// 用户完成授权登录（微信/微博/淘宝等）
	Conversion_AUTHORIZATION = "authorization"
	// 用户在咨询中成功留咨
	Conversion_CONSULT_CLUE = "consult_clue"
	// 用户付费金额/付费成本
	Conversion_PURCHASE_ROI = "purchase_roi"
	// 用户点击预约“游戏下载”
	Conversion_NOTIFY_DOWNLOAD = "notify_download"
	// 用户完成购买保险行为
	Conversion_PREMIUM_PAYMENT = "premium_payment"
	// 内广roi
	Conversion_UG_ROI = "ug_roi"
	// 用户在微信生态进入H5页面
	Conversion_IN_WECHAT_LOGIN = "in_wechat_login"
	// 用户在微信生态完成付费
	Conversion_IN_WECHAT_PAY = "in_wechat_pay"
	// 线索经联系确认是本人提交的信息，或者是本人有初步意向
	Conversion_CLUE_CONFIRM = "clue_confirm"
	// 线索和销售建立了交流，比如互加好友，建立联系，可以持续跟进
	Conversion_CLUE_INTERFLOW = "clue_interflow"
	// 线索有较强意向成交或者处于成交流程，尚未完结。
	Conversion_CLUE_HIGH_INTENTION = "clue_high_intention"
	// 在线索通表单提交联系方式成功之后，用户完成一笔在线支付行为
	Conversion_CLUE_PAY_SUCCEED = "clue_pay_succeed"
	// 线索客户产生到店行为
	Conversion_CLUE_ARRIVAL_LEAD = "arrival_lead"
	// 有效线索客户跟进失败
	Conversion_FAILED_LEAD = "failed_lead"
	// 适用于平台类商家客户拉活场景：沉默老用户进行app内访问后的次日留存，作为“app内访问”的深度转化目标组合投放，不可直投。属于T+1事件，使用“T+1模型”预估。
	Conversion_IN_APP_NEXT_DAY_OPEN = "in_app_next_day_open"
	// 用户完成首次保险支付短期内退保
	Conversion_WITHDRAW_INSURANCE = "withdraw_insurance"
	// 用户激活首日发生一次广告展示；发生多次广告展示的时候，多次回传。
	Conversion_IMPRESSION = "impression"
	// 用户激活首日广告展示次数达到预期，目前由广告主自定义；按历史经验，一般为5-20次
	Conversion_IPU_QUALIFY = "ipu_qualify"
	// 用户在搜索结果页发生点击
	Conversion_LU_CLICK = "lu_click"
	// 用户在搜索结果页上发生转化
	Conversion_LU_CONVERT = "lu_convert"
	// 用户完成首次保险支付60天内退保
	Conversion_WITHDRAW_M2 = "withdraw_m2"
	// 信贷客户为判定为次级的用户完成授信
	Conversion_LOW_LOAN_CREDIT = "low_loan_credit"
	// 信贷客户为判定为高价值的用户完成授信
	Conversion_HIGH_LOAN_CREDIT = "high_loan_credit"
	// 加微信成功后T日取消关注
	Conversion_UNFOLLOW_IN_WECHAT = "unfollow_in_wechat"
)
