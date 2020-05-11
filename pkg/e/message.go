package e

var messages = map[int]string{
	ERROR:         "服务器错误",
	NotFound:      "资源未找到",
	Success:       "操作成功",
	InvalidParams: "参数不合法",
	TokenExpired:  "token已过期",
	TokenInvalid:  "token验证失败",
}

func GetMessage(code int) string {
	if s, ok := messages[code]; ok {
		return s
	}
	return messages[ERROR]
}
