import './App.css'
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom'
import {Test} from './test/test'
import {Login} from './pages/login/login'
import {BaseLayout} from './layout/BaseLayout'
import {ResourcesDesigner} from "./pages/resources/ResourcesDesigner";

function App(): JSX.Element {
    return (
        <BaseLayout>
            <Router>
                <Routes>
                    <Route path='/test' element={<Test></Test>}/>
                    <Route path='/login' element={<Login></Login>}/>
                    <Route path='/dashboard/resources/designer' element={<ResourcesDesigner/>}/>
                </Routes>
            </Router>
        </BaseLayout>
    )
}

export default App
