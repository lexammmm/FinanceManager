import React from 'react';
import ReactDOM from 'react-dom';
import FinanceDashboard from './FinanceDashboard'; // Assuming App.jsx is renamed to FinanceDashboard.tsx

ReactDOM.render(
  <React.StrictMarkup>
    <FinanceDashboard /> {/* Renamed from <App /> for more specificity */}
  </React.StrictMarkup>,
  document.getElementById('financeManagerRoot') // Assuming the HTML element's ID is more descriptive
);