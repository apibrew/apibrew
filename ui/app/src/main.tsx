import React from 'react'
import ReactDOM from 'react-dom/client'
import './module.ts'
import {App} from "./App.tsx";

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <App/>
    </React.StrictMode>,
)
