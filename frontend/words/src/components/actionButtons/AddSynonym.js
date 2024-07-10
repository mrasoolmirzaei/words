import { useState } from "react";
import useAddSynonym from "../../services/hook/useAddSynonym";

const AddSynonym = () => {
  const { addSynonym } = useAddSynonym();
  const [word, setWord] = useState("");
  const [synonym, setSynonym] = useState("");

  const handleSynonym = (e) => {
    setSynonym(e.target.value);
  };
  const handleWord = (e) => {
    setWord(e.target.value);
  };
  const handleAddSynonym = () => {
    addSynonym({ word, synonym });
    setWord("");
    setSynonym("");
  };

  return (
    <div className="form-group my-3 w-100">
      <label htmlFor="wordInput">
        <strong>Add Synonym</strong>
      </label>
      <div className="input-group">
        <button className="input-group-text" onClick={handleAddSynonym}>
          +
        </button>
        <input
          type="text"
          className="form-control"
          placeholder="Word"
          value={word}
          onChange={handleWord}
        />
        <input
          type="text"
          className="form-control"
          placeholder="Synonym"
          value={synonym}
          onChange={handleSynonym}
        />
      </div>
    </div>
  );
};

export default AddSynonym;
