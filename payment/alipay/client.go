//author xinbing
//time 2018/10/17 9:56
//发起支付宝请求的client
package alipay

import "github.com/smartwalle/alipay"

type AlipayClient alipay.AliPay

func GetAliPayClient(appId, privateKey, publicKey string) *AlipayClient {
	client := alipay.New(
		appId,
		"",
		publicKey,
		privateKey,
		true,
	)
	alipayClient := AlipayClient(*client)
	return &alipayClient
}
