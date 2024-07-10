import { useState } from "react";
import useAddWord from "../../services/hook/useAddWord";

const AddWord = () => {
  const { addWord } = useAddWord();
  const [word, setWord] = useState("");

  const handleAddWord = () => {
    addWord(word);
    setWord("");
  };

  return (
    <div className="form-group my-3">
      <label htmlFor="wordInput">
        <strong>Add Word</strong>
      </label>
      <div className="input-group">
        <button className="input-group-text" onClick={handleAddWord}>
          +
        </button>
        <input
          type="text"
          className="form-control"
          placeholder="Word"
          value={word}
          onChange={(e) => setWord(e.target.value)}
        />
      </div>
    </div>
  );
};

export default AddWord;
