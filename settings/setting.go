package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = new(ConfigStruct)

type ConfigStruct struct {
	*AppConfig       `mapstructure:"app"`
	*AuthConfig      `mapstructure:"auth"`
	*LogConfig       `mapstructure:"log"`
	*MySQLConfig     `mapstructure:"mysql"`
	*RedisConfig     `mapstructure:"redis"`
	*SnowflakeConfig `mapstructure:"snowflake"`
}

type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
}

type AuthConfig struct {
	TokenExpireDurationNum  int64  `mapstructure:"token_expire_duration_num"`
	TokenExpireDurationUnit string `mapstructure:"token_expire_duration_unit"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

type SnowflakeConfig struct {
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

func Init() (err error) {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")

	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("Fatal error viper.ReadInConfig: %s \n", err)
		return
	}

	if err = viper.Unmarshal(Config); err != nil {
		fmt.Printf("Fatal error viper.Unmarshal: %s \n", err)
	}
	PrintConfig()
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Config); err != nil {
			fmt.Printf("Fatal error viper.Unmarshal: %s \n", err)
		}
		PrintConfig()
	})
	return
}

func PrintConfig() {
	fmt.Printf("load config: app=%#v, auth=%#v, log=%#v, mysql=%#v, redis=%#v,snowflake=%#v \n",
		Config.AppConfig, Config.AuthConfig, Config.LogConfig, Config.MySQLConfig, Config.RedisConfig, Config.SnowflakeConfig)
}
