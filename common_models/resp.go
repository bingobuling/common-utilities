//author xinbing
//time 2018/9/13 11:24
//
package common_models

type Resp struct {
	Code 		int 				`json:"code"`
	Message 	string 				`json:"message"`
	Data    	interface{} 		`json:"data"` //data 待废弃
}

var BadRequest = Resp{}.FailedWithCode(400, "bad request")

func (p Resp) Success(msg string, data interface{}) *Resp {
	p.Code = 0
	p.Message = msg
	p.Data = data
	return &p
}

func (p Resp) Failed(msg string) *Resp {
	p.Code = -1
	p.Message = msg
	return &p
}

func (p Resp) FailedWithCode(code int, msg string) *Resp {
	p.Code = code
	p.Message = msg
	return &p
}