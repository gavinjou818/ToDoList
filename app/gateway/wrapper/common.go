package wrapper

import (
	"sync"
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

// 其实就是以前 hystrix 的配置
type Config struct {
	Namespace              string
	Timeout                time.Duration
	MaxConcurrentRequests  int           // 最大并发数
	RequestVolumeThreshold uint64        // 水位
	SleepWindow            time.Duration // 滑动窗口
	ErrorPercentThreshold  int           // 错误比率
}

type Group struct {
	mu        sync.RWMutex
	namespace string
	settings  map[string]bool
	conf      *Config
}

var (
	_mu   sync.RWMutex
	_conf = &Config{
		Namespace:              "default",
		Timeout:                time.Duration(hystrix.DefaultTimeout),
		MaxConcurrentRequests:  hystrix.DefaultMaxConcurrent,
		RequestVolumeThreshold: uint64(hystrix.DefaultVolumeThreshold),
		SleepWindow:            time.Duration(hystrix.DefaultSleepWindow),
		ErrorPercentThreshold:  hystrix.DefaultErrorPercentThreshold,
	}
)

func (conf *Config) fix() {
	if conf.Namespace == "" {
		conf.Namespace = "default"
	}
	if conf.Timeout <= 0 {
		conf.Timeout = time.Duration(hystrix.DefaultTimeout)
	}
	if conf.MaxConcurrentRequests <= 0 {
		conf.MaxConcurrentRequests = hystrix.DefaultMaxConcurrent
	}
	if conf.RequestVolumeThreshold <= 0 {
		conf.RequestVolumeThreshold = uint64(hystrix.DefaultVolumeThreshold)
	}
	if conf.SleepWindow == 0 {
		conf.SleepWindow = time.Duration(hystrix.DefaultSleepWindow)
	}
	if conf.ErrorPercentThreshold <= 0 {
		conf.ErrorPercentThreshold = hystrix.DefaultErrorPercentThreshold
	}
}

func NewGroup(conf *Config) *Group {
	if conf == nil {
		// 加读锁是因为 _conf 允许你读，但是有写操作的时候就不允许读了
		_mu.RLock()
		conf = _conf
		_mu.RUnlock()
	} else {
		conf.fix()
	}
	return &Group{
		namespace: conf.Namespace,
		settings:  make(map[string]bool),
		conf:      conf,
	}
}

func (g *Group) Reload(conf *Config) {
	if conf == nil {
		return
	}
	conf.fix()
	g.mu.Lock()
	g.conf = conf
	g.mu.Unlock()
}

// 触发之后运行的动作
func (g *Group) Do(name string, run func() error) (err error) {
	name = g.namespace + "-" + name
	g.setBreakerConfig(name)
	return hystrix.Do(name, func() error {
		return run()
	}, nil)
}

func (g *Group) setBreakerConfig(name string) {
	if _, ok := g.settings[name]; !ok {
		hystrix.ConfigureCommand(name, hystrix.CommandConfig{
			Timeout:                int(time.Duration(g.conf.Timeout) / time.Millisecond),
			MaxConcurrentRequests:  g.conf.MaxConcurrentRequests,
			RequestVolumeThreshold: int(g.conf.RequestVolumeThreshold),
			SleepWindow:            int(time.Duration(g.conf.SleepWindow) / time.Millisecond),
			ErrorPercentThreshold:  g.conf.ErrorPercentThreshold,
		})

		copy := make(map[string]bool)
		for key, val := range g.settings {
			copy[key] = val
		}
		copy[name] = true
		g.settings = copy
	}
}
