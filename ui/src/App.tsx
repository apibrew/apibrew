import './App.css'
import { BrowserRouter as Router, Route, Routes, useNavigate } from 'react-router-dom'
import { Test } from './test/test'
import { Login } from './pages/login/login'
import { DashboardLayout } from './layout/dashboard-layout'
import { useEffect } from 'react'
import { BaseLayout } from './layout/BaseLayout'
import { TokenService } from './service/token'
import { Migrate } from './pages/migrate'
import { AppDesigner } from './pages/app-designer'

function Dashboard(): JSX.Element {
    const isLoggedIn = TokenService.isLoggedIn()
    const navigate = useNavigate()

    useEffect(() => {
        if (!isLoggedIn) {
            navigate('/login')
        }
    })
    if (!isLoggedIn) {
        return <></>
    }

    return <>
        <DashboardLayout>
            <Routes>
                <Route path='test' element={<Test></Test>}/>
                <Route path='app-designer' element={<AppDesigner/>}/>
                <Route path='migrate' element={<Migrate/>}/>
            </Routes>
        </DashboardLayout>
    </>
}

function App(): JSX.Element {
    return (
        <BaseLayout>
            <Router>
                <Routes>
                    <Route path='/dashboard/*' element={<Dashboard></Dashboard>}/>
                    <Route path='/login' element={<Login></Login>}/>
                </Routes>
            </Router>
        </BaseLayout>
    )
}

export default App
