import React, {ReactNode, useContext, useEffect, useState} from "react";
import {Point} from "./point";

export interface MovingOps {
    moving: boolean
    movingIdx: number
    moveLoc?: Point
}

export const MovingContext = React.createContext<MovingOps>({} as MovingOps)

export interface MovableProps extends React.SVGProps<SVGGElement> {
    children: ReactNode
}

export function Movable(props: MovableProps) {
    const [moving, setMoving] = useState<boolean>(false)
    const [movingIdx, setMovingIdx] = useState<number>(0)

    const [moveLoc, setMoveLoc] = useState<Point | undefined>({x: 0, y: 0})
    const [beginLoc, setBeginLoc] = useState<Point>({x: 0, y: 0})

    return <g {...props}>
        <MovingContext.Provider value={{
            moving: moving,
            movingIdx: movingIdx,
            moveLoc: moveLoc
        }}>
            <g onPointerDown={(e) => {
                console.log('start moving')
                setMoving(true)
                setBeginLoc({x: e.clientX, y: e.clientY})
                setMovingIdx(movingIdx + 1)
            }}
               onPointerMove={(e) => {
                   if (moving) {
                       setMoveLoc({x: e.clientX - beginLoc.x, y: e.clientY - beginLoc.y})
                       setMovingIdx(movingIdx + 1)
                   }
               }}
               onPointerUp={(e) => {
                   console.log('stop moving')
                   setMoveLoc(undefined)
                   setMoving(false)
                   setMovingIdx(movingIdx + 1)
               }}
            >
                {props.children}
            </g>
        </MovingContext.Provider>
    </g>
}

export interface MovableComponentProps extends React.SVGProps<SVGGElement> {
    children: ReactNode
}

export function MovableComponent(props: MovableComponentProps) {
    const movingContext = useContext(MovingContext)
    const [hover, setHover] = useState<boolean>(false)
    const [loc, setLoc] = useState<Point>({x: 0, y: 0})
    const [matLoc, setMatLoc] = useState<Point>({x: 0, y: 0})

    useEffect(() => {
        if (hover) {
            if (movingContext.moveLoc) {
                setLoc({x: movingContext.moveLoc.x, y: movingContext.moveLoc.y})
            } else {
                setMatLoc({x: matLoc.x + loc.x, y: matLoc.y + loc.y})
                setLoc({x: 0, y: 0})
            }
        }

        console.log(hover, movingContext.movingIdx)

    }, [hover, movingContext.movingIdx])

    return <g {...props} transform={`translate(${loc.x + matLoc.x}, ${loc.y + matLoc.y})`}
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