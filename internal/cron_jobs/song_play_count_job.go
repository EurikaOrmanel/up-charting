package cronjobs

import (
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/repositories"
	"EurikaOrmanel/up-charter/internal/schemas"
	audiomackServices "EurikaOrmanel/up-charter/internal/services/audiomack"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
)

func checkAllSongsChart(repoDb repositories.DB) {
	//TODO: join queries to fetch songs and platform urls and then get their stream counts
	pageCounts := schemas.PaginationQuery{Count: 10, Page: 1}
	currentSongPlatforms := repoDb.FindSongPlatformCountNLastDate("audiomack", pageCounts)
	for {

		for i := 0; i <= len(currentSongPlatforms)-1; i++ {
			currentSongPlatform := currentSongPlatforms[i]
			fmt.Println(*currentSongPlatform)

			audiomackData, err := audiomackServices.AudiomackSongInfo(currentSongPlatform.Url)
			if err != nil {
				log.Println(err)
				continue
			}
			countIncrement := audiomackData.Stats.Plays - currentSongPlatform.TotalCount
			if countIncrement > 0 {
				songPlays := []*models.SongDailyPlay{{
					Count:          countIncrement,
					SongID:         currentSongPlatform.SongID,
					SongPlatformID: currentSongPlatform.PlatformID,
				}}
				err = repoDb.AddSongDailyPlays(songPlays)
				if err != nil {
					log.Println(err)
					continue
				}

				currentChart := repoDb.GetFromChartBySongID(currentSongPlatform.SongID.String())

				fmt.Println("songTitle:", audiomackData.Title, "currentChart:", currentChart)

				if currentChart.ID == uuid.Nil {
					songInfo := repoDb.FindSongByID(currentSongPlatform.SongID.String())
					fmt.Println("SongInfo:", songInfo)
					repoDb.AddSongToChart(&models.Top100Chart{
						SongID:  currentSongPlatform.SongID,
						GenreID: songInfo.GenreID,
					})
				} else {
					repoDb.UpdateChartPosition(&currentChart)
				}

			}
		}
		pageCounts.Page += 1
		currentSongPlatforms = repoDb.FindSongPlatformCountNLastDate("audiomack", pageCounts)
		if len(currentSongPlatforms) == 0 {
			break
		}
	}
	fmt.Println(strings.Repeat("*", 100))
}
