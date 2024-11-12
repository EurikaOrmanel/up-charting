package cronjobs

import (
	"EurikaOrmanel/up-charter/config"
	"fmt"

	"github.com/robfig/cron/v3"
)

func ConfigJobs(appConfig config.AppConfig) {
	c := cron.New()
	c.AddFunc("0 0 0 * * *", func() {
		fmt.Println("Ah")
		checkAllSongsChart(appConfig.RepoDb)
	})
	c.Start()
}
