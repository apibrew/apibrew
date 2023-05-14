import App from "../App";
import { Namespace } from "../model";
import { AppDesignerBoard } from "../model/app-designer";

export const DefaultAppDesignerBoard: AppDesignerBoard = {
    name: 'Default',
    description: 'The default board.',
    resourceSelector: ['*'],
    resourceVisuals: [],
}