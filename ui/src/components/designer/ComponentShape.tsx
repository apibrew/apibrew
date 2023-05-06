import {useEffect, useRef, useState} from "react";

export interface ComponentShapeProps {
    paper: joint.dia.Paper
    graph: joint.dia.Graph
    children: React.ReactNode
    element: joint.dia.Element
}

export function ComponentShape(props: ComponentShapeProps): JSX.Element {
    const containerRef = useRef<HTMLDivElement>(null);

    const {paper, graph, children, element} = props

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

        const cellView = paper.findViewByModel(element);
        container.style.position = 'absolute';
        container.style.pointerEvents = 'none'
        container.style.userSelect = 'none'
        container.style.boxSizing = 'border-box'
        updatePosition()

        element.on('change:position', () => {
            updatePosition()
        });
    }, [])

    return <div ref={containerRef}>{props.children}</div>
}