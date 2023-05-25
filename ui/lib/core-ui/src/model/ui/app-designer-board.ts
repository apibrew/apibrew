// Sub Types

export interface ResourceVisual {
    resource: string;
    allowRecordsOnBoard: boolean;
    location: {

        x: number;
        y: number;
    };

}

// Resource Type
export interface AppDesignerBoard {
    id: string;
    description: string;
    name: string;
    namespaceSelector: string[];
    version: number;
    resourceVisuals: ResourceVisual[];
    resourceSelector: string[];

}

// Resource and Property Names
export const AppDesignerBoardName = "AppDesignerBoard";

export const AppDesignerBoardIdName = "Id";

export const AppDesignerBoardDescriptionName = "Description";

export const AppDesignerBoardNameName = "Name";

export const AppDesignerBoardNamespaceSelectorName = "NamespaceSelector";

export const AppDesignerBoardVersionName = "Version";

export const AppDesignerBoardResourceVisualsName = "ResourceVisuals";

export const AppDesignerBoardResourceSelectorName = "ResourceSelector";


