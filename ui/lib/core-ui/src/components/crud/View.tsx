import {Box, Button} from "@mui/material"
import {PageLayout} from "../../layout/PageLayout"
import {Cancel} from "@mui/icons-material"
import {Resource} from "../../model"
import {useNavigate, useParams} from "react-router-dom"
import {Form} from "./Form"
import {Record, RecordService} from "../../service/record"
import React, {useEffect} from "react"
import {Crud} from "../../model/ui/crud.ts";

export interface ViewProps {
  resource: Resource
  crudConfig: Crud
}

export function View(props: ViewProps): JSX.Element {
  const navigate = useNavigate()
  const [record, setRecord] = React.useState<Record>()

  const params = useParams<{ id: string }>()

  useEffect(() => {
    RecordService.get<Record>(props.resource.namespace ?? 'default', props.resource.name, params.id!)
      .then((record) => {
        setRecord(record)
      })
  }, [params.id])

  if (!record) {
    return <>Loading...</>
  }

  return (
    <PageLayout pageTitle={props.resource.name} actions={<React.Fragment>
      <Box sx={{display: 'flex'}}>
        <Box m={0.5}>
          <Button variant={'outlined'}
                  color='primary'
                  size='small'
                  onClick={() => {
                    navigate('../')
                  }}
                  startIcon={<Cancel/>}>Cancel</Button>
        </Box>
      </Box>
    </React.Fragment>}>
      <React.Fragment>
        <Form resource={props.resource} readOnly={true} record={record} setRecord={setRecord}
              formConfig={props.crudConfig.formConfig}/>
      </React.Fragment>
    </PageLayout>
  )
}
