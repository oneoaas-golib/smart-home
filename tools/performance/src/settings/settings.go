package settings

import (
	"os"
	"time"
	"github.com/astaxie/beego/config"
	"fmt"
	"strconv"
	"io/ioutil"
	"path/filepath"
	"log"
)

const (
	CONF_NAME string = "performance.conf"
	APP_NAME string = "performance"
	APP_MAJOR = 0
	APP_MINOR = 1
	APP_PATCH = 0
	permMode os.FileMode = 0666
)

// Singleton
var instantiated *Settings = nil
var iter uint8

func SettingsPtr() *Settings {
	if instantiated == nil {
		instantiated = new(Settings)
	}
	return instantiated
}

type Settings struct {
	Iterations	int
	DeviceList	[]string
	StartTime	time.Time
	UpTime		time.Duration
	Baud		int
	Timeout		time.Duration
	Cachetime	time.Duration
	StopBits	int
	Command		string
	cfg 		config.IniConfigContainer
	dir		string
}

func (s *Settings) Init() *Settings {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	s.StartTime = time.Now()
	s.dir = fmt.Sprintf("%s/", dir)
	s.Baud = 19200
	s.Timeout = 500
	s.Cachetime = 3600
	s.StopBits = 2
	s.Iterations = 1000000
	s.Command = "010300000005"

	s.Load()

	return s
}

func (s *Settings) Load() (*Settings, error) {

	iter++

	if _, err := os.Stat(s.dir + CONF_NAME); os.IsNotExist(err) {
		return s.Save()
	}

	// read config file
	cfg, err := config.NewConfig("ini", s.dir + CONF_NAME)
	if err != nil {
		return s, err
	}

	if cfg.String("app_version") != s.AppVresion() {
		s.Save()
		if iter > 2 {
			return s, fmt.Errorf("Не известная ошибка")
		}

		return s.Load()
	}

	s.Baud, _ = cfg.Int("baud")
	timeout, _ := cfg.Int("timeout")
	s.Timeout = time.Duration(timeout)
	cachetime, _ := cfg.Int("cachetime")
	s.Cachetime = time.Duration(cachetime)
	s.StopBits, _ = cfg.Int("stopbits")
	s.Iterations, _ = cfg.Int("iterations")
	s.Command = cfg.String("command")

	return s, err
}

func (s *Settings) Save() (*Settings, error) {

	if _, err := os.Stat(s.dir + CONF_NAME); os.IsNotExist(err) {
		ioutil.WriteFile(s.dir + CONF_NAME, []byte{}, permMode)
	}

	cfg, err := config.NewConfig("ini", s.dir + CONF_NAME)
	if err != nil {
		return s, err
	}

	cfg.Set("app_version", s.AppVresion())
	cfg.Set("baud", strconv.Itoa(s.Baud))
	cfg.Set("timeout", strconv.Itoa(int(s.Timeout)))
	cfg.Set("cachetime", strconv.Itoa(int(s.Cachetime)))
	cfg.Set("stopbits", strconv.Itoa(s.StopBits))
	cfg.Set("iterations", strconv.Itoa(s.Iterations))
	cfg.Set("command", s.Command)

	if err := cfg.SaveConfigFile(s.dir + CONF_NAME); err != nil {
		fmt.Printf("err with create conf file: %s\n", s.dir + CONF_NAME)
		return s, err
	}

	return s, nil
}

func (s *Settings) AppVresion() string {
	return fmt.Sprintf("%d.%d.%d", APP_MAJOR, APP_MINOR, APP_PATCH)
}

func init() {
	instantiated = new(Settings)
}