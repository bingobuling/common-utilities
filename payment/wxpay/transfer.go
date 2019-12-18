package wxpay

import (
	"bytes"
	"encoding/xml"
	"github.com/bingobuling/common-utilities/utilities"
	"github.com/chanxuehong/util"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"gopkg.in/chanxuehong/wechat.v2/mch/core"
	"gopkg.in/chanxuehong/wechat.v2/mch/mmpaymkttransfers"
	"gopkg.in/chanxuehong/wechat.v2/mch/mmpaymkttransfers/promotion"
	"strconv"
	"strings"
)

const (
	success_code = "SUCCESS"
)

// 微信转款到零钱
func TransferToChange(req *TransferToChangeReq, client *core.Client) (*TransferToChangeResp, error) {
	if len(req.OpenID) == 0 {
		return nil, errors.New("openid cannot empty")
	}
	if len(req.PartnerTradeNo) == 0 {
		return nil, errors.New("partner_trade_no cannot empty")
	}
	if len(req.Desc) == 0 {
		return nil, errors.New("desc cannot empty")
	}
	if len(req.IP) == 0 {
		return nil, errors.New("ip cannot empty")
	}
	if req.Amount < 100 {
		return nil, errors.New("amount cannot lower than 100 cent")
	}
	nonce := utilities.GetRandomStr(32)
	params := map[string]string{
		"mch_appid":        client.AppId(),
		"mchid":            client.MchId(),
		"openid":           req.OpenID,
		"partner_trade_no": req.PartnerTradeNo,
		"nonce_str":        nonce,
		"amount":           strconv.Itoa(req.Amount),
		"desc":             req.Desc,
		"spbill_create_ip": req.IP,
	}
	receiverName := strings.Trim(req.ReceiverName, " ")
	if len(receiverName) == 0 {
		params["check_name"] = "NO_CHECK"
	} else {
		params["check_name"] = "FORCE_CHECK"
		params["re_user_name"] = receiverName
	}
	respMap, err := promotion.Transfers(client, params)
	resp := &TransferToChangeResp{}
	if err != nil {
		if err2 := xml.Unmarshal([]byte(err.Error()), resp); err2 == nil {
			return resp, nil
		}
		return nil, errors.New("TransferToChange Transfers error:" + err.Error())
	}
	err = toRespFromMap(&respMap, resp, "TransferToChange")
	return resp, nil
}

// 查询
func GetTransferToChangeInfo(req *GetTransferToChangeInfoReq, client *core.Client) (*GetTransferToChangeInfoResp, error) {
	nonce := utilities.GetRandomStr(32)
	params := map[string]string{
		"appid":            client.AppId(),
		"mch_id":           client.MchId(),
		"nonce_str":        nonce,
		"partner_trade_no": req.PartnerTradeNo,
	}
	respMap, err := mmpaymkttransfers.GetTransferInfo(client, params)
	resp := &GetTransferToChangeInfoResp{}
	if err != nil {
		if err2 := xml.Unmarshal([]byte(err.Error()), resp); err2 == nil {
			return resp, nil
		}
		return nil, errors.New("GetTransferToChangeInfo GetTransferInfo error:" + err.Error())
	}
	err = toRespFromMap(&respMap, resp, "GetTransferToChangeInfo")
	return resp, err
}

func toRespFromMap(respMap *map[string]string, pointer interface{}, logFlag string) error {
	buffer := bytes.NewBuffer(make([]byte, 0, 4<<10))
	err := util.EncodeXMLFromMap(buffer, *respMap, "xml")
	if err != nil {
		glog.Errorln()
		return errors.New(logFlag + " EncodeXMLFromMap err:" + err.Error())
	}
	err = xml.Unmarshal(buffer.Bytes(), pointer)
	if err != nil {
		return errors.New(logFlag + " xml.Unmarshal error:" + err.Error())
	}
	return nil
}
