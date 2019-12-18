//author xinbing
//time 2018/11/15 14:21
//
package paypal

import (
	"testing"
	"fmt"
)

const (
	youqian_client_id = "ATOPNgfccc0uAPWPirhVgioEwf9Ognq7tsGGO-xm3Kt3GJ04nwieZkJjg-OLRT7hlwgjGsLTfNx-M2vH"

	sandbox_api = "https://svcs.sandbox.paypal.com/AdaptivePayments/Pay"
	sandbox_payout = "https://api.sandbox.paypal.com/v1/payments/payouts"
	sendbox_sender = "bing.xin-facilitator@outlook.com"
	sendbox_sender_token = "access_token$sandbox$y6md5p96xxt63ktp$7ee04882aa0775e6f1e39fbc1bd805a6"
	send_app = "test facilitators Test Store"
	send_signature = "AA5-MECTA9ewcabfxYBNqYXkb1uvAD0r9KEnA9XNCNNg-CdCrdRjNzKU"
	send_pwd = "GU8BUVG28DYUJ6CE"

	sandbox_sender_1 = "bing.xin-facilitator-1@outlook.com"
	sendbox_access_token_1 = "access_token$sandbox$49jzsjzqtzvgtbqk$c7da5ab61a05bf7b34489976a42084f1"
	send_1_app = "test youqians Test Store"
	send_1_signature = "A8fHY0MF6LIr1huvo-q8qjl7joROAM2kb8vmzeN0sOwTF0fIW8kVh1zV"
	//send_1_pwd = "ZAN2VJF2DUSRMCLS"
	send_1_pwd = "ddup25321990"
	buyer_0 = "bing.xin-buyer@outlook.com"
	buyer_1 = "bing.xin-buyer-1@outlook.com"
)
var cc = NewPaypalClient(youqian_client_id,
"EESzMGliVFsmeKGQ0NZAef5UdPcdz9mQTrOu8uSQRen5amLoHSRW_cOyTEpkAKvP-ldmM_l3rn4zpUh9", false)

func TestTransfer(t *testing.T) {
	//batchId := utilities.GetRandomStr(16)
	batchId := "11111111111111111111111111111111111114"
	req := &PayoutReq{
		SenderBatchHeader: &SenderBatchHeader{
			EmailSubject: "i love u",
			SenderBatchId: batchId,
		},
		Items: []*PayoutItem{
			{
				RecipientType: RecipientType_Email,
				Receiver: "bing.xin-buyer@outlook.com",
				Amount: &Amount{
					Value: "1.0",
					Currency: "USD",
				},
				Note: "i will give you some money",
				SenderItemId: batchId + "0",
			},
		},
	}

	resp, err := cc.CreatePayout(req)
	fmt.Println(resp.BatchHeader, err)
}

func TestPaypalClient_GetPayout(t *testing.T) {
	//AK8S3QBDJJVAU
	resp, err := cc.GetPayout("AK8S3QBDJJVAU") //20190119——AK8S3QBDJJVAU
	fmt.Println(resp, err)
}

func TestPaypalClient_GetPayoutItem(t *testing.T) {
	resp, err := cc.GetPayoutItem("")
	fmt.Println(resp, err)
}

func TestPaypalClient_CancelUnClaimedPayoutItem(t *testing.T) {
	resp, err := cc.CancelUnClaimedPayoutItem("")
	fmt.Println(resp, err)
}