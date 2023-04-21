package e

var MsgFlags = map[uint]string{
	Success:      "success",
	Error:        "fail",
	InvalidParam: "param invalid",
}

// GetMsg	获取状态码对应信息
func GetMsg(code uint) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
