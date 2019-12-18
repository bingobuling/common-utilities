//author xinbing
//time 2018/10/17 9:50
// 支付相关接口
package alipay

import (
	"common-utilities/payment/alipay/alipayModels"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/smartwalle/alipay"
)

//查询支付接口，使用的支付宝的流水单号，这个接口不能够处查询转账给用户的交易流水号
func QueryAliTrade(tradeNo string, alipayClient *AlipayClient) (*alipayModels.QueryAliTradeResp, error) {
	reqBody := alipay.AliPayTradeQuery{
		TradeNo: tradeNo,
	}
	client := alipay.AliPay(*alipayClient)
	resp, err := client.TradeQuery(reqBody)
	logrus.Info("TransToAccount return：", resp, err)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		logrus.Error("TransToAccount Error", resp)
		return nil, errors.New(resp.AliPayTradeQuery.SubMsg)
	}
	bytes, err := json.Marshal(resp.AliPayTradeQuery)
	if err != nil {
		return nil, err
	}
	returnResp := &alipayModels.QueryAliTradeResp{}
	err = json.Unmarshal(bytes, returnResp)
	if err != nil {
		return nil, err
	}
	return returnResp, nil
}

