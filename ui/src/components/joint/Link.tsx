import {useEffect} from "react";
import {useGraph} from "./context";
import {shapes} from "jointjs";

export interface LinkProps {
    source: string
    target: string
}

export function Link(props: LinkProps) {
    const graph = useGraph()

    useEffect(() => {
        var link = new shapes.standard.Link();
        link.source(graph.attributes['cell_' + props.source]);
        link.target(graph.attributes['cell_' + props.target]);

        console.log('link', link)

        link.addTo(graph);
    }, [props.source, props.target])
    return <></>
}