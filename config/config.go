/**
* @Author:Tristan
* @Date: 2022/1/4 5:08 下午
 */

package config

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/patrickmn/go-cache"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"log"
	"os"
)

var Conf Config
var DB *gorm.DB
var LogFileDir = Conf.LogDir
var K8sClient *kubernetes.Clientset
var CacheLocal *cache.Cache
var RedisGO *redis.Client

type LogName struct {
	StatLogFile        string `yaml:"stat_log_file"`
	ErrorLogFile       string `yaml:"error_log_file"`
	PanicLogFile       string `yaml:"panic_log_file"`
	StderrPanicLogFile string `yaml:"stderr_panic_log_file"`
}

type MysqlConfig struct {
	Dsn      string `yaml:"dsn"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}

type RedisConfig struct {
	Hostname string `yaml:"hostname"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
}

type Config struct {
	LogDir     string      `yaml:"log_dir"`
	LogName    LogName     `yaml:"log_name"`
	Mysql      MysqlConfig `yaml:"mysql"`
	Redis      RedisConfig `yaml:"redis"`
	TokenCache string      `yaml:"token_cache"`
}

type Option func(*Config)

func AddMysqlDns(addr string) Option {
	return func(o *Config) {
		o.Mysql.Dsn = addr
	}
}

func AddMysqlPass(p string) Option {
	return func(o *Config) {
		o.Mysql.Password = p
	}
}

func AddMysqlUserName(u string) Option {
	return func(o *Config) {
		o.Mysql.Username = u
	}
}

// 获取环境变量
func (conf *Config) getOptionFromEnv() {
	var opts []Option

	mysqlDns := os.Getenv("MYSQL_DNS")
	if len(mysqlDns) > 0 {
		opts = append(opts, AddMysqlDns(mysqlDns))
	}

	mysqlUsername := os.Getenv("MYSQL_USERNAME")
	if len(mysqlDns) > 0 {
		opts = append(opts, AddMysqlUserName(mysqlUsername))
	}

	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	if len(mysqlDns) > 0 {
		opts = append(opts, AddMysqlPass(mysqlPass))
	}

	conf.configure(opts)
}

func (conf *Config) configure(opts []Option) {
	for _, o := range opts {
		o(conf)
	}
}

// 初始化环境变量，先读取 yaml 配置文件，在读取环境变量
func NewConfig(filename string) error {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(buf, &Conf)
	if err != nil {
		return fmt.Errorf("fileName %q err: %v", filename, err)
	}
	Conf.getOptionFromEnv()
	return nil
}

func init() {
	err := NewConfig("./config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
}
