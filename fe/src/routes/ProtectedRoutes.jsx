import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../hooks/useAuth";

export default function ProtectedRoute({ children }) {
  const { user, loading } = useAuth();
  if (loading) {
    return <div className="h-screen flex justify-center items-center">Loading...</div>;
  }
  if (!user) {
    return <Navigate to="/" replace />;
  } 
  return children ? children : <Outlet />;
}
