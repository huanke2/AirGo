package initialize

import (
	"AirGo/global"
	"AirGo/service"
	"github.com/robfig/cron/v3"
)

// 秒级操作
// cron.WithSeconds()
// 函数没执行完就跳过本次函数
// cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger))
// 打印任务日志
// cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags)))

// c := cron.New(cron.WithSeconds(), cron.WithChain(cron.SkipIfStillRunning(cron.DefaultLogger)), cron.WithLogger(
// 	cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))

// 添加一个任务
//_, err := c.AddFunc("*/5 * * * * *", func() {
// 	fmt.Println("开始定时任务")
// })

// 开始执行（每个任务会在自己的 goroutine 中执行）
//c.Start()

// 允许往正在执行的 cron 中添加任务
//c.AddFunc("@daily", func() { fmt.Println("Every day") })

// 检查上一个和下一个任务执行的时间
//inspect(c.Entries())

//c.Stop()  // 停止调度，但正在运行的作业不会被停止

func InitCrontab() {
	global.Logrus.Info("用户流量有效期定时任务")
	UserCrontab()
}

//用户流量，有效期 任务

func UserCrontab() {
	c := cron.New()
	_, err := c.AddFunc("*/2 * * * *", func() {
		err := service.UserExpiryCheck()
		if err != nil {
			global.Logrus.Error("service.UserExpiryCheck error:", err)
		}
	})
	if err != nil {
		return
	}
	c.Start()
}

// casbin 更新
