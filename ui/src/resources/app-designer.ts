import { Resource } from "../model";

export const AppDesignerBoardResource: Resource = {
    name: 'AppDesignerBoard',
    namespace: 'ui',
    properties: [
        {
            name: 'name',
            type: 'STRING',
            unique: true,
            required: true,
            length: 255,
            description: 'The unique identifier of the board.'
        },
        {
            name: 'description',
            type: 'STRING',
            required: false,
            length: 255,
            description: 'The description of the board.'
        },
        {
            name: 'resourceSelector',
            type: 'LIST',
            subType: 'STRING',
    ],
    description: 'A board is a collection of widgets that are displayed together.',
    version: 1,
    virtual: false,
}
