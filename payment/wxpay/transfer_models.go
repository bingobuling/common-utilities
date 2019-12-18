//author xin-bing
//time 11/25/2018 21:07
//desc 微信转账models
package wxpay

import (
	"errors"
	"github.com/bingobuling/common-utilities/payment/alipay/alipayModels"
)

//商户账号appid 	mch_appid 	是 	wx8888888888888888 	String 	申请商户号的appid或商户号绑定的appid
//商户号 	mchid 	是 	1900000109 	String(32) 	微信支付分配的商户号
//设备号 	device_info 	否 	013467007045764 	String(32) 	微信支付分配的终端设备号
//随机字符串 	nonce_str 	是 	5K8264ILTKCH16CQ2502SI8ZNMTM67VS 	String(32) 	随机字符串，不长于32位
//签名 	sign 	是 	C380BEC2BFD727A4B6845133519F3AD6 	String(32) 	签名，详见签名算法
//商户订单号 	partner_trade_no 	是 	10000098201411111234567890 	String 	商户订单号，需保持唯一性
//(只能是字母或者数字，不能包含有符号)
//用户openid 	openid 	是 	oxTWIuGaIt6gTKsQRLau2M0yL16E 	String 	商户appid下，某用户的openid
//校验用户姓名选项 	check_name 	是 	FORCE_CHECK 	String 	NO_CHECK：不校验真实姓名
//FORCE_CHECK：强校验真实姓名
//收款用户姓名 	re_user_name 	可选 	王小王 	String 	收款用户真实姓名。
//如果check_name设置为FORCE_CHECK，则必填用户真实姓名
//金额 	amount 	是 	10099 	int 	企业付款金额，单位为分
//企业付款描述信息 	desc 	是 	理赔 	String 	企业付款操作说明信息。必填。
//Ip地址 	spbill_create_ip 	是 	192.168.0.1 	String(32) 	该IP同在商户平台设置的IP白名单中的IP没有关联，该IP可传用户端或者服务端的IP。
// 转款到零钱请求
type TransferToChangeReq struct {
	OpenID         string //用户的OpenID，必填
	PartnerTradeNo string //订单编号，要确保一个商户下唯一 必填
	Amount         int    //金额，单位为分 必填
	ReceiverName   string //收款人真实姓名,可不用填写，也不建议填写
	Desc           string //描述 必填
	IP             string //必填，该IP可传用户端或者服务端的IP。
}

// 转款到零钱响应
type TransferToChangeResp struct {
	ReturnCode string `xml:"return_code"` //SUCCESS/FAIL 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string `xml:"return_msg"`  //返回信息，如非空，为错误原因 签名失败 参数格式校验错误
	// SUCCESS/FAIL，注意：当状态为FAIL时，存在业务结果未明确的情况，
	// 所以如果状态FAIL，请务必再请求一次查询接口[请务必关注错误代码（err_code字段），通过查询查询接口确认此次付款的结果。]，以确认此次付款的结果。
	ResultCode string `xml:"result_code"`
	ErrCode    string `xml:"err_code"`     //错误码
	ErrCodeDes string `xml:"err_code_des"` //错误码文字描述
	MchAppID   string `xml:"mch_appid"`    //申请商户号的appid或商户号绑定的appid（企业号corpid即为此appId）
	MchID      string `xml:"mchid"`        //微信支付分配的商户号
	DeviceInfo string `xml:"device_info"`  // 微信支付分配的终端设备号，
	NonceStr   string `xml:"nonce_str"`    //随机字符串，不长于32位
	//以下字段在return_code 和result_code都为SUCCESS的时候有返回
	PartnerTradeNo string `xml:"partner_trade_no"` //
	PaymentNo      string `xml:"payment_no"`
	PaymentTime    string `xml:"payment_time"`
}

// 是否转账成功
func (p *TransferToChangeResp) IsSuccess() bool {
	return p.ReturnCode == success_code && p.ResultCode == success_code
}

func (p *TransferToChangeResp) GetTransferTypeErr() (alipayModels.TransferErrType, error) {
	switch p.ErrCode {
	case "OPENID_ERROR":
		return alipayModels.ReceiverNotExist, nil
	case "NOTENOUGH":
		return alipayModels.PayerBalanceNotEnough, nil
	case "NAME_MISMATCH":
		return alipayModels.PayerBalanceNotEnough, nil
	default:
		return alipayModels.UnHandledErrType, errors.New("unHandled aliPay transfer err, code:" + p.ErrCode + ",msg:" + p.ErrCodeDes)
	}
}

//随机字符串 	nonce_str 	是 	5K8264ILTKCH16CQ2502SI8ZNMTM67VS 	String(32) 	随机字符串，不长于32位
//签名 	sign 	是 	C380BEC2BFD727A4B6845133519F3AD6 	String(32) 	生成签名方式查看3.2.1节
//商户订单号 	partner_trade_no 	是 	10000098201411111234567890 	String(28) 	商户调用企业付款API时使用的商户订单号
//商户号 	mch_id 	是 	10000098 	String(32) 	微信支付分配的商户号
//Appid 	appid 	是 	wxe062425f740d30d8 	String(32) 	商户号的appid
// 查询转打款到零钱息请求
type GetTransferToChangeInfoReq struct {
	PartnerTradeNo string //必填
}

type GetTransferToChangeInfoResp struct {
	ReturnCode string `xml:"return_code"` //SUCCESS/FAIL 此字段是通信标识，非付款标识，付款是否成功需要查看result_code来判断
	ReturnMsg  string `xml:"return_msg"`  //返回信息，如非空，为错误原因 签名失败 参数格式校验错误
	//以下信息会在ReturnCode为SUCCESS时返回
	ResultCode     string `xml:"result_code"`      //SUCCESS/FAIL ，非付款标识，付款是否成功需要查看status字段来判断
	ErrorCode      string `xml:"err_code"`         //错误码信息
	ErrCodeDes     string `xml:"err_code_des"`     //错误信息描述
	AppID          string `xml:"appid"`            //微信AppID
	MchID          string `xml:"mch_id"`           //微信支付分配的商户号
	DetailID       string `xml:"detail_id"`        //调用企业付款API时，微信系统内部产生的单号
	PartnerTradeNo string `xml:"partner_trade_no"` //商户单号 商户使用查询API填写的单号的原路返回
	Status         string `xml:"status"`           //转账状态 SUCCESS:转账成功 FAILED:转账失败 PROCESSING:处理中
	Reason         string `xml:"reason"`           //失败原因，如失败则有个失败原因，文字描述
	PaymentAmount  int    `xml:"payment_amount"`   //付款金额 单位为分
	OpenID         string `xml:"openid"`           //收款人的openid
	TransferName   string `xml:"transfer_name"`    //收款人用户姓名
	TransferTime   string `xml:"transfer_time"`    //发起转账的时间
	Desc           string `xml:"desc"`             //付款时的描述
}

func (p *GetTransferToChangeInfoResp) IsSuccess() bool {
	return p.ReturnCode == success_code && p.ResultCode == success_code
}
func (p *GetTransferToChangeInfoResp) IsUnTransferOrder() bool {
	return p.ErrorCode == "NOT_FOUND"
}
