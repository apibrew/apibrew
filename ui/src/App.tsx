import './App.css'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import { Test } from './test/test'
import { Login } from './pages/login/login'
import { DashboardLayout } from './layout/dashboard-layout'

function Dashboard(): JSX.Element {
    return <>
        <DashboardLayout>
            <Routes>
                <Route path='test' element={<Test></Test>} />
            </Routes>
        </DashboardLayout>
    </>
}
import { BaseLayout } from './layout/BaseLayout'

function App(): JSX.Element {
    return (
        <BaseLayout>
            <Router>
                <Routes>
                    <Route path='/dashboard/*' element={<Dashboard></Dashboard >}/>
                    <Route path='/login' element={<Login></Login>}/>
                </Routes>
            </Router>
        </BaseLayout>
    )
}

export default App
