package storage

type GenreRepository interface {
	GetAllGenres() ([]Genre, error)
	GetGenreByID(id uint) (*Genre, error)
	CreateGenre(req CreateGenreRequest) (*Genre, error)
	UpdateGenre(id uint, req UpdateGenreRequest) (*Genre, error)
	DeleteGenre(id uint) error
}

type ActorRepository interface {
	GetAllActors() ([]Actor, error)
	GetActorByID(id uint) (*Actor, error)
	CreateActor(req CreateActorRequest) (*Actor, error)
	UpdateActor(id uint, req UpdateActorRequest) (*Actor, error)
	DeleteActor(id uint) error
}

type HomeRepository interface {
	GetMoviesAndSeries() ([]MovieSerie, error)
}

type MovieSerieRepository interface {
	GetAllMoviesAndSeries() ([]MovieSerie, error)
	GetMovieSerieByID(id uint) (*MovieSerie, error)
	CreateMovieSerie(req CreateMovieSerieRequest) (*MovieSerie, error)
	UpdateMovieSerie(id uint, req UpdateMovieSerieRequest) (*MovieSerie, error)
	DeleteMovieSerie(id uint) error
	GetGenresByMsID(msID uint) ([]Genre, error)
	GetActorsByMsID(msID uint) ([]Actor, error)
	GetEpisodesByMsID(msID uint) ([]Episode, error)
	GetRatingDataByMsID(msID uint) (RatingData, error)
	UpsertRating(msID uint, userEmail string, rating int) error
	ReplaceMovieSerieGenres(msID uint, genreIDs []uint) error
	ReplaceMovieSerieActors(msID uint, actorIDs []uint) error
	GetCommentsByMsIDPaginated(msID uint, page, limit int) ([]Comment, int64, error)
	GetUserRating(msID uint, userEmail string) (int, error)
	CreateComment(msID uint, userEmail, comment string) error
	GetCommentByID(commentID uint) (*Comment, error)
	DeleteComment(commentID uint) error
}
