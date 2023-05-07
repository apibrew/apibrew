import {useEffect, useState} from "react";
import {useGraph} from "./context";
import {createPortal} from "react-dom";
import * as joint from "jointjs";
import {dia, shapes} from "jointjs";

export interface ReactMarkupContainerProps extends dia.Element.GenericAttributes<shapes.basic.RectSelectors> {
    children: React.ReactNode
}

export function ReactMarkupContainer(props: ReactMarkupContainerProps): JSX.Element {
    const graph = useGraph()
    const [container, setContainer] = useState<HTMLElement>()

    const elemId = 'react-markup-' + Math.random().toString(36)

    props = {
        id: elemId,
        type: 'react-markup-container',
        ...props,
    }

    useEffect(() => {
        const elem = new joint.shapes.basic.Rect(props)

        elem.markup = `<g id="${elemId}"></g>`

        graph.addCell(elem)

        setContainer(document.getElementById(elemId)!)

        return () => {
            elem.remove()
        }
    }, [])

    return <>{container && createPortal(props.children, container)}</>
}