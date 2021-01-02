package models

import (
	"errors"
	"fmt"
	"time"

	util "github.com/Komefumi/YourFuncsImaged_API/app/util"
)

var(
	// ErrFailedFuncCreate indicates failure to create Usable Function
	ErrFailedFuncCreate = errors.New("Failed To Create New Func")
	ErrFailedFuncFind = errors.New("Failed to find functions for user")
)

// UsableFunction struct to model UsableFunction
type UsableFunction struct {
	ID uint `gorm:"primaryKey;column:id;autoIncrement"`
	Content string `gorm:"column:content"`
	UserID uint `gorm:"column:user_id"`
	User User
	CreatedAt time.Time `gorm:"column:created_at"`
}

// CreateUsableFunc creates a new UsableFunction
func CreateUsableFunc(userID uint, funcString string) (UsableFunction, error) {
	newFunc := UsableFunction{ UserID: userID, Content: funcString }
	db := util.DBAccessorFunc()
	if dbc:= db.Create(&newFunc); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return newFunc, ErrFailedFuncCreate
	}
	return newFunc, nil
}

// FindUsableFunctionsForUser returns for userID provided, or error if failed
func FindUsableFunctionsForUser(userID uint) ([]UsableFunction, error) {
	var foundFuncs []UsableFunction;
	db := util.DBAccessorFunc()
	db = db.Table("usable_functions").Select("id", "content", "created_at").Where("user_id = ?", userID).Find(&foundFuncs)
	if db.Error != nil {
		return foundFuncs, ErrFailedFuncFind
	}
	return foundFuncs, nil
}