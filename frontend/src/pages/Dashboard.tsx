import { useState, useEffect } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBook, faPlus } from '@fortawesome/free-solid-svg-icons';
import {
  BookOpen,
  User,
  LogOut,
  Menu,
  X,
  Search,
  List,
  Grid,
} from 'lucide-react';
import BookCard from '../components/BookCard';
import AddBookModal from '../components/AddBookModal';
import Pagination from '../components/Pagination';
import FilterBar from '../components/FilterBar';
import { useNavigate } from 'react-router-dom';
import SearchBookCard from '../components/SearchBookCard';

const ITEMS_PER_PAGE = 12;

const Dashboard = () => {
  const navigate = useNavigate();
  const [searchScope, setSearchScope] = useState('my-collections');
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);
  const [viewMode, setViewMode] = useState('grid');
  const [pendingIsbnQuery, setPendingIsbnQuery] = useState('');
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const [searchQuery, setSearchQuery] = useState('');
  const [currentPage, setCurrentPage] = useState(1);
  const [showProfileMenu, setShowProfileMenu] = useState(false);
  const [books, setBooks] = useState<Book[]>([]);
  const [searchedBooks, setSearchedBooks] = useState<
    {
      title: string;
      author: string;
      isbn: string;
    }[]
  >([]);
  const [filters, setFilters] = useState({
    status: '',
    rating: '',
    sortBy: 'created_at',
    sortOrder: 'desc' as 'asc' | 'desc',
  });
  const [metadata, setMetadata] = useState({
    results: 0,
    current_page: 1,
    page_limit: 12,
  });

  useEffect(() => {
    const fetchBooks = async () => {
      try {
        let url = '/api/v1/books';

        if (searchScope === 'openlibrary') {
          url = `/api/v1/books/search?page=${currentPage}&isbn=${searchQuery}`;
        }

        const response = await fetch(url, {
          method: 'GET',
          headers: {
            'Content-Type': 'application/json',
          },
        });

        if (!response.ok) {
          throw new Error(`Error: ${response.status}`);
        }

        const result = await response.json();

        if (searchScope === 'my-collections') {
          const { data, metadata } = result;
          setBooks(data);
          setMetadata(metadata);
          setCurrentPage(metadata.current_page);
        } else {
          setSearchedBooks(
            result.data.docs.map(
              (book: {
                title: string;
                author_name: string[];
                isbn: string;
              }) => ({
                title: book.title,
                author: book.author_name?.join(', '),
                isbn:
                  book.isbn && book.isbn.length > 0 ? book.isbn[0] : undefined,
              })
            )
          );
        }
      } catch (error) {
        console.error('Error fetching books:', error);
      }
    };

    fetchBooks();
  }, [searchScope, searchQuery, currentPage]);

  // Handle page change
  const handlePageChange = async (newPage: number) => {
    try {
      const response = await fetch(`/api/v1/books?page=${newPage}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        throw new Error(`Error: ${response.status}`);
      }

      const result = await response.json();
      const { data, metadata } = result;

      // If the page is empty and not the first page, fetch the previous page
      if (data.length === 0 && newPage > 1) {
        await handlePageChange(newPage - 1);
      } else {
        setBooks(data);
        setMetadata(metadata);
      }
    } catch (error) {
      console.error('Error fetching books for page:', error);
    }
  };

  //   Initialize the books state with an empty array
  type Book = {
    id: number;
    title: string;
    author: string;
    isbn?: string;
    status: string;
    rating: number;
    notes?: string;
    created_at: string;
  };

  const handleAddBook = async (bookData: any) => {
    try {
      const response = await fetch('/api/v1/books', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(bookData),
      });

      const result = await response.json();

      if (!response.ok) {
        throw new Error(result['message']);
      }

      //   Update the books state with the new book
      setBooks((prevBooks) => [...prevBooks, result['data']]);
    } catch (error) {
      alert(`${error}`);
    }
  };

  const handleEditBook = async (bookData: any, id: number) => {
    try {
      const response = await fetch(`/api/v1/books/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(bookData),
      });

      const result = await response.json();

      if (!response.ok) {
        throw new Error(result.message || 'Failed to update the book.');
      }

      // Update the books state with the edited book
      setBooks((prevBooks) =>
        prevBooks.map((book) =>
          book.id === id ? { ...book, ...result.data } : book
        )
      );
    } catch (error) {
      alert(error);
    }
  };

  const handleDeleteBook = async (id: number) => {
    if (window.confirm('Are you sure you want to delete this book?')) {
      const response = await fetch(`/api/v1/books/${id}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (response.status !== 204) {
        throw new Error('Failed to delete the book');
      }

      if (response.status !== 204) {
        throw new Error('Failed to delete the book');
      }

      const updatedBooks = books.filter((book) => book.id !== id);
      const isLastItemonPage = updatedBooks.length === 0 && currentPage > 1;

      const newPage = isLastItemonPage ? currentPage - 1 : currentPage;

      setMetadata((prev) => ({ ...prev, current_page: newPage }));

      await handlePageChange(newPage);
    }
  };

  const handleViewNotes = (id: number) => {
    const book = books.find((b) => b.id === id);
    if (book?.notes) {
      alert(book.notes); // Replace with a proper modal in production
    }
  };

  const handleLogout = () => {
    // TODO: Implement proper logout logic
    navigate('/login');
  };

  const filteredBooks = books
    .filter((book) => {
      if (searchQuery) {
        const query = searchQuery.toLowerCase();
        if (
          !book.title.toLowerCase().includes(query) &&
          !book.author.toLowerCase().includes(query) &&
          !book.isbn?.includes(query)
        ) {
          return false;
        }
      }

      if (filters.status && book.status !== filters.status) {
        return false;
      }

      if (filters.rating && book.rating < parseInt(filters.rating)) {
        return false;
      }

      return true;
    })
    .sort((a, b) => {
      const order = filters.sortOrder === 'asc' ? 1 : -1;
      switch (filters.sortBy) {
        case 'title':
          return order * a.title.localeCompare(b.title);
        case 'author':
          return order * a.author.localeCompare(b.author);
        case 'rating':
          return order * (b.rating - a.rating);
        case 'created_at':
          return (
            order * new Date(b.created_at).getTime() -
            new Date(a.created_at).getTime()
          );
        default:
          return 0;
      }
    });

  const totalPages = Math.ceil(metadata.results / metadata.page_limit);
  const paginatedBooks = filteredBooks.slice(
    (currentPage - 1) * ITEMS_PER_PAGE,
    currentPage * ITEMS_PER_PAGE
  );

  const handleSearch = () => {
    if (searchScope === 'openlibrary') {
      setSearchQuery(pendingIsbnQuery);
      setCurrentPage(1);
    }
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter' && searchScope === 'openlibrary') {
      handleSearch();
    }
  };

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm sticky top-0 z-10">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex items-center justify-between h-16">
            {/* Logo */}
            <div className="flex items-center">
              <BookOpen className="h-6 w-6 text-primary" />
              <span className="ml-2 text-xl font-bold hidden sm:block">
                BookShelf
              </span>
            </div>

            {/* Desktop Navigation */}
            <div className="hidden md:flex items-center space-x-4">
              {/* Profile Menu */}
              <div className="relative">
                <button
                  className="flex items-center space-x-2 text-gray-600 hover:text-gray-900 transition-colors"
                  onClick={() => setShowProfileMenu(!showProfileMenu)}
                >
                  <User className="h-5 w-5" />
                  <span>Profile</span>
                </button>

                {showProfileMenu && (
                  <div className="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 ring-1 ring-black ring-opacity-5">
                    <button
                      onClick={handleLogout}
                      className="flex items-center w-full px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                    >
                      <LogOut className="h-4 w-4 mr-2" />
                      Sign out
                    </button>
                  </div>
                )}
              </div>
            </div>

            {/* Mobile Menu Button */}
            <div className="md:hidden">
              <button
                onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
                className="p-2 text-gray-600 hover:text-gray-900 transition-colors"
              >
                {isMobileMenuOpen ? (
                  <X className="h-6 w-6" />
                ) : (
                  <Menu className="h-6 w-6" />
                )}
              </button>
            </div>
          </div>

          {/* Mobile Menu */}
          {isMobileMenuOpen && (
            <div className="md:hidden py-4 border-t">
              <div className="space-y-4">
                <button
                  onClick={handleLogout}
                  className="flex items-center w-full px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                >
                  <LogOut className="h-4 w-4 mr-2" />
                  Sign out
                </button>
              </div>
            </div>
          )}
        </div>
      </header>

      {/* Search Bar */}
      <div className="bg-white border-b">
        <div className="max-w-4xl mx-auto px-3 sm:px-6 lg:px-8 py-4">
          <div className="relative">
            <select
              className="absolute left-0 top-0 h-full border-r border-gray-300 bg-gray-50 rounded-l-md text-sm px-10"
              value={searchScope}
              onChange={(e) => setSearchScope(e.target.value)}
            >
              <option value="my-collections">My Collections</option>
              <option value="openlibrary">OpenLibrary</option>
            </select>
            <input
              type="text"
              placeholder={
                searchScope === 'my-collections'
                  ? 'Search by title, author or ISBN'
                  : 'Search by ISBN'
              }
              value={
                searchScope === 'my-collections'
                  ? searchQuery
                  : pendingIsbnQuery
              }
              onChange={(e) => {
                if (searchScope === 'my-collections') {
                  setSearchQuery(e.target.value);
                  setCurrentPage(1);
                } else {
                  setPendingIsbnQuery(e.target.value);
                }
              }}
              onKeyDown={handleKeyDown}
              className="w-full pl-52 pr-8 py-3 border rounded-md text-sm"
            />
            <Search
              onClick={handleSearch}
              className="absolute right-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400"
            />
          </div>
        </div>
      </div>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* View Controls */}
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-2xl font-bold">
            {searchScope === 'my-collections'
              ? 'My Collections'
              : 'Search Results'}
          </h1>
          <button
            onClick={() => setViewMode(viewMode === 'grid' ? 'list' : 'grid')}
            className="p-2 border rounded-md hover:bg-gray-50"
            title={
              viewMode === 'grid'
                ? 'Switch to list view'
                : 'Switch to grid view'
            }
          >
            {viewMode === 'grid' ? (
              <List className="h-5 w-5" />
            ) : (
              <Grid className="h-5 w-5" />
            )}
          </button>
        </div>

        {/* Filters */}
        <FilterBar filters={filters} onFilterChange={setFilters} />

        {/* Books Grid/List */}
        <div
          className={`grid ${
            viewMode === 'grid'
              ? 'grid-cols-1 sm:grid-cols-2 lg:grid-cols-3'
              : 'grid-cols-1'
          } gap-6`}
        >
          {searchScope === 'openlibrary'
            ? searchedBooks.map((book, index) => (
                <SearchBookCard key={index} book={book} />
              ))
            : paginatedBooks.map((book) => (
                <BookCard
                  key={book.id}
                  book={book}
                  onEdit={(updatedBookData) =>
                    handleEditBook(updatedBookData, book.id)
                  }
                  onDelete={handleDeleteBook}
                  onViewNotes={handleViewNotes}
                />
              ))}
        </div>

        {/* Empty State */}
        {((searchScope === 'my-collections' && filteredBooks.length === 0) ||
          (searchScope === 'openlibrary' && searchedBooks.length === 0)) && (
          <div className="text-center py-12">
            <FontAwesomeIcon
              icon={faBook}
              className="h-12 w-12 text-gray-400 mb-4"
            />
            <h3 className="text-lg font-medium text-gray-900 mb-2">
              No books found
            </h3>
            <p className="text-gray-500">
              {searchQuery
                ? 'Try adjusting your search or filters'
                : 'Start by adding some books to your collection'}
            </p>
          </div>
        )}

        {/* Pagination */}
        {filteredBooks.length > 0 && (
          <Pagination
            currentPage={metadata.current_page}
            totalPages={totalPages}
            onPageChange={handlePageChange}
          />
        )}

        {/* Add Book Button */}
        <button
          onClick={() => setIsAddModalOpen(true)}
          className="fixed right-8 bottom-8 bg-primary text-white p-4 rounded-full shadow-lg hover:bg-secondary transition-colors duration-200"
        >
          <FontAwesomeIcon icon={faPlus} className="h-6 w-6" />
        </button>

        {/* Add Book Modal */}
        <AddBookModal
          isOpen={isAddModalOpen}
          onClose={() => setIsAddModalOpen(false)}
          onAdd={handleAddBook}
        />
      </main>
    </div>
  );
};

export default Dashboard;
