import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import {Route, BrowserRouter as Router, Routes} from "react-router-dom";
import {Sdk} from "./components/sdk/Sdk.tsx";
import { Swagger } from './components/swagger/Swagger.tsx';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <CssBaseline/>
        <Box display='flex' sx={{
            height: '100vh',
            width: '100%',
            background: 'rgb(240,240,240)',
        }}>
            <Router>
                <Routes>
                    <Route path='/sdk' element={<Sdk/>}/>
                    <Route path='/swagger' element={<Swagger/>}/>
                </Routes>
            </Router>
        </Box>
    </React.StrictMode>,
)
