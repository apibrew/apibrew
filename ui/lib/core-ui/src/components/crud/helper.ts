import {Crud as CrudModel, FormItem, CrudName, GridColumnConfig} from "../../model/ui/crud";
import {RecordService} from "../../service/record";
import {Resource} from "../../model";
import {isSimpleProperty, isSpecialProperty} from "../../util/property";
import {not} from "../../util/lambda";

export async function resetCrudForm(resource: Resource): Promise<CrudModel> {
    const name = `ResourceCrud-${resource.namespace}-${resource.name}`

    const gridColumns: GridColumnConfig[] = [
        {
            name: 'id',
            title: 'ID',
            width: 300,
        },
        ...(resource.properties.filter(not(isSpecialProperty)).filter(isSimpleProperty).map(item => {
            return {
                name: item.name,
                title: item.title,
                sortable: true,
                filterable: true,
            } as GridColumnConfig
        }))
    ]

    const newCrudConfig: CrudModel = {
        id: '',
        version: 1,
        name: name,
        resource: resource.name,
        namespace: resource.namespace ?? 'default',
        gridConfig: {
            columns: gridColumns,
            actions: [],
            disableDefaultActions: false
        },
        formConfig: {
            children: [
                {
                    kind: 'section',
                    title: 'Details',
                    children: resource.properties.filter(not(isSpecialProperty)).map(item => {
                        return {
                            kind: 'property',
                            propertyPath: item.name,
                            children: [
                                {
                                    kind: 'input',
                                }
                            ]
                        } as FormItem
                    })
                } as FormItem,
                {
                    kind: 'section',
                    title: 'System Properties',
                    children: resource.properties.filter(isSpecialProperty).map(item => {
                        return {
                            kind: 'property',
                            propertyPath: item.name,
                            children: [
                                {
                                    kind: 'input',
                                    params: {
                                        readOnly: true
                                    },
                                }
                            ]
                        } as FormItem
                    })
                } as FormItem
            ]
        }
    }

    await RecordService.apply<CrudModel>('ui', CrudName, newCrudConfig)

    return newCrudConfig
}
