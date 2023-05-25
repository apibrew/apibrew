import React from 'react'
import ReactDOM from 'react-dom/client'


import {PageLayout} from '.'
import {DynamicComponent} from "./components/dynamic-component/DynamicComponent.tsx";

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <PageLayout>
            <DynamicComponent moduleName='Skeleton' modulePackage='Test' componentName='Component1'/>
        </PageLayout>
    </React.StrictMode>,
)
