import {BACKEND_URL} from '../config'
import {PageLayout} from '../layout/PageLayout'
import Button from '@mui/material/Button'
import {PlusOneOutlined} from '@mui/icons-material'
import React, {JSX} from 'react'

export function Test(): JSX.Element {
    const componentCode = `
import Button from '@mui/material/Button'

export default function Component1() {
  return (
    <div>
    test
    <Button onClick={() => {
        alert('test')
    }} >Button1</Button>
    </div>
  );
}
`
    const DynamicComponent = React.lazy(() => {
        const imported = import("../components/dynamic/DynamicComponent");

        return imported.then((module) => ({ default: module.DynamicComponent }));
    });
    return (
        <PageLayout pageTitle={'Test Page'} actions={<>
            <Button variant={'contained'} color='success' startIcon={<PlusOneOutlined/>}>New Item</Button>
        </>}>
            <>
                Hello World {BACKEND_URL}
                <br/>
                <DynamicComponent componentCode={componentCode}/>
            </>
        </PageLayout>
    )
}
