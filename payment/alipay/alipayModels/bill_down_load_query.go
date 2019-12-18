//author xinbing
//time 2018/12/3 20:38
package alipayModels

type BillDownloadQueryReq struct {
	BillType BillDownloadQueryType `json:"bill_type"`
	BillDate string                `json:"bill_date"` //日期yyyy-MM-dd 或者 yyyy-MM，不传dd则按月查询
}

type BillDownloadQueryResp struct {
	BillDownloadUrl string `json:"bill_download_url"`
}

type BillDownloadQueryType string

const (
	BillDownloadQueryType_Trade        = "trade"        // 账单类型，商户通过接口或商户经开放平台授权后其所属服务商通过接口可以获取以下账单类型
	BillDownloadQueryType_SignCustomer = "signcustomer" // 指商户基于支付宝交易收单的业务账单 signcustomer是指基于商户支付宝余额收入及支出等资金变动的帐务账单
)
