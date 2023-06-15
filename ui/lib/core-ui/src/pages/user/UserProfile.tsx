import {Button, Card, CardActions, CardContent, CardHeader} from "@mui/material";
import {Form} from "../../components/crud/Form.tsx";
import {useResourceByName} from "../../hooks/resource.ts";
import {useRecordBy, useRecordByName} from "../../hooks/record.ts";
import {Crud} from "../../model/ui/crud.ts";
import {Loading} from "../../components/basic/Loading.tsx";
import {RecordService, TokenService} from "@apibrew/core-lib";
import {useContext, useEffect, useState} from "react";
import {LayoutContext} from "../../context/layout-context.ts";
import {User} from "../../model";
import {useErrorHandler} from "../../hooks/error-handler.tsx";
import {AuthorizationService} from "@apibrew/core-lib";

export interface UserProfileProps {

}

export function UserProfile(props: UserProfileProps): JSX.Element {
    const resource = useResourceByName('user', 'system')
    const crudConfig = useRecordByName<Crud>('Crud', 'ui', 'ResourceCrud-system-user')
    const layoutContext = useContext(LayoutContext)
    const errorHandler = useErrorHandler()

    const uid = TokenService.getUid()

    const [record, setRecord] = useState<User>()

    useEffect(() => {
        RecordService.get<User>('system', 'user', uid).then((record) => {
            setRecord(record)
        }, errorHandler)
    }, [uid])

    if (!record || !resource || !crudConfig) {
        return <Loading/>
    }

    const formConfig = crudConfig.formConfig

    return <Card>
        <CardHeader title={'User: ' + record.username}/>
        <CardContent>
            {resource && formConfig && <Form resource={resource}
                                             record={record}
                                             setRecord={setRecord}
                                             formConfig={formConfig}></Form>}
        </CardContent>
        <CardActions>
            <Button onClick={() => {
                const updateFilteredRecord = AuthorizationService.filterRecordForUpdate(resource, record)

                RecordService.update('system', 'user', updateFilteredRecord).then(() => {
                    layoutContext.showAlert({
                        severity: 'success',
                        message: 'User updated'
                    });
                }, errorHandler)
            }}>Save</Button>
        </CardActions>
    </Card>
}