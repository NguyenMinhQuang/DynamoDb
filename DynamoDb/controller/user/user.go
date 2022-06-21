package usercontroller

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/dynamodb/entities/user"

	"github.com/dynamodb/repository/database"
	"github.com/google/uuid"
)

type UserController struct {
	repository database.IDatabase
}

type IUserController interface {
	GetOneUser(Id uuid.UUID) (user.User, error)
	GetListUser() ([]user.User, error)
	InsertUser(entity *user.User) error
	UpdatetUser(id uuid.UUID, entity *user.User) error
	DeleteUser(id uuid.UUID) error
}

func NewUserController(repository database.IDatabase) IUserController {
	return &UserController{
		repository: repository,
	}
}

func (c *UserController) GetOneUser(Id uuid.UUID) (user.User, error) {
	u := user.User{}
	u.ID = Id

	reponse, err := c.repository.FindOne(u.GetFilterId(), u.TableName())
	// error when get user
	if err != nil {
		return u, err
	}
	// Could not found
	if reponse.Item == nil {
		return u, fmt.Errorf("Could not find user id: %s", Id)
	}
	uR := user.User{}
	err = dynamodbattribute.UnmarshalMap(reponse.Item, &uR)
	if err != nil {
		return u, fmt.Errorf("Failed to unmarshal Record, %v", err)
	}

	if u.ID != uR.ID {
		return u, fmt.Errorf("Error when find user id: %s", Id)
	}
	return uR, nil
}
func (c *UserController) GetListUser() ([]user.User, error) {
	return []user.User{}, nil
}
func (c *UserController) InsertUser(entity *user.User) error {
	entity.CreatedAt = time.Now()
	_, err := c.repository.CreateOrUpdate(entity.UserInfoCreate(), entity.TableName())
	if err != nil {
		return fmt.Errorf("Error create user id: %s", entity.ID)
	}
	return nil
}
func (c *UserController) UpdatetUser(id uuid.UUID, entity *user.User) error {
	return nil
}
func (c *UserController) DeleteUser(id uuid.UUID) error {
	return nil
}
