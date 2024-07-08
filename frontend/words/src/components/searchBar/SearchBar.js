import { useState } from "react";

const SearchBar = ({onSearch}) => {

  const [searchQuery, setSearchQuery] = useState("");

  const handleSearchBar = (e) => {
    setSearchQuery(e.target.value);
  };
  const submitSearch = () => {
    onSearch(searchQuery)
  }
  return (
    <div className="input-group mb-3">
      <input
        type="search"
        className="form-control"
        placeholder="Search for a word"
        aria-label="Search"
        described="search"
        value={searchQuery}
        onChange={handleSearchBar}
      />
      <button className="input-group-text" onClick={submitSearch}>
        Search
      </button>
    </div>
  );
};

export default SearchBar;
