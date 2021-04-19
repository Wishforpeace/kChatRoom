package message

const (
	MsgTypeLogin       = "TypeLogin"       //用户登陆
	MsgTypeLoginRes    = "TypeLoginRes"    //用户登陆返回
	MsgTypeRegister    = "TypeRegister"    //用户注册
	MsgTypeRegisterRes = "TypeRegisterRes" //用户注册返回
	MsgTypeOnline      = "TypeOnline"      //用户上线
	MsgTypeOnlineRes   = "TypeOnlineRes"   //用户上线通知
	MsgTypeSms         = "TypeSms"         //群发
	MsgTypeSmsOne      = "TypeSmsOne"      //私发
)

type Message struct {
	Code int `json:"code"`
}
