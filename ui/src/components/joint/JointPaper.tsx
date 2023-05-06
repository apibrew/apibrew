import {ReactNode, useEffect, useRef, useState} from "react";
import * as joint from "jointjs";
import {PaperContext, useGraph} from "./context";

export interface JointPaperProps {
    options: joint.dia.Paper.Options
    children: ReactNode
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

    return <div style={{position: "relative", overflow: 'hidden'}} ref={paperRef}>
        {paper && <PaperContext.Provider value={paper}>
            {props.children}
        </PaperContext.Provider>}
    </div>
}