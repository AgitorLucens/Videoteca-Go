import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { loginService } from "../../services/login.service";
import InputIcon from "../inputicon/InputIcon.jsx";
import ErrorPopup from "../errorpopup/ErrorPopup.jsx";
import LoadingSpinner from "../loadingspinner/LoadinSpinner.jsx";
import { IconArroa, IconUser, IconLock } from "../sprite/Sprite.jsx";
import { jwtDecode } from "jwt-decode";
import { useAuth } from "../../hooks/useAuth.jsx";
import './LoginCard.css';

export default function LoginCard() {
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const { login } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    try {
      const response = await loginService({ username, password });
      const claim = jwtDecode(response.token);

      const user = {
        ...claim,
        token: response.token
      };
      await login(user);

      if (claim.role === "superadmin") {
        navigate("/superadmin", { replace: true });
      } else {
        navigate("/home", { replace: true });
      }

    } catch (error) {
      setError("Login failed. Please check your credentials.");
      console.error("Login failed:", error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="w-full h-full bg-[#2a2b38] rounded-lg p-6 flex flex-col justify-center items-center ">
      <h4 className="text-2xl font-semibold mb-4 text-[#c4c3ca]">Welcome Back</h4>
      <form className="w-full flex flex-col gap-4" onSubmit={handleSubmit}>
        <InputIcon  type="text"
                    placeholder="Your Username"
                    icon={IconArroa}
                    onChange={(e) => setUsername(e.target.value)}/> 
        <InputIcon  type="password"
                    placeholder="Your Password"
                    icon={IconLock}
                    onChange={(e) => setPassword(e.target.value)}/> 
        <button
          type="submit"
          className="bg-[#ffeba7] text-[#102770] 
              hover:bg-[#102770] hover:text-[#ffeba7] 
              font-semibold py-2 rounded-md 
              h-[44px] w-[118px]
              mx-auto
              text-transform: uppercase"
        >
          Login
        </button>
      </form>
      {/* Loading Spinner */}
      {loading && (
        <LoadingSpinner className="animate-spin h-12 w-12 bg-[#ffeba7] text-blue-500"/>
      )}
       {/* Error Popup */}
      <ErrorPopup message={error} onClose={() => setError("")} />
      <div className="w-full flex justify-center p-8">
        <Link to="/forgot-password" className="text-sm text-[#c4c3ca] underline hover:text-[#ffeba7] hover:no-underline transition-all duration-200"> Forgot your password? </Link>
      </div>
    </div>
  );
}