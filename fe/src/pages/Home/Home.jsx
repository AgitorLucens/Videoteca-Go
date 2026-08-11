import { useState, useEffect } from "react";
import Carousel from "../../components/carousel/Carousel.jsx";
import Loading from "../../components/loading/Loading.jsx";
import { useAuth } from "../../hooks/useAuth";
import { homeService } from "../../services/home.service.js";

export default function Home() {
  const { user } = useAuth();
  const [movies, setMovies] = useState([]);
  const [loadingContent, setLoadingContent] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchHomeContent = async () => {
      if (!user) {
        setLoadingContent(false);
        return;
      }

      try {
        const data = await homeService();
        setMovies(data.data || []);
        setError(null);
      } catch (err) {
        console.error("Error fetching content:", err);
        setError("Could not load media library.");
      } finally {
        setLoadingContent(false);
      }
    };

    fetchHomeContent();
  }, [user]);

  if (!user) {
    return <Loading message="Authenticating" />;
  }

  if (loadingContent) {
    return <Loading message="Fetching content from the server" />;
  }

  if (error) {
    return <h1 className="text-red-500 p-6">Error: {error}</h1>;
  }

  if (movies.length === 0) {
    return (
      <div className="p-4 bg-[#1f2029] min-h-screen text-white text-center">
        <h1 className="text-3xl mb-4">Welcome, {user.email}!</h1>
        <div className="bg-gray-800 p-10 rounded-lg max-w-lg mx-auto mt-20">
          <p className="text-xl text-yellow-500 mb-4">Library Empty</p>
          <p>
            We couldn't find any movies or series to display right now.
            Please check back later or contact support.
          </p>
        </div>
      </div>
    );
  }

  return (
    <div className="bg-[#1f2029] min-h-screen text-white">
      <div className="p-4 sm:p-6">
        <h1 className="text-xl sm:text-2xl md:text-3xl mb-6">Welcome, {user.email}!</h1>

        <h2 className="text-xl font-semibold mb-4 text-[#ffeba7]">Recently Added</h2>
        <Carousel data={movies} />
      </div>
    </div>
  );
}
