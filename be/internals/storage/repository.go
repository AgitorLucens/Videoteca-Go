package storage

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) GetAllMoviesAndSeries() ([]MovieSerie, error) {
	var ms []MovieSerie
	if err := r.DB.Table("movie_series").Find(&ms).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

func (r *Repository) GetMoviesAndSeries() ([]MovieSerie, error) {
	ms := []MovieSerie{}
	if err := r.DB.Raw("SELECT * FROM ObtenerUltimas10Peliculas()").Scan(&ms).Error; err != nil {
		return nil, err
	}
	fmt.Println(ms)
	return ms, nil
}

func (r *Repository) GetMovieSerieByID(id uint) (*MovieSerie, error) {
	var ms MovieSerie
	if err := r.DB.Table("movie_series").Where("ms_id = ?", id).First(&ms).Error; err != nil {
		return nil, err
	}
	return &ms, nil
}

func (r *Repository) CreateMovieSerie(req CreateMovieSerieRequest) (*MovieSerie, error) {
	ms := MovieSerie{
		Title:            req.Title,
		Synopsis:         req.Synopsis,
		ReleaseYear:      req.ReleaseYear,
		ClassificationMS: req.ClassificationMS,
		Director:         req.Director,
		Cover:            req.Cover,
		DateAdded:        time.Now().Format("2006-01-02"),
		Duration:         req.Duration,
		MSType:           req.MSType,
		Trailer:          req.Trailer,
	}
	if err := r.DB.Table("movie_series").Create(&ms).Error; err != nil {
		return nil, err
	}
	if len(req.GenreIDs) > 0 {
		if err := r.ReplaceMovieSerieGenres(ms.ID, req.GenreIDs); err != nil {
			return &ms, err
		}
	}
	if len(req.ActorIDs) > 0 {
		if err := r.ReplaceMovieSerieActors(ms.ID, req.ActorIDs); err != nil {
			return &ms, err
		}
	}
	return &ms, nil
}

func (r *Repository) UpdateMovieSerie(id uint, req UpdateMovieSerieRequest) (*MovieSerie, error) {
	var ms MovieSerie
	if err := r.DB.Table("movie_series").Where("ms_id = ?", id).First(&ms).Error; err != nil {
		return nil, errors.New("movie/serie not found")
	}

	updates := map[string]interface{}{}
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Synopsis != "" {
		updates["synopsis"] = req.Synopsis
	}
	if req.ReleaseYear != 0 {
		updates["release_year"] = req.ReleaseYear
	}
	if req.ClassificationMS != "" {
		updates["classification_ms"] = req.ClassificationMS
	}
	if req.Director != "" {
		updates["director"] = req.Director
	}
	if req.Cover != "" {
		updates["cover"] = req.Cover
	}
	if req.Duration != 0 {
		updates["duration"] = req.Duration
	}
	if req.MSType != "" {
		updates["ms_type"] = req.MSType
	}
	if req.Trailer != "" {
		updates["trailer"] = req.Trailer
	}

	if len(updates) > 0 {
		if err := r.DB.Table("movie_series").Where("ms_id = ?", id).Updates(updates).Error; err != nil {
			return nil, err
		}
	}

	if req.GenreIDs != nil {
		if err := r.ReplaceMovieSerieGenres(id, req.GenreIDs); err != nil {
			return nil, err
		}
	}
	if req.ActorIDs != nil {
		if err := r.ReplaceMovieSerieActors(id, req.ActorIDs); err != nil {
			return nil, err
		}
	}

	r.DB.Table("movie_series").Where("ms_id = ?", id).First(&ms)
	return &ms, nil
}

func (r *Repository) DeleteMovieSerie(id uint) error {
	if err := r.DB.Table("movie_serie_actors").Where("ms_id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	if err := r.DB.Table("movie_serie_genres").Where("ms_id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	if err := r.DB.Table("comments").Where("ms_id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	if err := r.DB.Table("ratings").Where("ms_id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	if err := r.DB.Table("episodes").Where("ms_id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	if err := r.DB.Table("movie_series").Where("ms_id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllGenres() ([]Genre, error) {
	var genres []Genre
	if err := r.DB.Table("genres").Find(&genres).Error; err != nil {
		return nil, err
	}
	return genres, nil
}

func (r *Repository) GetGenresByMsID(msID uint) ([]Genre, error) {
	var genres []Genre
	if err := r.DB.Table("genres").
		Joins("JOIN movie_serie_genres ON genres.genre_id = movie_serie_genres.genre_id").
		Where("movie_serie_genres.ms_id = ?", msID).
		Find(&genres).Error; err != nil {
		return nil, err
	}
	return genres, nil
}

func (r *Repository) GetActorsByMsID(msID uint) ([]Actor, error) {
	var actors []Actor
	if err := r.DB.Table("actors").
		Joins("JOIN movie_serie_actors ON actors.actor_id = movie_serie_actors.actor_id").
		Where("movie_serie_actors.ms_id = ?", msID).
		Find(&actors).Error; err != nil {
		return nil, err
	}
	return actors, nil
}

func (r *Repository) GetEpisodesByMsID(msID uint) ([]Episode, error) {
	var episodes []Episode
	if err := r.DB.Table("episodes").Where("ms_id = ?", msID).Find(&episodes).Error; err != nil {
		return nil, err
	}
	return episodes, nil
}

func (r *Repository) GetCommentsByMsID(msID uint) ([]Comment, error) {
	var comments []Comment
	if err := r.DB.Table("comments").Where("ms_id = ?", msID).Order("comment_id DESC").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *Repository) GetGenreByID(id uint) (*Genre, error) {
	var genre Genre
	if err := r.DB.Table("genres").Where("genre_id = ?", id).First(&genre).Error; err != nil {
		return nil, err
	}
	return &genre, nil
}

func (r *Repository) CreateGenre(req CreateGenreRequest) (*Genre, error) {
	genre := Genre{GenreName: req.Name}
	if err := r.DB.Table("genres").Create(&genre).Error; err != nil {
		return nil, err
	}
	return &genre, nil
}

func (r *Repository) UpdateGenre(id uint, req UpdateGenreRequest) (*Genre, error) {
	var genre Genre
	if err := r.DB.Table("genres").Where("genre_id = ?", id).First(&genre).Error; err != nil {
		return nil, errors.New("genre not found")
	}
	if err := r.DB.Table("genres").Where("genre_id = ?", id).Update("genre_name", req.Name).Error; err != nil {
		return nil, err
	}
	r.DB.Table("genres").Where("genre_id = ?", id).First(&genre)
	return &genre, nil
}

func (r *Repository) DeleteGenre(id uint) error {
	if err := r.DB.Table("movie_serie_genres").Where("genre_id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	if err := r.DB.Table("genres").Where("genre_id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAllActors() ([]Actor, error) {
	var actors []Actor
	if err := r.DB.Table("actors").Find(&actors).Error; err != nil {
		return nil, err
	}
	return actors, nil
}

func (r *Repository) GetActorByID(id uint) (*Actor, error) {
	var actor Actor
	if err := r.DB.Table("actors").Where("actor_id = ?", id).First(&actor).Error; err != nil {
		return nil, err
	}
	return &actor, nil
}

func (r *Repository) CreateActor(req CreateActorRequest) (*Actor, error) {
	actor := Actor{
		ActorFirstName: req.FirstName,
		ActorLastName:  req.LastName,
		ActorPhoto:     req.Photo,
	}
	if err := r.DB.Table("actors").Create(&actor).Error; err != nil {
		return nil, err
	}
	return &actor, nil
}

func (r *Repository) UpdateActor(id uint, req UpdateActorRequest) (*Actor, error) {
	var actor Actor
	if err := r.DB.Table("actors").Where("actor_id = ?", id).First(&actor).Error; err != nil {
		return nil, errors.New("actor not found")
	}
	updates := map[string]interface{}{
		"actor_first_name": req.FirstName,
		"actor_last_name":  req.LastName,
	}
	if req.Photo != nil {
		updates["actor_photo"] = *req.Photo
	}
	if err := r.DB.Table("actors").Where("actor_id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}
	r.DB.Table("actors").Where("actor_id = ?", id).First(&actor)
	return &actor, nil
}

func (r *Repository) DeleteActor(id uint) error {
	if err := r.DB.Table("movie_serie_actors").Where("actor_id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	if err := r.DB.Table("actors").Where("actor_id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) SearchMovies(query string) ([]MovieSerie, error) {
	var ms []MovieSerie
	if err := r.DB.Table("movie_series").
		Where("title ILIKE ?", "%"+query+"%").
		Limit(10).
		Find(&ms).Error; err != nil {
		return nil, err
	}
	return ms, nil
}

func (r *Repository) SearchActors(query string) ([]Actor, error) {
	var actors []Actor
	if err := r.DB.Table("actors").
		Where("actor_first_name ILIKE ? OR actor_last_name ILIKE ? OR (actor_first_name || ' ' || actor_last_name) ILIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Limit(10).
		Find(&actors).Error; err != nil {
		return nil, err
	}
	return actors, nil
}

func (r *Repository) GetRatingDataByMsID(msID uint) (RatingData, error) {
	var rd RatingData
	var avg float64
	var count int64
	r.DB.Table("ratings").Where("ms_id = ?", msID).Count(&count)
	r.DB.Table("ratings").Where("ms_id = ?", msID).Select("COALESCE(AVG(rating), 0)").Scan(&avg)

	var fiveStarCount int64
	r.DB.Table("ratings").Where("ms_id = ? AND rating = 5", msID).Count(&fiveStarCount)

	percentage := "0"
	if count > 0 {
		pct := float64(fiveStarCount) / float64(count) * 100
		percentage = fmt.Sprintf("%.0f", pct)
	}

	rd = RatingData{
		Average:   avg,
		Votes:     int(count),
		Percentage: percentage,
	}
	return rd, nil
}

func (r *Repository) GetUserRating(msID uint, userEmail string) (int, error) {
	var rating int
	err := r.DB.Table("ratings").Where("ms_id = ? AND app_user = ?", msID, userEmail).Select("rating").Scan(&rating).Error
	if err != nil {
		return 0, err
	}
	return rating, nil
}

func (r *Repository) ReplaceMovieSerieGenres(msID uint, genreIDs []uint) error {
	if err := r.DB.Table("movie_serie_genres").Where("ms_id = ?", msID).Delete(nil).Error; err != nil {
		return err
	}
	for _, genreID := range genreIDs {
		msg := MovieSerieGenre{MsID: msID, GenreID: genreID}
		if err := r.DB.Table("movie_serie_genres").Create(&msg).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) UpsertRating(msID uint, userEmail string, rating int) error {
	var existing Rating
	result := r.DB.Table("ratings").Where("ms_id = ? AND app_user = ?", msID, userEmail).First(&existing)
	if result.Error == nil {
		return r.DB.Table("ratings").Where("rating_id = ?", existing.RatingID).Update("rating", rating).Error
	}
	entry := Rating{MsID: msID, AppUser: userEmail, Rating1: rating}
	return r.DB.Table("ratings").Create(&entry).Error
}

func (r *Repository) CreateComment(msID uint, userEmail, comment string) error {
	record := Comment{MsID: msID, AppUser: userEmail, Comment1: comment}
	return r.DB.Table("comments").Create(&record).Error
}

func (r *Repository) GetCommentByID(commentID uint) (*Comment, error) {
	var comment Comment
	if err := r.DB.Table("comments").Where("comment_id = ?", commentID).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *Repository) DeleteComment(commentID uint) error {
	return r.DB.Table("comments").Where("comment_id = ?", commentID).Delete(nil).Error
}

func (r *Repository) GetCommentsByMsIDPaginated(msID uint, page, limit int) ([]Comment, int64, error) {
	var total int64
	r.DB.Table("comments").Where("ms_id = ?", msID).Count(&total)

	offset := (page - 1) * limit
	rows, err := r.DB.Table("comments").
		Select("comments.*, users.profile_picture").
		Joins("LEFT JOIN users ON comments.app_user = users.email").
		Where("comments.ms_id = ?", msID).
		Order("comments.comment_id DESC").
		Offset(offset).Limit(limit).
		Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var c Comment
		var photoBytes []byte
		if err := rows.Scan(&c.CommentID, &c.Comment1, &c.AppUser, &c.MsID, &photoBytes); err != nil {
			return nil, 0, err
		}
		if len(photoBytes) > 0 {
			c.Photo = "data:image/png;base64," + base64.StdEncoding.EncodeToString(photoBytes)
		}
		comments = append(comments, c)
	}

	return comments, total, nil
}

func (r *Repository) ReplaceMovieSerieActors(msID uint, actorIDs []uint) error {
	if err := r.DB.Table("movie_serie_actors").Where("ms_id = ?", msID).Delete(nil).Error; err != nil {
		return err
	}
	for _, actorID := range actorIDs {
		msa := MovieSerieActor{MsID: msID, ActorID: actorID}
		if err := r.DB.Table("movie_serie_actors").Create(&msa).Error; err != nil {
			return err
		}
	}
	return nil
}
