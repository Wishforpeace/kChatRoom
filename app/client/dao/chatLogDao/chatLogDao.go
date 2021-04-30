package chatLogDao

import (
	"gorm.io/gorm"
	"kChatRoom/app/client/model/chatLogModel"
	"kChatRoom/utils/db"
)

type chatLogDao struct {
	Db *gorm.DB
}

//NewChatLogDao new chat log dao
func NewChatLogDao() *chatLogDao {
	return &chatLogDao{
		Db: db.NewDB(),
	}
}

//SaveLog 保存聊天记录
func (t *chatLogDao) SaveLog(log *chatLogModel.ChatLogModel) error {
	Db := t.Db
	res := Db.Create(log)
	return res.Error
}

type ChatLogView struct {
	chatLogModel.ChatLogModel
	UserName string `json:"username"  gorm:"index;column:username" form:"username"`
}

//GetChatLog 获取聊天记录
//pageSize limit
//page  页码
func (t *chatLogDao) GetChatLog(page, pageSize int) []*ChatLogView {
	Db := t.Db
	logs := make([]*ChatLogView, 0)
	Db = Db.Table("`c_chat_log` as l").Joins("left join `c_users` as u on l.mail = u.mail").Select("l.*,u.username")

	//分页操作
	if page > 0 && pageSize > 0 {
		Db = Db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	_ = Db.Order("l.`id` desc").Find(&logs)
	return logs
}
