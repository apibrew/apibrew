import {Crud as CrudModel, CrudFormItem, CrudName} from "../../model/schema";
import {RecordService} from "../../service/record";
import {Resource} from "../../model";
import {isSpecialProperty} from "../../util/property";
import {not} from "../../util/lambda";

export async function resetCrudForm(resource: Resource): Promise<CrudModel> {
    const name = `ResourceCrud-${resource.namespace}-${resource.name}`

    const newCrudConfig: CrudModel = {
        id: '',
        version: 1,
        name: name,
        resource: resource.name,
        namespace: resource.namespace ?? 'default',
        gridConfig: {},
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
                        } as CrudFormItem
                    })
                } as CrudFormItem,
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
                        } as CrudFormItem
                    })
                } as CrudFormItem
            ]
        }
    }

    await RecordService.apply<CrudModel>('ui', CrudName, newCrudConfig)

    return newCrudConfig
}