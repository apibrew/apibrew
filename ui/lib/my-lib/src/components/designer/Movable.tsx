import React, { type ReactNode, useContext, useEffect, useState } from 'react'
import { type Point } from './point'
import { ScaleContext } from './Scale'

export interface MovingOps {
    moving: boolean
    movingIdx: number
    moveLoc?: Point
}

export const MovingContext = React.createContext<MovingOps>({
    moving: false,
    movingIdx: 0
})

export interface MovableProps extends React.SVGProps<SVGGElement> {
    children: ReactNode
}

export function Movable(props: MovableProps) {
    const scale = useContext(ScaleContext)
    const [moving, setMoving] = useState<boolean>(false)
    const [movingIdx, setMovingIdx] = useState<number>(0)

    const [moveLoc, setMoveLoc] = useState<Point | undefined>({ x: 0, y: 0 })
    const [beginLoc, setBeginLoc] = useState<Point>({ x: 0, y: 0 })

    return <MovingContext.Provider value={{
        moving,
        movingIdx,
        moveLoc
    }}>
        <g onPointerDown={(e) => {
            setMoving(true)
            setBeginLoc({ x: e.clientX, y: e.clientY })
            setMovingIdx(movingIdx + 1)
        }}
        onPointerMove={(e) => {
            if (moving) {
                setMoveLoc({
                    x: (e.clientX - beginLoc.x) * (1 / scale),
                    y: (e.clientY - beginLoc.y) * (1 / scale)
                })
                setMovingIdx(movingIdx + 1)
            }
        }}
        onPointerUp={(e) => {
            if (moving) {
                setMoveLoc(undefined)
                setMoving(false)
                setMovingIdx(movingIdx + 1)
            }
        }}
        onMouseLeave={(e) => {
            if (moving) {
                setMoveLoc(undefined)
                setMoving(false)
                setMovingIdx(movingIdx + 1)
            }
        }}
        >
            {props.children}
        </g>
    </MovingContext.Provider>
}

export interface MovableComponentProps extends React.SVGProps<SVGGElement> {
    children: ReactNode
    location: Point
    updateLocation: (loc: Point) => void
}

export function MovableComponent(props: MovableComponentProps) {
    const movingContext = useContext(MovingContext)
    const [hover, setHover] = useState<boolean>(false)
    const [loc, setLoc] = useState<Point>({ x: 0, y: 0 })

    useEffect(() => {
        if (hover) {
            if (movingContext.moveLoc) {
                setLoc({ x: movingContext.moveLoc.x, y: movingContext.moveLoc.y })
            } else {
                props.updateLocation({ x: props.location.x + loc.x, y: props.location.y + loc.y })
                setLoc({ x: 0, y: 0 })
            }
        }
    }, [hover, movingContext.movingIdx])

    return <g {...props} transform={`translate(${props.location.x + loc.x},${props.location.y + loc.y})`}
        className={'movable-component'}
        onMouseEnter={(e) => {
            setHover(true)
        }}
        onMouseLeave={(e) => {
            setHover(false)
        }}
    >
        {props.children}
    </g>
}
