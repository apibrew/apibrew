import {Arrow} from "./Arrow";
import React, {useContext, useEffect, useState} from "react";
import {Point} from "./point";
import {SvgContainerContext} from "./SvgContainer";
import {MovingContext} from "./Movable";

export interface LinkProps {
    sourceSelector: string
    targetSelector: string
}

export function Link(props: LinkProps) {
    const [startPoint, setStartPoint] = useState<Point>({x: 0, y: 0})
    const [endPoint, setEndPoint] = useState<Point>({x: 0, y: 0})
    const container = useContext(SvgContainerContext)
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

        setStartPoint({x: sourceRect.left - container.x + sourceRect.width, y: sourceRect.top - container.y + sourceRect.height / 2})
        setEndPoint({x: targetRect.left - container.x, y: targetRect.top - container.y + targetRect.height / 2})

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