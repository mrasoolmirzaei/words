import { useState } from "react";
import { addWord } from "../api/word";
import { toast } from "react-toastify";

const useAddWord = () => {
  const [loading, setLoading] = useState(false);

  const addWordHandler = async (word) => {
    setLoading(true);
    const result = await addWord(word);
    if (result.ok) {
      toast.success("Word added successfully!");
    } else {
      toast.error(result.statusText);
    }
    setLoading(false);
  };

  return { loading, addWord: addWordHandler };
};

export default useAddWord;
