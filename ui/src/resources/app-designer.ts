import { type Resource } from '../model'

export const AppDesignerBoardResource: Resource = {
    name: 'AppDesignerBoard',
    namespace: 'ui',
    properties: [
        {
            name: 'name',
            type: 'STRING',
            immutable: true,
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
            subProperty: {
                name: '',
                type: 'STRING'
            }
        },
        {
            name: 'resourceVisuals',
            type: 'LIST',
            subProperty: {
                name: '',
                type: 'STRUCT',
                properties: [
                    {
                        name: 'resource',
                        type: 'STRING',
                        required: true,
                        length: 255
                    },
                    {
                        name: 'allowRecordsOnBoard',
                        type: 'BOOL',
                        required: true
                    },
                    {
                        name: 'location',
                        type: 'STRUCT',
                        properties: [
                            {
                                name: 'x',
                                type: 'FLOAT32',
                                required: true
                            },
                            {
                                name: 'y',
                                type: 'FLOAT32',
                                required: true
                            }
                        ]
                    }
                ]
            }
        }
    ],
    description: 'A board is a collection of widgets that are displayed together.',
    version: 1,
    virtual: false
}
