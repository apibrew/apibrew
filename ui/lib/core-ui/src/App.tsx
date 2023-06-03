import { BrowserRouter as Router, Route, Routes, useNavigate } from 'react-router-dom'
import { Test } from './test/test'
import { Login } from './pages/login/login'
import React, { JSX, useEffect } from 'react'
import { AppDesigner } from './pages/app-designer'
import { CrudPage } from './pages/crud-page/CrudPage'
import {TokenService} from "./service";
import {BaseLayout, DashboardLayout} from "./layout";

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
        <Route path='/crud/:namespace/:resource/*' element={(
            <CrudPage/>
        )}/>
      </Routes>
    </DashboardLayout>
  </>
}

export function App (): JSX.Element {
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

