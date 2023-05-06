import React from 'react';
import {JointPaper} from "../joint/JointPaper";
import {JointGraph} from "../joint/JointGraph";
import {RectContainer} from "../joint/RectContainer";

// React component to render the diagram
export const Designer: React.FC = () => {
    return <div>
        <JointGraph>
            <JointPaper options={{
                width: '100%',
                height: '600px',
                gridSize: 10,
            }}>
                <RectContainer position={{x: 10, y: 10}}
                               size={{width: 250, height: 300}}
                               attrs={{rect: {fill: '#EEEEEE', stroke: 'black', 'stroke-width': 0}}}>
                    <h6>Hello world</h6>
                </RectContainer>
            </JointPaper>
        </JointGraph>
    </div>
}

