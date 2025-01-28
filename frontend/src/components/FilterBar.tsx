import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFilter } from '@fortawesome/free-solid-svg-icons';

interface FilterBarProps {
  onFilterChange: (filters: any) => void;
  filters: {
    status: string;
    rating: string;
    sortBy: string;
    sortOrder: 'asc' | 'desc';
  };
}

const FilterBar = ({ onFilterChange, filters }: FilterBarProps) => {
  return (
    <div className="bg-white p-4 rounded-lg shadow-sm mb-6">
      <div className="flex flex-wrap items-center gap-4">
        <div className="flex items-center">
          <FontAwesomeIcon icon={faFilter} className="text-gray-400 mr-2" />
          <span className="text-sm font-medium">Filters:</span>
        </div>

        <select
          className="border rounded-md px-3 py-1.5 text-sm"
          value={filters.status}
          onChange={(e) =>
            onFilterChange({ ...filters, status: e.target.value })
          }
        >
          <option value="">All Status</option>
          <option value="unread">Unread</option>
          <option value="reading">Reading</option>
          <option value="completed">Completed</option>
        </select>

        <select
          className="border rounded-md px-3 py-1.5 text-sm"
          value={filters.rating}
          onChange={(e) =>
            onFilterChange({ ...filters, rating: e.target.value })
          }
        >
          <option value="">All Ratings</option>
          <option value="5">5 Stars</option>
          <option value="4">4+ Stars</option>
          <option value="3">3+ Stars</option>
        </select>

        <div className="flex-1" />

        <div className="flex items-center space-x-2">
          <select
            className="border rounded-md px-3 py-1.5 text-sm"
            value={filters.sortBy}
            onChange={(e) =>
              onFilterChange({ ...filters, sortBy: e.target.value })
            }
          >
            <option value="created_at">Date Added</option>
            <option value="title">Title</option>
            <option value="author">Author</option>
            <option value="rating">Rating</option>
          </select>

          <button
            className="border rounded-md px-3 py-1.5 text-sm hover:bg-gray-50"
            onClick={() =>
              onFilterChange({
                ...filters,
                sortOrder: filters.sortOrder === 'asc' ? 'desc' : 'asc',
              })
            }
          >
            {filters.sortOrder === 'asc' ? '↑' : '↓'}
          </button>
        </div>
      </div>
    </div>
  );
};

export default FilterBar;
