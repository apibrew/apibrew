import Box from "@mui/material/Box";
import {useResource} from "../../context/resource.ts";
import {useValue} from "../../context/value.ts";
import {FormElement} from "./FormElement.tsx";
import {isSpecialProperty} from "../../util/property.ts";
import {not} from "../../util/lambda.ts";

export interface ResourceAutoFormProps {
    readOnly?: boolean
}

export function ResourceAutoForm(props: ResourceAutoFormProps) {
    const resource = useResource()
    const valueProvider = useValue()

    return <Box sx={{display: 'flex'}} flexDirection='column'>
       <Box>
           {resource.properties.filter(not(isSpecialProperty)).map((property) => {
               return <Box key={property.name}>
                   <FormElement resource={resource} property={property} readOnly={props.readOnly}
                                value={valueProvider.value[property.name]} setValue={updatedValue => {
                       valueProvider.onChange({
                           ...valueProvider.value,
                           [property.name]: updatedValue
                       })
                   }}/>
               </Box>
           })}
       </Box>
        <hr/>
        <Box>
            <h3>System Properties</h3>
            {resource.properties.filter(isSpecialProperty).map((property) => {
                return <Box key={property.name}>
                    <FormElement resource={resource} property={property} readOnly={true}
                                 value={valueProvider.value[property.name]} setValue={updatedValue => {
                        valueProvider.onChange({
                            ...valueProvider.value,
                            [property.name]: updatedValue
                        })
                    }}/>
                </Box>
            })}
        </Box>
    </Box>
}
