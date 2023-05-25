import React from 'react'
import ReactDOM from 'react-dom/client'
import { BrowserRouter as Router, Route, Routes, useNavigate } from 'react-router-dom'


import {Crud, PageLayout, TokenService} from '.'

TokenService.setToken({
    content: 'sample-token',
})

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <PageLayout>
            <Router>
                <Routes>
                    <Route path='/crud/function/*' element={<Crud namespace='extensions' resource='Function'/>}/>
                </Routes>
            </Router>
        </PageLayout>
    </React.StrictMode>,
)
