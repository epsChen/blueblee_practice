package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"github.com/epsChen/bluebell/models"
	"github.com/epsChen/bluebell/utils"
	"github.com/go-sql-driver/mysql"
)

const secret = "secret"

// 疑点1 first返回值到底是不是查询数据
// 疑点2 这样查询语句和逻辑对吗
// 可以参考之前的demo
//func CheckUser(username string) error {
//	//根据用户名查找数据 如果不为空 就说明存在 返回错误
//	user := &models.User{}
//	result := db.Where("user_name = ?", username).First(user)
//	if result != nil {
//		return errors.New("用户已存在")
//	}
//
//	return nil
//}

// 当用户名一致时会直接返回错误
func InsertUser(p *models.ParamSignUp) (err error) {
	//2.为用户生成一个id
	uid := utils.GenID()
	password := encryptPassword(p.Password)
	user := &models.User{
		UserID:   uid,
		UserName: p.Username,
		Password: password,
		Gender:   0,                //先都默认为女
		Email:    sql.NullString{}, //先都默认为空
	}
	result := db.Create(&user)
	if result.Error != nil {
		//对错误进行判断 如果是mysql错误 并且错误码为1062（唯一性约束）
		//返回用户已存在
		if sqlError, ok := result.Error.(*mysql.MySQLError); ok && sqlError.Number == 1062 {
			err = models.ErrUserExist
			return
		}
		err = result.Error
		return
	}
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
