import React from 'react'
import ReactDOM from 'react-dom/client'
import {BrowserRouter as Router, Route, Routes, useParams} from 'react-router-dom'
import {BaseLayout, Crud, TokenService} from '.'
import {Designer} from './components/designer/Designer'

TokenService.setToken({
    content: 'sample-token',
})

export function CrudPage() {
    const params = useParams()
    console.log(params)
    return <Crud namespace={params.namespace} resource={params.resource}/>
}

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <BaseLayout>
            <Router>
                <Routes>
                    <Route path='/designer/*' element={(
                        <Designer name='Default'/>
                    )}/>
                    <Route path='/crud/:namespace/:resource/*' element={(
                        <CrudPage/>
                    )}/>
                </Routes>
            </Router>
        </BaseLayout>
    </React.StrictMode>,
)
