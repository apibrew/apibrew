import React from 'react'
import ReactDOM from 'react-dom/client'
import {TokenService} from '.'
import {App} from "./App.tsx";

TokenService.setToken({
    content: 'sample-token',
})

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <App/>
    </React.StrictMode>,
)
