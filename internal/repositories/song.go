package repositories

import (
	"EurikaOrmanel/up-charter/internal/models"
	"EurikaOrmanel/up-charter/internal/schemas"
)

func (db DB) AddSong(song *models.Song) error {
	return db.Create(song).Error
}

func (db DB) FindSongByID(id string) models.Song {
	song := new(models.Song)
	db.First(song, "id = ?", id)
	return *song

}
func (db DB) FindSongByArtistIdNTitle(artistId string, title string) models.Song {
	song := new(models.Song)

	db.First(song, "artist_id = ? AND title = ?", artistId, title)
	return *song

}

// func (db DB) SongPlatformNChart(query schemas.PaginationQuery) {
// 	songs := []models.Song{}
// 	db.Preload("platforms").Preload("play_counts").Table("song")
// }

func (db DB) SongsTotalChartNLastChartDate(query schemas.PaginationQuery) {
	//as the name reads, that's what we're gon do.

}
func (db DB) SongNChartByID(id string) models.SongNCount {
	songNCount := new(models.SongNCount)
	db.
		Table("songs").
		Select("songs.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_id = songs.id").
		Order("total_count DESC").
		Limit(1).
		Find(songNCount)

	return *songNCount
}
func (db DB) SongNCharts(query schemas.SongChartQuery) []*models.SongNCount {
	songs := []*models.SongNCount{}

	db.
		Scopes(paginate(query.PaginationQuery, db.DB)).
		Table("songs").
		Select("songs.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_platform_id = song_platforms.id").
		Find(&songs).
		Order("total_count DESC")

	return songs
}

func (db DB) SongsInChartNCount() {}
func (db DB) SongNChartsByIDs(songIds []string) []*models.SongNCount {
	songs := []*models.SongNCount{}

	db.
		Table("songs").
		Select("songs.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_platform_id = song_platforms.id").
		Where("songs.id in (?)", songIds).
		Find(&songs).
		Order("total_count DESC")

	return songs

}
func (db DB) SongNChartByGr8ThanCount(count int, query schemas.SongChartQuery) []*models.SongNCount {
	songs := []*models.SongNCount{}

	db.
		Scopes(paginate(query.PaginationQuery, db.DB)).
		Table("songs").
		Select("songs.*,COALESCE(SUM(song_daily_plays.count),0) as total_count").
		Where("total_count > ?", count).
		Joins("LEFT JOIN song_daily_plays ON song_daily_plays.song_platform_id = song_platforms.id").
		Find(&songs).
		Order("total_count DESC")

	return songs
}
