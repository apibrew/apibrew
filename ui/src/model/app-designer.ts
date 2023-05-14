export interface AppDesignerBoard {
    name: string;
    description?: string;
    resourceSelector: string[];
    resourceVisuals: {
        resource: string;
        allowRecordsOnBoard: boolean;
        location: {
            x: number;
            y: number;
        };
    }[];
}