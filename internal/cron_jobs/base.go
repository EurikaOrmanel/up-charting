package cronjobs

import (
	"EurikaOrmanel/up-charter/config"
	"fmt"

	"github.com/robfig/cron/v3"
)

func ConfigJobs(appConfig config.AppConfig) {
	c := cron.New()
	cronPeriod := "0 0 0 * * *"
	cronPeriod = "*/1 * * * *"
	c.AddFunc(cronPeriod, func() {
		fmt.Println("Ah")
		checkAllSongsChart(appConfig)
	})
	c.Start()
}
