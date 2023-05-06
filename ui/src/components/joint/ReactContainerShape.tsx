import {useEffect, useRef} from "react";
import {useGraph, usePaper} from "./context";

export interface ComponentShapeProps {
    name?: string
    children: React.ReactNode
    element: joint.dia.Element
}

export function ReactContainerShape(props: ComponentShapeProps): JSX.Element {
    const graph = useGraph()
    const paper = usePaper()
    const element = props.element

    const containerRef = useRef<HTMLDivElement>(null);

    const updatePosition = () => {
        const container = containerRef.current!
        const cellView = paper.findViewByModel(element);
        const bbox = cellView.getBBox();
        container.style.top = `${bbox.y}px`;
        container.style.left = `${bbox.x}px`;
        container.style.width = `${bbox.width}px`;
        container.style.height = `${bbox.height}px`;
    }

    useEffect(() => {
        const container = containerRef.current!
        graph.addCell(element)

        if (props.name) {
            graph.attributes['cell_' + props.name] = element
        }

        const cellView = paper.findViewByModel(element);
        container.style.position = 'absolute';
        container.style.pointerEvents = 'none'
        container.style.userSelect = 'none'
        container.style.boxSizing = 'border-box'
        updatePosition()

        element.on('change:position', () => {
            updatePosition()
        });

        return () => {
            element.remove()
            container.remove()
        }
    }, [])

    return <div ref={containerRef}>{props.children}</div>
}