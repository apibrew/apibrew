import React from 'react';
import { Client } from '@apibrew/client';
export declare const ClientContext: React.Context<Client | undefined>;
export declare const ClientProvider: React.Provider<Client | undefined>;
export declare const ClientConsumer: React.Consumer<Client | undefined>;
export declare function useClient(): Client;
