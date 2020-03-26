package entity

// ChatHistEntity モデル
type ChatHistEntity struct {
	// MyFlg 自分の送信かどうか
	MyFlg bool `json:"my_flg" column:"my_flg"`
	// UserID ユーザID
	UserID string `json:"userId" column:"user_id"`
	// UserName ユーザ名
	UserName string `json:"userName" column:"user_name"`
	// SendDate 権限ID
	SendDate string `json:"sendDate" column:"send_date"`
	// MsgStr 送信メッセージ(文言)
	MsgStr string `json:"msgStr" column:"msg_str"`
	// MsgStmp 送信メッセージ(スタンプ)
	MsgStmp string `json:"msgStmp" column:"msg_stmp"`
	// ChkRead 既読有無
	ChkRead string `json:"chkRead" column:"chk_read"`
}
