import React from "react";
import {Designer} from "../../components";
import {PageLayout} from "../../layout";

export function AppDesigner(): JSX.Element {
    return <PageLayout pageTitle={'App Designer'}>
        <Designer name='Default'/>
    </PageLayout>
}
