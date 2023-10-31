import {Event} from '../model/extension';

export class EventHelper {
    public static shortInfo(event: Event): string {
        return `[${event.action}]${event.resource.namespace.name}/${event.resource.name}/${event.id}`;
    }
}