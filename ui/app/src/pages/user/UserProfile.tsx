import { Button, Card, CardActions, CardContent, CardHeader } from "@mui/material";
import { Form } from "../../components/crud/Form.tsx";
import { useResourceByName } from "../../hooks/resource.ts";
import { useRecordByName } from "../../hooks/record.ts";
import { Crud } from "../../model/ui/crud.ts";
import { Loading } from "../../components/basic/Loading.tsx";
import { RecordService, TokenService } from "@apibrew/ui-lib";
import { useContext, useEffect, useState } from "react";
import { useErrorHandler } from "../../hooks/error-handler.tsx";
import { AuthorizationService, LayoutContext } from "@apibrew/ui-lib";
import { User, Record } from "@apibrew/client";

export interface UserProfileProps {

}

export function UserProfile(): JSX.Element {
    const resource = useResourceByName('user', 'system')
    const crudConfig = useRecordByName<Crud>('Crud', 'ui', 'ResourceCrud-system-user')
    const layoutContext = useContext(LayoutContext)
    const errorHandler = useErrorHandler()

    const uid = TokenService.getUid()

    const [record, setRecord] = useState<Record<User>>()

    useEffect(() => {
        RecordService.get<User>('system', 'user', uid).then((record) => {
            setRecord(record)
        }, errorHandler)
    }, [uid])

    if (!record || !resource || !crudConfig) {
        return <Loading />
    }

    const formConfig = crudConfig.formConfig

    return <Card>
        <CardHeader title={'User: ' + record.username} />
        <CardContent>
            {resource && formConfig && <Form resource={resource}
                record={record}
                setRecord={record => {
                    setRecord(record as Record<User>)
                }}
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