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
  successMsg: string;
}

const categories = ['Salary', 'Investments', 'Other'];

const IncomeForm: React.FC = () => {
  const [incomeData, setIncomeData] = useState<IncomeData>({ source: '', amount: 0, date: '', category: '' });
  const [formStatus, setFormStatus] = useState<FormStatus>({ isLoading: false, errorMsg: '', successMsg: '' });
  const dispatch = useDispatch();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    const updatedValue = name === 'amount' ? parseFloat(value) || 0 : value;
    updateIncomeData(name, updatedValue);
  };

  const updateIncomeData = (key: string, value: string | number) => {
    setIncomeData({ ...incomeData, [key]: value });
  };

  const validateForm = (): boolean => {
    if (incomeData.amount <= 0) {
      updateFormStatus('Please enter a valid amount.', '', false);
      return false;
    }
    return true;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if(!validateForm()) return;
    
    updateFormStatus('', '', true); // Reset status with loading true
    dispatch(addIncome(incomeData));
    
    try {
      await submitIncomeData();
      updateFormStatus('', 'Income entry added successfully.', false);
    } catch (error: any) {
      console.error('Error saving the income: ', error.message);
      updateFormStatus('Failed to save the income entry. Please try again.', '', false);
    }
  };

  const submitIncomeData = () => {
    return axios.post(`${process.env.REACT_APP_BACKEND_URL}/addIncome`, incomeData);
  }

  const updateFormStatus = (errorMsg: string, successMsg: string, isLoading: boolean) => {
    setFormStatus({ isLoading, errorMsg, successMsg });
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