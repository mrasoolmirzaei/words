import { useState } from "react";
import useSearchWord from "../../services/hook/useSearchWord";

const SearchBar = () => {
  const [searchQuery, setSearchQuery] = useState("");
  const { handleSearch } = useSearchWord();
  const handleSearchBar = (e) => {
    setSearchQuery(e.target.value);
  };
  const submitSearch = () => {
    handleSearch(searchQuery)
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
      <button className="input-group-text" id="search" onClick={submitSearch}>
        Search
      </button>
    </div>
  );
};

export default SearchBar;
