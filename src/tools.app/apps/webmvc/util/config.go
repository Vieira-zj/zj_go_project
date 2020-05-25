package util

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/unknwon/goconfig"
)

type configMap map[string]map[string]string

// Configs 项目配置
type Configs struct {
	filePath string
	cfgmap   configMap
}

// Parse 解析配置
func (config *Configs) Parse(fpath string) error {
	config.filePath = fpath
	cfg, err := goconfig.LoadConfigFile(fpath)
	if err != nil {
		return err
	}

	config.cfgmap = make(configMap)
	for _, sec := range cfg.GetSectionList() {
		config.cfgmap[sec] = make(map[string]string)
		for _, key := range cfg.GetKeyList(sec) {
			config.cfgmap[sec][key], _ = cfg.GetValue(sec, key)
		}
	}
	return nil
}

// GetAllCfg 返回全部配置
func (config *Configs) GetAllCfg() (c map[string]map[string]string) {
	return config.cfgmap
}

// GetSection 返回配置项
func (config *Configs) GetSection(sec string) map[string]string {
	return config.cfgmap[sec]
}

// ReloadAllCfg 刷新配置文件
func (config *Configs) ReloadAllCfg() error {
	return config.Parse(config.filePath)
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
