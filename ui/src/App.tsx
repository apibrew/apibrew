import React from 'react';
import './App.css';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Test } from './test/test'

function App(): JSX.Element {
  return (
    <Router>
      <Routes>
        <Route path='/test' element={<Test></Test>} />
      </Routes>
    </Router>
  )

}

export default App;
