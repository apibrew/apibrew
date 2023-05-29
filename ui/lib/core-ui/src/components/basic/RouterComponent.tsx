import {Route, Routes} from "react-router-dom";
import React from "react";

export interface RouterComponentProps {
    children: React.ReactNode[]
    debug: boolean

    [index: number]: string //paths
}

export function RouterComponent(props: RouterComponentProps) {
    return <Routes>
        {props.children.map((item, index) => (
            <Route key={props[index]} path={props[index]} element={<>
                {props.debug && <h4>Route: {props[index]}</h4>}
                {item}
            </>}/>
        ))}
    </Routes>
}
