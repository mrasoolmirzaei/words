import React from 'react';
import useAddWord from '../../services/hook/useAddWord';
import useAddSynonym from '../../services/hook/useAddSynonym';
const AddInput = ({add,type}) => <div className="input-group">
<div className="input-group-prepend">
  <button className="input-group-text" onClick={add}>Add</button>
</div>
<input type="text" className="form-control" id="inlineFormInputGroupUsername" placeholder={type} />
</div>
const ActionButtons = () => {

  const { addWord } = useAddWord();
  const { addSynonym } = useAddSynonym();

  return (
    <div className="d-flex justify-content-between mt-3 gap-2">
      <AddInput add={addSynonym} type='synonym'/>
      <AddInput add={addWord} type='word'/>
    </div>
  );
};

export default ActionButtons;
