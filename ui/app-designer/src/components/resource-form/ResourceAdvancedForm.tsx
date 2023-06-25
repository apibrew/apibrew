import {type Resource} from '../../model'
import React from "react";

export interface ResourceAdvancedProps {
  resource: Resource
  onChange: (resource: Resource) => void
}

export function ResourceAdvancedForm(props: ResourceAdvancedProps): JSX.Element {
  console.log(props)
  return <React.Fragment>
    <div>Resource Advanced Form</div>
  </React.Fragment>
}
