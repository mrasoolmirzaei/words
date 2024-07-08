import { useState } from 'react';
import { addSynonym } from '../api/synonym';

const useAddSynonym = () => {
  const [synonym, setSynonym] = useState('');

  const addSynonymHandler = async (word, synonym) => {
    try {
      const result = await addSynonym(word, synonym);
      if (result.success) {
        setSynonym(result.data);
      } else {
        console.error({ error: result.error });
      }
    } catch (error) {
      console.error(`addSynonymHandler failed: ${error.message}`);
    }
  };

  return { synonym, addSynonym: addSynonymHandler };
};

export default useAddSynonym;
