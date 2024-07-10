import { useState } from "react";
import useAddSynonym from "../../services/hook/useAddSynonym";
import { lettersOnlyPattern } from "../../constants/regex";

const AddSynonym = () => {
  const { addSynonym } = useAddSynonym();
  const [word, setWord] = useState("");
  const [synonym, setSynonym] = useState("");
  const [wordValidationError, setWordValidationError] = useState("");
  const [synonymValidationError, setSynonymValidationError] = useState("");

  const handleChangeSynonym = (e) => {
    const { value } = e.target;
    setSynonym(value);
    if (lettersOnlyPattern.test(value)) {
      setSynonymValidationError("");
    } else {
      setSynonymValidationError("Please enter only letters.");
    }
  };
  const handleChangeWord = (e) => {
    const { value } = e.target;
    setWord(value);
    if (lettersOnlyPattern.test(value)) {
      setWordValidationError("");
    } else {
      setWordValidationError("Please enter only letters.");
    }
  };
  const handleAddSynonym = () => {
    addSynonym({ word, synonym });
    setWord("");
    setSynonym("");
  };

  return (
    <div className="form-group min-h-7rem w-100">
      <strong>Add Synonym</strong>
      <div className="input-group">
        <button
          className="btn btn-primary h-min-content"
          disabled={wordValidationError || synonymValidationError}
          onClick={handleAddSynonym}
        >
          +
        </button>
        <div>
          <input
            type="text"
            className="form-control"
            placeholder="Word"
            value={word}
            onChange={handleChangeWord}
          />
          {wordValidationError && (
            <p className="text-danger">{wordValidationError}</p>
          )}
        </div>

        <div>
          <input
            type="text"
            className="form-control"
            placeholder="Synonym"
            value={synonym}
            onChange={handleChangeSynonym}
          />
          {synonymValidationError && (
            <p className="text-danger">{synonymValidationError}</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default AddSynonym;
