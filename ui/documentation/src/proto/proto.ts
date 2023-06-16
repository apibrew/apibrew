import {Root} from "./image";
import Axios from 'axios'

let isLoaded = false

const waiters: any[] = []

const loaded = () => {
    isLoaded = true
    waiters.forEach(waiter => waiter())
}

export const loadPromise$ = new Promise((resolve: any) => {
    if (!isLoaded) {
        waiters.push(resolve)
    } else {
        resolve()
    }
})

let imageData: Root | undefined = undefined


Axios.get('/image.json.gz').then(({data}) => {
    console.log(data)
    loaded()
    imageData = data
})

export const GetProtoImage = () => {
    if (!imageData) {
        throw new Error('imageData not loaded')
    }

    return imageData as Root
}