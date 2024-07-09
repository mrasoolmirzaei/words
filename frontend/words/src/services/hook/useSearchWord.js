import { useState } from "react";
import { searchSynonym } from "../api/synonym";

const useSearchWord = () => {
  const [searchResults, setSearchResults] = useState(null);
  const [loading, setLoading] = useState(false);

  const searchWordHandler = async (query) => {
    if (!query) {
      setSearchResults(null);
      return;
    }
    setLoading(true);
    const result = await searchSynonym(query);
    if (result.ok) {
      const data = await result.json();
      setSearchResults(data);
    }
    setLoading(false);
  };

  return { searchResults, loading, handleSearch: searchWordHandler };
};

export default useSearchWord;
