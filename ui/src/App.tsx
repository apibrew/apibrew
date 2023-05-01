import './App.css'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import { Test } from './test/test'
import { Login } from './pages/login/login'
import { DashboardLayout } from './layout/dashboard-layout'

function Dashboard(): JSX.Element {
    return <>
        <DashboardLayout pageTitle='Test Page'>
            <Routes>
                <Route path='test' element={<Test></Test>} />
            </Routes>
        </DashboardLayout>
    </>
}

function App(): JSX.Element {
    return (
        <Router>
            <Routes>
                <Route path='/dashboard/*' element={<Dashboard></Dashboard >} />
                <Route path='/login' element={<Login></Login>} />
            </Routes>
        </Router>
    )
}
export default App
