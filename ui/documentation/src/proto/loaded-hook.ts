import {useState} from "react";
import {loadPromise$} from "./proto.ts";

export function useProtoImageLoading() {
    const [loaded, setLoaded] = useState(false)

    loadPromise$.then(() => {
        setLoaded(true)
    })

    return loaded
}