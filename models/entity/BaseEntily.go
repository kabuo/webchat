package entity

import "time"

// BaseModel モデル [共通]
type BaseModel struct {
	// DelFlg 削除フラグ
	DelFlg bool `default:"false" column:"del_flg"`
	// ExecDate 実施日時
	ExecDate time.Time `gorm:"-" column:"exec_date"`
	// ExecUserID 実施ユーザ
	ExecUserID string `column:"exec_user_id"`
}
