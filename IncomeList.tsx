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

const IncomeList: React.FC<IncomeListProps> = () => {
  const dispatch = useDispatch();
  const incomeEntries = useSelector((state: RootState) => state.income.entries);
  const [editForm, setEditForm] = useState<EditFormData | null>(null);

  useEffect(() => {
    dispatch(fetchIncomeEntries());
  }, [dispatch]);

  const handleDelete = (id: string) => {
    dispatch(deleteIncome(id));
  };

  const startEdit = (entry: IncomeEntry) => {
    setEditForm({ ...entry });
  };

  const handleEditChange = (e: React.ChangeEvent<HTMLInputElement>, field: keyof EditFormData) => {
    setEditForm(prev => ({
      ...prev,
      [field]: e.target.value,
    } as EditFormData));
  };

  const submitEdit = () => {
    if (editForm) {
      dispatch(updateIncome(editForm)); 
      setEditForm(null);
    }
  };

  return (
    <div>
      {incomeEntries.length > 0 ? (
        <ul>
          {incomeEntries.map((entry: IncomeTEntry) => (
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
                  <button onClick={() => setEditMediaPlan(null)}>Cancel</button>
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