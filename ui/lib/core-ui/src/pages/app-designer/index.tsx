import React from "react";
import {Designer} from "../../components";
import {PageLayout} from "../../layout";

export function AppDesigner(): JSX.Element {
    return <PageLayout>
        <Designer name='Default'/>
    </PageLayout>
}
