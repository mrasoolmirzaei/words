import { useState } from 'react';
import { searchSynonym } from '../api/synonym';

const useSearchWord = () => {
  const [searchResults, setSearchWord] = useState('');

  const searchWordHandler = async (query) => {
    const result = await searchSynonym(query);
    setSearchWord(result);
  };

  return { searchResults, handleSearch: searchWordHandler };
};

export default useSearchWord;
