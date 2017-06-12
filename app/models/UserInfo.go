package models

import (
	_ "github.com/google/uuid"
)

type UserInfo struct {
	Id            string `json:"id" xorm:"pk 'id'"`
	LoginName     string `json:"loginName" xorm:"varchar(64) notnull 'login_name'"`
	LoginPassword string `json:"loginPassword" xorm:"varchar(128) notnull 'login_password'"`
	Salt          string `json:"salt" xorm:"varchar(123) notnull 'salt'"`

	UserName string `json:"userName" xorm:"varchar(64) notnull 'user_name'"`

	Email string `json:"email" xorm:"varchar(128) 'email'"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

var DefaultUserInfo = UserInfo{
	Id:            "init",
	LoginName:     "anyone",
	LoginPassword: "9fafdfffa1222faddec96c954ed67618aa78808842c8753d6d788e77258ada82",
	Salt:          "init",
	UserName:      "anyone",
	Email:         "anyone@cali.io",
}