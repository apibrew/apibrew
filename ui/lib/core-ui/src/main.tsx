import React from 'react'
import ReactDOM from 'react-dom/client'
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'


import {DashboardLayout, TokenService} from '.'
import {DynamicComponent} from "./components/dynamic/DynamicComponent.tsx";

TokenService.setToken({
    content: 'sample-token',
})

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <Router>
            <Routes>
                <Route path='/crud/function/*' element={(
                    <DashboardLayout>
                        <DynamicComponent component={'Layout/CrudFunction'}/>
                    </DashboardLayout>
                )}/>
            </Routes>
        </Router>
    </React.StrictMode>,
)
