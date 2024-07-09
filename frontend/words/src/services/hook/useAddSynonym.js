import { useState } from 'react';
import { addSynonym } from '../api/synonym';
import { toast } from 'react-toastify';

const useAddSynonym = () => {
  const [loading, setLoading] = useState(false);

  const addSynonymHandler = async ({ word, synonym: wordSynonym }) => {
    setLoading(true);
    try {
      const result = await addSynonym(word, wordSynonym);
      if (result.ok) {
        toast.success('Synonym added successfully!');
      }else {
        toast.error(result.statusText);
      }
    } finally {
      setLoading(false);
    }
  };

  return { loading, addSynonym: addSynonymHandler };
};

export default useAddSynonym;
