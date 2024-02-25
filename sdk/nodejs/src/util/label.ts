import {Resource} from "../model";
import {Entity} from "../entity";

export function recordLabel(resource: Resource, record: Entity): string {
  const recordFields = record as any

  const nameProps = ['title', 'name', 'label', 'id'];

  for (const nameProp of nameProps) {
    if (resource.properties[nameProp] != null && recordFields[nameProp]) {
      return recordFields[nameProp]
    }
  }

  throw new Error('No name prop found for record: ' + JSON.stringify(record))
}
