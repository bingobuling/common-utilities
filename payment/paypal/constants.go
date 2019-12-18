//author xinbing
//time 2018/12/14 14:09
package paypal

const apiSandBoxBase = "https://api.sandbox.paypal.com"	//沙盒环境连接地址
const apiLiveBase = "https://api.paypal.com"	//线上环境连接地址

const requsetAccessTokenBeforeExpired = 60 * 2 //过期2分钟前处理

const (
	RecipientType_Email = "EMAIL" //邮箱格式
	RecipientType_PHONE = "PHONE"	//手机号码格式，全球的手机号码格式，极特殊的情况下可能会包含字母？
	RecipientType_PAYPAL_ID	= "PAYPAL_ID" //paypalId 示例：5ZF799AX8X3CL
)