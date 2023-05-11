import React, {ReactNode, useEffect, useRef, useState} from "react";

export interface SvgContainerOps {
    x: number
    y: number
}

export const SvgContainerContext = React.createContext<SvgContainerOps>({} as SvgContainerOps)

export interface SvgContainerProps extends React.SVGProps<SVGGElement> {
    children: ReactNode
}

export function SvgContainer(props: SvgContainerProps) {
    const ref = useRef<SVGGElement>(null)

    const [container, setContainer] = useState<SvgContainerOps>()

    useEffect(() => {
        if (!ref.current) {
            return
        }

        const boundingClientRect = ref.current.getBoundingClientRect()

        setContainer({
            x: boundingClientRect.left,
            y: boundingClientRect.top
        })
    }, [ref.current])

    return <g ref={ref} {...props}>
        {container && <SvgContainerContext.Provider value={container}>
            {props.children}
        </SvgContainerContext.Provider>}
    </g>

}
