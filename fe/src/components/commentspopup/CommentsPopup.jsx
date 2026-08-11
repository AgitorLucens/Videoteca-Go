import { useState, useEffect } from "react";
import { getComments, createComment, deleteComment } from "../../services/comment.service";
import { useAuth } from "../../hooks/useAuth";

const DEFAULT_AVATAR = "https://www.kindpng.com/picc/m/24-248253_user-profile-default-image-png-clipart-png-download.png";

export default function CommentsPopup({ msId, onClose }) {
  const { user } = useAuth();
  const [comments, setComments] = useState([]);
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0);
  const [loading, setLoading] = useState(true);
  const [newComment, setNewComment] = useState("");
  const [submitting, setSubmitting] = useState(false);

  const fetch = async (p) => {
    setLoading(true);
    try {
      const res = await getComments(msId, p);
      setComments(res.comments || []);
      setTotalPages(res.totalPages || 0);
    } catch {
      setComments([]);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetch(page);
  }, [msId, page]);

  const handleSubmit = async () => {
    if (!newComment.trim()) return;
    setSubmitting(true);
    try {
      await createComment(msId, newComment.trim());
      setNewComment("");
      setPage(1);
      fetch(1);
    } catch {
      alert("Failed to add comment");
    } finally {
      setSubmitting(false);
    }
  };

  const handleDelete = async (commentId) => {
    try {
      await deleteComment(commentId);
      fetch(page);
    } catch {
      alert("Failed to delete comment");
    }
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center z-50 bg-black/50" onClick={onClose}>
      <div
        className="bg-[#2a2b38] rounded-xl shadow-lg p-6 w-full max-w-lg max-h-[80vh] flex flex-col"
        onClick={(e) => e.stopPropagation()}
      >
        <div className="flex justify-between items-center mb-4">
          <h2 className="text-xl font-bold text-white">Comments</h2>
          <button onClick={onClose} className="text-gray-400 hover:text-white text-xl">&times;</button>
        </div>

        {user && user.role === "user" && (
          <div className="flex gap-2 mb-4">
            <input
              value={newComment}
              onChange={(e) => setNewComment(e.target.value)}
              placeholder="Write a comment..."
              className="flex-1 bg-[#1f2029] text-white rounded-md px-3 py-2 text-sm outline-none border border-gray-700 focus:border-[#ffeba7]"
              onKeyDown={(e) => { if (e.key === "Enter") handleSubmit(); }}
            />
            <button
              onClick={handleSubmit}
              disabled={submitting || !newComment.trim()}
              className="bg-[#ffeba7] text-[#102770] font-semibold px-4 py-2 rounded-md text-sm disabled:opacity-40 hover:bg-[#e6d492] transition-colors"
            >
              {submitting ? "..." : "Post"}
            </button>
          </div>
        )}

        <div className="flex-1 overflow-y-auto space-y-3 min-h-0">
          {loading ? (
            <p className="text-gray-400 text-center py-8">Loading...</p>
          ) : comments.length === 0 ? (
            <p className="text-gray-500 text-center py-8">No comments yet.</p>
          ) : (
            comments.map((c) => (
              <div key={c.id} className="bg-[#1f2029] rounded-lg p-3 flex items-start justify-between gap-2">
                <div className="flex items-start gap-3 min-w-0">
                  <img
                    src={c.photo || DEFAULT_AVATAR}
                    alt=""
                    className="w-7 h-7 rounded-full object-cover flex-shrink-0 mt-0.5"
                  />
                  <div>
                    <p className="text-xs text-[#ffeba7] mb-1">{c.appUser}</p>
                    <p className="text-gray-200 text-sm break-words">{c.comment}</p>
                  </div>
                </div>
                {user && (user.role === "admin" || user.role === "superadmin") && (
                  <button
                    onClick={() => handleDelete(c.id)}
                    className="text-red-400 hover:text-red-300 text-xs whitespace-nowrap mt-1 flex-shrink-0"
                  >
                    Delete
                  </button>
                )}
              </div>
            ))
          )}
        </div>

        {totalPages > 1 && (
          <div className="flex justify-center items-center gap-3 mt-4 pt-4 border-t border-gray-700">
            <button
              onClick={() => setPage((p) => Math.max(1, p - 1))}
              disabled={page <= 1}
              className="px-3 py-1 rounded bg-gray-700 text-white text-sm disabled:opacity-40 hover:bg-gray-600 transition"
            >
              Prev
            </button>
            <span className="text-sm text-gray-400">
              {page} / {totalPages}
            </span>
            <button
              onClick={() => setPage((p) => Math.min(totalPages, p + 1))}
              disabled={page >= totalPages}
              className="px-3 py-1 rounded bg-gray-700 text-white text-sm disabled:opacity-40 hover:bg-gray-600 transition"
            >
              Next
            </button>
          </div>
        )}
      </div>
    </div>
  );
}
