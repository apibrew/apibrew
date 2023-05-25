import {Function} from "./function";


// Sub Types

// Resource Type
export interface Schedule {
    id: string;
    name: string;
    schedule: string;
    function: Function;
    version: number;

}

// Resource and Property Names
export const ScheduleName = "Schedule";

export const ScheduleIdName = "Id";

export const ScheduleNameName = "Name";

export const ScheduleScheduleName = "Schedule";

export const ScheduleFunctionName = "Function";

export const ScheduleVersionName = "Version";


