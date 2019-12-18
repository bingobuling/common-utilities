//author xinbing
//time 2018/12/14 11:36
package paypal

import (
	"common-utilities/payment"
	"time"
)

type GetAccessTokenResp struct {
	Token 	AccessToken 	`json:"token"`
}

type AccessToken struct {
	Scope	string `json:"scope"`
	Nonce 	string `json:"nonce"`
	AccessToken string `json:"access_token"`
	TokenType 	string `json:"token_type"`
	AppId string `json:"app_id"`
	ExpiresIn	int64 `json:"expires_in"`
}

type PayoutReq struct {
	SenderBatchHeader *SenderBatchHeader `json:"sender_batch_header"`
	Items	[]*PayoutItem `json:"items"`
}

type PayoutResp struct {
	PaypalError
	BatchHeader BaseRespBatchHeader `json:"batch_header"`
}
type GetPayoutResp struct {
	PaypalError
	BatchHeader GetPayoutRespBatchHeader `json:"batch_header"`
	Items []GetPayoutRespItem `json:"items"`
}
type GetPayoutRespBatchHeader struct {
	BaseRespBatchHeader
	Amount Amount `json:"amount"` //总金额
	Fees Amount `json:"fees"`	//总手续费
	TimeCreated time.Time `json:"time_created"`
	TimeCompleted time.Time `json:"time_completed"`
}
type GetPayoutRespItem struct {
	PayoutItemId		string `json:"payout_item_id"`
	TransactionId		string `json:"transaction_id"`
	TransactionStatus	string `json:"transaction_status"`
	PayoutItemFee		Amount `json:"payout_item_fee"`
	Amount 				Amount `json:"amount"`
	TimeProcessed		time.Time `json:"time_processed"`
	PayoutBatchId 		string `json:"payout_batch_id"`
	PayoutItem 			PayoutItem `json:"payout_item"`
}

type GetPayoutItemResp struct {
	PaypalError
}

type CancelPayoutItemResp struct {
	PaypalError
}

type BaseRespBatchHeader struct {
	PayoutBatchId string `json:"payout_batch_id"`
	BatchStatus string `json:"batch_status"`
	SenderBatchHeader SenderBatchHeader `json:"sender_batch_header"`
}

type SenderBatchHeader struct {
	EmailSubject string `json:"email_subject"`
	SenderBatchId string `json:"sender_batch_id"`
}

type PayoutItem struct {
	RecipientType	string `json:"recipient_type"`
	Receiver		string `json:"receiver"`
	Amount 			*Amount `json:"amount"`
	Note			string `json:"note"`
	SenderItemId	string `json:"sender_item_id"`
}

type Amount struct {
	Value	string 	`json:"value"`
	Currency payment.Currency `json:"currency"`
}


type PaypalError struct {
	ErrorCode 	string `json:"name"`
	Message		string `json:"message"`
	Details     []PayoutErrorDetail `json:"details"`
}
type PayoutErrorDetail struct {
	Field string `json:"field"`
	Issue string `json:"issue"`
}

func (p *PaypalError) IsSuccess() bool {
	return len(p.ErrorCode) == 0
}