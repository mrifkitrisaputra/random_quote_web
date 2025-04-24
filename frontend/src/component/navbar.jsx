import React from 'react';
import { Link } from 'react-router-dom';

const Navbar = () => {
  return (
    <nav className="bg-blue-600 p-4 text-white flex justify-between items-center">
      <Link to="/" className="text-lg font-bold">QuickPoll</Link>
      <div className="space-x-4">
        <Link to="/" className="hover:text-gray-300">Home</Link>
        <Link to="/create" className="hover:text-gray-300">Create Poll</Link>
      </div>
    </nav>
  );
};

export default Navbar;