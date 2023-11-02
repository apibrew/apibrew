
export interface Author {
    id: string
    name: string
    description?: string
    book_count: number
    version: number
}

export const AuthorEntityInfo = {
    namespace: "default",
    resource: "Author",
    restPath: "author",
}


