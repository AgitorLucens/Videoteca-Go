import { useState, useEffect } from "react";
import { Link, useLocation } from "react-router-dom";
import { useAuth } from "../../hooks/useAuth.jsx";
import { getProfile } from "../../services/profile.service.js";

const DEFAULT_AVATAR = "https://www.kindpng.com/picc/m/24-248253_user-profile-default-image-png-clipart-png-download.png";

export default function AccountDropdown() {
  const [open, setOpen] = useState(false);
  const { user, logout } = useAuth();
  const isAdmin = user?.role === "admin";
  const { pathname } = useLocation();
  const [profilePic, setProfilePic] = useState("");

  useEffect(() => {
    setOpen(false);
  }, [pathname]);

  useEffect(() => {
    if (user) {
      getProfile()
        .then((data) => setProfilePic(data.profile_picture || ""))
        .catch(() => {});
    }
  }, [user]);

  return (
    <div className="relative">
      <button
        onClick={() => setOpen(!open)}
        className="flex items-center gap-2 rounded-full bg-[#ffbf00] text-[#102770] px-3 py-2 hover:bg-[#102770] hover:text-[#ffeba7]"
      >
        <img
          src={profilePic || DEFAULT_AVATAR}
          alt="User avatar"
          className="h-8 w-8 rounded-full object-cover"
        />
        <span className="hidden sm:block font-medium">My Account</span>
      </button>

      {open && (
        <div className="absolute right-0 mt-2 w-48 rounded-xl bg-[#ffbf00] text-[#102770] shadow-lg ring-1 ring-gray-200 z-50">
          <ul className="py-2 text-sm text-gray-700">

            <li>
              <Link
                to="/profile"
                onClick={() => setOpen(false)}
                className="block px-4 py-2 hover:bg-gray-100"
              >
                Profile
              </Link>
            </li>
            {isAdmin && (
              <li>
                <Link
                  to="/admin/movies"
                  onClick={() => setOpen(false)}
                  className="block px-4 py-2 hover:bg-gray-100"
                >
                  Movies & Series
                </Link>
              </li>
            )}
            <li>
              <button
                onClick={() => {
                  setOpen(false);
                  logout();
                }}
                className="w-full text-left block px-4 py-2 hover:bg-gray-100"
              >
                Logout
              </button>
            </li>
          </ul>
        </div>
      )}
    </div>
  );
}
