package main

import (
	"fmt"
	"kiga-tool/cmd"
	"kiga-tool/pkg/mysql"
	"kiga-tool/utils"
	"log"
	"runtime"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

func main() {
	cpu := 4
	if runtime.NumCPU()/2 > 4 {
		cpu = runtime.NumCPU() / 2
	} else if runtime.NumCPU() < 4 {
		cpu = runtime.NumCPU()
	}
	runtime.GOMAXPROCS(cpu)

	go DoMonitor(15, func(stat *MonitorStat) {
		fmt.Printf("==== stat %s === \nalloc:\t%d\ttotal:\t%d\tsys:\t%d\tmallocs: %d\tfrees:\t%d\npause:\t%d\tgc:\t%d\tgoroutine: %d\nheap: %d / %d / %d\tstack: %d\n",
			time.Now().String(),
			stat.Alloc, stat.TotalAlloc, stat.Sys, stat.Mallocs, stat.Frees,
			stat.PauseTotalNs, stat.NumGC, stat.NumGoroutine,
			stat.HeapIdle, stat.HeapInuse, stat.HeapReleased, stat.StackInuse,
		)
	})

	viper.SetConfigName("kiga-tool.toml") //配置文件名字
	viper.SetConfigType("toml")           //配置文件类型
	viper.AddConfigPath(".")              //配置文件搜索路径

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: ", err)
	}

	fmt.Println(mysql.GetMysqlConfig())
	spew.Config = *spew.NewDefaultConfig()
	spew.Config.ContinueOnMethod = true

	cmd.AppName = "Message"
	cmd.AppVersion = "1.0.1"
	cmd.Execute()

	fmt.Println(utils.GetAllIPv4Address())
	select {}
}

// MonitorStat is the state of the runtime
type MonitorStat struct {
	runtime.MemStats
	LiveObjects  uint64 `json:"live_objects,omitempty"`  // Live objects = Mallocs - Frees
	NumGoroutine int    `json:"num_goroutine,omitempty"` // Number of goroutines
}

// DoMonitor start a loop for monitor
func DoMonitor(duration int, callback func(*MonitorStat)) {
	interval := time.Duration(duration) * time.Second
	timer := time.Tick(interval)
	for range timer {
		var rtm runtime.MemStats
		runtime.ReadMemStats(&rtm)
		callback(&MonitorStat{
			MemStats:     rtm,
			NumGoroutine: runtime.NumGoroutine(),
			LiveObjects:  rtm.Mallocs - rtm.Frees,
		})
	}
}
