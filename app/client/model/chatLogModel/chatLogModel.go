package chatLogModel

import "time"

//ChatLogModel 聊天记录
type ChatLogModel struct {
	ID        uint      `json:"id" gorm:"primaryKey column:id" form:"id"`
	Type      string    `json:"type"  gorm:"column:type" form:"type"`
	Mail      string    `json:"mail" gorm:"column:mail" form:"mail"`
	Msg       string    `json:"msg" gorm:"column:msg"`
	Head      string    `json:"head" gorm:"column:head" form:"head"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" `
}

// TableName  自定义表明
func (t *ChatLogModel) TableName() string {
	return "c_chat_log"
}
