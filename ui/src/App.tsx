import React from 'react'
import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Test } from './test/test'
import { Login } from './pages/login/login'

function App(): JSX.Element {
    return (
        <Router>
            <Routes>
                <Route path='/test' element={<Test></Test>} />
                <Route path='/login' element={<Login></Login>} />
            </Routes>
        </Router>
    )
}

export default App
