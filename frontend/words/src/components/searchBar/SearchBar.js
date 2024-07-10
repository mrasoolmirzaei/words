import { useState } from "react";
import PropTypes from "prop-types";
import { lettersOnlyPattern } from "../../constants/regex";

const SearchBar = ({ onSearch }) => {
  const [searchQuery, setSearchQuery] = useState("");
  const [validationError, setValidationError] = useState("");

  const handleSearchBar = (e) => {
    const { value } = e.target;
    if (!value) onSearch("");
    setSearchQuery(value);
    if (lettersOnlyPattern.test(value)) {
      setValidationError("");
    } else {
      setValidationError("Please enter only letters.");
    }
  };
  const submitSearch = () => {
    onSearch("");
    onSearch(searchQuery);
  };
  return (
    <div className="form-group w-100">
      <div className="input-group">
        <input
          type="search"
          className="form-control"
          placeholder="Search for a word"
          aria-label="Search"
          value={searchQuery}
          onChange={handleSearchBar}
        />
        <button className="btn btn-primary" disabled={validationError} onClick={submitSearch}>
          Search
        </button>
      </div>
      {validationError && <p className="text-danger">{validationError}</p>}
    </div>
  );
};

SearchBar.propTypes = {
  onSearch: PropTypes.func.isRequired,
};

export default SearchBar;
