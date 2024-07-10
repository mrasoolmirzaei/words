import PropTypes from 'prop-types';

const SearchResults = ({ results }) => {
  if (!results) {
    return null;
  }
  return (
    <div className="card">
      <div className="card-header">Synonyms:</div>
      <div className="d-flex max-w-28rem flex-wrap">
        {results.synonyms?.map((item) => (
          <span key={item.id} className="badge bg-secondary fs-6 m-2">{item.title}</span>
        ))}
      </div>
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
