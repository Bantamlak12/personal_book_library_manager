import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBook } from '@fortawesome/free-solid-svg-icons';

interface SearchBookCardProps {
  book: {
    title: string;
    author: string;
    isbn: string;
  };
  onAddToCollection: (book: {
    title: string;
    author: string;
    isbn: string;
  }) => void;
}

const SearchBookCard = ({ book, onAddToCollection }: SearchBookCardProps) => {
  return (
    <div className="bg-white rounded-lg shadow-sm p-4 hover:shadow-md transition-shadow duration-200">
      <div className="flex justify-between items-start mb-4">
        <div className="flex-1">
          <h3 className="font-semibold text-gray-900 text-lg mb-1">
            {book.title}
          </h3>
          {book.author && (
            <p className="text-gray-600 text-sm mb-2">{book.author}</p>
          )}
          {book.isbn && book.isbn.length > 0 && (
            <p className="text-gray-500 text-xs mb-2">ISBN: {book.isbn}</p>
          )}
        </div>
        <FontAwesomeIcon icon={faBook} className="text-gray-400 h-6 w-6" />
      </div>

      <button
        onClick={() => onAddToCollection(book)}
        className="w-full py-2 px-4 bg-blue-500 text-white rounded-lg hover:bg-secondary transition duration-300"
      >
        Add to My Collection
      </button>
    </div>
  );
};

export default SearchBookCard;
