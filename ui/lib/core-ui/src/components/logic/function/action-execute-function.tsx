import {ActionComponent} from "../../../model/component-interfaces";
import {RecordService} from "../../../service";
import {Function} from "../../../model/extensions/function";
import {LayoutOptions} from "../../../context/layout-context.ts";
import {Button, Card, CardActions, CardContent, CardHeader} from "@mui/material";
import Box from "@mui/material/Box";
import {FunctionExecution} from "../../../model/extensions/function-execution.ts";
import {useState} from "react";

export interface ExecuteFunctionFormProps {
    functionRecord: Function
    cancel: () => void
    execute: (execution: FunctionExecution) => void
}

function ExecuteFunctionForm(props: ExecuteFunctionFormProps) {
    const [execution, setExecution] = useState<FunctionExecution>({
        function: props.functionRecord,
        id: '',
        version: 1,
    })

    return <Card>
        <CardHeader title={'Execute function ' + props.functionRecord.name}/>
        <CardContent>
            <input/>
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
        const functionRecord = await RecordService.get<Function>('extensions', 'Function', functionId)

        if (!functionRecord) {
            throw new Error(`Function ${functionId} not found`)
        }

        function runFunction() {

        }

        const modalConfig = layoutContext.showModal({
            content: <Box sx={{
                position: 'absolute' as 'absolute',
                top: '50%',
                left: '50%',
                transform: 'translate(-50%, -50%)',
                width: 800
            }}>
                <ExecuteFunctionForm
                    execute={(execution) => {
                        console.log(execution)
                        modalConfig.close()
                    }}
                    cancel={() => {
                        modalConfig.close()
                    }}
                    functionRecord={functionRecord} />
            </Box>,
        })

        return null;
    }
}
