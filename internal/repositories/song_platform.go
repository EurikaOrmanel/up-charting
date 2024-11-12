package repositories

import (
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
)

func (db DB) FindSongPlatformCountNLastDate(platforName string, query schemas.PaginationQuery) []*models.SongPlatTotCountNLastCounted {
	songPlatforms := []*models.SongPlatTotCountNLastCounted{}
	/*TODO: check song plays and get the ones that have been updated
	within 24 hours in the sub-query*/

	/*TODO: fetch songs that haven't been updated in the past 24 hrs
	based on the ones that are in the top query
	then paginate the shit */
	songUpdatedOver24hrs := db.Table("song_daily_plays").
		Select("song_platform_id").
		Where("song_daily_plays.created_at < DATETIME('now','-1 day')").
		Order("created_at DESC")

	db.
		Scopes(paginate(query, db.DB)).
		Table("song_platforms").
		Select("song_platforms.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Where("song_platforms.id IN (?) ",
			songUpdatedOver24hrs).
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_platform_id = song_platforms.id").
		Find(&songPlatforms)

	return songPlatforms
}
