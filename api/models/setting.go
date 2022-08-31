package models

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


func (s *Setting) GetSettings() []Setting {
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