import { type Resource } from "@apibrew/ui-lib"
import React from "react";

export interface ResourceAdvancedProps {
  resource: Resource
  onChange: (resource: Resource) => void
}

export function ResourceAdvancedForm(props: ResourceAdvancedProps): JSX.Element {
  return <React.Fragment>
    <div>Resource Advanced Form</div>
  </React.Fragment>
}
