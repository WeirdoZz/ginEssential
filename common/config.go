package common

import (
	"github.com/spf13/viper"
)

var (
	ServerSetting *ServerSettingS
	DBSetting     *DBSettingS
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("application")
	vp.SetConfigType("yml")
	vp.AddConfigPath("config/")

	err := vp.ReadInConfig()
	if err != nil {
		panic(err)
	}

	s := &Setting{
		vp: vp,
	}
	return s, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

type ServerSettingS struct {
	Port string
}

type DBSettingS struct {
	DriverName string
	Host       string
	Port       string
	Database   string
	Username   string
	Password   string
	Charset    string
}
