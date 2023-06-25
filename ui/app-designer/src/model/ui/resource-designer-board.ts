


// Sub Types

export interface ResourceVisual {
     resource: string;
     allowRecordsOnBoard: boolean;
     location?: {

     x: number;
     y: number;
}
;

}

// Resource Type
export interface ResourceDesignerBoard {
    id: string;
description?: string;
name: string;
resourceVisuals?: ResourceVisual[];
resourceSelector?: string[];
namespaceSelector?: string[];
version: number;

}
// Resource and Property Names
export const ResourceDesignerBoardName = "ResourceDesignerBoard";

export const ResourceDesignerBoardIdName = "Id";

export const ResourceDesignerBoardDescriptionName = "Description";

export const ResourceDesignerBoardNameName = "Name";

export const ResourceDesignerBoardResourceVisualsName = "ResourceVisuals";

export const ResourceDesignerBoardResourceSelectorName = "ResourceSelector";

export const ResourceDesignerBoardNamespaceSelectorName = "NamespaceSelector";

export const ResourceDesignerBoardVersionName = "Version";


