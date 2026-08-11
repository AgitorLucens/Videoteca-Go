import { useState, useEffect } from "react";
import { getGenres, createGenre, updateGenre, deleteGenre } from "../../services/genre.service";
import Loading from "../../components/loading/Loading";
import ErrorPopup from "../../components/errorpopup/ErrorPopup";
import Message from "../../components/message/Message";

export default function TagList() {
  const [genres, setGenres] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");
  const [name, setName] = useState("");
  const [editingId, setEditingId] = useState(null);
  const [editName, setEditName] = useState("");
  const [submitting, setSubmitting] = useState(false);

  useEffect(() => {
    fetchGenres();
  }, []);

  const fetchGenres = async () => {
    setLoading(true);
    try {
      const res = await getGenres();
      setGenres(res.data || []);
    } catch {
      setError("Failed to load tags.");
    } finally {
      setLoading(false);
    }
  };

  const handleCreate = async (e) => {
    e.preventDefault();
    if (!name.trim()) return;
    setSubmitting(true);
    try {
      await createGenre({ name: name.trim() });
      setName("");
      setSuccess("Tag created successfully.");
      fetchGenres();
    } catch {
      setError("Failed to create tag.");
    } finally {
      setSubmitting(false);
    }
  };

  const handleEdit = (genre) => {
    setEditingId(genre.id);
    setEditName(genre.name);
  };

  const handleUpdate = async (id) => {
    if (!editName.trim()) return;
    setSubmitting(true);
    try {
      await updateGenre(id, { name: editName.trim() });
      setEditingId(null);
      setEditName("");
      setSuccess("Tag updated successfully.");
      fetchGenres();
    } catch {
      setError("Failed to update tag.");
    } finally {
      setSubmitting(false);
    }
  };

  const handleDelete = async (id, tagName) => {
    if (!window.confirm(`Are you sure you want to delete "${tagName}"?`)) return;
    try {
      await deleteGenre(id);
      setGenres((prev) => prev.filter((g) => g.id !== id));
      setSuccess(`"${tagName}" deleted successfully.`);
    } catch {
      setError("Failed to delete tag.");
    }
  };

  if (loading) return <Loading message="Loading tags" />;

  return (
    <div className="min-h-screen bg-[#1f2029] text-white p-6">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-3xl font-bold mb-6">Tags</h1>

        <ErrorPopup message={error} onClose={() => setError("")} />
        {success && (
          <Message type="success" dismissible onDismiss={() => setSuccess("")}>
            {success}
          </Message>
        )}

        <form onSubmit={handleCreate} className="mb-6 flex gap-3">
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="Tag name"
            className="flex-1 bg-[#2a2b38] border border-gray-600 rounded-md px-4 py-2 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7]"
            required
          />
          <button
            type="submit"
            disabled={submitting}
            className="bg-[#ffeba7] text-[#102770] font-semibold py-2 px-6 rounded-md hover:bg-[#102770] hover:text-[#ffeba7] transition-colors duration-200 disabled:opacity-50"
          >
            {submitting ? "Adding..." : "Add Tag"}
          </button>
        </form>

        {genres.length === 0 ? (
          <div className="text-center py-20 text-gray-400">
            <p className="text-xl mb-4">No tags found.</p>
          </div>
        ) : (
          <div className="overflow-x-auto rounded-lg">
            <table className="w-full text-left">
              <thead>
                <tr className="bg-[#2a2b38] text-[#ffeba7]">
                  <th className="p-3">ID</th>
                  <th className="p-3">Name</th>
                  <th className="p-3 text-right">Actions</th>
                </tr>
              </thead>
              <tbody>
                {genres.map((g) => (
                  <tr
                    key={g.id}
                    className="border-b border-gray-700 hover:bg-[#2a2b38]/50 transition-colors"
                  >
                    <td className="p-3 text-gray-400">{g.id}</td>
                    <td className="p-3">
                      {editingId === g.id ? (
                        <input
                          type="text"
                          value={editName}
                          onChange={(e) => setEditName(e.target.value)}
                          className="bg-[#1f2029] border border-gray-600 rounded px-3 py-1 text-white focus:outline-none focus:border-[#ffeba7]"
                          autoFocus
                          onKeyDown={(e) => {
                            if (e.key === "Enter") handleUpdate(g.id);
                            if (e.key === "Escape") setEditingId(null);
                          }}
                        />
                      ) : (
                        <span className="font-medium">{g.name}</span>
                      )}
                    </td>
                    <td className="p-3 text-right">
                      <div className="flex gap-2 justify-end">
                        {editingId === g.id ? (
                          <>
                            <button
                              onClick={() => handleUpdate(g.id)}
                              disabled={submitting}
                              className="bg-green-600 text-white px-3 py-1 rounded text-sm hover:bg-green-700 transition-colors disabled:opacity-50"
                            >
                              Save
                            </button>
                            <button
                              onClick={() => setEditingId(null)}
                              className="bg-gray-600 text-white px-3 py-1 rounded text-sm hover:bg-gray-700 transition-colors"
                            >
                              Cancel
                            </button>
                          </>
                        ) : (
                          <>
                            <button
                              onClick={() => handleEdit(g)}
                              className="bg-blue-600 text-white px-3 py-1 rounded text-sm hover:bg-blue-700 transition-colors"
                            >
                              Edit
                            </button>
                            <button
                              onClick={() => handleDelete(g.id, g.name)}
                              className="bg-red-600 text-white px-3 py-1 rounded text-sm hover:bg-red-700 transition-colors"
                            >
                              Delete
                            </button>
                          </>
                        )}
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
