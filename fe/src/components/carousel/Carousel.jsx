import { Link } from "react-router-dom";
import Slider from "react-slick";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

export default function Carousel({ data = [] }) {
  const settings = {
    className: "center",
    centerMode: true,
    infinite: data.length > 3,
    centerPadding: "60px",
    slidesToShow: Math.min(3, data.length),
    slidesToScroll: 1,
    speed: 500,
    responsive: [
      {
        breakpoint: 768,
        settings: {
          slidesToShow: 1,
          centerPadding: "40px",
        },
      },
    ],
  };

  if (data.length === 0) return null;

  return (
    <div className="slider-container">
      <Slider {...settings}>
        {data.map((item) => (
          <div key={item.id} className="px-2">
            <Link
              to={`/movies/${item.id}`}
              className="block relative group rounded-lg overflow-hidden bg-[#2a2b38]"
            >
              {item.cover ? (
                <img
                  src={item.cover}
                  alt={item.title}
                  className="w-full h-72 object-cover"
                />
              ) : (
                <div className="w-full h-72 bg-gray-700 flex items-center justify-center text-gray-400">
                  No Cover
                </div>
              )}
              <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/90 to-transparent p-4">
                <h3 className="text-white font-semibold text-lg truncate">
                  {item.title}
                </h3>
                <div className="flex items-center gap-2 text-sm text-gray-300">
                  <span>{item.releaseYear}</span>
                  {item.msType && (
                    <span
                      className={`px-2 py-0.5 rounded text-xs font-semibold ${
                        item.msType === "movie"
                          ? "bg-blue-600/30 text-blue-300"
                          : "bg-purple-600/30 text-purple-300"
                      }`}
                    >
                      {item.msType}
                    </span>
                  )}
                </div>
              </div>
            </Link>
          </div>
        ))}
      </Slider>
    </div>
  );
}
