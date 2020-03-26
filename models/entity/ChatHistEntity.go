package entity

// ChatHist モデル [tbl_chat_hist]
type ChatHist struct {
	// 共通
	BaseModel

	// UserIDFrom ユーザID(自分)
	UserIDFrom string `gorm:"primary_key;not null"       json:"userIDFrom" column:"user_id_from"`
	// UserIDTo ユーザID(相手)
	UserIDTo string `json:"userIDTo"       column:"user_id_to"`
	// SendDate 送信日時(yyyy-mm-dd-hh-mm-ss.fff)
	SendDate string `json:"sendDate"       column:"send_date"`
	// MsgStr メッセージ文言
	MsgStr string `json:"msgStr"       json:"userIdFrom" column:"msg_str"`
	// MsgStmp メッセージスタンプ
	MsgStmp string `json:"msgStmp"       column:"msg_stmp"`
	// ChkRead 既読有無
	ChkRead string `json:"chkRead" column:"chk_read"`
}

// TableName テーブル名 [tbl_chat_hist]
func (ChatHist) TableName() string {
	return "tbl_chat_hist"
}
