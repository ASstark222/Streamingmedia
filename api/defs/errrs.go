package defs

/*
定义错误信息
 */

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

//错误响应
type ErrResponse struct {
	HttpSC int
	Error Err
}

var (
	//解析错误
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400,Error: Err{Error: "Request body is not correct",ErrorCode: "001"}}
	//用户验证错误
	ErrorNotAuthUser = ErrResponse{HttpSC: 401,Error: Err{Error: "User authentication failed",ErrorCode: "002"}}
	//数据库操作错误
	ErrorDBError = ErrResponse{HttpSC: 500,Error:Err{Error: "DB operation failed",ErrorCode: "003"}}
	//json序列化错误
	ErrorInternalFaults = ErrResponse{HttpSC: 500,Error: Err{Error: "Internal service error ",ErrorCode: "004"}}
)