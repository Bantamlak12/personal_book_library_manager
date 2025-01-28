import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faStar,
  faEdit,
  faTrash,
  faBook,
  faBookOpen,
  faCheck,
} from '@fortawesome/free-solid-svg-icons';
import { useState } from 'react';
import EditBookModal from './EditBookModal';

interface BookCardProps {
  book: {
    id: number;
    title: string;
    author: string;
    status: string;
    rating: number;
    notes?: string;
    isbn?: string;
    created_at?: string;
  };
  onEdit: (id: number) => void;
  onDelete: (id: number) => void;
  onViewNotes: (id: number) => void;
}

const BookCard = ({ book, onEdit, onDelete, onViewNotes }: BookCardProps) => {
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);

  const getStatusIcon = (status: string) => {
    switch (status) {
      case 'reading':
        return faBookOpen;
      case 'completed':
        return faCheck;
      default:
        return faBook;
    }
  };

  const getStatusColor = (status: string) => {
    switch (status) {
      case 'reading':
        return 'text-blue-500';
      case 'completed':
        return 'text-green-500';
      default:
        return 'text-gray-500';
    }
  };

  return (
    <div className="bg-white rounded-lg shadow-sm p-4 hover:shadow-md transition-shadow duration-200">
      <div className="flex justify-between items-start">
        <div className="flex-1">
          <h3 className="font-semibold text-gray-900 text-lg mb-1">
            {book.title}
          </h3>
          <p className="text-gray-600 text-sm mb-2">{book.author}</p>
          {book.isbn && (
            <p className="text-gray-500 text-xs mb-2">ISBN: {book.isbn}</p>
          )}
          <div className="flex items-center space-x-2 mb-2">
            <FontAwesomeIcon
              icon={getStatusIcon(book.status)}
              className={`${getStatusColor(book.status)}`}
            />
            <span className="text-sm capitalize">{book.status}</span>
          </div>
          <div className="flex mb-2">
            {[...Array(5)].map((_, index) => (
              <FontAwesomeIcon
                key={index}
                icon={faStar}
                className={`h-4 w-4 ${
                  index < book.rating ? 'text-yellow-400' : 'text-gray-300'
                }`}
              />
            ))}
          </div>
          <p className="text-gray-500 text-xs">
            Added: {book.created_at?.split('T')[0]}
          </p>
        </div>
        <div className="flex flex-col space-y-2">
          <button
            onClick={() => setIsEditModalOpen(true)}
            className="p-2 text-blue-500 hover:bg-blue-50 rounded-full transition-colors duration-200"
            title="Edit book"
          >
            <FontAwesomeIcon icon={faEdit} />
          </button>
          <EditBookModal
            isOpen={isEditModalOpen}
            onClose={() => setIsEditModalOpen(false)}
            onSave={onEdit}
            book={book}
          />

          <button
            onClick={() => onViewNotes(book.id)}
            className="p-2 text-gray-500 hover:bg-gray-50 rounded-full transition-colors duration-200"
            title="View notes"
          >
            <FontAwesomeIcon icon={faBook} />
          </button>
          <button
            onClick={() => onDelete(book.id)}
            className="p-2 text-red-500 hover:bg-red-50 rounded-full transition-colors duration-200"
            title="Delete book"
          >
            <FontAwesomeIcon icon={faTrash} />
          </button>
        </div>
      </div>
    </div>
  );
};

export default BookCard;
