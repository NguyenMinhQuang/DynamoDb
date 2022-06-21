package user

import "github.com/dynamodb/entities"

type User struct {
	entities.Base
	UserName string `json:"username"`
	Address  string `json:"address"`
}

func (u *User) GetFilterId() map[string]interface{} {
	return map[string]interface{}{"_id": u.ID.String()}
}

func (u *User) TableName() string {
	return "User"
}

func (u *User) UserInfoCreate() map[string]interface{} {
	return map[string]interface{}{
		"_id":       u.ID.String(),
		"username":  u.UserName,
		"address":   u.Address,
		"createdAt": u.CreatedAt.Format(entities.GetTimeFormat()),
		"updatedAt": u.UpdatedAt.Format(entities.GetTimeFormat()),
	}
}
