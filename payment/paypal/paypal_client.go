//author xinbing
//time 2018/12/10 14:59
package paypal

import (
	"bytes"
	"common-utilities/http_utils"
	"encoding/base64"
	"encoding/json"
	"time"
	"github.com/pkg/errors"
	"sync"
)


type PaypalClient struct {
	ClientId 				string
	Secret	 				string
	baseApi	 				string
	lastGetAccessTokenTime	int64  //上一次获取accessTime，单位为秒
	token         			*AccessToken
	accessTokenLock			sync.Mutex
}

//同一个client id, pay pal client应当是单例的
func NewPaypalClient(clientId, secret string, prodEnv bool) *PaypalClient {
	baseApi := apiSandBoxBase
	if prodEnv {
		baseApi = apiLiveBase
	}
	return &PaypalClient{
		ClientId: clientId,
		Secret: secret,
		baseApi: baseApi,
		accessTokenLock: sync.Mutex{},
	}
}

// 获取paypal access token
func (p *PaypalClient) GetAccessToken() (resp *GetAccessTokenResp, err error) {
	url := p.baseApi + "/v1/oauth2/token"
	buf := bytes.NewReader([]byte("grant_type=client_credentials"))
	header := p.baseAuthHeader()
	header["Content-Type"] = "application/x-www-form-urlencoded"
	byt, err1 := http_utils.Post(url, buf, header)
	if err1 != nil {
		err = err1
		return
	}
	resp = &GetAccessTokenResp{}
	token := &AccessToken{}
	err = json.Unmarshal(byt, token)
	if err == nil {
		p.lastGetAccessTokenTime = time.Now().Unix()
		p.token = &AccessToken {
			Scope: token.Scope,
			Nonce: token.Nonce,
			AccessToken: token.AccessToken,
			TokenType: token.TokenType,
			AppId: token.AppId,
			ExpiresIn: token.ExpiresIn,
		}
	}
	return
}

// 创建支付
func (p *PaypalClient) CreatePayout(req *PayoutReq) (resp *PayoutResp, err error) {
	byt, err1 := json.Marshal(req)
	if err1 != nil {
		err = err1
		return
	}
	header, err1 := p.authHeader()
	if err1 != nil {
		return nil, errors.WithMessage(err1, "author error")
	}
	url := p.baseApi + "/v1/payments/payouts"
	byt ,err = http_utils.Post(url, bytes.NewReader(byt), header)
	if err != nil && len(byt) == 0 {
		return
	}
	resp = &PayoutResp{}
	err = json.Unmarshal(byt, resp)
	return
}

func (p *PaypalClient) GetPayout(batchId string) (resp *GetPayoutResp, err error){
	header, err1 := p.authHeader()
	if err1 != nil {
		return nil, errors.WithMessage(err1, "author error")
	}
	url := p.baseApi + "/v1/payments/payouts/" + batchId
	byt, err1 := http_utils.Get(url, header)
	if err1 != nil  && len(byt) == 0 {
		return nil, err1
	}
	resp = &GetPayoutResp{}
	err = json.Unmarshal(byt, resp)
	return
}

func (p *PaypalClient) GetPayoutItem(payoutItemId string) (resp *GetPayoutItemResp, err error) {
	header, err1 := p.authHeader()
	if err1 != nil {
		return nil, errors.WithMessage(err1, "author error")
	}
	url := p.baseApi + "/v1/payments/payouts-item/" + payoutItemId
	byt, err1 := http_utils.Get(url, header)
	if err1 != nil  && len(byt) == 0 {
		return nil, err1
	}
	resp = &GetPayoutItemResp{}
	err = json.Unmarshal(byt, resp)
	return
}

func (p *PaypalClient) CancelUnClaimedPayoutItem(payoutItemId string) (resp *CancelPayoutItemResp, err error){
	header, err1 := p.authHeader()
	if err1 != nil {
		return nil, errors.WithMessage(err1, "author error")
	}
	url := p.baseApi + "/v1/payments/payouts-item/" +payoutItemId+ "/cancel"
	byt, err1 := http_utils.Post(url, nil, header)
	if err1 != nil  && len(byt) == 0 {
		return nil, err1
	}
	resp = &CancelPayoutItemResp{}
	err = json.Unmarshal(byt, resp)
	return
}

func (p *PaypalClient) authHeader() (map[string]string, error) {
	if p.needRefreshAccessToken() {
		p.accessTokenLock.Lock()
		defer func(){
			p.accessTokenLock.Unlock()
		} ()
		if p.needRefreshAccessToken() {
			//需要重新获取accessToken
			_, err := p.GetAccessToken()
			if err != nil {
				return nil, errors.WithMessage(err, "get access token error")
			}
		}
	}
	return map[string]string {
		"Accept": "application/json",
		"Authorization": p.token.TokenType + " " + p.token.AccessToken,
		"Content-Type": "application/json",
		"Accept-Language": "en_US",
	}, nil
}

func (p *PaypalClient) needRefreshAccessToken() bool {
	return p.token == nil || time.Now().Unix() - p.lastGetAccessTokenTime - p.token.ExpiresIn <= requsetAccessTokenBeforeExpired
}

func (p *PaypalClient) baseAuthHeader() map[string]string {
	header := map[string]string{}
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(p.ClientId + ":" + p.Secret))
	header["Authorization"] = auth
	header["Accept-Language"] = "en_US"
	header["Accept"] = "application/json"
	return header
}