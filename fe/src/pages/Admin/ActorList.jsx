import { useState, useEffect, useRef } from "react";
import { getActors, createActor, updateActor, deleteActor } from "../../services/actor.service";
import Loading from "../../components/loading/Loading";
import ErrorPopup from "../../components/errorpopup/ErrorPopup";
import Message from "../../components/message/Message";

export default function ActorList() {
  const [actors, setActors] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [photo, setPhoto] = useState(null);
  const [photoPreview, setPhotoPreview] = useState(null);
  const [submitting, setSubmitting] = useState(false);
  const [editingId, setEditingId] = useState(null);
  const [editFirstName, setEditFirstName] = useState("");
  const [editLastName, setEditLastName] = useState("");
  const [editPhoto, setEditPhoto] = useState(null);
  const [editPhotoPreview, setEditPhotoPreview] = useState(null);
  const fileInputRef = useRef(null);
  const editFileInputRef = useRef(null);

  useEffect(() => {
    fetchActors();
  }, []);

  const fetchActors = async () => {
    setLoading(true);
    try {
      const res = await getActors();
      setActors(res.data || []);
    } catch {
      setError("Failed to load actors.");
    } finally {
      setLoading(false);
    }
  };

  const readFileAsBase64 = (file) => {
    return new Promise((resolve, reject) => {
      const reader = new FileReader();
      reader.onload = () => resolve(reader.result);
      reader.onerror = reject;
      reader.readAsDataURL(file);
    });
  };

  const handlePhotoChange = async (e, setPhotoFn, setPreviewFn) => {
    const file = e.target.files[0];
    if (!file) return;
    const base64 = await readFileAsBase64(file);
    setPhotoFn(base64);
    setPreviewFn(base64);
  };

  const resetForm = () => {
    setFirstName("");
    setLastName("");
    setPhoto(null);
    setPhotoPreview(null);
    if (fileInputRef.current) fileInputRef.current.value = "";
  };

  const handleCreate = async (e) => {
    e.preventDefault();
    if (!firstName.trim() || !lastName.trim()) return;
    setSubmitting(true);
    try {
      await createActor({
        firstName: firstName.trim(),
        lastName: lastName.trim(),
        photo: photo || null,
      });
      resetForm();
      setSuccess("Actor created successfully.");
      fetchActors();
    } catch {
      setError("Failed to create actor.");
    } finally {
      setSubmitting(false);
    }
  };

  const handleEdit = (actor) => {
    setEditingId(actor.id);
    setEditFirstName(actor.firstName);
    setEditLastName(actor.lastName);
    setEditPhoto(null);
    setEditPhotoPreview(actor.photo || null);
    if (editFileInputRef.current) editFileInputRef.current.value = "";
  };

  const handleUpdate = async (id) => {
    if (!editFirstName.trim() || !editLastName.trim()) return;
    setSubmitting(true);
    try {
      await updateActor(id, {
        firstName: editFirstName.trim(),
        lastName: editLastName.trim(),
        photo: editPhoto !== null ? editPhoto : editPhotoPreview,
      });
      setEditingId(null);
      setEditFirstName("");
      setEditLastName("");
      setEditPhoto(null);
      setEditPhotoPreview(null);
      setSuccess("Actor updated successfully.");
      fetchActors();
    } catch {
      setError("Failed to update actor.");
    } finally {
      setSubmitting(false);
    }
  };

  const handleDelete = async (id, name) => {
    if (!window.confirm(`Are you sure you want to delete "${name}"?`)) return;
    try {
      await deleteActor(id);
      setActors((prev) => prev.filter((a) => a.id !== id));
      setSuccess(`"${name}" deleted successfully.`);
    } catch {
      setError("Failed to delete actor.");
    }
  };

  if (loading) return <Loading message="Loading actors" />;

  return (
    <div className="min-h-screen bg-[#1f2029] text-white p-6">
      <div className="max-w-6xl mx-auto">
        <h1 className="text-3xl font-bold mb-6">Actors</h1>

        <ErrorPopup message={error} onClose={() => setError("")} />
        {success && (
          <Message type="success" dismissible onDismiss={() => setSuccess("")}>
            {success}
          </Message>
        )}

        <form onSubmit={handleCreate} className="mb-6 bg-[#2a2b38] rounded-lg p-4">
          <h2 className="text-lg font-semibold mb-3 text-[#ffeba7]">Add New Actor</h2>
          <div className="flex flex-wrap gap-3 items-end">
            <div className="flex-1 min-w-[150px]">
              <label className="block text-sm text-gray-400 mb-1">First Name</label>
              <input
                type="text"
                value={firstName}
                onChange={(e) => setFirstName(e.target.value)}
                className="w-full bg-[#1f2029] border border-gray-600 rounded-md px-3 py-2 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7]"
                required
              />
            </div>
            <div className="flex-1 min-w-[150px]">
              <label className="block text-sm text-gray-400 mb-1">Last Name</label>
              <input
                type="text"
                value={lastName}
                onChange={(e) => setLastName(e.target.value)}
                className="w-full bg-[#1f2029] border border-gray-600 rounded-md px-3 py-2 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7]"
                required
              />
            </div>
            <div className="flex-1 min-w-[200px]">
              <label className="block text-sm text-gray-400 mb-1">Photo</label>
              <input
                ref={fileInputRef}
                type="file"
                accept="image/*"
                onChange={(e) => handlePhotoChange(e, setPhoto, setPhotoPreview)}
                className="w-full text-sm text-gray-400 file:mr-3 file:py-2 file:px-3 file:rounded-md file:border-0 file:bg-[#ffeba7] file:text-[#102770] file:font-semibold file:text-sm hover:file:bg-[#102770] hover:file:text-[#ffeba7] file:transition-colors file:duration-200 file:cursor-pointer"
              />
            </div>
            <button
              type="submit"
              disabled={submitting}
              className="bg-[#ffeba7] text-[#102770] font-semibold py-2 px-6 rounded-md hover:bg-[#102770] hover:text-[#ffeba7] transition-colors duration-200 disabled:opacity-50"
            >
              {submitting ? "Adding..." : "Add Actor"}
            </button>
          </div>
          {photoPreview && (
            <div className="mt-3">
              <img src={photoPreview} alt="Preview" className="w-16 h-16 object-cover rounded" />
            </div>
          )}
        </form>

        {actors.length === 0 ? (
          <div className="text-center py-20 text-gray-400">
            <p className="text-xl mb-4">No actors found.</p>
          </div>
        ) : (
          <div className="overflow-x-auto rounded-lg">
            <table className="w-full text-left">
              <thead>
                <tr className="bg-[#2a2b38] text-[#ffeba7]">
                  <th className="p-3">Photo</th>
                  <th className="p-3">First Name</th>
                  <th className="p-3">Last Name</th>
                  <th className="p-3 text-right">Actions</th>
                </tr>
              </thead>
              <tbody>
                {actors.map((a) => (
                  <tr
                    key={a.id}
                    className="border-b border-gray-700 hover:bg-[#2a2b38]/50 transition-colors"
                  >
                    <td className="p-3">
                      {editingId === a.id ? (
                        <div className="flex items-center gap-2">
                          {(editPhotoPreview || editPhoto) ? (
                            <img
                              src={editPhoto || editPhotoPreview}
                              alt="Preview"
                              className="w-10 h-10 object-cover rounded"
                            />
                          ) : (
                            <div className="w-10 h-10 bg-gray-700 rounded flex items-center justify-center text-xs text-gray-400">
                              N/A
                            </div>
                          )}
                          <input
                            ref={editFileInputRef}
                            type="file"
                            accept="image/*"
                            onChange={(e) => handlePhotoChange(e, setEditPhoto, setEditPhotoPreview)}
                            className="text-sm text-gray-400 file:mr-2 file:py-1 file:px-2 file:rounded file:border-0 file:bg-[#ffeba7] file:text-[#102770] file:font-semibold file:text-xs hover:file:bg-[#102770] hover:file:text-[#ffeba7] file:transition-colors file:duration-200 file:cursor-pointer"
                          />
                        </div>
                      ) : a.photo ? (
                        <img src={a.photo} alt={`${a.firstName} ${a.lastName}`} className="w-10 h-10 object-cover rounded" />
                      ) : (
                        <div className="w-10 h-10 bg-gray-700 rounded flex items-center justify-center text-xs text-gray-400">
                          N/A
                        </div>
                      )}
                    </td>
                    <td className="p-3">
                      {editingId === a.id ? (
                        <input
                          type="text"
                          value={editFirstName}
                          onChange={(e) => setEditFirstName(e.target.value)}
                          className="bg-[#1f2029] border border-gray-600 rounded px-3 py-1 text-white focus:outline-none focus:border-[#ffeba7]"
                          autoFocus
                        />
                      ) : (
                        <span className="font-medium">{a.firstName}</span>
                      )}
                    </td>
                    <td className="p-3">
                      {editingId === a.id ? (
                        <input
                          type="text"
                          value={editLastName}
                          onChange={(e) => setEditLastName(e.target.value)}
                          className="bg-[#1f2029] border border-gray-600 rounded px-3 py-1 text-white focus:outline-none focus:border-[#ffeba7]"
                        />
                      ) : (
                        a.lastName
                      )}
                    </td>
                    <td className="p-3 text-right">
                      <div className="flex gap-2 justify-end">
                        {editingId === a.id ? (
                          <>
                            <button
                              onClick={() => handleUpdate(a.id)}
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
                              onClick={() => handleEdit(a)}
                              className="bg-blue-600 text-white px-3 py-1 rounded text-sm hover:bg-blue-700 transition-colors"
                            >
                              Edit
                            </button>
                            <button
                              onClick={() => handleDelete(a.id, `${a.firstName} ${a.lastName}`)}
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
