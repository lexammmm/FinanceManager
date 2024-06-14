import React, { useState, FormEvent } from 'react';
import { useDispatch } from 'react-redux';
import axios from 'axios';

interface ExpenseFormState {
  name: string;
  amount: number;
  date: string;
}

const ExpenseForm: React.FC = () => {
  const [formState, setFormState] = useState<ExpenseFormState>({
    name: '',
    amount: 0,
    date: '',
  });

  const dispatch = useDispatch();

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormState({
      ...formState,
      [name]: name === 'amount' ? parseFloat(value) : value,
    });
  };

  const handleFormSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    dispatch({
      type: 'ADD_EXPENSE', 
      payload: formState,
    });

    setFormState({ name: '', amount: 0, date: '' });
  };

  return (
    <form onSubmit={handleFormSubmit}>
      <div>
        <label htmlFor="name">Name:</label>
        <input
          id="name"
          name="name"
          type="text"
          value={formState.name}
          onChange={handleInputChange}
        />
      </div>
      <div>
        <label htmlFor="amount">Amount:</label>
        <input
          id="amount"
          name="amount"
          type="number"
          value={formState.amount}
          onChange={handleInputChange}
        />
      </div>
      <div>
        <label htmlFor="date">Date:</label>
        <input
          id="date"
          name="date"
          type="date"
          value={formState.date}
          onChange={handleInputChange}
        />
      </div>
      <button type="submit">Add Expense</button>
    </form>
  );
};

export default ExpenseForm;