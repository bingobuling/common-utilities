//author xinbing
//time 2018/12/10 16:59
package payment

type Currency string

const (
	Currency_CNY = "CNY"	//人民币，paypal不支持人民币
	Currency_HKD = "HKD"	//港币
	Currency_TWD = "TWD"	//台币
	Currency_JPY = "JPY"   //日元
	Currency_KRW = "KRW"   //韩元
	Currency_USD = "USD"   //american dollar
	Currency_EUR = "EUR"   //euro
	Currency_GBP = "GBP"   //英镑
	Currency_AUD = "AUD"  //澳大利亚元
	Currency_CAD = "CAD"  //加元
)