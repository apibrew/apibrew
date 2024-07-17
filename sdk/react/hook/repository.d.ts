import { Entity, EntityInfo, Repository } from '@apibrew/client';
export declare function useRepository<T extends Entity>(entityInfo: EntityInfo): Repository<T>;
