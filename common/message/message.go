package message

const (
	MsgTypeLogin    = "TypeLogin"    //用户登陆
	MsgTypeRegister = "TypeRegister" //用户注册
	MsgTypeOnline   = "TypeOnline"   //用户上线
	MsgTypeSms      = "TypeSms"      //群发
	MsgTypeSmsOne   = "TypeSmsOne"   //私发
)

//用户在线状态
const (
	UserStatusOnline = iota //在线
	UserStatusLeave         //离线
)

// Message 聊天通信消息题
type Message struct {
	Type  string `json:"type"`   //消息类型
	Mail  string `json:"mail"`   //邮箱
	ToUid int    `json:"to_uid"` //接受人
	Msg   string `json:"msg"`    //消息内容
}

// RequestMsg 界面ajax请求结构体
type RequestMsg struct {
	Code int         `json:"code"`
	Res  interface{} `json:"res"`
	Msg  string      `json:"msg"`
}
