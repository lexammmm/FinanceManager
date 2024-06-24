import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import axios from 'axios';
import { addIncome } from './actions/incomeActions';

interface IncomeData {
  source: string;
  amount: number;
  date: string;
  category: string;
}

interface FormStatus {
  isLoading: boolean;
  errorMsg: string;
  successMsg: string; // New addition for success message
}

const categories = ['Salary', 'Investments', 'Other']; // Example categories

const IncomeForm: React.FC = () => {
  const [incomeData, setIncomeData] = useState<IncomeData>({ source: '', amount: 0, date: '', category: '' });
  const [formStatus, setFormStatus] = useState<FormStatus>({ isLoading: false, errorMsg: '', successMsg: '' });
  const dispatch = useDispatch();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    setIncomeData({ ...incomeData, [e.target.name]: e.target.name === 'amount' ? parseFloat(e.target.value) || 0 : e.target.value });
  };

  const validateForm = () => {
    if (incomeData.amount <= 0) {
      setFormStatus({ ...formStatus, errorMsg: 'Please enter a valid amount.' });
      return false;
    }
    return true;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if(!validateForm()) return; // Form validation before proceeding
    
    setFormStatus({ ...formStatus, isLoading: true, errorMsg: '', successMsg: '' }); // Reset and show loading
    dispatch(addIncome(incomeData));

    try {
      await axios.post(`${process.env.REACT_APP_BACKEND_URL}/addIncome`, incomeData);
      setFormStatus({ ...form }
, isLoading: false, successMsg: 'Income entry added successfully.' });
    } catch (error: any) {
      console.error('There was an error saving the income entry: ', error.message);
      setFormStatus({ ...form }
, isLoading: false, errorMsg: 'Failed to save the income entry. Please try again.' });
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      {formStatus.errorMsg && <p>Error: {formStatus.errorMsg}</p>}
      {formStatus.successMsg && <p>Success: {formStatus.successMsg}</p>}
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
          value={incomeData.amount.toString()}
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
      <div>
        <label htmlFor="category">Category</label>
        <select
          id="category"
          name="category"
          value={incomeData.category}
          onChange={handleChange}
          required
        >
          <option value="">Select a category</option>
          {categories.map((category) => (
            <option key={category} value={category}>{category}</option>
          ))}
        </select>
      </div>
      <button type="submit" disabled={formStatus.isLoading}>Add Income</button>
    </form>
  );
};

export default IncomeForm;