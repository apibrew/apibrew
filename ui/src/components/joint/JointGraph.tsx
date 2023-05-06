import {useMemo, useState} from "react";
import {dia} from "jointjs";
import {GraphContext} from "./context";

export interface JointGraphProps {
    children: JSX.Element
}

export function JointGraph(props: JointGraphProps) {
    const graph = useState(new dia.Graph())[0]

    return <GraphContext.Provider value={graph}>
        {props.children}
    </GraphContext.Provider>
}
