import { useState } from 'react';
import useAddWord from "../../services/hook/useAddWord";

const AddWord = () => {
  const { addWord } = useAddWord();
  const [word, setWord] = useState('');

  const handleAddWord = () => {
    addWord(word);
    setWord('');
  };

  return (
    <div className="input-group">
      <div className="input-group-prepend">
        <button className="input-group-text" onClick={handleAddWord}>
          +
        </button>
      </div>
      <input
        type="text"
        className="form-control"
        placeholder='Word'
        value={word}
        onChange={(e) => setWord(e.target.value)}
      />
    </div>
  );
};

export default AddWord;
