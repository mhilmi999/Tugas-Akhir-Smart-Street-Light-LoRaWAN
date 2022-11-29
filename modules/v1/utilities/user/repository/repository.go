package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mhilmi999/Tugas-Akhir-Smart-Street-Light-LoRaWAN/modules/v1/utilities/user/models"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	// GetId(id string) (int, error)
	GetLatestCon(token string) (models.Received, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (n *repository) Save(user models.User) (models.User, error) {
	err := n.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (n *repository) FindByEmail(email string) (models.User, error) {
	var user models.User
	err := n.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// func (n *repository) GetId(id string) (int, error) {
// 	var idData int
// 	err := n.db.Table("device").Select("device_id").Where("antares_id = ?", id).Scan(&idData).Error
// 	return idData, err
// }

func (n *repository) GetLatestCon(token string) (models.Received, error) {
	data := models.Received{}

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://platform.antares.id:8443/~/antares-cse/antares-id/SmartStreetLight/tugasAkhir/la", nil)
	req.Header.Set("X-M2M-Origin", token)
	req.Header.Set("Content-Type", "application/json;ty=4")
	req.Header.Set("Accept", "application/json")
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	defer resp.Body.Close()
	isiBody, err := ioutil.ReadAll(resp.Body)
	//dataMentah := string(isiBody)
	json.Unmarshal(isiBody, &data)
	fmt.Println(data.First.Con)
	return data, err
}
