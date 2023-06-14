import React from 'react'
import ReactDOM from 'react-dom/client'
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import {Sdk} from "./components/sdk/Sdk.tsx";
import {Box} from "@mui/material";
import './index.css'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <Box display='flex' sx={{
            height: '100vh',
            width: '100%',
            background: 'rgb(240,240,240)',
        }}>
            <Router>
                <Routes>
                    <Route path='/sdk' element={<Sdk/>}/>
                </Routes>
            </Router>
        </Box>
    </React.StrictMode>,
)
