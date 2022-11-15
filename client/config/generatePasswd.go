package config

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

var (
	config Config
	maxLen = 100
)

func readConfig() {
	viper.AddConfigPath("../")
	// viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("------ReadInConfig error", err)
		panic(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println("------Unmarshal error", err)
		panic(err)
	}

	// BeginPort = config.BeginPort
	// PasswordArr = len(config.PasswordArr)
}

func getPasswd() {
	readConfig()
	for i := 0; i < maxLen; i++ {
		newPort := config.BeginPort + i
		_str := "xztProxy" + strconv.Itoa(newPort)
		bytes := sha256.Sum256([]byte(_str))     //计算哈希值，返回一个长度为32的数组
		hashCode := hex.EncodeToString(bytes[:]) //将数组转换成切片
		fmt.Println("  -", hashCode)
	}
}
