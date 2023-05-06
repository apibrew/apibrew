import {useMemo} from "react";
import {dia} from "jointjs";
import {GraphContext} from "./context";

export interface JointGraphProps {
    children: JSX.Element
}

export function JointGraph(props: JointGraphProps) {
    const graph = useMemo(() => new dia.Graph(), [])

    return <GraphContext.Provider value={graph}>
        {props.children}
    </GraphContext.Provider>
}
