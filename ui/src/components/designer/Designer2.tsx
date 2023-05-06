import React, {useEffect, useRef, useState} from 'react';
import * as joint from 'jointjs';
import {ComponentShape} from "./ComponentShape";

// Custom JointJS paper class with support for rendering React components
class ReactPaper extends joint.dia.Paper {
    constructor(options: joint.dia.Paper.Options) {
        super(options);
    }

    renderReactComponent(
        element: joint.dia.Element,
    ) {
        const cellView = this.findViewByModel(element);
        const bbox = cellView.getBBox();
        const container = document.createElement('div');
        container.style.position = 'absolute';
        container.style.top = `${bbox.y}px`;
        container.style.left = `${bbox.x}px`;
        container.style.width = `${bbox.width}px`;
        container.style.height = `${bbox.height}px`;
        container.style.pointerEvents = 'none'
        container.style.userSelect = 'none'
        container.style.boxSizing = 'border-box'

        this.el.appendChild(container);

        element.on('change:position', (...args) => {
            const bbox = cellView.getBBox();
            container.style.top = `${bbox.y}px`;
            container.style.left = `${bbox.x}px`;
            container.style.width = `${bbox.width}px`;
            container.style.height = `${bbox.height}px`;
        });

        return container
    }
}

// React component to render inside the JointJS shape
const CustomComponent: React.FC<{ title: string; properties: Record<string, any> }> = (props) => {
    return (
        <div>
            <h3>{props.title}</h3>
            {Object.entries(props.properties).map(([key, value]) => (
                <p key={key}>
                    {key}: {value}
                </p>
            ))}
        </div>
    );
};

// React component to render the diagram
export const Designer: React.FC = () => {
    const paperRef = useRef<HTMLDivElement>(null);
    const [graph, setGraph] = useState(new joint.dia.Graph())
    const [paper, setPaper] = useState<ReactPaper>()

    useEffect(() => {
        setPaper(new ReactPaper({
            el: paperRef.current!,
            width: '100%',
            height: '600px',
            gridSize: 10,
            model: graph,
            interactive: true,
        }))

        // Render the React component inside the JointJS shape
    }, []);

    return <div style={{position: "relative", overflow: 'hidden' }} ref={paperRef}>
        {paper && <div>
            <ComponentShape paper={paper}
                            graph={graph}
                            element={new joint.shapes.basic.Rect({
                                position: {x: 100, y: 100},
                                size: {width: 250, height: 300},
                                attrs: {rect: {fill: '#EEEEEE', stroke: 'black', 'stroke-width': 0}},
                            })}>
                <h1>Hello world</h1>
            </ComponentShape>
        </div>}
    </div>;
};

