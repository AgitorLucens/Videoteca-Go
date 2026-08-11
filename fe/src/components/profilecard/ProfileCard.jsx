import { useState, useEffect, useCallback } from "react";
import ProfileSlider from "../profileslider/ProfileSlider.jsx";
import { IconLeftArrow, IconSettings } from "../sprite/Sprite.jsx";
import { useAuth } from "../../hooks/useAuth.jsx";
import { getProfile } from "../../services/profile.service.js";

const DEFAULT_AVATAR = "https://www.kindpng.com/picc/m/24-248253_user-profile-default-image-png-clipart-png-download.png";

export default function ProfileCard() {
  const { user } = useAuth();
  const [leftOpen, setLeftOpen] = useState(false);
  const [rightOpen, setRightOpen] = useState(false);
  const [activeView, setActiveView] = useState(null);
  const [profileData, setProfileData] = useState(null);

  const fetchProfile = useCallback(async () => {
    try {
      const data = await getProfile();
      setProfileData(data);
    } catch {
      // ignore
    }
  }, []);

  useEffect(() => {
    fetchProfile();
  }, [fetchProfile]);

  const handleSelectView = (view) => {
    setActiveView(view);
    setLeftOpen(true);
  };

  const handleUpdate = () => {
    fetchProfile();
  };

  const p = profileData || user || {};

  return (
    <div className="flex justify-center pt-10">
      <div className="relative w-full max-w-[400px] px-4">
        <ProfileSlider
          isOpen={leftOpen}
          onClose={() => setLeftOpen(false)}
          side="left"
          activeView={activeView}
          onUpdate={handleUpdate}
        />

        <div className="bg-[#2a2b38] p-6 flex flex-col rounded-lg relative z-10">
          <div className="flex mb-4 justify-between">
            <button
              onClick={() => leftOpen && setLeftOpen(false)}
              className="p-1 rounded-full text-[#c4c3ca] hover:text-white transition duration-150 focus:outline-none focus:ring-2 focus:ring-[#c4c3ca]"
              aria-label="Close info panel"
            >
              <IconLeftArrow className="h-6 w-6" />
            </button>

            <button
              onClick={() => setRightOpen((p) => !p)}
              className="p-1 rounded-full text-[#c4c3ca] hover:text-white transition duration-150 focus:outline-none focus:ring-2 focus:ring-[#c4c3ca]"
              aria-label="Open settings panel"
            >
              <IconSettings className="h-6 w-6" />
            </button>
          </div>

          <div className="flex justify-center mb-4">
            <img
              src={p.profile_picture || DEFAULT_AVATAR}
              alt="User avatar"
              className="h-35 w-35 rounded-full"
            />
          </div>

          <ul className="text-sm text-white space-y-2">
            <li className="px-4 py-2"><span className="text-[#c4c3ca]">Name:</span> {p.name || "N/A"}</li>
            <li className="px-4 py-2"><span className="text-[#c4c3ca]">Username:</span> {p.username || p.sub || "N/A"}</li>
            <li className="px-4 py-2"><span className="text-[#c4c3ca]">Email:</span> {p.email || "N/A"}</li>
          </ul>
        </div>

        <ProfileSlider
          isOpen={rightOpen}
          onClose={() => setRightOpen(false)}
          side="right"
          onSelectView={handleSelectView}
        />
      </div>
    </div>
  );
}
