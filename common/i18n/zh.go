package i18n

const (
	ErrBind   string = "绑定失败"
	ErrServer string = "服务器忙碌，请稍后重试"
)

var ZhMessage = map[string]string{
	"RegisterRequest.Mobile.required":   "手机号不能为空",
	"RegisterRequest.Password.required": "密码不能为空",
}
