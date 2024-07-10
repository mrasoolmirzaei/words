import { useState } from "react";
import PropTypes from "prop-types";

const SearchBar = ({ onSearch }) => {
  const [searchQuery, setSearchQuery] = useState("");

  const handleSearchBar = (e) => {
    const { value } = e.target;
    setSearchQuery(value);
    if (!value) onSearch("");
  };
  const submitSearch = () => {
    onSearch("");
    onSearch(searchQuery);
  };
  return (
    <div className="input-group mb-3">
      <input
        type="search"
        className="form-control"
        placeholder="Search for a word"
        aria-label="Search"
        value={searchQuery}
        onChange={handleSearchBar}
      />
      <button className="input-group-text" onClick={submitSearch}>
        Search
      </button>
    </div>
  );
};

SearchBar.propTypes = {
  onSearch: PropTypes.func.isRequired,
};

export default SearchBar;
