package cronjobs

import (
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/repositories"
	"EurikaOrmanel/up-charter/internal/schemas"
	audiomackServices "EurikaOrmanel/up-charter/internal/services/audiomack"
)

func checkAllSongsChart(repoDb repositories.DB) {
	//TODO: join queries to fetch songs and platform urls and then get their stream counts
	pageCounts := schemas.PaginationQuery{Count: 10, Page: 1}
	currentSongPlatforms := repoDb.FindSongPlatformCountNLastDate("audiomack", pageCounts)
	for {

		for i := 0; i <= len(currentSongPlatforms)-1; i++ {
			currentSong := currentSongPlatforms[i]

			audiomackData, err := audiomackServices.AudiomackSongInfo(currentSong.Url)
			if err != nil {
				continue
			}
			countIncrement := audiomackData.Stats.Plays - currentSong.TotalCount
			if countIncrement > 0 {
				songPlays := []*models.SongDailyPlay{{
					Count:          countIncrement,
					SongID:         currentSong.SongID,
					SongPlatformID: currentSong.PlatformID,
				}}
				err = repoDb.AddSongDailyPlays(songPlays)
			}
		}
		pageCounts.Page += 1
		currentSongPlatforms = repoDb.FindSongPlatformCountNLastDate("audiomack", pageCounts)
		if len(currentSongPlatforms) == 0 {
			break
		}
	}

}
