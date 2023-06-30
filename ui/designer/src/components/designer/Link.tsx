import { Arrow } from './Arrow'
import { useContext, useEffect, useState } from 'react'
import { type Point } from './point'
import { MovingContext } from './Movable'
import { ScaleContext } from './Scale'

export interface LinkProps {
    sourceSelector: string
    targetSelector: string
}

export function Link(props: LinkProps) {
    const scale = useContext(ScaleContext)

    const [startPoint, setStartPoint] = useState<Point>({ x: 0, y: 0 })
    const [endPoint, setEndPoint] = useState<Point>({ x: 0, y: 0 })
    const movingContext = useContext(MovingContext)

    useEffect(() => {
        const sourceElem = document.querySelector(props.sourceSelector)
        const targetElem = document.querySelector(props.targetSelector)

        if (!sourceElem) {
            throw new Error('source element not found: ' + props.sourceSelector)
        }

        if (!targetElem) {
            throw new Error('target element not found: ' + props.sourceSelector)
        }

        const sourceRect = sourceElem.getBoundingClientRect()
        const targetRect = targetElem.getBoundingClientRect()
        const containerRect = document.querySelector('.designer-parent')!.getBoundingClientRect()

        setStartPoint({
            x: (sourceRect.left + sourceRect.width - containerRect.x) * (1 / scale),
            y: (sourceRect.top + sourceRect.height / 2 - containerRect.y) * (1 / scale)
        })
        setEndPoint({
            x: (targetRect.left - containerRect.x) * (1 / scale),
            y: (targetRect.top + targetRect.height / 2 - containerRect.y) * (1 / scale)
        })
    }, [props.sourceSelector, props.targetSelector, movingContext.movingIdx])

    return (
        <Arrow
            isHighlighted={true}
            startPoint={startPoint}
            endPoint={endPoint}
            showDebugGuideLines={false}
        />
    )
}
