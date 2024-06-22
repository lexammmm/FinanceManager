import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import axios from 'axios';
import { addIncome } from './actions/incomeActions';

interface IncomeData {
  source: string;
  amount: number;
  date: string;
}

interface FormStatus {
  isLoading: boolean;
  errorMsg: string;
}

const IncomeForm: React.FC = () => {
  const [incomeData, setIncomeData] = useState<IncomeData>({ source: '', amount: 0, date: '' });
  const [formStatus, setFormStatus] = useState<FormStatus>({ isLoading: false, errorMsg: '' });
  const dispatch = useDispatch();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setIncomeData({ ...incomeData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setFormStatus({ ...formWStatus, isLoading: true, errorMsg: '' }); // Reset and show loading
    dispatch(addIncome(incomeData));

    try {
      await axios.post(`${process.env.REACT_APP_BACKEND_URL}/addIncome`, incomeData);
      alert('Income entry added successfully.');
    } catch (error: any) {
      console.error('There was an error saving the income entry: ', error.message);
      setFormStatus({ ...formStatus, isLoading: false, errorMsg: 'Failed to save the income entry. Please try again.' });
    } finally {
      setFormStatus({ ...formStatus, isLoading: false }); // Reset loading state regardless of outcome
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      {formStatus.errorMsg && <p>Error: {formStatus.errorMsg}</p>}
      <div>
        <label htmlFor="source">Income Source</label>
        <input
          type="text"
          id="source"
          name="source"
          value={incomeData.source}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label htmlFor="amount">Amount</label>
        <input
          type="number"
          id="amount"
          name="amount"
          value={incomeData.amount}
          onChange={handleChange}
          required
        />
      </div>
      <div>
        <label htmlFor="date">Date</label>
        <input
          type="date"
          id="date"
          name="date"
          value={incomeData.date}
          onChange={handleChange}
          required
        />
      </div>
      <button type="submit" disabled={formStatus.isLoading}>Add Income</button>
    </form>
  );
};

export default IncomeForm;