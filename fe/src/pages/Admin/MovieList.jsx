import { useState, useEffect } from "react";
import { Link, useNavigate } from "react-router-dom";
import { getMovieSeries, deleteMovieSerie } from "../../services/movieserie.service";
import Loading from "../../components/loading/Loading";
import ErrorPopup from "../../components/errorpopup/ErrorPopup";
import Message from "../../components/message/Message";

export default function MovieList() {
  const [movies, setMovies] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    fetchMovies();
  }, []);

  const fetchMovies = async () => {
    setLoading(true);
    try {
      const res = await getMovieSeries();
      setMovies(res.data || []);
    } catch (err) {
      setError("Failed to load movies and series.");
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id, title) => {
    if (!window.confirm(`Are you sure you want to delete "${title}"?`)) return;
    try {
      await deleteMovieSerie(id);
      setMovies((prev) => prev.filter((m) => m.id !== id));
      setSuccess(`"${title}" deleted successfully.`);
    } catch (err) {
      setError("Failed to delete movie/serie.");
    }
  };

  if (loading) return <Loading message="Loading movies & series" />;

  return (
    <div className="min-h-screen bg-[#1f2029] text-white p-6">
      <div className="max-w-6xl mx-auto">
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-3xl font-bold">Movies & Series</h1>
          <Link
            to="/admin/movies/new"
            className="bg-[#ffeba7] text-[#102770] font-semibold py-2 px-4 rounded-md hover:bg-[#102770] hover:text-[#ffeba7] transition-colors duration-200"
          >
            + Add New
          </Link>
        </div>

        <ErrorPopup message={error} onClose={() => setError("")} />
        {success && (
          <Message type="success" dismissible onDismiss={() => setSuccess("")}>
            {success}
          </Message>
        )}

        {movies.length === 0 ? (
          <div className="text-center py-20 text-gray-400">
            <p className="text-xl mb-4">No movies or series found.</p>
            <Link
              to="/admin/movies/new"
              className="text-[#ffeba7] underline hover:no-underline"
            >
              Add the first one
            </Link>
          </div>
        ) : (
          <div className="overflow-x-auto rounded-lg">
            <table className="w-full text-left">
              <thead>
                <tr className="bg-[#2a2b38] text-[#ffeba7]">
                  <th className="p-3">Cover</th>
                  <th className="p-3">Title</th>
                  <th className="p-3">Type</th>
                  <th className="p-3">Year</th>
                  <th className="p-3">Classification</th>
                  <th className="p-3">Director</th>
                  <th className="p-3 text-right">Actions</th>
                </tr>
              </thead>
              <tbody>
                {movies.map((m) => (
                  <tr
                    key={m.id}
                    className="border-b border-gray-700 hover:bg-[#2a2b38]/50 transition-colors"
                  >
                    <td className="p-3">
                      {m.cover ? (
                        <img
                          src={m.cover}
                          alt={m.title}
                          className="w-12 h-16 object-cover rounded"
                        />
                      ) : (
                        <div className="w-12 h-16 bg-gray-700 rounded flex items-center justify-center text-xs text-gray-400">
                          N/A
                        </div>
                      )}
                    </td>
                    <td className="p-3 font-medium">{m.title}</td>
                    <td className="p-3">
                      <span
                        className={`px-2 py-1 rounded text-xs font-semibold ${
                          m.msType === "movie"
                            ? "bg-blue-600/20 text-blue-400"
                            : "bg-purple-600/20 text-purple-400"
                        }`}
                      >
                        {m.msType}
                      </span>
                    </td>
                    <td className="p-3">{m.releaseYear}</td>
                    <td className="p-3">{m.classificationMS}</td>
                    <td className="p-3">{m.director}</td>
                    <td className="p-3 text-right">
                      <div className="flex gap-2 justify-end">
                        <button
                          onClick={() => navigate(`/admin/movies/${m.id}/edit`)}
                          className="bg-blue-600 text-white px-3 py-1 rounded text-sm hover:bg-blue-700 transition-colors"
                        >
                          Edit
                        </button>
                        <button
                          onClick={() => handleDelete(m.id, m.title)}
                          className="bg-red-600 text-white px-3 py-1 rounded text-sm hover:bg-red-700 transition-colors"
                        >
                          Delete
                        </button>
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}
