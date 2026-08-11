import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { forgotPassword, resetPassword } from "../../services/forgotpassword.service";
import InputIcon from "../../components/inputicon/InputIcon.jsx";
import { IconArroa, IconLock } from "../../components/sprite/Sprite.jsx";
import ErrorPopup from "../../components/errorpopup/ErrorPopup.jsx";

export default function ForgotPassword() {
  const [step, setStep] = useState(1);
  const [email, setEmail] = useState("");
  const [newPassword, setNewPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleVerifyEmail = async (e) => {
    e.preventDefault();
    setError("");
    setLoading(true);
    try {
      await forgotPassword(email);
      setStep(2);
    } catch (err) {
      setError(err.response?.data?.error || "No account found with that email");
    } finally {
      setLoading(false);
    }
  };

  const handleResetPassword = async (e) => {
    e.preventDefault();
    setError("");
    if (newPassword !== confirmPassword) {
      setError("Passwords do not match");
      return;
    }
    setLoading(true);
    try {
      await resetPassword(email, newPassword);
      setSuccess("Password reset successfully! Redirecting to login...");
      setTimeout(() => navigate("/"), 2000);
    } catch (err) {
      setError(err.response?.data?.error || "Failed to reset password");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-[#1f2029] flex items-center justify-center px-4">
      <div className="w-full max-w-[400px] bg-[#2a2b38] rounded-lg p-6 flex flex-col items-center">
        <h4 className="text-2xl font-semibold mb-4 text-[#c4c3ca]">
          {step === 1 ? "Forgot Password" : "Set New Password"}
        </h4>

        {step === 1 ? (
          <form className="w-full flex flex-col gap-4" onSubmit={handleVerifyEmail}>
            <InputIcon
              type="email"
              placeholder="Your Email"
              icon={IconArroa}
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            <button
              type="submit"
              disabled={loading}
              className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-2 rounded-md h-[44px] w-[118px] mx-auto"
            >
              {loading ? "Checking..." : "Verify Email"}
            </button>
          </form>
        ) : (
          <form className="w-full flex flex-col gap-4" onSubmit={handleResetPassword}>
            <InputIcon
              type="password"
              placeholder="New Password"
              icon={IconLock}
              value={newPassword}
              onChange={(e) => setNewPassword(e.target.value)}
            />
            <InputIcon
              type="password"
              placeholder="Confirm Password"
              icon={IconLock}
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
            />
            <button
              type="submit"
              disabled={loading}
              className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-2 rounded-md h-[44px] w-[118px] mx-auto"
            >
              {loading ? "Resetting..." : "Reset Password"}
            </button>
          </form>
        )}

        <div className="w-full flex justify-center p-4">
          <Link to="/" className="text-sm text-[#c4c3ca] underline hover:text-[#ffeba7] hover:no-underline transition-all duration-200">
            Back to Login
          </Link>
        </div>

        <ErrorPopup message={error} onClose={() => setError("")} />

        {success && (
          <p className="text-green-400 text-sm mt-2">{success}</p>
        )}
      </div>
    </div>
  );
}
