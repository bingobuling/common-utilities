package alipayModels

import (
	"time"
	"errors"
)

//给用户转账model
type TransferToAccountReq struct {
	OrderId       string `json:"order_id"`//这笔转账的id，唯一
	Amount        float64 `json:"amount"`//转账金额，至少一毛钱，单位为元，只支持2位小数
	PayeeType     string `json:"payee_type"`//收款方账户类型，分为：ALIPAY_USERID（支付宝id），ALIPAY_LOGONID（支付宝账号，如手机号、邮箱），默认值为ALIPAY_LOGONID
	PayeeAccount  string `json:"payee_account"`//收款方账户，根据PayeeType不同而可能不同
	PayerShowName string `json:"payer_show_name"`//转账人姓名
	PayeeRealName string `json:"payee_real_name"`//收款人姓名，如果填写，那么会支付宝会校验账号的实名是否与该名称一致，如不一致打款失败
	Remark        string `json:"remark"`//转账备注 最多200字符
}

type TransferToAccountResp struct {
	Code    string 		`json:"code"`
	Msg     string 		`json:"msg"`
	SubCode string 		`json:"sub_code"`
	SubMsg  string 		`json:"sub_msg"`
	OutBizNo string 	`json:"out_biz_no"`  //商户的订单id
	TradeNo string 	   `json:"order_id"`    //支付宝上的这笔交易的交易单号,
	PayDate time.Time  `json:"pay_date"` //转款时间
}

// 是否转款成功
func (p *TransferToAccountResp) IsSuccess() bool{
	return p.Code == "10000"
}

// 获取转账失败的错误
func (p *TransferToAccountResp) GetTransferErrType() (TransferErrType, error) {
	switch p.SubCode {
	case "PAYEE_USER_INFO_ERROR":
		return ReceiverInfoNotMatch, nil
	case "PAYEE_NOT_EXIST":
		return ReceiverNotExist, nil
	case "PAYEE_ACC_OCUPIED":
		return ReceiverAmbiguous, nil
	case "PAYER_BALANCE_NOT_ENOUGH":
		return PayerBalanceNotEnough, nil
	case "PAYER_PAYEE_CANNOT_SAME":
		return PayerReceiverCannotSame, nil
	case "EXCEED_LIMIT_SM_MIN_AMOUNT":
		return ExceedLimitMinAmount, nil
	case "EXCEED_LIMIT_PERSONAL_SM_AMOUNT", "EXCEED_LIMIT_ENT_SM_AMOUNT", "EXCEED_LIMIT_UNRN_DM_AMOUNT":
		return ExceedLimitMaxAmount, nil
	case "EXCEED_LIMIT_DM_MAX_AMOUNT":
		return ExceedLimitDailyMaxAmount, nil
	default:
		return UnHandledErrType, errors.New("unHandled aliPay transfer err, code:"+p.SubCode+",msg:"+p.SubMsg)
	}
}

// 转账错误
type TransferErrType string
const (
	UnHandledErrType 		TransferErrType = "UnHandledErrType"	 //未处理的转账错误
	ReceiverNotExist		TransferErrType = "ReceiverNotExist"			//收款账号不存在
	ReceiverInfoNotMatch  TransferErrType = "ReceiverInfoNotMatch"		//收款账户信息不匹配
	ReceiverAmbiguous	TransferErrType = "ReceiverAmbiguous"			//收款人信息模糊不能确定，例如一个手机号对应多个支付宝，单传手机号不能确定，还需要传输真实姓名
	PayerBalanceNotEnough	TransferErrType = "PayerBalanceNotEnough"		//支付账户余额不足
	PayerReceiverCannotSame  TransferErrType = "PayerReceiverCannotSame"		//支付账户，收款账户不能一致
	ExceedLimitMinAmount	TransferErrType = "ExceedLimitMinAmount"		//低于最低支付限额
	ExceedLimitMaxAmount	TransferErrType = "ExceedLimitMaxAmount"		//超出单笔个人转账最大限额
	ExceedLimitDailyMaxAmount	TransferErrType = "ExceedLimitDailyMaxAmount"	//超出每日转账的最大金额
)

type QueryTransferToAccountReq struct {
	OutBizNo     string 	`json:"out_biz_no,omitempty"` // 与 OrderId 二选一
	OrderId      string 	`json:"order_id,omitempty"`   // 与 OutBizNo 二选一
}

type QueryTransferToAccountResp struct {
	Code           string    `json:"code"`
	Msg            string    `json:"msg"`
	SubCode        string    `json:"sub_code"`
	SubMsg         string    `json:"sub_msg"`
	FailReason     string 	  `json:"fail_reason"`      // 查询到的订单状态为FAIL失败或REFUND退票时，返回具体的原因。
	ErrorCode      string    `json:"error_code"`       // 查询失败时，本参数为错误代 码。 查询成功不返回。 对于退票订单，不返回该参数。
	OutBizNo        string 	  `json:"out_biz_no"`  // 发起转账来源方定义的转账单据号。 该参数的赋值均以查询结果中 的 out_biz_no 为准。 如果查询失败，不返回该参数
	TradeNo        string    `json:"order_id"` // 支付宝转账单据号，查询失败不返回。
	Status         string 	  `json:"status"`  // 转账单据状态
	PayDate        time.Time `json:"pay_date"` // 支付时间
	ArrivalTimeEnd time.Time `json:"arrival_time_end"`// 预计到账时间，转账到银行卡专用
	OrderFee       float64 	 `json:"order_fee"`// 预计收费金额（元），转账到银行卡专用
}

// 是否查询成功
func (p *QueryTransferToAccountResp) IsSuccess() bool{
	return p.Code == "10000"
}

// 是否未转账
func (p *QueryTransferToAccountResp) IsUnTransferOrder () bool{
	return p.SubCode == "ORDER_NOT_EXIST"
}