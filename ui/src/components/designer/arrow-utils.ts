/**
 * Copyright (c) ProductBoard, Inc.
 * All rights reserved.
 */

import { Point } from "./point";

const MAX_Y_CONTROL_POINT_SHIFT = 50;
const MAX_X_CONTROL_POINT_SHIFT = 10;

// Y coordinates of our control points are moved in case of low delta Y to prevent lines overlapping.
// Sign flips the curve depending on delta Y.
// Movement is described according to following function: `y=a\left(0.9^{1.2^{\frac{x}{10}}}\right)`
//    1 +--------------------------------------------------------------------+
//      |       ************** +           +          +           +          |
//      |                     *******         1*(0.9**(1.2**(x/10))) ******* |
//  0.8 |-+                          *****                                 +-|
//      |                                ****                                |
//      |                                    ****                            |
//      |                                       ***                          |
//  0.6 |-+                                        ***                     +-|
//      |                                            ***                     |
//      |                                              ***                   |
//  0.4 |-+                                              ***               +-|
//      |                                                   **               |
//      |                                                     **             |
//  0.2 |-+                                                     **         +-|
//      |                                                         ***        |
//      |                                                            ***     |
//      |                                                               *****|
//    0 |-+                                                                +-|
//      |                                                                    |
//      |           +          +           +          +           +          |
// -0.2 +--------------------------------------------------------------------+
// -100         -50         0           50        100         150        200
export const calculateLowDyControlPointShift = (
    dx: number,
    dy: number,
    maxShift = MAX_Y_CONTROL_POINT_SHIFT
) => {
    if (dx > 0) return 0;
    const sign = dy < 0 ? -1 : 1;
    const value = Math.round(
        maxShift * Math.pow(0.9, Math.pow(1.2, Math.abs(dy) / 10))
    );

    // prevent negative zero
    if (value === 0) return 0;

    return sign * value;
};

export const calculateDeltas = (
    startPoint: Point,
    endPoint: Point
): {
    dx: number;
    dy: number;
    absDx: number;
    absDy: number;
} => {
    const dx = endPoint.x - startPoint.x;
    const dy = endPoint.y - startPoint.y;
    const absDx = Math.abs(dx);
    const absDy = Math.abs(dy);

    return { dx, dy, absDx, absDy };
};

export const calculateCanvasDimensions = ({
                                              absDx,
                                              absDy,
                                              boundingBoxBuffer,
                                          }: {
    absDx: number;
    absDy: number;
    boundingBoxBuffer: { vertical: number; horizontal: number };
}): {
    canvasWidth: number;
    canvasHeight: number;
} => {
    const canvasWidth = absDx + 2 * boundingBoxBuffer.horizontal;
    const canvasHeight = absDy + 2 * boundingBoxBuffer.vertical;

    return { canvasWidth, canvasHeight };
};

// Curve flexure should remain on the same area no matter of absolute deltas, so we have to slightly shift X coordinates of our control points. It was created empirically, it's not based on a clear formula.
export const calculateFixedLineInflectionConstant = (
    absDx: number,
    absDy: number
) => {
    const WEIGHT_X = 4;
    const WEIGHT_Y = 0.8;

    return Math.round(Math.sqrt(absDx) * WEIGHT_X + Math.sqrt(absDy) * WEIGHT_Y);
};

export const calculateControlPointsWithoutBoundingBox = ({
                                                             absDx,
                                                             absDy,
                                                             dx,
                                                             dy,
                                                         }: {
    absDx: number;
    absDy: number;
    dx: number;
    dy: number;
}): {
    p1: Point;
    p2: Point;
    p3: Point;
    p4: Point;
} => {
    let leftTopX = 0;
    let leftTopY = 0;
    let rightBottomX = absDx;
    let rightBottomY = absDy;
    if (dx < 0) [leftTopX, rightBottomX] = [rightBottomX, leftTopX];
    if (dy < 0) [leftTopY, rightBottomY] = [rightBottomY, leftTopY];

    const fixedLineInflectionConstant = calculateFixedLineInflectionConstant(
        absDx,
        absDy
    );
    const lowDyYShift = calculateLowDyControlPointShift(dx, dy);
    const lowDyXShift = Math.abs(
        calculateLowDyControlPointShift(dx, dy, MAX_X_CONTROL_POINT_SHIFT)
    );

    const p1 = {
        x: leftTopX,
        y: leftTopY,
    };
    const p2 = {
        x: leftTopX + fixedLineInflectionConstant + lowDyXShift,
        y: leftTopY + lowDyYShift,
    };
    const p3 = {
        x: rightBottomX - fixedLineInflectionConstant - lowDyXShift,
        y: rightBottomY - lowDyYShift,
    };
    const p4 = {
        x: rightBottomX,
        y: rightBottomY,
    };

    return { p1, p2, p3, p4 };
};
export const calculateControlPoints = ({
                                           boundingBoxElementsBuffer,
                                           absDx,
                                           absDy,
                                           dx,
                                           dy,
                                       }: {
    boundingBoxElementsBuffer: number;
    absDx: number;
    absDy: number;
    dx: number;
    dy: number;
}): {
    p1: Point;
    p2: Point;
    p3: Point;
    p4: Point;
    boundingBoxBuffer: {
        vertical: number;
        horizontal: number;
    };
} => {
    const { p1, p2, p3, p4 } = calculateControlPointsWithoutBoundingBox({
        absDx,
        absDy,
        dx,
        dy,
    });

    const topBorder = Math.min(p1.y, p2.y, p3.y, p4.y);
    const bottomBorder = Math.max(p1.y, p2.y, p3.y, p4.y);
    const leftBorder = Math.min(p1.x, p2.x, p3.x, p4.x);
    const rightBorder = Math.max(p1.x, p2.x, p3.x, p4.x);

    const verticalBuffer =
        (bottomBorder - topBorder - absDy) / 2 + boundingBoxElementsBuffer;
    const horizontalBuffer =
        (rightBorder - leftBorder - absDx) / 2 + boundingBoxElementsBuffer;

    const boundingBoxBuffer = {
        vertical: verticalBuffer,
        horizontal: horizontalBuffer,
    };

    return {
        p1: {
            x: p1.x + horizontalBuffer,
            y: p1.y + verticalBuffer,
        },
        p2: {
            x: p2.x + horizontalBuffer,
            y: p2.y + verticalBuffer,
        },
        p3: {
            x: p3.x + horizontalBuffer,
            y: p3.y + verticalBuffer,
        },
        p4: {
            x: p4.x + horizontalBuffer,
            y: p4.y + verticalBuffer,
        },
        boundingBoxBuffer,
    };
};