import { Entity, EntityInfo, ListRecordParams } from '@apibrew/client';
export declare function useRecordByName<T extends Entity & {
    name: string;
}>(entityInfo: EntityInfo, name: string, wi?: number): T | undefined;
export declare function useRecordBy<T extends Entity>(entityInfo: EntityInfo, identifier: Partial<T>, wi?: number): T | undefined;
export declare function useRecords<T extends Entity>(entityInfo: EntityInfo, params?: ListRecordParams, wi?: number): T[] | undefined;
