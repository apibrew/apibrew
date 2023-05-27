import React from 'react'
import ReactDOM from 'react-dom/client'
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'


import {Crud, DashboardLayout, TokenService} from '.'

TokenService.setToken({
    content: 'sample-token',
})

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <Router>
            <Routes>
                <Route path='/crud/function/*' element={<DashboardLayout>
                    <Crud namespace='extensions' resource='Function'/>
                </DashboardLayout>}/>
            </Routes>
        </Router>
    </React.StrictMode>,
)
