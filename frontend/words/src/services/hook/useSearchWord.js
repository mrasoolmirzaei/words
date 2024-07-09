import { useState } from "react";
import { searchSynonym } from "../api/synonym";
import { toast } from "react-toastify";

const useSearchWord = () => {
  const [searchResults, setSearchResults] = useState(null);
  const [loading, setLoading] = useState(false);

  const searchWordHandler = async (query) => {
    if (!query) {
      setSearchResults(null);
      return;
    }
    setLoading(true);
    try {
      const result = await searchSynonym(query);
      if (result.ok) {
        const data = await result.json();
        setSearchResults(data);
      }else {
        toast.error(result.statusText);
      }
    } finally {
      setLoading(false);
    }
  };

  return { searchResults, loading, handleSearch: searchWordHandler };
};

export default useSearchWord;
