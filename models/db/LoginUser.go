package db

import (
	"webchat/models/entity"
)

// FindAllLoginUsers は ログインユーザテーブルのレコードを全件取得する
func FindAllLoginUsers() []entity.LoginUser {

	loginUsers := []entity.LoginUser{}

	// select all
	DbAccess.Order("user_name asc").Find(&loginUsers)

	return loginUsers
}

// FindLoginUser は ログインユーザテーブルのレコードを全件取得する
func FindLoginUser(userID string) entity.LoginUser {

	loginUser := entity.LoginUser{}

	// select
	DbAccess.Where("user_Id = ?", userID).First(&loginUser)

	return loginUser
}

// InsertLoginUser は ログインユーザテーブルにレコードを追加する
func InsertLoginUser(loginUser *entity.LoginUser) {

	loginUser.BaseModel.ExecUserID = loginUser.UserID
	DbAccess.Create(&loginUser)
}

// UpdateLoginUser は ログインユーザテーブルの指定したレコードを変更する
func UpdateLoginUser(loginUser *entity.LoginUser) {

	// update
	loginUsers := []entity.LoginUser{}
	DbAccess.Model(&loginUsers).Where("user_id = ?", loginUser.UserID).Update("password", loginUser.Password, "user_name", loginUser.UserName, "role_id", "01")
}

// DeleteLoginUser は ログインユーザテーブルの指定したレコードを削除する
func DeleteLoginUser(userID string) {

	// delete
	products := []entity.LoginUser{}
	DbAccess.Delete(&products, userID)
}
