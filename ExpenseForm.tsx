import React, { useState, ChangeEvent, FormEvent } from 'react';
import { useDispatch } from 'react-redux';
import axios from 'axios';

interface ExpenseFormState {
    name: string;
    amount: number;
    date: string;
}

const ExpenseForm: React.FC = () => {
    const [formState, setFormState] = useState<ExpenseDownInputState>({
        name: '',
        amount: 0,
        date: ''
    });

    const dispatch = useDispatch();

    const handleInputChange = (e: ChangeEvent<HTMLInputElement>) => {
        setFormState({
            ...formState,
            [e.target.name]: e.target.value
        });
    };

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();

        dispatch({
            type: 'ADD_EXPENSE',
            payload: formState
        });

        try {
            await axios.post(`${process.env.REACT_APP_BACKEND_URL}/expenses`, formState);
            alert('Expense added successfully!');
        } catch (error) {
            console.error('Error adding expense:', error);
            alert('Failed to add expense.');
        }

        setFormState({
            name: '',
            amount: 0,
            date: ''
        });
    };

    return (
        <form onSubmit={handleSubmit}>
            <div>
                <label htmlFor="name">Expense Name:</label>
                <input
                    type="text"
                    id="name"
                    name="name"
                    value={formState.name}
                    onChange={handleInputChange}
                />
            </div>
            <div>
                <label htmlFor="amount">Amount:</label>
                <input
                    type="number"
                    id="amount"
                    name="amount"
                    value={formState.amount}
                    onChange={handleInputChange}
                />
            </div>
            <div>
                <label htmlFor="date">Date:</label>
                <input
                    type="date"
                    id="date"
                    name="date"
                    value={formState.date}
                    onChange={handleInputChange}
                />
            </div>
            <button type="submit">Add Expense</button>
        </form>
    );
};

export default ExpenseForm;