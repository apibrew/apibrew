/**
 * Copyright (c) ProductBoard, Inc.
 * All rights reserved.
 */
import React from 'react'

import styled from 'styled-components'

import {
    calculateDeltas,
    calculateCanvasDimensions,
    calculateControlPoints
} from './arrow-utils'
import { type Point } from './point'

const CONTROL_POINTS_RADIUS = 5
const STRAIGHT_LINE_BEFORE_ARROW_HEAD = 5

interface ArrowConfig {
    arrowColor?: string
    arrowHighlightedColor?: string
    controlPointsColor?: string
    boundingBoxColor?: string
    dotEndingBackground?: string
    dotEndingRadius?: number
    arrowHeadEndingSize?: number
    hoverableLineWidth?: number
    strokeWidth?: number
}

interface Props {
    startPoint: Point
    endPoint: Point
    isHighlighted?: boolean
    showDebugGuideLines?: boolean
    onMouseEnter?: (e: React.MouseEvent) => void
    onMouseLeave?: (e: React.MouseEvent) => void
    onClick?: (e: React.MouseEvent) => void
    config?: ArrowConfig
    tooltip?: string
}

interface TranslateProps {
    $xTranslate: number
    $yTranslate: number
}

type LineProps = {
    $isHighlighted: boolean
    $showDebugGuideLines?: boolean
    $boundingBoxColor?: string
} & TranslateProps

// @ts-expect-error ignored error
const Line = styled.g.attrs(({ $xTranslate, $yTranslate }: LineProps) => ({
    style: { transform: `translate(${$xTranslate}px, ${$yTranslate}px)` }
}))<LineProps>``

const CurvedLine = styled(Line)`
  border: ${({ $showDebugGuideLines, $boundingBoxColor = 'black' }) =>
        $showDebugGuideLines ? `dashed 1px ${$boundingBoxColor}` : '0'};
`

const RenderedLine = styled.path`
  transition: stroke 300ms;
`

const Endings = styled(Line)`
  pointer-events: none;
  z-index: ${({ $isHighlighted }) => ($isHighlighted ? 11 : 10)};
`

// @ts-expect-error ignored error
const ArrowHeadEnding = styled.path.attrs(({ $xTranslate, $yTranslate }: TranslateProps) => ({
    style: { transform: `translate(${$xTranslate}px, ${$yTranslate}px)` }
})
)<TranslateProps>`
  transition: stroke 300ms;
`

const DotEnding = styled.circle`
  transition: stroke 300ms;
`

const HoverableLine = styled.path`
  cursor: default;
`

const HoverableArrowHeadEnding = styled(ArrowHeadEnding)`
  cursor: default;
`

const HoverableDotEnding = styled.circle`
  cursor: default;
`

const ControlPoints = ({
    p1,
    p2,
    p3,
    p4,
    color
}: {
    p1: Point
    p2: Point
    p3: Point
    p4: Point
    color: string
}) => {
    return (
        <>
            <circle
                cx={p2.x}
                cy={p2.y}
                r={CONTROL_POINTS_RADIUS}
                strokeWidth="0"
                fill={color}
            />
            <circle
                cx={p3.x}
                cy={p3.y}
                r={CONTROL_POINTS_RADIUS}
                strokeWidth="0"
                fill={color}
            />
            <line
                strokeDasharray="1,3"
                stroke={color}
                x1={p1.x}
                y1={p1.y}
                x2={p2.x}
                y2={p2.y}
            />
            <line
                strokeDasharray="1,3"
                stroke={color}
                x1={p3.x}
                y1={p3.y}
                x2={p4.x}
                y2={p4.y}
            />
        </>
    )
}

export const Arrow = ({
    startPoint,
    endPoint,
    isHighlighted = false,
    showDebugGuideLines = false,
    onMouseEnter,
    onMouseLeave,
    onClick,
    config,
    tooltip
}: Props) => {
    const defaultConfig = {
        arrowColor: '#bcc4cc',
        arrowHighlightedColor: '#4da6ff',
        controlPointsColor: '#ff4747',
        boundingBoxColor: '#ffcccc',
        dotEndingBackground: '#fff',
        dotEndingRadius: 3,
        arrowHeadEndingSize: 9,
        hoverableLineWidth: 15,
        strokeWidth: 1
    }
    const currentConfig = {
        ...defaultConfig,
        ...config
    }

    const {
        arrowColor,
        arrowHighlightedColor,
        controlPointsColor,
        boundingBoxColor,
        arrowHeadEndingSize,
        strokeWidth,
        hoverableLineWidth,
        dotEndingBackground,
        dotEndingRadius
    } = currentConfig

    const arrowHeadOffset = arrowHeadEndingSize / 2
    const boundingBoxElementsBuffer =
        strokeWidth +
        arrowHeadEndingSize / 2 +
        dotEndingRadius +
        CONTROL_POINTS_RADIUS / 2

    const { absDx, absDy, dx, dy } = calculateDeltas(startPoint, endPoint)
    const { p1, p2, p3, p4, boundingBoxBuffer } = calculateControlPoints({
        boundingBoxElementsBuffer,
        dx,
        dy,
        absDx,
        absDy
    })

    const { canvasWidth, canvasHeight } = calculateCanvasDimensions({
        absDx,
        absDy,
        boundingBoxBuffer
    })

    const canvasXOffset =
        Math.min(startPoint.x, endPoint.x) - boundingBoxBuffer.horizontal
    const canvasYOffset =
        Math.min(startPoint.y, endPoint.y) - boundingBoxBuffer.vertical

    const curvedLinePath = `
    M ${p1.x} ${p1.y}
    C ${p2.x} ${p2.y},
    ${p3.x} ${p3.y},
    ${p4.x - STRAIGHT_LINE_BEFORE_ARROW_HEAD} ${p4.y}
    L ${p4.x} ${p4.y}`

    const getStrokeColor = () => {
        if (isHighlighted) return arrowHighlightedColor

        return arrowColor
    }

    const strokeColor = getStrokeColor()

    return (
        <>
            <CurvedLine
                width={canvasWidth}
                height={canvasHeight}
                $isHighlighted={isHighlighted}
                $showDebugGuideLines={showDebugGuideLines}
                $boundingBoxColor={boundingBoxColor}
                $xTranslate={canvasXOffset}
                $yTranslate={canvasYOffset}
            >
                <RenderedLine
                    d={curvedLinePath}
                    strokeWidth={strokeWidth}
                    stroke={getStrokeColor()}
                    fill="none"
                />
                <HoverableLine
                    d={curvedLinePath}
                    strokeWidth={hoverableLineWidth}
                    stroke="transparent"
                    pointerEvents="all"
                    fill="none"
                    onMouseEnter={onMouseEnter}
                    onMouseLeave={onMouseLeave}
                    onClick={onClick}
                >
                    {tooltip && <title>{tooltip}</title>}
                </HoverableLine>
                <HoverableArrowHeadEnding
                    d={`
            M ${(arrowHeadEndingSize / 5) * 2} 0
            L ${arrowHeadEndingSize} ${arrowHeadEndingSize / 2}
            L ${(arrowHeadEndingSize / 5) * 2} ${arrowHeadEndingSize}`}
                    fill="none"
                    stroke="transparent"
                    strokeWidth={hoverableLineWidth}
                    strokeLinecap="round"
                    pointerEvents="all"
                    $xTranslate={p4.x - arrowHeadOffset * 2}
                    $yTranslate={p4.y - arrowHeadOffset}
                    onMouseEnter={onMouseEnter}
                    onMouseLeave={onMouseLeave}
                    onClick={onClick}
                >
                    {tooltip && <title>{tooltip}</title>}
                </HoverableArrowHeadEnding>
                <HoverableDotEnding
                    cx={p1.x}
                    cy={p1.y}
                    r={dotEndingRadius}
                    stroke="transparent"
                    strokeWidth={hoverableLineWidth}
                    fill="transparent"
                >
                    {tooltip && <title>{tooltip}</title>}
                </HoverableDotEnding>
            </CurvedLine>
            <Endings
                width={canvasWidth}
                height={canvasHeight}
                $isHighlighted={isHighlighted}
                $xTranslate={canvasXOffset}
                $yTranslate={canvasYOffset}
            >
                <DotEnding
                    cx={p1.x}
                    cy={p1.y}
                    r={dotEndingRadius}
                    stroke={strokeColor}
                    strokeWidth={strokeWidth}
                    fill={dotEndingBackground}
                />
                <ArrowHeadEnding
                    d={`
            M ${(arrowHeadEndingSize / 5) * 2} 0
            L ${arrowHeadEndingSize} ${arrowHeadEndingSize / 2}
            L ${(arrowHeadEndingSize / 5) * 2} ${arrowHeadEndingSize}`}
                    fill="none"
                    stroke={strokeColor}
                    strokeWidth={strokeWidth}
                    strokeLinecap="round"
                    $xTranslate={p4.x - arrowHeadOffset * 2}
                    $yTranslate={p4.y - arrowHeadOffset}
                />
                {showDebugGuideLines && (
                    <ControlPoints
                        p1={p1}
                        p2={p2}
                        p3={p3}
                        p4={p4}
                        color={controlPointsColor}
                    />
                )}
            </Endings>
        </>
    )
}
