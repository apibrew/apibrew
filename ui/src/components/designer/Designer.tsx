import React from 'react';
import {JointPaper} from "../joint/JointPaper";
import {JointGraph} from "../joint/JointGraph";
import {RectContainer, RectContainerProps} from "../joint/RectContainer";
import {Link} from "../joint/Link";
import Button from "@mui/material/Button";

// React component to render the diagram
export const Designer: React.FC = () => {
    const [rects, setRects] = React.useState<RectContainerProps[]>([])

    return <div>
        <Button onClick={() => {
            setRects([...rects, {
                name: 'item-' + rects.length,
                position: {x: rects.length * 300, y: 10},
                size: {width: 250, height: 300},
                attr: {rect: {fill: '#EEEEEE', stroke: 'black', 'stroke-width': 0}},
                children: <h6>Hello world</h6>,
            }])
        }} variant={'contained'}>Add</Button>
        <Button onClick={() => {
            setRects(rects.slice(0, rects.length - 1))
        }} variant={'contained'}>Delete</Button>
        <JointGraph>
            <JointPaper options={{
                width: '100%',
                height: '600px',
                gridSize: 10,
            }}>
                {rects.map(item => <RectContainer key={item.name} {...item} children={item.children}/>)}
                <Link source={'item-1'} target={'item-2'}/>
            </JointPaper>
        </JointGraph>
    </div>
}

