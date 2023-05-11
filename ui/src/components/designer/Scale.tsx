import React, { type ReactNode } from 'react'

export interface ScaleProps extends React.SVGProps<SVGGElement> {
    level: number
    children: ReactNode
}

export const ScaleContext = React.createContext<number>(1)

export function Scale(props: ScaleProps) {
    return <g transform={`scale(${props.level})`} {...props}>
        <ScaleContext.Provider value={props.level}>
            {props.children}
        </ScaleContext.Provider>
    </g>
}
