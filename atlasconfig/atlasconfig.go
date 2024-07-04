package atlasconfig

import (
	"github.com/spf13/viper"
)

type Option struct {
	TagValue 	string 		`mapstructure:"tagvalue"`
	DisplayText	string		`mapstructure:"displaytext"`
}

type Filter struct {
	TagName 	string 		`mapstructure:"tagname"`
	DisplayText string 		`mapstructure:"displaytext"`
	Options		[]Option 	`mapstructure:"options"`
}

type ATLASConfig struct {
	Filters 	[]Filter 	`mapstructure:"filters"`
}

var Config ATLASConfig

func ReadATLASConfig () error {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err	
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return err
	}

	return nil

}


