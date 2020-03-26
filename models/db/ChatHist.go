package db

import (
	"webchat/models/entity"
)

// FindChatHist は 指定したユーザID(自分と相手)のチャット履歴レコードを取得する
func FindChatHist(myUserID string, opUserID string) []entity.ChatHistEntity {

	chatHists := []entity.ChatHistEntity{}

	// select
	//DbAccess.Where("user_Id = ?", userID).First(&loginUser)

	return chatHists
}

/*
// InsertLoginUser は ログインユーザテーブルにレコードを追加する
func InsertChatHist(chatHistEntity *entity.ChatHistEntity) {

	//chatHistEntity.BaseModel.ExecUserId = chatHistEntity.UserId
	//DbAccess.Create(&loginUser)
}

// DeleteLoginUser は ログインユーザテーブルの指定したレコードを削除する
func DeleteChatHist(userID string) {

	// delete
	//products := []entity.LoginUser{}
	//DbAccess.Delete(&products, userID)
}
*/
