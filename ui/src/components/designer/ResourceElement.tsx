import {Resource} from "../../model";
import React from "react";
import {ReactMarkupContainer} from "../joint/ReactMarkupContainer";
import {dia} from "jointjs";

export interface ResourceElementProps {
    resource: Resource
    position?: dia.Point;
}

export function ResourceElement(props: ResourceElementProps) {
    return <ReactMarkupContainer position={props.position}>
        <g>
            <rect stroke="#000" id="svg_1" height="153.99999" width="190" y="9" x="5.32258" fill="#fff"/>
            <text textAnchor="start" fontFamily="Noto Sans JP" fontSize="24" id="svg_2" y="38.5" x="107" strokeWidth="0"
                  stroke="#000" fill="#000000">gdffdfdf
            </text>
            <rect stroke="#000" id="svg_3" height="13.22581" width="8.06452" y="46.77419" x="1.29032" fill="#fff"/>
            <rect stroke="#000" id="svg_4" height="13.22581" width="8.06452" y="79.35484" x="1.29032" fill="#fff"/>
            <rect stroke="#000" id="svg_5" height="13.22581" width="8.06452" y="63.22581" x="1.29032" fill="#fff"/>
        </g>
    </ReactMarkupContainer>
}
