package schemas

type AddAlbumInput struct {
	Title     string             `validate:"string|required" json:"title"`
	Platforms SongPlatformsInput `validate:"required|list" json:"platforms"`
	ArtistId  string             `validate:"uuid4|required" json:"artist_id"`
	GenreId   string             `validate:"required|uuid4" json:"genre_id"`
}
