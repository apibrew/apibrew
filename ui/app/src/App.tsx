import './App.css'
import { BrowserRouter as Router, Route, Routes, useNavigate } from 'react-router-dom'
import { Test } from './test/test'
import { Login } from './pages/login/login'
import { JSX, useEffect } from 'react'
import { AppDesigner } from './pages/app-designer'
import { CrudPage } from './pages/crud-page/CrudPage'
import { BaseLayout, DashboardLayout, TokenService } from '@apibrew/core-ui'

function Dashboard (): JSX.Element {
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
        <Route path="test" element={<Test></Test>}/>
        <Route path="app-designer" element={<AppDesigner/>}/>
        {/* Cruds */}
        <Route path="country/*" element={<CrudPage namespace="default" resource="country"/>}/>
        <Route path="city/*" element={<CrudPage namespace="default" resource="city"/>}/>
      </Routes>
    </DashboardLayout>
  </>
}

function App (): JSX.Element {
  return (
    <BaseLayout>
      <Router>
        <Routes>
          <Route path="/dashboard/*" element={<Dashboard></Dashboard>}/>
          <Route path="/login" element={<Login></Login>}/>
        </Routes>
      </Router>
    </BaseLayout>
  )
}

export default App
