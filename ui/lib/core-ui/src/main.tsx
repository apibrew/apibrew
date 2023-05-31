import React from 'react'
import ReactDOM from 'react-dom/client'
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'
import {BaseLayout, Crud, TokenService} from '.'
import {Designer} from './components/designer/Designer'

TokenService.setToken({
    content: 'sample-token',
})

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <BaseLayout>
            <Router>
                <Routes>
                    <Route path='/crud/designer/*' element={(
                        <Designer/>
                    )}/>
                    <Route path='/crud/function/*' element={(
                        <Crud namespace={'extensions'} resource={'Function'}/>
                    )}/>
                    <Route path='/crud/country/*' element={(
                        <Crud namespace={'default'} resource={'country'}/>
                    )}/>
                </Routes>
            </Router>
        </BaseLayout>
    </React.StrictMode>,
)
