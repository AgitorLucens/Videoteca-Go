import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import {
  createMovieSerie,
  updateMovieSerie,
  getMovieSerieById,
} from "../../services/movieserie.service";
import { getGenres } from "../../services/genre.service";
import { getActors } from "../../services/actor.service";
import InputIcon from "../../components/inputicon/InputIcon";
import ErrorPopup from "../../components/errorpopup/ErrorPopup";
import Loading from "../../components/loading/Loading";
import { IconArroa, IconUser, IconLock, IconFilm } from "../../components/sprite/Sprite";

export default function MovieSerieForm() {
  const { id } = useParams();
  const isEdit = Boolean(id);
  const navigate = useNavigate();

  const [loading, setLoading] = useState(isEdit);
  const [error, setError] = useState("");
  const [allGenres, setAllGenres] = useState([]);
  const [allActors, setAllActors] = useState([]);
  const [form, setForm] = useState({
    title: "",
    synopsis: "",
    releaseYear: "",
    classificationMS: "",
    director: "",
    cover: "",
    duration: "",
    msType: "movie",
    trailer: "",
    genreIds: [],
    actorIds: [],
  });

  useEffect(() => {
    const fetchData = async () => {
      try {
        const [genreData, actorData] = await Promise.all([
          getGenres(),
          getActors(),
        ]);
        setAllGenres(Array.isArray(genreData?.data) ? genreData.data : []);
        setAllActors(Array.isArray(actorData?.data) ? actorData.data : []);
      } catch (err) {
        setError("Failed to load genres or actors.");
      }
    };
    fetchData();
  }, []);

  useEffect(() => {
    if (!isEdit) return;
    const fetchMovie = async () => {
      try {
        const data = await getMovieSerieById(id);
        const ms = data.movieSerie || data;
        setForm({
          title: ms.title || "",
          synopsis: ms.synopsis || "",
          releaseYear: ms.releaseYear || "",
          classificationMS: ms.classificationMS || "",
          director: ms.director || "",
          cover: ms.cover || "",
          duration: ms.duration || "",
          msType: ms.msType || "movie",
          trailer: ms.trailer || "",
          genreIds: data.genreIds || [],
          actorIds: data.actorIds || [],
        });
      } catch (err) {
        setError("Failed to load movie/serie data.");
      } finally {
        setLoading(false);
      }
    };
    fetchMovie();
  }, [id, isEdit]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const toggleGenre = (genreId) => {
    setForm((prev) => {
      const exists = prev.genreIds.includes(genreId);
      return {
        ...prev,
        genreIds: exists
          ? prev.genreIds.filter((id) => id !== genreId)
          : [...prev.genreIds, genreId],
      };
    });
  };

  const toggleActor = (actorId) => {
    setForm((prev) => {
      const exists = prev.actorIds.includes(actorId);
      return {
        ...prev,
        actorIds: exists
          ? prev.actorIds.filter((id) => id !== actorId)
          : [...prev.actorIds, actorId],
      };
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError("");

    const payload = {
      title: form.title,
      synopsis: form.synopsis,
      releaseYear: parseInt(form.releaseYear, 10) || 0,
      classificationMS: form.classificationMS,
      director: form.director,
      cover: form.cover,
      duration: parseFloat(form.duration) || 0,
      msType: form.msType,
      trailer: form.trailer,
      genreIds: form.genreIds,
      actorIds: form.actorIds,
    };

    try {
      if (isEdit) {
        await updateMovieSerie(id, payload);
      } else {
        await createMovieSerie(payload);
      }
      navigate("/admin/movies");
    } catch (err) {
      setError(
        isEdit ? "Failed to update movie/serie." : "Failed to create movie/serie."
      );
    }
  };

  if (loading) return <Loading message="Loading movie data" />;

  return (
    <div className="min-h-screen bg-[#1f2029] text-white p-6">
      <div className="max-w-2xl mx-auto">
        <h1 className="text-3xl font-bold mb-6">
          {isEdit ? "Edit" : "Add"} Movie / Series
        </h1>

        <ErrorPopup message={error} onClose={() => setError("")} />

        <form
          onSubmit={handleSubmit}
          className="bg-[#2a2b38] rounded-lg p-6 flex flex-col gap-4"
        >
          <InputIcon
            icon={IconFilm}
            type="text"
            placeholder="Title"
            name="title"
            value={form.title}
            onChange={handleChange}
          />

          <textarea
            name="synopsis"
            placeholder="Synopsis"
            value={form.synopsis}
            onChange={handleChange}
            rows={4}
            className="w-full p-3 rounded-md bg-[#1f2029] text-white placeholder-gray-400 outline-none resize-none"
          />

          <div className="grid grid-cols-2 gap-4">
            <InputIcon
              icon={IconArroa}
              type="number"
              placeholder="Release Year"
              name="releaseYear"
              value={form.releaseYear}
              onChange={handleChange}
            />
            <InputIcon
              icon={IconUser}
              type="text"
              placeholder="Classification (e.g. PG)"
              name="classificationMS"
              value={form.classificationMS}
              onChange={handleChange}
            />
          </div>

          <InputIcon
            icon={IconUser}
            type="text"
            placeholder="Director"
            name="director"
            value={form.director}
            onChange={handleChange}
          />

          <InputIcon
            icon={IconArroa}
            type="text"
            placeholder="Cover Image URL"
            name="cover"
            value={form.cover}
            onChange={handleChange}
          />

          <div className="grid grid-cols-2 gap-4">
            <InputIcon
              icon={IconFilm}
              type="number"
              placeholder="Duration (minutes)"
              name="duration"
              value={form.duration}
              onChange={handleChange}
            />
            <div className="flex flex-col">
              <label className="text-sm text-gray-400 mb-1">Type</label>
              <select
                name="msType"
                value={form.msType}
                onChange={handleChange}
                className="w-full p-3 rounded-md bg-[#1f2029] text-white outline-none"
              >
                <option value="movie">Movie</option>
                <option value="serie">Series</option>
              </select>
            </div>
          </div>

          <InputIcon
            icon={IconArroa}
            type="text"
            placeholder="Trailer URL"
            name="trailer"
            value={form.trailer}
            onChange={handleChange}
          />

          <div className="flex flex-col gap-2">
            <label className="text-sm text-gray-400 font-semibold">Genres / Tags</label>
            <div className="flex flex-wrap gap-2">
              {allGenres.map((genre) => (
                <button
                  key={genre.id}
                  type="button"
                  onClick={() => toggleGenre(genre.id)}
                  className={`px-3 py-1.5 rounded-full text-sm font-medium border transition-colors duration-150 ${
                    form.genreIds.includes(genre.id)
                      ? "bg-[#ffeba7] text-[#102770] border-[#ffeba7]"
                      : "bg-transparent text-gray-300 border-gray-600 hover:border-gray-400"
                  }`}
                >
                  {genre.name}
                </button>
              ))}
              {allGenres.length === 0 && (
                <span className="text-gray-500 text-sm">No genres available. Add some in the Tags page.</span>
              )}
            </div>
          </div>

          <div className="flex flex-col gap-2">
            <label className="text-sm text-gray-400 font-semibold">Actors</label>
            <div className="flex flex-wrap gap-2 max-h-48 overflow-y-auto">
              {allActors.map((actor) => (
                <button
                  key={actor.id}
                  type="button"
                  onClick={() => toggleActor(actor.id)}
                  className={`px-3 py-1.5 rounded-full text-sm font-medium border transition-colors duration-150 ${
                    form.actorIds.includes(actor.id)
                      ? "bg-[#ffeba7] text-[#102770] border-[#ffeba7]"
                      : "bg-transparent text-gray-300 border-gray-600 hover:border-gray-400"
                  }`}
                >
                  {actor.firstName} {actor.lastName}
                </button>
              ))}
              {allActors.length === 0 && (
                <span className="text-gray-500 text-sm">No actors available. Add some in the Actors page.</span>
              )}
            </div>
          </div>

          <div className="flex gap-4 mt-4">
            <button
              type="submit"
              className="bg-[#ffeba7] text-[#102770] font-semibold py-2 px-6 rounded-md hover:bg-[#102770] hover:text-[#ffeba7] transition-colors duration-200"
            >
              {isEdit ? "Save Changes" : "Create"}
            </button>
            <button
              type="button"
              onClick={() => navigate("/admin/movies")}
              className="bg-gray-600 text-white font-semibold py-2 px-6 rounded-md hover:bg-gray-700 transition-colors duration-200"
            >
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
