package config

import "github.com/caarlos0/env/v6"

type WebConfig struct {
	Host string `json:"host" env:"WEB_HOST" envDefault:""`
	Port string `json:"port" env:"WEB_PORT" envDefault:"8080"`
	//Access-Control-Allow-Originの設定
	AllowOrigin string `json:"allow_origin" env:"WEB_ALLOW_ORIGIN" envDefault:"*"`
}

type DBConfig struct {
	DBType string `json:"db_type" env:"DB_TYPE" envDefault:"sqlite"`
	DBFile string `json:"db_file" env:"DB_FILE" envDefault:"test.db"`
}

type JWTConfig struct {
	LifeTimeHour int    `env:"JWT_TOKEN_HOUR_LIFESPAN" envDefault:"24"`    //トークンの有効期限 時間
	Secret       string `env:"JWT_API_SECRET" envDefault:"auth-go-sample"` //トークンのシークレットキー 任意の文字列
	Pepper       string `env:"AUTH_PEPPER" envDefault:"auth-go-sample"`
}

var Web WebConfig
var DB DBConfig
var JWT JWTConfig

// 設定の初期化
func Init() {
	Web = WebConfig{}
	if err := env.Parse(&Web); err != nil {
		panic(err)
	}
	DB = DBConfig{}
	if err := env.Parse(&DB); err != nil {
		panic(err)
	}
	JWT = JWTConfig{}
	if err := env.Parse(&JWT); err != nil {
		panic(err)
	}
}
