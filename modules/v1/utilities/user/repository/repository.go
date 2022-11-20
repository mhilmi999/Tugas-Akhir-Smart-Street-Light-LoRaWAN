package repository

import (
	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/models"
	"gorm.io/gorm"
)

type Repository interface{
	Save(user models.User) (models.User, error)
	FindByEmail(email string)(models.User, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (n *repository) Save(user models.User) (models.User, error){
	err := n.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (n *repository) FindByEmail(email string)(models.User, error){
	var user models.User
	err := n.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}