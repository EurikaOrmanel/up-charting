package cronjobs

import (
	"EurikaOrmanel/up-charter/config"

	"github.com/robfig/cron/v3"
)

func ConfigJobs(appConfig config.AppConfig) {
	c := cron.New()
	c.AddFunc("0 0 0 * * *", func() {
		checkAllSongsChart(appConfig.RepoDb)
	})
	c.Start()
}
