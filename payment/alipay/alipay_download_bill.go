//author xinbing
//time 2018/10/17 9:57
//下载账单
package alipay

import (
	"github.com/smartwalle/alipay"
	"common-utilities/payment/alipay/alipayModels"
)

// 下载bill
func DownloadBill(params *alipayModels.BillDownloadQueryReq ,alipayClient *AlipayClient) (*alipayModels.BillDownloadQueryResp, error){
	client := alipay.AliPay(*alipayClient)
	query := alipay.BillDownloadURLQuery{
		BillDate: params.BillDate,
		BillType: string(params.BillType),
	}
	resp, err := client.BillDownloadURLQuery(query)
	if err != nil {
		return nil, err
	}
	return &alipayModels.BillDownloadQueryResp{
		BillDownloadUrl: resp.AliPayDataServiceBillDownloadURLQueryResponse.BillDownloadUrl,
	}, err
}