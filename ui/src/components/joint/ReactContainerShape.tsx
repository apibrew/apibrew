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
        console.log('updatePosition triggered')
        const container = containerRef.current!
        const bbox = element.getBBox();
        const paperArea = paper.getArea()
        const pw = paperArea.width
        const ph = paperArea.height

        const xd = pw * paper.scale().sx * (1 - paper.scale().sx) / 2
        const yd = ph * paper.scale().sy * (1 - paper.scale().sy) / 2
        const t = bbox.y * paper.scale().sy + yd
        const l = bbox.x * paper.scale().sx + xd

        console.log('pw, ph', pw, ph)
        console.log('xd, yd', xd, yd)
        console.log('l, t', l, t)
        console.log(paper.scale().sx)

        container.style.transform = `scale(${paper.scale().sx})`
        container.style.transformOrigin = 'top left'

        container.style.left = `${l}px`;
        container.style.top = `${t}px`;
        container.style.width = `${bbox.width}px`;
        container.style.height = `${bbox.height}px`;

        console.log('container.style.top', container.style.top)
    }

    useEffect(() => {
        const container = containerRef.current!
        graph.addCell(element)

        if (props.name) {
            graph.attributes['cell_' + props.name] = element
        }

        container.style.position = 'absolute';
        container.style.pointerEvents = 'none'
        container.style.userSelect = 'none'
        container.style.boxSizing = 'border-box'
        container.style.overflow = 'hidden'
        updatePosition()

        element.on('change:position', () => {
            updatePosition()
        });

        paper.on('change:size change:scale', () => {
            updatePosition()
        });

        return () => {
            element.remove()
            container.remove()
        }
    }, [])

    useEffect(() => {
        updatePosition()
    }, [paper && paper.scale().sx])

    return <div ref={containerRef}>{props.children}</div>
}