package userModel

import "time"

//UserModel 用户信息
type UserModel struct {
	ID        uint      `json:"id" gorm:"primaryKey column:id" form:"id"`
	UserName  string    `json:"username"  gorm:"index;column:username" form:"username"`
	Password  string    `json:"password" gorm:"column:password" form:"password"`
	Mail      string    `json:"mail" gorm:"column:mail" form:"mail"`
	LoginIp   string    `json:"login_ip" gorm:"column:login_ip"`
	LoginNum  int64     `json:"login_num" gorm:"column:login_num"`
	Desc      string    `json:"desc"  gorm:"column:desc" form:"desc"`
	Status    byte      `json:"status" gorm:"index;column:status" form:"status"`
	Head      string    `json:"head" gorm:"column:head" form:"head"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" `
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" `
}

// TableName  自定义表明
func (t *UserModel) TableName() string {
	return "c_users"
}
