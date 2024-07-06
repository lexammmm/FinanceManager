import React, { useEffect, useState } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { RootState } from './store';
import { fetchIncomeEntries, deleteIncome, updateIncome } from './incomeSlice';

interface IncomeEntry {
  id: string;
  source: string;
  amount: number;
  date: string;
}

interface IncomeListProps {}

interface EditFormData {
  id: string;
  source: string;
  amount: number;
  date: string;
}

const logAction = (message: string) => {
  console.log(message);
};

const IncomeList: React.FC<IncomeListProps> = () => {
  const dispatch = useDispatch();
  const incomeEntries = useSelector((state: RootState) => state.income.entries);
  const [editForm, setEditForm] = useState<EditFormData | null>(null);
  const [minimumAmountFilter, setMinimumAmountFilter] = useState<number>(0);

  useEffect(() => {
    dispatch(fetchIncomeEntries());
    logAction('Fetching income entries...');
  }, [dispatch]);

  const handleDelete = (id: string) => {
    dispatch(deleteIncome(id));
    logAction(`Deleting income entry with ID: ${id}`);
  };

  const startEdit = (entry: IncomeEntry) => {
    setEditForm({ ...entry });
    logAction(`Starting to edit income entry with ID: ${entry.id}`);
  };

  const handleEditChange = (e: React.ChangeEvent<HTMLInputElement>, field: keyof EditFormData) => {
    setEditForm(prev => {
      const updatedForm = { ...prev, [field]: e.target.type === 'number' ? parseInt(e.target.value) : e.target.value } as EditFormData;
      return updatedForm;
    });
    logAction(`Editing field: ${field}`);
  };

  const submitEdit = () => {
    if(editForm) {
      dispatch(updateIncome(editForm));
      logOperation(`Submitted edit for income entry with ID: ${editForm.id}`);
      setEditForm(null);
    }
  };

  const cancelEdit = () => {
    logAction('Edit canceled.');
    setEditForm(null);
  };
  
  const handleFilterChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setMinimumAmountFilter(parseInt(e.target.value));
  };
  
  const filteredIncomeEntries = incomeEntries.filter(entry => entry.amount >= minimumAmountFilter);

  return (
    <div>
      <input
        type="number"
        value={minimumAmountFilter}
        onChange={handleFilterChange}
        placeholder="Minimum Amount"
      />
      {filteredIncomeEntries.length > 0 ? (
        <ul>
          {filteredIncomeEntries.map((entry: IncomeEntry) => (
            <li key={entry.id}>
              {editForm && editForm.id === entry.id ? (
                <div>
                  <input
                    type="text"
                    value={editForm.source}
                    onChange={e => handleEditChange(e, 'source')}
                  />
                  <input
                    type="number"
                    value={editForm.amount}
                    onChange={e => handleEditChange(e, 'amount')}
                  />
                  <input
                    type="date"
                    value={editForm.date}
                    onChange={e => handleEditChange(e, 'date')}
                  />
                  <button onClick={submitEdit}>Save</button>
                  <button onClick={cancelEdit}>Cancel</button>
                </div>
              ) : (
                <>
                  {entry.source} - ${entry.amount} - {entry.date}
                  <button onClick={() => handleDelete(entry.id)}>Delete</button>
                  <button onClick={() => startEdit(entry)}>Edit</button>
                </>
              )}
            </li>
          ))}
        </ul>
      ) : (
        <p>No income entries found.</p>
      )}
    </div>
  );
};

export default IncomeList;