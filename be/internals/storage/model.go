package storage

type RegistrationRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role"`
}

type HomeResponse struct {
	Status          string       `json:"status"`
	MoviesAndSeries []MovieSerie `json:"data"`
}

type UserInformation struct {
	User  ApplicationUser `json:"user"`
	Roles []string        `json:"roles"`
}

type ApplicationUser struct {
	ID             string  `json:"id" db:"id"`
	Username       string  `json:"username" db:"username"`
	Email          string  `json:"email" db:"email"`
	Name           string  `json:"name" db:"name"`
	ProfilePicture *string `json:"profile_picture,omitempty" db:"profile_picture"`
	PasswordHash   string  `json:"-" db:"password_hash"`
	CreatedAt      string  `json:"created_at" db:"created_at"`
}

type Role struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type UserRole struct {
	UserID string `db:"user_id"`
	RoleID string `db:"role_id"`
}

type MovieSerie struct {
	ID               uint    `json:"id" gorm:"primaryKey;column:ms_id"`
	Title            string  `json:"title" gorm:"column:title"`
	Synopsis         string  `json:"synopsis" gorm:"column:synopsis"`
	ReleaseYear      int     `json:"releaseYear" gorm:"column:release_year"`
	ClassificationMS string  `json:"classificationMS" gorm:"column:classification_ms"`
	Director         string  `json:"director" gorm:"column:director"`
	Cover            string  `json:"cover" gorm:"column:cover"`
	DateAdded        string  `json:"dateAdded" gorm:"column:date_added"`
	Duration         float64 `json:"duration" gorm:"column:duration"`
	MSType           string  `json:"msType" gorm:"column:ms_type"`
	Trailer          string  `json:"trailer" gorm:"column:trailer"`
}

type CreateMovieSerieRequest struct {
	Title            string `json:"title" binding:"required"`
	Synopsis         string `json:"synopsis" binding:"required"`
	ReleaseYear      int    `json:"releaseYear" binding:"required"`
	ClassificationMS string `json:"classificationMS" binding:"required"`
	Director         string `json:"director" binding:"required"`
	Cover            string `json:"cover" binding:"required"`
	Duration         float64 `json:"duration"`
	MSType           string `json:"msType" binding:"required"`
	Trailer          string `json:"trailer"`
	GenreIDs         []uint  `json:"genreIds"`
	ActorIDs         []uint  `json:"actorIds"`
}

type UpdateMovieSerieRequest struct {
	Title            string `json:"title"`
	Synopsis         string `json:"synopsis"`
	ReleaseYear      int    `json:"releaseYear"`
	ClassificationMS string `json:"classificationMS"`
	Director         string `json:"director"`
	Cover            string `json:"cover"`
	Duration         float64 `json:"duration"`
	MSType           string `json:"msType"`
	Trailer          string `json:"trailer"`
	GenreIDs         []uint  `json:"genreIds"`
	ActorIDs         []uint  `json:"actorIds"`
}

type Genre struct {
	GenreID   uint   `json:"id" gorm:"primaryKey;column:genre_id"`
	GenreName string `json:"name" gorm:"column:genre_name"`
}

type Actor struct {
	ActorID        uint    `json:"id" gorm:"primaryKey;column:actor_id"`
	ActorFirstName string  `json:"firstName" gorm:"column:actor_first_name"`
	ActorLastName  string  `json:"lastName" gorm:"column:actor_last_name"`
	ActorPhoto     *string `json:"photo,omitempty" gorm:"column:actor_photo"`
}

type Episode struct {
	MsID      uint    `json:"msId" gorm:"primaryKey;column:ms_id"`
	SeasonID  int     `json:"seasonId" gorm:"primaryKey;column:season_id"`
	EpisodeID int     `json:"episodeId" gorm:"primaryKey;column:episode_id"`
	Title     string  `json:"title" gorm:"column:title"`
	Duration  float64 `json:"duration" gorm:"column:duration"`
}

type Comment struct {
	CommentID uint   `json:"id" gorm:"primaryKey;column:comment_id"`
	Comment1  string `json:"comment" gorm:"column:comment"`
	AppUser   string `json:"appUser" gorm:"column:app_user"`
	MsID      uint   `json:"msId" gorm:"column:ms_id"`
	Photo     string `json:"photo,omitempty" gorm:"-"`
}

type Rating struct {
	RatingID uint   `json:"id" gorm:"primaryKey;column:rating_id"`
	Rating1  int    `json:"rating" gorm:"column:rating"`
	AppUser  string `json:"appUser" gorm:"column:app_user"`
	MsID     uint   `json:"msId" gorm:"column:ms_id"`
}

type CreateGenreRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateGenreRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateActorRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Photo     *string `json:"photo"`
}

type UpdateActorRequest struct {
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Photo     *string `json:"photo"`
}

type MovieSerieGenre struct {
	MsID    uint `json:"msId" gorm:"primaryKey;column:ms_id"`
	GenreID uint `json:"genreId" gorm:"primaryKey;column:genre_id"`
}

type MovieSerieActor struct {
	MsID    uint `json:"msId" gorm:"primaryKey;column:ms_id"`
	ActorID uint `json:"actorId" gorm:"primaryKey;column:actor_id"`
}

type RatingData struct {
	Average    float64 `json:"average"`
	Votes      int     `json:"votes"`
	Percentage string  `json:"percentage"`
}

type MovieOrSerieData struct {
	MovieSerie MovieSerie `json:"movieSerie"`
	Actors     []Actor    `json:"actors"`
	Genres     string     `json:"genres"`
	RatingData RatingData `json:"ratingData"`
	DurationHM string     `json:"durationHM"`
	Episodes   []Episode  `json:"episodes"`
}

type HomePageModel struct {
	LastTen   []MovieSerie `json:"lastTen"`
	Genre1    Genre        `json:"genre1"`
	Genre2    Genre        `json:"genre2"`
	Genre3    Genre        `json:"genre3"`
	Carousel1 []MovieSerie `json:"carousel1"`
	Carousel2 []MovieSerie `json:"carousel2"`
	Carousel3 []MovieSerie `json:"carousel3"`
}
