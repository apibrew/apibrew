import { Entity } from "../dh-client/client";
export interface Country2 extends Entity {
    id: string;
    name: string;
    description: string;
    version: number;
}
export interface Store extends Entity {
    id: string;
    name: string;
    description: string;
    version: number;
}
export interface Category extends Entity {
    id: string;
    name: string;
    description: string;
    version: number;
}
export interface Product extends Entity {
    id: string;
    name: string;
    description: string;
    price: number;
    store: Store;
    category: Category;
    quantity: number;
    version: number;
}
export interface Customer extends Entity {
    id: string;
    name: string;
    description: string;
    email: string;
    phone: string;
    version: number;
}
export interface Order extends Entity {
    id: string;
    customer: Customer;
    product: Product;
    quantity: number;
    price: number;
    status: string;
    version: number;
}
export interface RichTest3995 extends Entity {
    object: object;
    id: string;
    createdBy: string;
    int32O: number;
    int32: number;
    int64: number;
    text: string;
    float: number;
    time: string;
    timestamp: string;
    string: string;
    updatedBy: string;
    double: number;
    uuid: string;
    bool: boolean;
    bytes: string;
    date: string;
    updatedOn: string;
    version: number;
    createdOn: string;
}
export interface Country extends Entity {
    id: string;
    name: string;
    description: string;
    population: number;
    area: number;
    version: number;
}
export interface City extends Entity {
    id: string;
    name: string;
    country: Country;
    description: string;
    version: number;
}
export interface Income extends Entity {
    id: string;
    country: Country;
    city: City;
    grossIncome: number;
    tax: number;
    netIncome: number;
    version: number;
}
export interface TaxRate extends Entity {
    id: string;
    name: string;
    country: Country;
    city: City;
    order: number;
    until: number;
    rate: number;
    version: number;
}
export interface VirtualResource extends Entity {
    description: string;
    version: number;
    id: string;
    name: string;
}
export interface Namespace extends Entity {
    id: string;
    version: number;
    createdBy: string;
    updatedBy: string;
    createdOn: string;
    updatedOn: string;
    name: string;
    description: string;
    details: object;
    securityContext: object;
}
export interface User extends Entity {
    id: string;
    version: number;
    createdBy: string;
    updatedBy: string;
    createdOn: string;
    updatedOn: string;
    username: string;
    password: string;
    securityContext: object;
    details: object;
}
export interface DataSource extends Entity {
    id: string;
    version: number;
    createdBy: string;
    updatedBy: string;
    createdOn: string;
    updatedOn: string;
    name: string;
    description: string;
    backend: number;
    optionsPostgresUsername: string;
    optionsPostgresPassword: string;
    optionsPostgresHost: string;
    optionsPostgresPort: number;
    optionsPostgresDbName: string;
    optionsPostgresDefaultSchema: string;
    optionsMysqlUsername: string;
    optionsMysqlPassword: string;
    optionsMysqlHost: string;
    optionsMysqlPort: number;
    optionsMysqlDbName: string;
    optionsMysqlDefaultSchema: string;
    optionsRedisAddr: string;
    optionsRedisPassword: string;
    optionsRedisDb: number;
    optionsMongoUri: string;
    optionsMongoDbName: string;
}
export interface Extension extends Entity {
    id: string;
    version: number;
    createdBy: string;
    updatedBy: string;
    createdOn: string;
    updatedOn: string;
    name: string;
    description: string;
    namespace: string;
    resource: string;
    before: object;
    after: object;
    instead: object;
}
export interface Resource extends Entity {
    id: string;
    version: number;
    createdBy: string;
    updatedBy: string;
    createdOn: string;
    updatedOn: string;
    name: string;
    namespace: Namespace;
    virtual: boolean;
    immutable: boolean;
    abstract: boolean;
    dataSource: DataSource;
    entity: string;
    catalog: string;
    annotations: object;
    indexes: object;
    securityContext: object;
    title: string;
    description: string;
}
