import { useState } from 'react';
import { addWord } from '../api/word';

const useAddWord = () => {
  const [word, setWord] = useState('');

  const addWordHandler = async (word) => {
    const result = await addWord(word);
    setWord(result);
  };

  return { word, addWord: addWordHandler };
};

export default useAddWord;
