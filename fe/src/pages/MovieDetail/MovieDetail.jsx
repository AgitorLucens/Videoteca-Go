import { useState, useEffect } from "react";
import { useParams, useNavigate } from "react-router-dom";
import Slider from "react-slick";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import { getMovieSerieById } from "../../services/movieserie.service";
import { rateMovie } from "../../services/rating.service";
import Loading from "../../components/loading/Loading";
import { useAuth } from "../../hooks/useAuth";
import CommentsPopup from "../../components/commentspopup/CommentsPopup";

function getYouTubeEmbedUrl(url) {
  if (!url) return null;
  let videoId = "";
  const patterns = [
    /(?:youtube\.com\/watch\?v=|youtu\.be\/|youtube\.com\/embed\/)([a-zA-Z0-9_-]{11})/,
    /^([a-zA-Z0-9_-]{11})$/,
  ];
  for (const p of patterns) {
    const m = url.match(p);
    if (m) {
      videoId = m[1];
      break;
    }
  }
  return videoId ? `https://www.youtube.com/embed/${videoId}` : null;
}

function StarRating({ value, onChange }) {
  const [hover, setHover] = useState(0);
  return (
    <div className="flex items-center gap-1">
      {[1, 2, 3, 4, 5].map((star) => (
        <button
          key={star}
          type="button"
          onClick={() => onChange?.(star)}
          onMouseEnter={() => setHover(star)}
          onMouseLeave={() => setHover(0)}
          className="text-2xl transition-colors duration-150"
        >
          <span
            className={
              star <= (hover || value)
                ? "text-[#ffeba7]"
                : "text-gray-600"
            }
          >
            &#9733;
          </span>
        </button>
      ))}
    </div>
  );
}

export default function MovieDetail() {
  const { id } = useParams();
  const navigate = useNavigate();
  const { user } = useAuth();
  const [movie, setMovie] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [userRating, setUserRating] = useState(0);
  const [submitting, setSubmitting] = useState(false);
  const [showComments, setShowComments] = useState(false);

  useEffect(() => {
    const fetchMovie = async () => {
      try {
        const data = await getMovieSerieById(id);
        setMovie(data);
        if (data.userRating) setUserRating(data.userRating);
      } catch {
        setError("Failed to load movie/serie details.");
      } finally {
        setLoading(false);
      }
    };
    fetchMovie();
  }, [id]);

  const handleRate = async (star) => {
    setSubmitting(true);
    try {
      const res = await rateMovie(id, star);
      setMovie((prev) => ({ ...prev, ratingData: res.ratingData }));
      setUserRating(star);
    } catch (err) {
      console.error("Failed to submit rating:", err);
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) return <Loading message="Loading details" />;

  if (error) {
    return (
      <div className="min-h-screen bg-[#1f2029] text-white p-6 flex flex-col items-center justify-center gap-4">
        <p className="text-red-400 text-lg">{error}</p>
        <button
          onClick={() => navigate(-1)}
          className="bg-[#ffeba7] text-[#102770] font-semibold py-2 px-6 rounded-md"
        >
          Go Back
        </button>
      </div>
    );
  }

  const ms = movie?.movieSerie || {};
  const actors = movie?.actors || [];
  const episodes = movie?.episodes || [];
  const ratingData = movie?.ratingData || {};
  const embedUrl = getYouTubeEmbedUrl(ms.trailer);

  const actorSliderSettings = {
    slidesToShow: Math.min(5, actors.length),
    slidesToScroll: 1,
    infinite: actors.length > 5,
    centerMode: false,
    arrows: true,
    responsive: [
      { breakpoint: 1024, settings: { slidesToShow: Math.min(3, actors.length) } },
      { breakpoint: 640, settings: { slidesToShow: Math.min(2, actors.length) } },
    ],
  };

  return (
    <div className="min-h-screen bg-[#1f2029] text-white">
      <div className="relative">
        {ms.cover && (
          <div className="h-48 sm:h-72 md:h-96 overflow-hidden">
            <img
              src={ms.cover}
              alt={ms.title}
              className="w-full h-full object-cover"
            />
            <div className="absolute inset-0 bg-gradient-to-t from-[#1f2029] via-[#1f2029]/60 to-transparent" />
          </div>
        )}

        <div className="absolute top-4 left-4">
          <button
            onClick={() => navigate(-1)}
            className="bg-black/50 hover:bg-black/70 text-white px-4 py-2 rounded-md transition-colors"
          >
            &larr; Back
          </button>
        </div>

        <div className="relative px-4 sm:px-6 pb-6 -mt-16 sm:-mt-24 md:-mt-32">
          <h1 className="text-2xl sm:text-3xl md:text-4xl font-bold mb-2">{ms.title}</h1>

          <div className="flex flex-wrap items-center gap-3 text-sm text-gray-300 mb-4">
            {ms.releaseYear && <span>{ms.releaseYear}</span>}
            {ms.msType && (
              <span
                className={`px-2 py-0.5 rounded text-xs font-semibold ${ms.msType === "movie"
                  ? "bg-blue-600/30 text-blue-300"
                  : "bg-purple-600/30 text-purple-300"
                  }`}
              >
                {ms.msType}
              </span>
            )}
            {ms.classificationMS && (
              <span className="border border-gray-600 px-2 py-0.5 rounded text-xs">
                {ms.classificationMS}
              </span>
            )}
            {ms.duration > 0 && (
              <span>{ms.duration} min</span>
            )}
          </div>

          {movie?.genres && (
            <p className="text-sm text-[#ffeba7] mb-4">{movie.genres}</p>
          )}

          <div className="flex flex-wrap items-center gap-4">
            {ratingData.votes > 0 ? (
              <div className="flex items-center gap-2">
                <span className="text-3xl font-bold text-[#ffeba7]">
                  {ratingData.average.toFixed(1)}
                </span>
                <span className="text-gray-400">/ 5</span>
                <span className="text-gray-400">({ratingData.votes} votes)</span>
              </div>
            ) : (
              <p className="text-gray-500 text-sm">No ratings yet — be the first!</p>
            )}

            {user && (
              <div className="flex items-center gap-2 ml-auto">
                <span className="text-sm text-gray-400">Your rating:</span>
                <StarRating value={userRating} onChange={handleRate} />
                {submitting && <span className="text-xs text-gray-500">...</span>}
              </div>
            )}
          </div>
        </div>
      </div>

      <div className="px-4 sm:px-6 pb-8 max-w-6xl mx-auto">
        {(ms.synopsis || ms.director) && embedUrl ? (
          <div className="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
            <div className="space-y-6">
              {ms.synopsis && (
                <section>
                  <h2 className="text-xl font-semibold mb-2 text-[#ffeba7]">Synopsis</h2>
                  <p className="text-gray-300 leading-relaxed">{ms.synopsis}</p>
                </section>
              )}
              {ms.director && (
                <section>
                  <h2 className="text-xl font-semibold mb-2 text-[#ffeba7]">Director</h2>
                  <p className="text-gray-300">{ms.director}</p>
                </section>
              )}
              <section>
                <button
                  onClick={() => setShowComments(true)}
                  className="bg-[#ffeba7] text-[#102770] font-semibold py-2 px-6 rounded-md hover:bg-[#e6d492] transition-colors"
                >
                  View Comments
                </button>
              </section>
            </div>
            {embedUrl && (
              <section>
                <h2 className="text-xl font-semibold mb-2 text-[#ffeba7]">Trailer</h2>
                <div className="relative w-full" style={{ paddingBottom: "56.25%" }}>
                  <iframe
                    src={embedUrl}
                    title="Trailer"
                    className="absolute top-0 left-0 w-full h-full rounded-md"
                    allowFullScreen
                    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                  />
                </div>
              </section>
            )}
          </div>
        ) : (
          <div className="space-y-6 mb-8">
            {ms.synopsis && (
              <section>
                <h2 className="text-xl font-semibold mb-2 text-[#ffeba7]">Synopsis</h2>
                <p className="text-gray-300 leading-relaxed">{ms.synopsis}</p>
              </section>
            )}
            {ms.director && (
              <section>
                <h2 className="text-xl font-semibold mb-2 text-[#ffeba7]">Director</h2>
                <p className="text-gray-300">{ms.director}</p>
              </section>
            )}
            {embedUrl && (
              <section>
                <h2 className="text-xl font-semibold mb-2 text-[#ffeba7]">Trailer</h2>
                <div className="relative w-full max-w-2xl" style={{ paddingBottom: "56.25%" }}>
                  <iframe
                    src={embedUrl}
                    title="Trailer"
                    className="absolute top-0 left-0 w-full h-full rounded-md"
                    allowFullScreen
                    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                  />
                </div>
              </section>
            )}
              <section>
                <button
                  onClick={() => setShowComments(true)}
                  className="bg-[#ffeba7] text-[#102770] font-semibold py-2 px-6 rounded-md hover:bg-[#e6d492] transition-colors"
                >
                  View Comments
                </button>
              </section>
          </div>
        )}

        {episodes.length > 0 && (
          <section className="mb-8">
            <h2 className="text-xl font-semibold mb-2 text-[#ffeba7]">Episodes</h2>
            <div className="space-y-2">
              {episodes.map((ep) => (
                <div
                  key={`${ep.seasonId}-${ep.episodeId}`}
                  className="bg-[#2a2b38] p-3 rounded-md"
                >
                  <p className="text-sm text-gray-400">
                    S{ep.seasonId}E{ep.episodeId}
                  </p>
                  <p className="text-white font-medium">{ep.title}</p>
                  {ep.duration > 0 && (
                    <p className="text-sm text-gray-400">{ep.duration} min</p>
                  )}
                </div>
              ))}
            </div>
          </section>
        )}
      </div>



      {showComments && <CommentsPopup msId={id} onClose={() => setShowComments(false)} />}

      {actors.length > 0 && (
        <div className="pb-8 px-4 sm:px-6 max-w-6xl mx-auto">
          <h2 className="text-xl font-semibold mb-4 text-[#ffeba7]">Actors</h2>
          <Slider {...actorSliderSettings}>
            {actors.map((actor) => (
              <div key={actor.id} className="px-2">
                <div className="flex flex-col items-center bg-[#2a2b38] rounded-lg p-4">
                  {actor.photo ? (
                    <img
                      src={actor.photo}
                      alt={`${actor.firstName} ${actor.lastName}`}
                      className="w-20 h-20 rounded-full object-cover mb-2"
                    />
                  ) : (
                    <div className="w-20 h-20 rounded-full bg-gray-700 flex items-center justify-center text-gray-400 text-2xl font-bold mb-2">
                      {actor.firstName?.[0]}{actor.lastName?.[0]}
                    </div>
                  )}
                  <p className="text-sm text-gray-200 text-center truncate w-full">
                    {actor.firstName} {actor.lastName}
                  </p>
                </div>
              </div>
            ))}
          </Slider>
        </div>
      )}
    </div>
  );
}
