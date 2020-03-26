package controllers

import (
	"webchat/models/db"
	"webchat/models/entity"

	// Gin
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// FindAllLoginUsers は 全てのログインユーザ情報を取得する
func FindAllLoginUsers(c *gin.Context) {

	result := db.FindAllLoginUsers()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, result)
}

// FindLoginUser は 指定したユーザIDのログインユーザ情報を取得する
func FindLoginUser(c *gin.Context) {

	userID := c.Query("userId")
	result := db.FindLoginUser(userID)

	if result.Password != "" {
		LoginSession(c, result)
	}

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, result)
}

// AddLoginUser は ログインユーザ情報をDBへ登録する
func AddLoginUser(c *gin.Context) {

	userID := c.PostForm("userId")
	password := c.PostForm("password")
	userName := c.PostForm("userName")

	var model = entity.LoginUser{
		UserID:   userID,
		Password: password,
		UserName: userName,
		RoleID:   "01",
	}

	db.InsertLoginUser(&model)
}

// ChangeLoginUser は ログインユーザ情報を変更する
func ChangeLoginUser(c *gin.Context) {

	userID := c.PostForm("userId")
	password := c.PostForm("password")
	userName := c.PostForm("userName")
	roleID := c.PostForm("roleId")

	var model = entity.LoginUser{
		UserID:   userID,
		Password: password,
		UserName: userName,
		RoleID:   roleID,
	}

	db.UpdateLoginUser(&model)
}

// DeleteLoginUser は ログインユーザ情報をDBから削除する
func DeleteLoginUser(c *gin.Context) {

	userID := c.PostForm("userId")
	db.DeleteLoginUser(userID)
}

// LoginSession ログインセッション情報を保持する(ログイン)
func LoginSession(c *gin.Context, user entity.LoginUser) {

	/*var model = entity.SessionInfo{
		UserId:   user.UserId,
		UserName: user.UserName,
		IsAlive:  true,
	}*/

	session := sessions.Default(c)
	session.Set("alive", true)
	session.Set("userID", user.UserID)
	session.Set("password", user.Password)
	session.Set("userName", user.UserName)
	session.Set("roleId", user.RoleID)
	session.Save()
}

// ClearSession セッション情報をクリアする(ログアウト)
func ClearSession(c *gin.Context) {

	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
