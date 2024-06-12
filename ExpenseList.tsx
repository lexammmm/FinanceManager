import React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { AppState } from './store';
import { Expense } from './types';
import * as expenseActions from './expenseActions';

type ExpenseItemProps = {
    expense: Expense;
    onEdit: (expense: Expense) => void;
    onDelete: (id: string) => void;
};

const ExpenseItem: React.FC<ExpenseItemProps> = ({ expense, onEdit, onDelete }) => (
    <div>
        <div>Name: {expense.name}</div>
        <div>Amount: ${expense.amount}</div>
        <div>Date: {new Date(expense.date).toLocaleDateString()}</div>
        <button onClick={() => onEdit(expense)}>Edit</button>
        <button onClick={() => onDelete(expense.id)}>Delete</button>
    </div>
);

const ExpenseList: React.FC = () => {
    const dispatch = useDispatch();
    const expenses = useSelector((state: AppState) => state.expense);

    const handleDelete = (id: string) => {
        dispatch(expenseActions.deleteExpense(id));
    };

    const handleEdit = (expense: Expense) => {
        console.log('Edit logic here...', expense);
    };

    return (
        <div>
            {expenses.map(expense => (
                <ExpenseItem
                    key={expense.id}
                    expense={expense}
                    onEdit={handleEdit}
                    onDelete={handleDelete}
                />
            ))}
        </div>
    );
};

export default ExpenseList;