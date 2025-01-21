// Header.tsx
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBook } from '@fortawesome/free-solid-svg-icons';
import { Link } from 'react-router-dom';

const Header = () => {
  return (
    <header className="bg-white shadow-sm">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
        <div className="flex justify-between items-center">
          <div className="flex items-center">
            <FontAwesomeIcon icon={faBook} className="text-primary h-8 w-8" />
            <span className="ml-2 text-xl font-bold">BookShelf</span>
          </div>
          <nav className="hidden md:flex space-x-8">
            <a href="#features" className="text-gray-600 hover:text-gray-900">
              Features
            </a>
            <a href="#about" className="text-gray-600 hover:text-gray-900">
              About
            </a>
            <a href="#contact" className="text-gray-600 hover:text-gray-900">
              Contact
            </a>
          </nav>
          <div className="flex space-x-4">
            <Link to="/login" className="text-primary hover:text-secondary">
              Login
            </Link>
            <Link
              to="/signup"
              className="bg-primary text-white px-4 py-2 rounded-md hover:bg-secondary"
            >
              Sign Up
            </Link>
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;
