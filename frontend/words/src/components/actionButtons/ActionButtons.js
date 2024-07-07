import React from 'react';
import useAddWord from '../../services/hook/useAddWord';
import useAddSynonym from '../../services/hook/useAddSynonym';

const ActionButtons = () => {

  const { addWord } = useAddWord();
  const { addSynonym } = useAddSynonym();

  return (
    <div className="d-flex justify-content-between mt-3 gap-2">
      <button type="button" className="btn btn-outline-primary" onClick={addWord}>
        Add Word
      </button>
      <button type="button" className="btn btn-outline-info" onClick={addSynonym}>
        Add Synonym
      </button>
    </div>
  );
};

export default ActionButtons;
