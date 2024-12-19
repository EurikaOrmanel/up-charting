package cronjobs

import (
	"EurikaOrmanel/up-charter/config"
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
	audiomackServices "EurikaOrmanel/up-charter/internal/services/audiomack"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

func checkAllSongsChart(appConfig config.AppConfig) {
	//TODO: join queries to fetch songs and platform urls and then get their stream counts
	pageCounts := schemas.PaginationQuery{Count: 10, Page: 1}
	currentSongPlatforms := appConfig.RepoDb.FindSongPlatformCountNLastDate("audiomack", pageCounts)
	for {

		for i := 0; i <= len(currentSongPlatforms)-1; i++ {
			currentSongPlatform := currentSongPlatforms[i]
			audiomackData, err := audiomackServices.AudiomackSongInfo(currentSongPlatform.Url)
			if err != nil {
				log.Println(err)
				// continue
			}

			fmt.Println("currentSongPlatform.TotalCount:", currentSongPlatform.TotalCount, " ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;")
			countIncrement := audiomackData.Stats.Plays - currentSongPlatform.TotalCount
			if countIncrement > 0 {
				songPlays := []*models.SongDailyPlay{{
					Count:          countIncrement,
					SongID:         currentSongPlatform.SongID,
					SongPlatformID: currentSongPlatform.PlatformID,
				}}
				err = appConfig.RepoDb.AddSongDailyPlays(songPlays)
				if err != nil {
					log.Println(err)
					// continue
				}
			}

			currentChart := appConfig.RepoDb.GetFromChartBySongID(currentSongPlatform.SongID.String())

			if currentChart.ID == uuid.Nil {
				songInfo := appConfig.RepoDb.FindSongByID(currentSongPlatform.SongID.String())
				err := appConfig.RepoDb.AddSongToChart(&models.Top100Chart{
					SongID:  currentSongPlatform.SongID,
					GenreID: songInfo.GenreID,
				})
				if err != nil {
					log.Println(err)
				}
			} else {
				err := appConfig.RepoDb.UpdateChartPosition(&currentChart)
				if err != nil {
					log.Println(err, " ;;;;----;;;---;;;---;;;")
				}
			}

		}
		pageCounts.Page += 1
		currentSongPlatforms = appConfig.RepoDb.FindSongPlatformCountNLastDate("audiomack", pageCounts)
		if len(currentSongPlatforms) == 0 {
			break
		}
	}
	fmt.Println(strings.Repeat("*", 100))
	chart := appConfig.RepoDb.GetChart100NSong()
	appConfig.CacheConfig.SetTop100Chart(chart)

	//TODO:fetch new chart and cache to redis for later usage
}
