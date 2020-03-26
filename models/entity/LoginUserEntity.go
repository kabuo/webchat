package entity

// LoginUser モデル [tbl_login_users]
type LoginUser struct {
	// 共通
	BaseModel

	// UserId ユーザID
	UserID string `gorm:"primary_key;not null"       json:"userId" column:"user_id"`
	// Password パスワード
	Password string `json:"password" column:"password"`
	// UserName ユーザ名
	UserName string `json:"userName" column:"user_name"`
	// RoleId 権限ID
	RoleID string `json:"roleId" column:"role_id"`
	// Icon アイコン
	Icon string `json:"icon" column:"icon"`
}

// TableName テーブル名 [tbl_login_users]
func (LoginUser) TableName() string {
	return "tbl_login_users"
}
