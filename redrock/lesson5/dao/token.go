package dao

import "lesson5/model"

func StoreRefreshToken(userID uint, tokenString string, expiresAT int64) error {
	userToken := model.UserToken{
		UserID:    userID,
		Token:     tokenString,
		ExpiresAt: expiresAT,
		Revoked:   false,
	}
	return DB.Create(&userToken).Error
}

func CheckRefreshToken(tokenString string) (bool, error) {
	var userToekn model.UserToken
	err := DB.Where("token=?", tokenString).First(&userToekn).Error
	if err != nil {
		return false, err
	}
	if userToekn.Revoked == true {
		return false, nil
	}
	return true, nil
}

func RevokeToken(tokenString string) error {
	return DB.Model(&model.UserToken{}).Where("token = ?", tokenString).Update("revoked", true).Error
}
