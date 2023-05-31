import React from 'react'
import ReactDOM from 'react-dom/client'
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'


import {BaseLayout, Crud, TokenService} from '.'

TokenService.setToken({
    content: 'sample-token',
})

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <BaseLayout>
            <Router>
                <Routes>
                    <Route path='/crud/function/*' element={(
                        <Crud namespace={'extensions'} resource={'Function'}/>
                    )}/>
                </Routes>
            </Router>
        </BaseLayout>
    </React.StrictMode>,
)
