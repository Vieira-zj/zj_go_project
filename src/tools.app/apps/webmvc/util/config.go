package util

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/unknwon/goconfig"
)

// Configs 项目配置
type Configs struct {
	filePath string
	cfg      *goconfig.ConfigFile
	cfgmap   map[string]map[string]string
}

// NewConfigs 项目配置实例
func NewConfigs() *Configs {
	return &Configs{
		cfgmap: make(map[string]map[string]string),
	}
}

// Parse 解析配置
func (config *Configs) Parse(fpath string) (map[string]map[string]string, error) {
	config.filePath = fpath
	cfg, err := goconfig.LoadConfigFile(fpath)
	for _, sec := range cfg.GetSectionList() {
		config.cfgmap[sec] = make(map[string]string)
		for _, key := range cfg.GetKeyList(sec) {
			config.cfgmap[sec][key], _ = cfg.GetValue(sec, key)
		}
	}
	return config.cfgmap, err
}

// GetAllCfg 返回全部配置
func (config *Configs) GetAllCfg() (c map[string]map[string]string) {
	return config.cfgmap
}

// ReloadAllCfg 刷新配置文件
func (config *Configs) ReloadAllCfg() (c map[string]map[string]string, err error) {
	return config.Parse(config.filePath)
}

// GetSection 返回配置项
func (config *Configs) GetSection(sec string) map[string]string {
	return config.cfgmap[sec]
}

// WatchConfig 监听事件, 自动刷新配置
func (config *Configs) WatchConfig() {
	go func() {
		watch, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
			return
		}
		defer watch.Close()

		if err = watch.Add(config.filePath); err != nil {
			log.Fatal(err)
			return
		}
		log.Println("WatchConfig:", config.filePath)

		for {
			select {
			case event := <-watch.Events:
				{
					if event.Op&fsnotify.Write == fsnotify.Write {
						config.ReloadAllCfg()
					}
					log.Printf("WatchConfig Op=%v, Name=%s\n", event.Op, event.Name)
				}
			case err := <-watch.Errors:
				{
					log.Fatal(err)
					return
				}
			}
		}
	}()
}
