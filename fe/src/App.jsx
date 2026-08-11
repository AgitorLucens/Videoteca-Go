import './App.css'
import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom'
import { AuthProvider } from "./hooks/useAuth";
import MainLayout from './components/mainlayout/MainLayout.jsx';
import ProtectedRoute from "./routes/ProtectedRoutes.jsx";

import Auth from './pages/Auth/Auth.jsx'
import ForgotPassword from './pages/ForgotPassword/ForgotPassword.jsx'
import Home from './pages/Home/Home.jsx'
import Main from "./pages/Main/Main.jsx"
import Profile from "./pages/Profile/Profile.jsx"
import MovieList from "./pages/Admin/MovieList.jsx"
import MovieSerieForm from "./pages/Admin/MovieSerieForm.jsx"
import TagList from "./pages/Admin/TagList.jsx"
import ActorList from "./pages/Admin/ActorList.jsx"
import SuperAdmin from "./pages/SuperAdmin/SuperAdmin.jsx"
import MovieDetail from "./pages/MovieDetail/MovieDetail.jsx"

function App() {
  return (
    <BrowserRouter>
      <AuthProvider>
        <Routes>
          <Route path="/" element={<Auth />} />
          <Route path="/forgot-password" element={<ForgotPassword />} />

          <Route element={
            <ProtectedRoute>
              <MainLayout />
            </ProtectedRoute>
          }>
            <Route path="/home" element={<Home />} />
            <Route path="/profile" element={<Profile />} />
            <Route path="/main" element={<Main />} />
            <Route path="/admin/movies" element={<MovieList />} />
            <Route path="/admin/movies/new" element={<MovieSerieForm />} />
            <Route path="/admin/movies/:id/edit" element={<MovieSerieForm />} />
<Route path="/admin/tags" element={<TagList />} />
<Route path="/admin/actors" element={<ActorList />} />
            <Route path="/superadmin" element={<SuperAdmin />} />
            <Route path="/movies/:id" element={<MovieDetail />} />
          </Route>

          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  );
}

export default App;
