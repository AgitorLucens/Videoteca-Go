import { useState, useEffect } from "react";
import { getAdmins, createAdmin, updateAdminPassword } from "../../services/superadmin.service";
import Loading from "../../components/loading/Loading";
import ErrorPopup from "../../components/errorpopup/ErrorPopup";
import Message from "../../components/message/Message";

export default function SuperAdmin() {
  const [admins, setAdmins] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");
  const [submitting, setSubmitting] = useState(false);
  const [form, setForm] = useState({ name: "", email: "", username: "", password: "", passwordconfirm: "" });
  const [passwordModal, setPasswordModal] = useState(null);
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  useEffect(() => {
    fetchAdmins();
  }, []);

  const fetchAdmins = async () => {
    setLoading(true);
    try {
      const res = await getAdmins();
      setAdmins(res.data || []);
    } catch {
      setError("Failed to load admins.");
    } finally {
      setLoading(false);
    }
  };

  const handleChange = (e) => {
    setForm((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const handleCreate = async (e) => {
    e.preventDefault();
    if (form.password !== form.passwordconfirm) {
      setError("Passwords do not match.");
      return;
    }
    setSubmitting(true);
    try {
      await createAdmin({
        name: form.name,
        email: form.email,
        username: form.username,
        password: form.password,
        passwordconfirm: form.passwordconfirm,
      });
      setForm({ name: "", email: "", username: "", password: "", passwordconfirm: "" });
      setSuccess("Admin created successfully.");
      fetchAdmins();
    } catch {
      setError("Failed to create admin.");
    } finally {
      setSubmitting(false);
    }
  };

  const handlePasswordChange = async (e) => {
    e.preventDefault();
    if (newPassword !== confirmPassword) {
      setError("Passwords do not match.");
      return;
    }
    setSubmitting(true);
    try {
      await updateAdminPassword(passwordModal.ID, newPassword);
      setSuccess(`Password updated for ${passwordModal.Name}.`);
      setPasswordModal(null);
      setNewPassword("");
      setConfirmPassword("");
    } catch {
      setError("Failed to update password.");
    } finally {
      setSubmitting(false);
    }
  };

  if (loading) return <Loading message="Loading admins" />;

  return (
    <div className="min-h-screen bg-[#1f2029] text-white p-6">
      <div className="max-w-4xl mx-auto">
        <h1 className="text-3xl font-bold mb-6">Super Admin</h1>

        <ErrorPopup message={error} onClose={() => setError("")} />
        {success && (
          <Message type="success" dismissible onDismiss={() => setSuccess("")}>
            {success}
          </Message>
        )}

        <form onSubmit={handleCreate} className="mb-8 bg-[#2a2b38] rounded-lg p-5">
          <h2 className="text-lg font-semibold mb-4 text-[#ffeba7]">Create Admin</h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm text-gray-400 mb-1">Name</label>
              <input
                name="name"
                value={form.name}
                onChange={handleChange}
                className="w-full bg-[#1f2029] border border-gray-600 rounded-md px-3 py-2 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7]"
                required
              />
            </div>
            <div>
              <label className="block text-sm text-gray-400 mb-1">Username</label>
              <input
                name="username"
                value={form.username}
                onChange={handleChange}
                className="w-full bg-[#1f2029] border border-gray-600 rounded-md px-3 py-2 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7]"
                required
              />
            </div>
            <div>
              <label className="block text-sm text-gray-400 mb-1">Email</label>
              <input
                name="email"
                type="email"
                value={form.email}
                onChange={handleChange}
                className="w-full bg-[#1f2029] border border-gray-600 rounded-md px-3 py-2 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7]"
                required
              />
            </div>
            <div>
              <label className="block text-sm text-gray-400 mb-1">Password</label>
              <input
                name="password"
                type="password"
                value={form.password}
                onChange={handleChange}
                className="w-full bg-[#1f2029] border border-gray-600 rounded-md px-3 py-2 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7]"
                required
              />
            </div>
            <div>
              <label className="block text-sm text-gray-400 mb-1">Confirm Password</label>
              <input
                name="passwordconfirm"
                type="password"
                value={form.passwordconfirm}
                onChange={handleChange}
                className="w-full bg-[#1f2029] border border-gray-600 rounded-md px-3 py-2 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7]"
                required
              />
            </div>
          </div>
          <button
            type="submit"
            disabled={submitting}
            className="mt-4 bg-[#ffeba7] text-[#102770] font-semibold py-2 px-6 rounded-md hover:bg-[#102770] hover:text-[#ffeba7] transition-colors duration-200 disabled:opacity-50"
          >
            {submitting ? "Creating..." : "Create Admin"}
          </button>
        </form>

        <h2 className="text-xl font-semibold mb-4 text-[#ffeba7]">Existing Admins</h2>
        {admins.length === 0 ? (
          <div className="text-center py-10 text-gray-400">
            <p className="text-xl">No admins found.</p>
          </div>
        ) : (
          <div className="overflow-x-auto rounded-lg">
            <table className="w-full text-left">
              <thead>
                <tr className="bg-[#2a2b38] text-[#ffeba7]">
                  <th className="p-3">ID</th>
                  <th className="p-3">Name</th>
                  <th className="p-3">Username</th>
                  <th className="p-3">Email</th>
                  <th className="p-3">Actions</th>
                </tr>
              </thead>
              <tbody>
                {admins.map((a) => (
                  <tr
                    key={a.ID}
                    className="border-b border-gray-700 hover:bg-[#2a2b38]/50 transition-colors"
                  >
                    <td className="p-3 text-gray-400">{a.ID}</td>
                    <td className="p-3 font-medium">{a.Name}</td>
                    <td className="p-3">{a.UserName}</td>
                    <td className="p-3">{a.Email}</td>
                    <td className="p-3">
                      <button
                        onClick={() => {
                          setPasswordModal(a);
                          setNewPassword("");
                          setConfirmPassword("");
                        }}
                        className="text-sm bg-[#ffeba7] text-[#102770] font-semibold px-3 py-1 rounded-md hover:bg-[#102770] hover:text-[#ffeba7] transition-colors"
                      >
                        Change Password
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}

        {passwordModal && (
          <div className="fixed inset-0 bg-black/50 flex items-center justify-center z-50" onClick={() => setPasswordModal(null)}>
            <div className="bg-[#2a2b38] rounded-lg p-6 w-full max-w-sm mx-4" onClick={(e) => e.stopPropagation()}>
              <h3 className="text-lg font-semibold mb-2 text-[#ffeba7]">Change Password</h3>
              <p className="text-sm text-gray-400 mb-4">for {passwordModal.Name} ({passwordModal.UserName})</p>
              <form onSubmit={handlePasswordChange} className="flex flex-col gap-3">
                <input
                  type="password"
                  placeholder="New Password"
                  value={newPassword}
                  onChange={(e) => setNewPassword(e.target.value)}
                  className="w-full bg-[#1f2029] border border-gray-600 rounded-md px-3 py-2 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7]"
                  required
                />
                <input
                  type="password"
                  placeholder="Confirm Password"
                  value={confirmPassword}
                  onChange={(e) => setConfirmPassword(e.target.value)}
                  className="w-full bg-[#1f2029] border border-gray-600 rounded-md px-3 py-2 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7]"
                  required
                />
                <div className="flex gap-2 justify-end mt-2">
                  <button
                    type="button"
                    onClick={() => setPasswordModal(null)}
                    className="px-4 py-2 text-sm text-gray-400 hover:text-white transition-colors"
                  >
                    Cancel
                  </button>
                  <button
                    type="submit"
                    disabled={submitting}
                    className="bg-[#ffeba7] text-[#102770] font-semibold px-4 py-2 rounded-md hover:bg-[#102770] hover:text-[#ffeba7] transition-colors disabled:opacity-50"
                  >
                    {submitting ? "Updating..." : "Update"}
                  </button>
                </div>
              </form>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
