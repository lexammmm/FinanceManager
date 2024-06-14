import React, 'useState' from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { AppState } from './store';
import { Expense } from './types';
import * as expenseActions from './expenseActions';

type ExpenseItemProps = {
    expense: Expense;
    onEdit: (id: string, expense: Expense) => void;
    onDelete: (id: string) => void;
};

const ExpenseItem: React.FC<ExpenseitemProps> = ({ expense, onEdit, onDelete }) => {
    const [isEditing, setIsEditing] = useState(false);
    const [editedExpense, setEditedExpense] = useState(expense);

    const handleSave = () => {
        onEdit(editedExpense.id, editedExpense);
        setIsEditing(false);
    };

    if (isEditing) {
        return (
            <div>
                <input type="text" value={editedExpense.name} onChange={e => setEditedExpense({ ...editedExpense, name: e.target.value })} />
                <input type="number" value={editedExpense.amount} onChange={e => setEditedExpense({ ...editedExpense, amount: Number(e.target.value) })} />
                <input type="date" value={editedExpense.date} onChange={e => setEditedExpense({ ...editedExpense, date: e.target.value })} />
                <button onClick={handleSave}>Save</button>
                <button onClick={() => setIsEditing(false)}>Cancel</button>
            </div>
        );
    }

    return (
        <div>
            <div>Name: {expense.name}</div>
            <div>Amount: ${expense.amount}</div>
            <div>Date: {new Date(expense.date).toLocaleDateString()}</div>
            <button onClick={() => setIsEditing(true)}>Edit</button>
            <button onClick={() => onDelete(expense.id)}>Delete</button>
        </div>
    );
};

const ExpenseList: React.FC = () => {
    const dispatch = useDispatch();
    const expenses = useSelector((state: AppState) => state.expense);

    const handleDelete = (id: string) => {
        dispatch(expenseActions.deleteExpense(id));
    };

    const handleEdit = (id: string, updatedExpense: Expense) => {
        dispatch(expenseActions.updateExpense(id, updatedExpense));
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