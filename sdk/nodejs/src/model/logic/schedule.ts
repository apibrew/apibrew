

import { Function } from "./function";


export const ScheduleResource = {
    resource: "Schedule",
    namespace: "logic",
};

// Sub Types

// Resource Type
export interface Schedule {
    id: string;
name: string;
schedule: string;
function: Function;
annotations?: object;
createdBy?: string;
updatedBy?: string;
createdOn?: string;
updatedOn?: string;
version: number;

}
// Resource and Property Names
export const ScheduleName = "Schedule";

export const ScheduleIdName = "Id";

export const ScheduleNameName = "Name";

export const ScheduleScheduleName = "Schedule";

export const ScheduleFunctionName = "Function";

export const ScheduleAnnotationsName = "Annotations";

export const ScheduleCreatedByName = "CreatedBy";

export const ScheduleUpdatedByName = "UpdatedBy";

export const ScheduleCreatedOnName = "CreatedOn";

export const ScheduleUpdatedOnName = "UpdatedOn";

export const ScheduleVersionName = "Version";


