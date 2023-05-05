import Box from "@mui/material/Box";
import 'jointjs/dist/joint.css'
import 'jquery'
import 'lodash'
import 'backbone'
import * as joint from 'jointjs'
import { Component, ReactNode, createRef, useEffect, useRef } from "react";
import Button from "@mui/material/Button";
import { random } from "lodash";


export interface DesignerProps {

}

function handleModelChange(changes: any) {
    console.log('handleModelChange', changes)
}

export class Designer extends Component {
    readonly graph: joint.dia.Graph<joint.dia.Graph.Attributes, joint.dia.ModelSetOptions> = new joint.dia.Graph({}, { cellNamespace: joint.shapes })
    myRef: any;

    constructor(props: {}) {
        super(props);
        this.myRef = createRef();
    }

    componentDidMount(): void {
        if (this.myRef.current.className.indexOf('joint-paper') !== -1) {
            return
        }

        console.log('triggered')

        new joint.dia.Paper({
            el: this.myRef.current,
            model: this.graph,
            width: 1000,
            height: 500,
            gridSize: 1,
        });

        var rect = new joint.shapes.standard.Rectangle();
        rect.position(100, 30);
        rect.resize(100, 40);
        rect.attr({
            body: {
                fill: 'blue'
            },
            label: {
                text: 'Hello',
                fill: 'white'
            }
        });
        rect.addTo(this.graph);

        var rect2 = rect.clone();
        rect2.translate(300, 0);
        rect2.attr('label/text', 'World!');
        rect2.addTo(this.graph);

        var link = new joint.shapes.standard.Link();
        link.source(rect);
        link.target(rect2);
        link.addTo(this.graph);


    }

    render(): ReactNode {
        return <Box>
            <Button variant="contained" onClick={() => {
                var rect = new joint.shapes.standard.Rectangle();
                rect.position(100, 300 * Math.random());
                rect.resize(100, 40);
                rect.attr({
                    body: {
                        fill: 'blue'
                    },
                    label: {
                        text: 'Hello',
                        fill: 'white'
                    }
                });
                rect.addTo(this.graph);
            }}>Add New</Button>
            <div ref={this.myRef}></div>
        </Box >
    }
}
