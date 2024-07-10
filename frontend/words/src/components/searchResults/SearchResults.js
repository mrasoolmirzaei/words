import PropTypes from 'prop-types';

const SearchResults = ({ results }) => {
  if (!results) {
    return null;
  }
  return (
    <div className="card">
      <div className="card-header">Search Results:</div>
      <ul className="list-group list-group-flush">
        {results.synonyms?.map((item) => (
          <li key={item.id} className="list-group-item">
            {item.title}
          </li>
        ))}
      </ul>
    </div>
  );
};

SearchResults.propTypes = {
  results: PropTypes.shape({
    synonyms: PropTypes.arrayOf(
      PropTypes.shape({
        id: PropTypes.oneOfType([PropTypes.string, PropTypes.number]).isRequired,
        title: PropTypes.string.isRequired,
      })
    ),
  }),
};

export default SearchResults;
