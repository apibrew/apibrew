import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import { Route, BrowserRouter as Router, Routes } from "react-router-dom";
import { ResourceDesigner } from './components/resource-designer/ResourceDesigner';
import { TokenService } from '@apibrew/ui-lib';
import { LogicDesigner } from './components/logic-designer/LogicDesigner';

TokenService.storeAccessToken({
    "term": "VERY_LONG",
    "content": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tL2FwaWJyZXcvYXBpYnJldyIsInN1YiI6ImFkbWluIiwiYXVkIjpbImdpdGh1Yi5jb20vYXBpYnJldy9hcGlicmV3Il0sImV4cCI6MTc1MDc1Mzg4MSwibmJmIjoxNjg3NjgxODgxLCJpYXQiOjE2ODc2ODE4ODEsImp0aSI6ImE4Njg3NTUyLWJjNTUtNDczYi1iZmU0LTkzZDlmMjE1OGUwNCIsInNlY3VyaXR5Q29uc3RyYWludHMiOlt7Im5hbWVzcGFjZSI6IioiLCJyZXNvdXJjZSI6IioiLCJwcm9wZXJ0eSI6IioiLCJvcGVyYXRpb24iOiJGVUxMIiwicm9sZSI6InJvb3QifV0sInVzZXJuYW1lIjoiYWRtaW4iLCJyb2xlcyI6WyJyb290Il0sInVpZCI6IjI5MDA1NTY2LWZlYjktMTFlZC1hNGM5LWM2YWFjNjRmMTliMiJ9.SYyYpg2KvoDZB3ow64lrOkdFLENaNt-gNr7rnjyHrhys7OVgHUo4sfe-lZ4YfGwJswBC08xBhHwg8Xsh13DVHjfREIioK3VLYFSBKQ0eU1ZtPikdBYlbK0g0Z0FQuypRjqbTcbkMtXZHQ3n4tPdTXyBvxTSjQB4AgS49wmrNJ20",
    "expiration": "2025-06-24T08:31:21.010421Z"
})

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <CssBaseline />
        <Box display='flex' sx={{
            height: '100vh',
            width: '100%',
            background: 'rgb(240,240,240)',
        }}>
            <Router>
                <Routes>
                    <Route path='/resource-designer' element={<ResourceDesigner name='Default' />} />
                    <Route path='/logic-designer' element={<LogicDesigner />} />
                </Routes>
            </Router>
        </Box>
    </React.StrictMode>,
)
