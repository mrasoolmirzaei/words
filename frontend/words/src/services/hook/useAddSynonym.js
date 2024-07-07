import { useState } from 'react';
import { addSynonym } from '../api/synonym';

const useAddSynonym = () => {
  const [synonym, setSynonym] = useState('');

  const addSynonymHandler = async (word, synonym) => {
    const result = await addSynonym(word, synonym);
    setSynonym(result);
  };

  return { synonym, addSynonym: addSynonymHandler };
};

export default useAddSynonym;
