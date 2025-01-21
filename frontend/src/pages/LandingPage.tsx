import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faBook,
  faSearch,
  faStar,
  faChartLine,
} from '@fortawesome/free-solid-svg-icons';
import Header from '../components/layout/Header';
import Footer from '../components/layout/Footer';

// min-h-screen flex flex-col
const LandingPage = () => {
  return (
    <div className="min-h-screen flex flex-col">
      {/* Header */}
      <Header />

      {/* Hero Section */}
      <section className="flex-grow bg-gradient-to-b from-white to-gray-100">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-20">
          <div className="text-center">
            <h1 className="text-4xl md:text-6xl font-bold text-gray-900 mb-6">
              Organize Your Book Collection Effortlessly!
            </h1>
            <p className="text-xl text-gray-600 mb-8">
              Manage your personal library and discover new books with ease.
            </p>
            <div className="flex justify-center space-x-4">
              <button className="bg-primary text-white px-8 py-3 rounded-md hover:bg-secondary">
                Get Started
              </button>
              <button className="border border-primary text-primary px-8 py-3 rounded-md hover:bg-primary hover:text-white">
                Explore Features
              </button>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section id="features" className="py-20 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <h2 className="text-3xl font-bold text-center mb-12">Features</h2>
          <div className="grid md:grid-cols-4 gap-8">
            <div className="text-center p-6">
              <FontAwesomeIcon
                icon={faBook}
                className="h-12 w-12 text-primary mb-4"
              />
              <h3 className="text-xl font-semibold mb-2">
                Organize your collection
              </h3>
              <p className="text-gray-600">
                Keep track of all your books in one place
              </p>
            </div>
            <div className="text-center p-6">
              <FontAwesomeIcon
                icon={faSearch}
                className="h-12 w-12 text-primary mb-4"
              />
              <h3 className="text-xl font-semibold mb-2">Discover new books</h3>
              <p className="text-gray-600">Find your next favorite read</p>
            </div>
            <div className="text-center p-6">
              <FontAwesomeIcon
                icon={faStar}
                className="h-12 w-12 text-primary mb-4"
              />
              <h3 className="text-xl font-semibold mb-2">Rate and review</h3>
              <p className="text-gray-600">Share your thoughts on books</p>
            </div>
            <div className="text-center p-6">
              <FontAwesomeIcon
                icon={faChartLine}
                className="h-12 w-12 text-primary mb-4"
              />
              <h3 className="text-xl font-semibold mb-2">Track progress</h3>
              <p className="text-gray-600">Monitor your reading journey</p>
            </div>
          </div>
        </div>
      </section>

      {/* Footer */}
      <Footer />
    </div>
  );
};

export default LandingPage;
