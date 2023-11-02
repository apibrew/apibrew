import {Author} from './author';

export interface Book {
    id: string
    name: string
    description?: string
    author?: Author
    version: number
}

export const BookEntityInfo = {
    namespace: "default",
    resource: "Book",
    restPath: "book",
}


