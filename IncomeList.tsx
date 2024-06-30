import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { RootState } from './store';
import { fetchIncomeEntries, deleteIncomeEntry } from './incomeSlice';

interface IncomeEntry {
  id: string;
  source: string;
  amount: number;
  date: string;
}

interface IncomeListProps {}

const IncomeList: React.FC<IncomeListProps> = () => {
  const dispatch = useDispatch();
  const incomeEntries = useSelector((state: RootState) => state.income.entries);

  useEffect(() => {
    dispatch(fetchIncomeEntries());
  }, [dispatch]);

  const handleDelete = (id: string) => {
    dispatch(deleteIncomeEntry(id));
  };

  return (
    <div>
      {incomeEntries.length > 0 ? (
        <ul>
          {incomeEntries.map((entry: IncomeEntry) => (
            <li key={entry.id}>
              {entry.source} - ${entry.amount} - {entry.date}
              <button onClick={() => handleDelete(entry.id)}>Delete</button>
              <button>Edit</button>
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