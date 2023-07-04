import { type Resource } from '../model';
import { ServiceConfig } from './config';
export declare namespace ResourceService {
    function list(config: ServiceConfig): Promise<Resource[]>;
    function create(config: ServiceConfig, resource: Resource): Promise<Resource>;
    function update(config: ServiceConfig, resource: Resource): Promise<Resource>;
    function remove(config: ServiceConfig, resource: Resource, forceMigrate: boolean): Promise<void>;
    function get(config: ServiceConfig, resourceId: string): Promise<Resource>;
    function getByName(config: ServiceConfig, resourceName: string, namespace?: string): Promise<Resource>;
    function save(config: ServiceConfig, resource: Resource): Promise<Resource>;
    function migrate(config: ServiceConfig, resource: Resource): Promise<Resource>;
}
