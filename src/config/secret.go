package config

import (
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Secret struct {
	App struct {
		Name     string `yaml:"name"`
		Url      string `yaml:"url"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		TimeZone string `yaml:"timezone"`
	} `yaml:"app"`
	Db struct {
		SourcePgSql string `yaml:"sourcepgsql"`
		SourceMySql string `yaml:"sourcemysql"`
	} `yaml:"db"`
	Redis struct {
		Url string `yaml:"url"`
	} `yaml:"redis"`
	DragonPay struct {
		Url        string `yaml:"url"`
		MerchantId string `yaml:"merchantid"`
		Password   string `yaml:"password"`
		Postback   string `yaml:"postback"`
	} `yaml:"dragonpay"`
	JazzCash struct {
		Url            string `yaml:"url"`
		MerchantId     string `yaml:"merchantid"`
		Password       string `yaml:"password"`
		IntegeritySalt string `yaml:"integeritysalt"`
	} `yaml:"jazzcash"`
	Midtrans struct {
		Url        string `yaml:"url"`
		MerchantId string `yaml:"merchantid"`
		ClientKey  string `yaml:"clientkey"`
		ServerKey  string `yaml:"serverkey"`
	} `yaml:"midtrans"`
	Momo struct {
		Url         string `yaml:"url"`
		PartnerCode string `yaml:"partnercode"`
		AccessKey   string `yaml:"accesskey"`
		SecretKey   string `yaml:"secretkey"`
	} `yaml:"momo"`
	Nicepay struct {
		Url         string `yaml:"url"`
		MerchantId  string `yaml:"merchantid"`
		MerchantKey string `yaml:"merchantkey"`
	} `yaml:"nicepay"`
	Razer struct {
		Url        string `yaml:"url"`
		UrlApi     string `yaml:"urlApi"`
		MerchantId string `yaml:"merchantid"`
		VerifyKey  string `yaml:"verifykey"`
		SecretKey  string `yaml:"secretkey"`
	} `yaml:"razer"`
	Ximpay struct {
		UrlTsel   string `yaml:"urlTsel"`
		UrlHti    string `yaml:"urlHti"`
		UrlXl     string `yaml:"urlXl"`
		UrlIsat   string `yaml:"urlIsat"`
		UrlSf     string `yaml:"urlSf"`
		PartnerId string `yaml:"partnerId"`
		SecretKey string `yaml:"secretKey"`
		Username  string `yaml:"username"`
	} `yaml:"ximpay"`
	Log struct {
		Path string `yaml:"path"`
	}
}

func LoadSecret(path string) (*Secret, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadSecretFromBytes(data)
}

func LoadSecretFromBytes(data []byte) (*Secret, error) {
	fang := viper.New()
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()
	fang.SetEnvPrefix("GO")
	fang.SetConfigType("yaml")

	if err := fang.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return nil, err
	}
	var creds Secret
	err := fang.Unmarshal(&creds)
	if err != nil {
		log.Fatalf("Error loading creds: %v", err)
	}
	return &creds, nil
}
