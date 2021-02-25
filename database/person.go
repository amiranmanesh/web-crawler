package database

import (
	"github.com/jinzhu/gorm"
	b64 "encoding/base64"
)

type Person struct {
	gorm.Model
	FName           string `gorm:"type:varchar(100);" json:"first_name"`
	LName           string `gorm:"type:varchar(100);" json:"last_name"`
	Phone           string `gorm:"type:varchar(100);" json:"phone"`
	Address         string `gorm:"type:varchar(5000);" json:"address"`
	AddressGoogle   string `gorm:"type:varchar(5000);" json:"address_google"`
	AddressLocality string `gorm:"type:varchar(500);" json:"address_locality"`
}

func (p *Person) Save() error {

	if err := DB.GetDB().Create(&p).Error; err != nil {
		p.Address = b64.StdEncoding.EncodeToString([]byte(p.Address))
		p.AddressGoogle = b64.StdEncoding.EncodeToString([]byte(p.AddressGoogle))
		p.AddressLocality = b64.StdEncoding.EncodeToString([]byte(p.AddressLocality))
		if err := DB.GetDB().Create(&p).Error; err != nil {
			return err
		}
	}

	return nil

}
