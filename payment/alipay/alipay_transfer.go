//author xinbing
//time 2018/10/17 10:50
//转账相关接口
package alipay

import (
	"github.com/bingobuling/common-utilities/payment/alipay/alipayModels"
	"github.com/bingobuling/common-utilities/utilities"
	"github.com/pkg/errors"
	"github.com/smartwalle/alipay"
	"strconv"
	"time"
)

//转账给用户
func TransferToAccount(model alipayModels.TransferToAccountReq, alipayClient *AlipayClient) (*alipayModels.TransferToAccountResp, error) {
	if model.PayeeType != ALIPAYUERID {
		model.PayeeType = ALIPAYLOGONID
	}
	reqBody := alipay.AliPayFundTransToAccountTransfer{}
	reqBody.OutBizNo = model.OrderId
	reqBody.Amount = strconv.FormatFloat(model.Amount, 'f', 2, 64)
	reqBody.PayeeType = model.PayeeType
	reqBody.PayerShowName = model.PayerShowName
	reqBody.PayeeRealName = model.PayeeRealName
	reqBody.PayeeAccount = model.PayeeAccount
	reqBody.Remark = model.Remark
	if len(reqBody.OutBizNo) == 0 {
		return nil, errors.New("out_biz_no cannot empty")
	}
	if utilities.Compare(model.Amount, 0.1) < 0 {
		return nil, errors.New("amount must greater than 0.1")
	}
	if len(reqBody.PayeeAccount) == 0 {
		return nil, errors.New("payee_account cannot empty")
	}
	if len(reqBody.PayerShowName) == 0 {
		return nil, errors.New("payer_show_name cannot empty")
	}
	client := alipay.AliPay(*alipayClient)
	resp, err := client.FundTransToAccountTransfer(reqBody)
	if err != nil {
		return nil, err
	}
	return &alipayModels.TransferToAccountResp{
		Code:     resp.Body.Code,
		Msg:      resp.Body.Msg,
		SubCode:  resp.Body.SubCode,
		SubMsg:   resp.Body.SubMsg,
		OutBizNo: resp.Body.OutBizNo,
		TradeNo:  resp.Body.OrderId,
		PayDate:  utilities.ParseDateTimeStrWithDefault(resp.Body.PayDate, time.Now()),
	}, nil
}

//查询转账给用户的交易，可以使用商家的，也可以使用支付宝的订单单号
func QueryTransferToAccount(queryReq alipayModels.QueryTransferToAccountReq, alipayClient *AlipayClient) (*alipayModels.QueryTransferToAccountResp, error) {
	req := alipay.AliPayFundTransOrderQuery{
		OutBizNo: queryReq.OutBizNo, //二者有一即可
		OrderId:  queryReq.OrderId,
	}
	client := alipay.AliPay(*alipayClient)
	resp, err := client.FundTransOrderQuery(req)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return &alipayModels.QueryTransferToAccountResp{
			Code:       resp.Body.Code,
			Msg:        resp.Body.Msg,
			SubCode:    resp.Body.SubCode,
			SubMsg:     resp.Body.SubMsg,
			ErrorCode:  resp.Body.ErrorCode,
			FailReason: resp.Body.FailReason,
		}, nil
	}

	OrderFee, _ := strconv.ParseFloat(resp.Body.OrderFree, 64)
	return &alipayModels.QueryTransferToAccountResp{
		Code:           resp.Body.Code,
		Msg:            resp.Body.Msg,
		SubCode:        resp.Body.SubCode,
		SubMsg:         resp.Body.SubMsg,
		ErrorCode:      resp.Body.ErrorCode,
		FailReason:     resp.Body.FailReason,
		OutBizNo:       resp.Body.OutBizNo,
		TradeNo:        resp.Body.OrderId,
		Status:         resp.Body.Status,
		PayDate:        utilities.ParseDateTimeStrWithDefault(resp.Body.PayDate, time.Now()),
		ArrivalTimeEnd: utilities.ParseDateTimeStrWithDefault(resp.Body.ArrivalTimeEnd, time.Now()),
		OrderFee:       OrderFee,
	}, nil
}
