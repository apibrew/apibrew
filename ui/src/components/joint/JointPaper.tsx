import {ReactNode, useEffect, useRef, useState} from "react";
import * as joint from "jointjs";
import {PaperContext, useGraph} from "./context";

export interface JointPaperProps {
    options: joint.dia.Paper.Options
    children: ReactNode
    preventCollision?: boolean
}

export function JointPaper(props: JointPaperProps) {
    const graph = useGraph()

    const paperRef = useRef<HTMLDivElement>(null);
    const [paper, setPaper] = useState<joint.dia.Paper>()

    useEffect(() => {
        setPaper(new joint.dia.Paper({
            el: paperRef.current!,
            model: graph,
            interactive: true,
            ...props.options,
        }))

        // Render the React component inside the JointJS shape
    }, [graph]);

    useEffect(() => {
        if (paper && props.preventCollision) {
            paper.on({
                'element:pointerdown': (elementView, evt) => {
                    // store the position before the user starts dragging
                    evt.data = { startPosition: elementView.model.position() };
                },
                'element:pointerup': (elementView, evt) => {
                    const { model: element } = elementView;
                    const { model: graph } = paper;

                    if (element.getBBox().x < 0) {
                        // the element is dragged outside of the paper
                        // move to left corner
                        const { y } = evt.data.startPosition;
                        element.position(0, y);
                    }

                    if (element.getBBox().y < 0) {
                        // the element is dragged outside of the paper
                        // move to left corner
                        const { x } = evt.data.startPosition;
                        element.position(x, 0);
                    }

                    const elementsUnder = graph.findModelsInArea(element.getBBox()).filter(el => el !== element);
                    if (elementsUnder.length > 0) {
                        // an overlap found, revert the position
                        const { x, y } = evt.data.startPosition;
                        element.position(x, y);
                    }
                }
            });
        }
    }, [paper, props.preventCollision])

    return <div style={{position: "relative", overflow: 'hidden'}} ref={paperRef}>
        {paper && <PaperContext.Provider value={paper}>
            {props.children}
        </PaperContext.Provider>}
    </div>
}