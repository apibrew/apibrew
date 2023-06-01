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
                    <Route path='/designer/*' element={(
                        <Designer name='Default' />
                    )}/>
                    <Route path='/crud/function/*' element={(
                        <Crud namespace={'logic'} resource={'Function'}/>
                    )}/>
                    <Route path='/crud/resource-rule/*' element={(
                        <Crud namespace={'logic'} resource={'ResourceRule'}/>
                    )}/>
                    <Route path='/crud/resource-trigger/*' element={(
                        <Crud namespace={'logic'} resource={'FunctionTrigger'}/>
                    )}/>
                    <Route path='/crud/country/*' element={(
                        <Crud namespace={'default'} resource={'country'}/>
                    )}/>
                    <Route path='/crud/person/*' element={(
                        <Crud namespace={'default'} resource={'person'}/>
                    )}/>
                </Routes>
            </Router>
        </BaseLayout>
    </React.StrictMode>,
)
