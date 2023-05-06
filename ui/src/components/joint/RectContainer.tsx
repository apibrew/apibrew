import {ReactContainerShape} from "./ReactContainerShape";
import * as joint from "jointjs";
import {dia} from "jointjs";

export interface RectContainerProps extends dia.Element.GenericAttributes<joint.shapes.basic.RectSelectors> {
    name?: string
    children: React.ReactNode
}

export function RectContainer(props: RectContainerProps) {
    return <ReactContainerShape name={props.name} element={new joint.shapes.basic.Rect(props)}>{props.children}</ReactContainerShape>
}