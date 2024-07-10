import { useState } from "react";
import useAddWord from "../../services/hook/useAddWord";
import { lettersOnlyPattern } from "../../constants/regex";

const AddWord = () => {
  const { addWord } = useAddWord();
  const [word, setWord] = useState("");
  const [validationError, setValidationError] = useState("");

  const handleAddWord = () => {
    addWord(word);
    setWord("");
  };
  const handleChangeWord = (e) => {
    const { value } = e.target;
    setWord(value);
    if (lettersOnlyPattern.test(value)) {
      setValidationError("");
    } else {
      setValidationError("Please enter only letters.");
    }
  };

  return (
    <div className="form-group min-h-7rem w-100">
      <strong>Add Word</strong>
      <div className="input-group w-100">
        <input
          type="text"
          className="form-control"
          placeholder="Word"
          value={word}
          onChange={handleChangeWord}
        />
        <button
          className="btn btn-primary"
          disabled={validationError}
          onClick={handleAddWord}
        >
          +
        </button>
      </div>
      {validationError && <p className="text-danger">{validationError}</p>}
    </div>
  );
};

export default AddWord;
