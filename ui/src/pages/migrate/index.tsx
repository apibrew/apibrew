import { useEffect } from 'react';
import * as resources from '../../resources'
import { ResourceService } from '../../service/resource';

export function Migrate(): JSX.Element {
    useEffect(() => {
        resources.Resources.forEach(resource => {
            ResourceService.migrate(resource)
        })
    }, [])

    return <></>;
}