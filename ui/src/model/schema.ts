
export interface Crud {
    id: string;
    resource: string;
    namespace: string;
    gridConfig: {
    
};
    formConfig: {
    
};
    version: number;
    
}

export interface AppDesignerBoard {
    id: string;
    name: string;
    description: string;
    resourceSelector: string[];
    resourceVisuals: {
    resource: string;allowRecordsOnBoard: boolean;location: {
    x: number;y: number;
};
}[];
    version: number;
    
}

