import { useState } from "react";
import { useAuth } from "../../hooks/useAuth.jsx";
import { updateEmail, updatePassword, updateUsername, updatePhoto } from "../../services/profile.service.js";

function ChangePictureForm({ onUpdate }) {
  const [file, setFile] = useState(null);
  const [msg, setMsg] = useState("");
  const [err, setErr] = useState("");

  const handleFile = (e) => {
    const f = e.target.files[0];
    if (f) setFile(f);
  };

  const handleSubmit = async () => {
    if (!file) { setErr("Select an image"); return; }
    setErr(""); setMsg("");
    const reader = new FileReader();
    reader.onload = async () => {
      try {
        await updatePhoto(reader.result);
        setMsg("Photo updated!");
        onUpdate?.();
      } catch (e) {
        setErr(e.response?.data?.error || "Failed to update photo");
      }
    };
    reader.readAsDataURL(file);
  };

  return (
    <div>
      <h4 className="text-lg font-semibold text-[#c4c3ca] mb-4">Change Picture</h4>
      <input type="file" accept="image/*" onChange={handleFile} className="w-full text-gray-400 text-sm mb-3 file:mr-3 file:py-1.5 file:px-3 file:rounded file:border-0 file:bg-[#ffeba7] file:text-[#102770] file:font-semibold file:cursor-pointer hover:file:bg-[#102770] hover:file:text-[#ffeba7]" />
      <button onClick={handleSubmit} className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-2 px-4 rounded-md transition-all text-sm">Upload</button>
      {msg && <p className="text-green-400 text-sm mt-2">{msg}</p>}
      {err && <p className="text-red-400 text-sm mt-2">{err}</p>}
    </div>
  );
}

function ChangePasswordForm({ onUpdate }) {
  const [current, setCurrent] = useState("");
  const [newPw, setNewPw] = useState("");
  const [confirm, setConfirm] = useState("");
  const [msg, setMsg] = useState("");
  const [err, setErr] = useState("");

  const handleSubmit = async () => {
    setErr(""); setMsg("");
    if (!current || !newPw) { setErr("Fill all fields"); return; }
    if (newPw.length < 8) { setErr("New password must be at least 8 characters"); return; }
    if (newPw !== confirm) { setErr("Passwords do not match"); return; }
    try {
      await updatePassword(current, newPw);
      setMsg("Password updated!");
      setCurrent(""); setNewPw(""); setConfirm("");
      onUpdate?.();
    } catch (e) {
      setErr(e.response?.data?.error || "Failed to update password");
    }
  };

  return (
    <div>
      <h4 className="text-lg font-semibold text-[#c4c3ca] mb-4">Change Password</h4>
      <input type="password" placeholder="Current password" value={current} onChange={(e) => setCurrent(e.target.value)} className="w-full pl-3 pr-4 py-2.5 rounded-md bg-[#1f2029] border border-gray-600 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7] transition-colors text-sm mb-2" />
      <input type="password" placeholder="New password" value={newPw} onChange={(e) => setNewPw(e.target.value)} className="w-full pl-3 pr-4 py-2.5 rounded-md bg-[#1f2029] border border-gray-600 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7] transition-colors text-sm mb-2" />
      <input type="password" placeholder="Confirm new password" value={confirm} onChange={(e) => setConfirm(e.target.value)} className="w-full pl-3 pr-4 py-2.5 rounded-md bg-[#1f2029] border border-gray-600 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7] transition-colors text-sm mb-3" />
      <button onClick={handleSubmit} className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-2 px-4 rounded-md transition-all text-sm">Update</button>
      {msg && <p className="text-green-400 text-sm mt-2">{msg}</p>}
      {err && <p className="text-red-400 text-sm mt-2">{err}</p>}
    </div>
  );
}

function ChangeUsernameForm({ onUpdate, user }) {
  const [username, setUsername] = useState("");
  const [msg, setMsg] = useState("");
  const [err, setErr] = useState("");

  const handleSubmit = async () => {
    setErr(""); setMsg("");
    if (!username.trim()) { setErr("Username is required"); return; }
    if (username.trim().length < 3) { setErr("Username must be at least 3 characters"); return; }
    try {
      await updateUsername(username.trim());
      setMsg("Username updated!");
      onUpdate?.();
    } catch (e) {
      setErr(e.response?.data?.error || "Failed to update username");
    }
  };

  return (
    <div>
      <h4 className="text-lg font-semibold text-[#c4c3ca] mb-4">Change Username</h4>
      <p className="text-gray-400 text-sm mb-2">Current: {user?.username || "N/A"}</p>
      <input type="text" placeholder="New username" value={username} onChange={(e) => setUsername(e.target.value)} className="w-full pl-3 pr-4 py-2.5 rounded-md bg-[#1f2029] border border-gray-600 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7] transition-colors text-sm mb-3" />
      <button onClick={handleSubmit} className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-2 px-4 rounded-md transition-all text-sm">Update</button>
      {msg && <p className="text-green-400 text-sm mt-2">{msg}</p>}
      {err && <p className="text-red-400 text-sm mt-2">{err}</p>}
    </div>
  );
}

function ChangeEmailForm({ onUpdate, user }) {
  const [email, setEmail] = useState("");
  const [msg, setMsg] = useState("");
  const [err, setErr] = useState("");

  const handleSubmit = async () => {
    setErr(""); setMsg("");
    if (!email.trim()) { setErr("Email is required"); return; }
    try {
      await updateEmail(email.trim());
      setMsg("Email updated!");
      onUpdate?.();
    } catch (e) {
      setErr(e.response?.data?.error || "Failed to update email");
    }
  };

  return (
    <div>
      <h4 className="text-lg font-semibold text-[#c4c3ca] mb-4">Change Email</h4>
      <p className="text-gray-400 text-sm mb-2">Current: {user?.email || "N/A"}</p>
      <input type="email" placeholder="New email" value={email} onChange={(e) => setEmail(e.target.value)} className="w-full pl-3 pr-4 py-2.5 rounded-md bg-[#1f2029] border border-gray-600 text-white placeholder-gray-400 focus:outline-none focus:border-[#ffeba7] transition-colors text-sm mb-3" />
      <button onClick={handleSubmit} className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-2 px-4 rounded-md transition-all text-sm">Update</button>
      {msg && <p className="text-green-400 text-sm mt-2">{msg}</p>}
      {err && <p className="text-red-400 text-sm mt-2">{err}</p>}
    </div>
  );
}

export default function ProfileSlider({ isOpen, onClose, side = "right", activeView, onSelectView, onUpdate }) {
  const isLeft = side === "left";
  const { user } = useAuth();

  const renderLeftContent = () => {
    switch (activeView) {
      case "picture": return <ChangePictureForm onUpdate={onUpdate} />;
      case "password": return <ChangePasswordForm onUpdate={onUpdate} />;
      case "username": return <ChangeUsernameForm onUpdate={onUpdate} user={user} />;
      case "email": return <ChangeEmailForm onUpdate={onUpdate} user={user} />;
      default:
        return (
          <div>
            <h4 className="text-lg font-semibold text-[#c4c3ca] mb-4">Select an option</h4>
            <p className="text-gray-400 text-sm">Choose an option from the settings panel to get started.</p>
          </div>
        );
    }
  };

  return (
    <div
      className={`absolute top-0 left-0 h-full w-64 sm:w-72 md:w-80 bg-[#2a2b38] shadow-2xl z-0 transition-transform duration-300 ease-in-out overflow-y-auto rounded-lg ${
        isOpen
          ? isLeft
            ? "-translate-x-[260px] sm:-translate-x-[290px] md:-translate-x-[320px]"
            : "translate-x-[260px] sm:translate-x-[290px] md:translate-x-[400px]"
          : "translate-x-0"
      }`}
    >
      <div className="p-6">
        {isLeft ? (
          <>
            <h3 className="text-xl font-bold text-[#c4c3ca] mb-6 border-b border-gray-700 pb-2">Account Info</h3>
            {renderLeftContent()}
          </>
        ) : (
          <>
            <h3 className="text-xl font-bold text-[#c4c3ca] mb-6 border-b border-gray-700 pb-2">&#9889;&#65039; Account Settings</h3>
            <div className="flex flex-col gap-3">
              <button onClick={() => onSelectView("picture")} className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-2 rounded-md transition-all">Change Picture</button>
              <button onClick={() => onSelectView("password")} className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-2 rounded-md transition-all">Change Password</button>
              <button onClick={() => onSelectView("username")} className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-2 rounded-md transition-all">Change Username</button>
              <button onClick={() => onSelectView("email")} className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-2 rounded-md transition-all">Change Email</button>
            </div>
          </>
        )}
      </div>
    </div>
  );
}
