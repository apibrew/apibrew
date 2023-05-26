import {ActionComponent} from "../../../model/component-interfaces.ts";

export class ActionExecuteFunction implements ActionComponent<any> {
    execute(...args: any): any {
        console.log('called!', args)
        return null;
    }
}
