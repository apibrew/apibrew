
export interface Crud {
    id: string;
    resource: string;
    namespace: string;
    version: number;
    gridConfig: {
        defaultColumnConfig: {
            width: number; height: number;
        }; columns: {
            header: string; visible: boolean; width: number; sortable: boolean; filterable: boolean;
        }[];
    };
    formConfig: {
        fields: {
            name: string; label: string; type: string; required: boolean; readOnly: boolean; visible: boolean; defaultValue: string; options: {
                label: string; value: string; selected: boolean; disabled: boolean; visible: boolean; readOnly: boolean; defaultValue: string;
            }[];
        }[];
    };

}

export interface AppDesignerBoard {
    id: string;
    name: string;
    description: string;
    resourceSelector: string[];
    resourceVisuals: {
        resource: string; allowRecordsOnBoard: boolean; location: {
            x: number; y: number;
        };
    }[];
    version: number;

}

