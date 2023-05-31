import {ActionComponent} from "../../../model/component-interfaces";
import {RecordService} from "../../../service";
import {Function as Function$} from "../../../model/extensions/function";
import {LayoutOptions} from "../../../context/layout-context.ts";
import {Button, Card, CardActions, CardContent, CardHeader} from "@mui/material";
import Box from "@mui/material/Box";
import {FunctionExecution} from "../../../model/extensions/function-execution.ts";
import {useState} from "react";
import {Form} from "../../crud/Form.tsx";
import {useResource} from "../../../hooks/resource.ts";
import {FormConfig} from "../../../model/ui/crud.ts";

export interface ExecuteFunctionFormProps {
    functionRecord: Function$
    cancel: () => void
    execute: (execution: FunctionExecution) => void
}

export function ExecuteFunctionForm(props: ExecuteFunctionFormProps) {
    const defaultInput = {}
    const args = props.functionRecord.args ?? []

    args.forEach(arg => {
        defaultInput[arg.name] = ''
    })

    const [execution, setExecution] = useState<FunctionExecution>({
        function: props.functionRecord,
        id: '',
        version: 1,
        input: defaultInput
    })

    const resource = useResource('FunctionExecution', 'extensions')

    const formConfig: FormConfig = {
        children: [{
            kind: 'property',
            propertyPath: 'input',
            children: args.map(arg => {
                return {
                    kind: 'property',
                    propertyPath: arg.name,
                    children: [
                        {
                            kind: 'input',
                            title: arg.label ?? arg.name,
                        }
                    ]
                }
            })
        }]
    }

    return <Card>
        <CardHeader title={'Execute function ' + props.functionRecord.name}/>
        <CardContent>
            {resource && formConfig && <Form resource={resource}
                                             record={execution}
                                             setRecord={setExecution}
                                             formConfig={formConfig}></Form>}
        </CardContent>
        <CardActions>
            <Button onClick={() => {
                props.cancel()
            }}>Cancel</Button>
            <Button onClick={() => {
                props.execute(execution)
            }}>Run</Button>
        </CardActions>
    </Card>
}

export class ActionExecuteFunction implements ActionComponent<any> {
    async execute(functionId: string, layoutContext: LayoutOptions): Promise<any> {
        const functionRecord = await RecordService.get<Function$>('extensions', 'Function', functionId)

        if (!functionRecord) {
            throw new Error(`Function ${functionId} not found`)
        }

        async function runFunction(execution: FunctionExecution) {
            await RecordService.create<FunctionExecution>('extensions', 'FunctionExecution', execution).then(resp => {
                if (resp.status == 'error') {
                    layoutContext.showAlert({
                        severity: 'error',
                        message: JSON.stringify(resp.error)
                    })
                }
                if (resp.status == 'success') {
                    layoutContext.showAlert({
                        severity: 'success',
                        message: JSON.stringify(resp.output)
                    })
                } else {
                    layoutContext.showAlert({
                        severity: 'error',
                        message: `Function is not executed by engine`
                    })
                }
            }, err => {
                console.error(err)
                layoutContext.showAlert({
                    severity: 'error',
                    message: `Function execution failed`
                })
            })
        }

        const modalConfig = layoutContext.showModal({
            content: <Box sx={{
                position: 'absolute',
                top: '50%',
                left: '50%',
                transform: 'translate(-50%, -50%)',
                width: 800
            }}>
                <ExecuteFunctionForm
                    execute={(execution) => {
                        runFunction(execution)
                        modalConfig.close()
                    }}
                    cancel={() => {
                        modalConfig.close()
                    }}
                    functionRecord={functionRecord}/>
            </Box>,
        })

        return null;
    }
}
