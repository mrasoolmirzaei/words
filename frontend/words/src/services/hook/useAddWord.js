import { useState } from 'react';
import { addWord } from '../api/word';

const useAddWord = () => {
  const [word, setWord] = useState('');

  const addWordHandler = async (word) => {
    try {
      const result = await addWord(word);
      if (result.success) {
        setWord(result.data);
      } else {
        console.error({ error: result.error });
      }
    } catch (error) {
      console.error(`addWordHandler failed: ${error.message}`);
    }
  };

  return { word, addWord: addWordHandler };
};

export default useAddWord;
