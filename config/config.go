package config

import (
	"fmt"
	"io/fs"
	"reflect"

	"github.com/spf13/viper"
)

var Config = new(AppConfig)

type AppConfig struct {
	Name      string       `mapstructure:"name"`       // 项目名称
	Mode      string       `mapstructure:"mode"`       // 模式
	Verstion  string       `mapstructure:"version"`    // 项目版本信息
	Port      int16        `mapstructure:"port"`       // 项目启动端口
	StartTime string       `mapstructure:"start_time"` // 项目开始时间
	MachineId int64        `mapstructure:"machine_id"` // 雪花算法创建新节点值
	Mysql     *MysqlConfig `mapstructure:"mysql"`
	Redis     *RedisConfig `mapstructure:"redis"`
	Log       *LogConfig   `mapstructure:"log"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`           // 数据库连接地址
	Port         int16  `mapstructure:"port"`           // 数据库连接端口
	User         string `mapstructure:"user"`           // 数据库账户名
	Password     string `mapstructure:"password"`       // 数据库账户密码
	DbName       string `mapstructure:"db_name"`        // 连接数据库名
	MaxOpenConns int32  `mapstructure:"max_open_conns"` // 数据库最大连接数
	MaxIdleConns int32  `mapstructure:"max_idle_conns"` // 数据库最大空闲数
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`      // 数据库连接地址
	Port     int16  `mapstructure:"port"`      // 数据库连接端口
	Password string `mapstructure:"password"`  // 数据库账户密码
	Db       int    `mapstructure:"db"`        // redis 默认选择数据库
	PoolSize int32  `mapstructure:"pool_size"` // redis 最大连接池
}

type LogConfig struct {
	Level      string `mapStructure:"level"`       // 日志等级 zapcore.Level
	Filename   string `mapStructure:"filename"`    // 日志大小
	MaxSize    int    `mapStructure:"max_size"`    // 日志大小
	MaxAge     int    `mapStructure:"max_age"`     // 日志日志存储时长
	MaxBackups int    `mapStructure:"max_backups"` // 日志最多存储文件
	Compress   bool   `mapStructure:"compress"`    // 日志是否压缩
}

// filePath 支持读取指定地址文件
// 如果未识别到指定位置地址，则读取默认地址
// 默认读取项目根目录下 ./config/config
func Init(filePath ...string) (err error) {
	viper.AddConfigPath("./config")

	if len(filePath) > 1 {
		viper.SetConfigFile(filePath[1])
	} else {
		viper.SetConfigName("config")
	}

	// 配置读取失败
	if err = viper.ReadInConfig(); err != nil {
		var pathErr *fs.PathError
		var unsupportedConfigError viper.UnsupportedConfigError

		if (reflect.TypeOf(err).AssignableTo(reflect.TypeOf(pathErr)) || reflect.TypeOf(err).AssignableTo(reflect.TypeOf(unsupportedConfigError))) && len(filePath) > 1 {
			// 重置
			viper.Reset()

			// 读取默认地址
			return Init()
		} else {
			return
		}
	}

	if err = viper.Unmarshal(Config); err != nil {
		fmt.Printf("viper.Unmarshal faild, err:%v\n", err)
		return
	}

	return
}
