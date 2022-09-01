package models

import (
	// "errors"

	"github.com/jinzhu/gorm"
)

type Setting struct {
	Id uint64 `json:"id"`
	Identifier string `json:"identifier"`
	Enabled bool `json:"enabled"`
	Value string `json:"value"`
}

var settings = []Setting {
	{
		Id: 12,
		Identifier: "test",
		Enabled: true,
		Value: "this is a test",	
	},
	{
		Id: 22,
		Identifier: "test 22",
		Enabled: true,
		Value: "this is a test 22",	
	},
	
}


func (s *Setting) GetAll() []Setting {
	return settings
}

func (s *Setting) GetSettingById(id uint64) Setting {
	var data Setting

	for _, setting := range settings {
		if setting.Id == id {
			data = setting
		}
	}

	return data
}

func (s *Setting) Create(db *gorm.DB) (*Setting, error) {
	var err error
	err = db.Debug().Create(&s).Error
	if err != nil {
		return &Setting{}, err
	}
	return s, nil
}

func (s *Setting) Update() Setting {
	return settings[0]	
}

func (s *Setting) Delete() Setting {
	return settings[0]
}