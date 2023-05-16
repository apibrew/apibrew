export interface AppDesignerBoard {
    id?: string
    name: string
    description?: string
    resourceSelector: string[]
    resourceVisuals: Array<{
        resource: string
        allowRecordsOnBoard: boolean
        location: {
            x: number
            y: number
        }
    }>
}
