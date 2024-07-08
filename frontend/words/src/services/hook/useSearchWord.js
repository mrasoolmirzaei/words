import { useState } from 'react';
import { searchSynonym } from '../api/synonym';

const useSearchWord = () => {
  const [searchResults, setSearchResults] = useState(null);

  const searchWordHandler = async (query) => {
    try {
      const result = await searchSynonym(query);
      if (result.success) {
        setSearchResults(result.data);
      } else {
        console.error({ error: result.error });
      }
    } catch (error) {
      console.error(`searchWordHandler failed: ${error.message}`);
    }
  };

  return { searchResults, handleSearch: searchWordHandler };
};

export default useSearchWord;
