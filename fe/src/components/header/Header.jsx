import { useState } from "react";
import { Link } from "react-router-dom";
import { useAuth } from "../../hooks/useAuth.jsx";
import AccountDropdown from "../accountdropdown/AccountDropdown";
import SearchBar from "../searchbar/SearchBar";

export default function Header() {
  const { user, logout } = useAuth();
  const isAdmin = user?.role === "admin";
  const isSuperAdmin = user?.role === "superadmin";
  const [menuOpen, setMenuOpen] = useState(false);

  return (
    <header className="sticky top-0 z-40 w-full bg-[#2a2b38] text-white shadow-md p-2">
      <div className="flex items-center justify-between gap-2">
        <div className="flex items-center gap-2 sm:gap-5">
          <h1 className="text-lg sm:text-xl font-bold whitespace-nowrap">
            <Link to={isSuperAdmin ? "/superadmin" : "/home"}>Videoteca</Link>
          </h1>

          <button
            className="sm:hidden p-1 text-[#c4c3ca] hover:text-white"
            onClick={() => setMenuOpen(!menuOpen)}
            aria-label="Toggle menu"
          >
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              {menuOpen ? (
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
              ) : (
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
              )}
            </svg>
          </button>

          <nav className="hidden sm:flex gap-2 lg:gap-4 flex-wrap">
            {isSuperAdmin ? (
              <Link
                to="/superadmin"
                className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-0.5 px-3 rounded-md h-[30px] leading-[30px] text-sm transition-colors"
              >
                Create Admin
              </Link>
            ) : (
              <>
                <Link
                  to="/home"
                  className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-0.5 px-3 rounded-md h-[30px] leading-[30px] text-sm transition-colors"
                >
                  Home
                </Link>
                {isAdmin && (
                  <>
                    <Link to="/admin/movies" className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-0.5 px-3 rounded-md h-[30px] leading-[30px] text-sm transition-colors">Movies & Series</Link>
                    <Link to="/admin/tags" className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-0.5 px-3 rounded-md h-[30px] leading-[30px] text-sm transition-colors">Tags</Link>
                    <Link to="/admin/actors" className="bg-[#ffeba7] text-[#102770] hover:bg-[#102770] hover:text-[#ffeba7] font-semibold py-0.5 px-3 rounded-md h-[30px] leading-[30px] text-sm transition-colors">Actors</Link>
                  </>
                )}
              </>
            )}
          </nav>
        </div>

        <div className="hidden sm:block flex-1 max-w-xs ml-auto">
          {!isSuperAdmin && <SearchBar />}
        </div>

        <AccountDropdown user={user} onLogout={logout} />
      </div>

      {menuOpen && (
        <div className="sm:hidden mt-2 pt-2 border-t border-gray-700 space-y-2">
          {!isSuperAdmin && <SearchBar />}
          <nav className="flex flex-col gap-1">
            {isSuperAdmin ? (
              <Link to="/superadmin" className="block bg-[#ffeba7] text-[#102770] font-semibold py-1 px-3 rounded-md text-sm" onClick={() => setMenuOpen(false)}>Create Admin</Link>
            ) : (
              <>
                <Link to="/home" className="block bg-[#ffeba7] text-[#102770] font-semibold py-1 px-3 rounded-md text-sm" onClick={() => setMenuOpen(false)}>Home</Link>
                {isAdmin && (
                  <>
                    <Link to="/admin/movies" className="block bg-[#ffeba7] text-[#102770] font-semibold py-1 px-3 rounded-md text-sm" onClick={() => setMenuOpen(false)}>Movies & Series</Link>
                    <Link to="/admin/tags" className="block bg-[#ffeba7] text-[#102770] font-semibold py-1 px-3 rounded-md text-sm" onClick={() => setMenuOpen(false)}>Tags</Link>
                    <Link to="/admin/actors" className="block bg-[#ffeba7] text-[#102770] font-semibold py-1 px-3 rounded-md text-sm" onClick={() => setMenuOpen(false)}>Actors</Link>
                  </>
                )}
              </>
            )}
          </nav>
        </div>
      )}
    </header>
  );
}
