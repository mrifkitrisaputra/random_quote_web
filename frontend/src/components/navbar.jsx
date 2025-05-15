import React from "react";
import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <nav className="bg-blue-600 text-white shadow-md">
      <div className="container mx-auto px-4 py-3 flex justify-between items-center">
        {/* Logo / Nama Aplikasi */}
        <Link
          to="/"
          className="text-xl font-bold hover:text-blue-200 transition-colors"
        >
          QuickPoll
        </Link>

        {/* Menu Navigasi */}
        <ul className="flex space-x-6 items-center">
          <li>
            <Link to="/" className="hover:text-blue-200 transition-colors">
              Beranda
            </Link>
          </li>
          <li>
            <Link
              to="/create"
              className="hover:text-blue-200 transition-colors"
            >
              Buat Polling
            </Link>
          </li>
          <li>
            <Link
              to="/dashboard"
              className="hover:text-blue-200 transition-colors"
            >
              Dashboard
            </Link>
          </li>
        </ul>
      </div>
    </nav>
  );
};

export default Navbar;
