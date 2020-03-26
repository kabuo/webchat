package controllers

import (
	"webchat/models/db"

	// Gin
	"github.com/gin-gonic/gin"
)

// FindChatHist は 指定したユーザID(自分と相手)のチャット履歴情報を取得する
func FindChatHist(c *gin.Context) {

	myUserID := c.Query("myUserId")
	opUserID := c.Query("opUserId")
	result := db.FindChatHist(myUserID, opUserID)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, result)
}

/*
// AddChatHist は 指定したユーザIDのチャット送信情報をDBへ登録する
func AddChatHist(c *gin.Context) {

	userID := c.PostForm("userId")
	password := c.PostForm("password")
	userName := c.PostForm("userName")

	var model = entity.LoginUser{
		UserId:   userID,
		Password: password,
		UserName: userName,
		RoleId:   "01",
	}

	db.InsertLoginUser(&model)
}

// DeleteChatHist は ユーザIDのチャット送信情報をDBから削除する
func DeleteChatHist(c *gin.Context) {

	userID := c.PostForm("userId")
	db.DeleteLoginUser(userID)
}
*/
