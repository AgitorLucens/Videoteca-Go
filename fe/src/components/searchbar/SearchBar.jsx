import { useState, useEffect, useRef } from "react";
import { searchMovies } from "../../services/search.service";

export default function SearchBar() {
  const [query, setQuery] = useState("");
  const [results, setResults] = useState([]);
  const [loading, setLoading] = useState(false);
  const [showDropdown, setShowDropdown] = useState(false);
  const wrapperRef = useRef(null);

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (wrapperRef.current && !wrapperRef.current.contains(event.target)) {
        setShowDropdown(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  useEffect(() => {
    const search = async () => {
      if (query.trim().length < 2) {
        setResults([]);
        return;
      }

      setLoading(true);
      try {
        const response = await searchMovies(query);
        setResults(response.data || []);
        setShowDropdown(true);
      } catch (error) {
        console.error("Search error:", error);
        setResults([]);
      } finally {
        setLoading(false);
      }
    };

    const timeoutId = setTimeout(search, 300);
    return () => clearTimeout(timeoutId);
  }, [query]);

  const renderTypeBadge = (item) => {
    if (item.type === "actor") {
      return <span className="text-xs px-2 py-0.5 rounded ml-2 bg-green-600/30 text-green-300">actor</span>;
    }
    return (
      <span className={`text-xs px-2 py-0.5 rounded ml-2 ${
        item.type === "movie" ? "bg-blue-600/30 text-blue-300" : "bg-purple-600/30 text-purple-300"
      }`}>
        {item.type}
      </span>
    );
  };

  return (
    <div ref={wrapperRef} className="relative">
      <input
        type="text"
        placeholder="Search movies, series, actors..."
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        className="w-full px-4 py-2 rounded-md bg-[#1f2029] text-white placeholder-gray-400 outline-none border border-[#2a2b38] focus:border-[#ffeba7]"
      />
      {loading && (
        <div className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 text-sm">
          loading...
        </div>
      )}
      {showDropdown && results.length > 0 && (
        <ul className="absolute z-50 w-full mt-1 bg-[#2a2b38] rounded-md shadow-lg max-h-64 overflow-y-auto">
          {results.map((item, idx) => (
            <li
              key={`${item.type}-${item.id || idx}`}
              className="px-4 py-2 hover:bg-[#1f2029] cursor-pointer flex justify-between items-center text-white"
            >
              <span className="truncate">{item.title}</span>
              {renderTypeBadge(item)}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
