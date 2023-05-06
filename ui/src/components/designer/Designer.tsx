import Box from "@mui/material/Box";
import 'jointjs/dist/joint.css'
import 'jquery'
import 'lodash'
import 'backbone'
import * as joint from 'jointjs'
import {Component, createRef, ReactNode} from "react";
import Button from "@mui/material/Button";
import {TheClass2} from "./def1";


export interface DesignerProps {

}

function handleModelChange(changes: any) {
    console.log('handleModelChange', changes)
}

export class Designer extends Component {
    readonly graph: joint.dia.Graph<joint.dia.Graph.Attributes, joint.dia.ModelSetOptions> = new joint.dia.Graph({}, {cellNamespace: joint.shapes})
    myRef: any;
    private shape2?: TheClass2;

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

        var shape1 = new joint.shapes.standard.Rectangle()

        shape1.resize(200, 240)
        shape1.position(200, 100)

        shape1.addTo(this.graph)

        var shape2 = new TheClass2({
            name: 'Class1',
            properties: [
                {
                    name: 'attr-1',
                    type: 'STRING',
                    required: true,
                },
                {
                    name: 'attr-1',
                    type: 'STRING',
                    required: true,
                },
                {
                    name: 'attr-1',
                    type: 'STRING',
                    required: true,
                },
                {
                    name: 'attr-1',
                    type: 'STRING',
                    required: true,
                }
            ]
        })

        shape2.resize(300, 440)
        shape2.position(50, 20)

        shape2.abc = 'tttt'

        shape2.addTo(this.graph)

        this.shape2 = shape2
        this.shape2.redrawResource()
    }

    render(): ReactNode {
        return <Box>
            <Button variant="contained" onClick={() => {
                this.shape2!.resource.properties = [(this.shape2!.resource.properties![0])];
                this.shape2!.redrawResource()
            }}>Add New</Button>
            <div ref={this.myRef}></div>
        </Box>
    }
}
