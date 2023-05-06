import {createContext, useContext} from "react";

export const GraphContext = createContext<joint.dia.Graph | undefined>(undefined)

export const PaperContext = createContext<joint.dia.Paper | undefined>(undefined)

export function useGraph() {
    const graph = useContext(GraphContext)

    if (!graph) {
        throw new Error('not in a graph context')
    }

    return graph
}

export function usePaper() {
    const paper = useContext(PaperContext)

    if (!paper) {
        throw new Error('not in a paper context')
    }

    return paper
}