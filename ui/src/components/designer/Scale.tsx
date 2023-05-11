import {ReactNode} from "react";

export interface ScaleProps extends React.SVGProps<SVGGElement>{
    level: number
    children: ReactNode
}

export function Scale(props: ScaleProps) {
    return <g transform={`scale(${props.level})`} {...props}>
        {props.children}
    </g>
}

