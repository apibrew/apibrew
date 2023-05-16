import { useEffect } from 'react'
import * as resources from '../../resources'
import { ResourceService } from '../../service/resource'
import { NamespaceService } from '../../service/namespace'
import { uiNamespace } from '../../data/ui-namespace'
import { RecordService } from '../../service/record'
import { DefaultAppDesignerBoard } from '../../data/default-app-designer-board'
import { AppDesignerBoardResource } from '../../resources/app-designer'

export function Migrate(): JSX.Element {
    useEffect(() => {
        const migrate = async () => {
            await NamespaceService.migrate(uiNamespace)

            for (const resource of resources.Resources) {
                await ResourceService.migrate(resource)
            }

            await RecordService.apply(AppDesignerBoardResource.namespace!, AppDesignerBoardResource.name, DefaultAppDesignerBoard)
        }

        migrate()
    }, [])

    return <></>
}
