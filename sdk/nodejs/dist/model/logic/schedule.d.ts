import { Function } from "./function";
export declare const ScheduleResource: {
    resource: string;
    namespace: string;
};
export interface Schedule {
    id: string;
    name: string;
    schedule: string;
    function: Function;
    version: number;
}
export declare const ScheduleName = "Schedule";
export declare const ScheduleIdName = "Id";
export declare const ScheduleNameName = "Name";
export declare const ScheduleScheduleName = "Schedule";
export declare const ScheduleFunctionName = "Function";
export declare const ScheduleVersionName = "Version";
