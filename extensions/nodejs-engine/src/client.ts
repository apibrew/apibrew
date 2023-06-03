import {createGrpcTransport} from "@bufbuild/connect-node";
import {APBR_ADDR} from "./config";
import {createPromiseClient} from "@bufbuild/connect";
import {Extension} from "./gen/stub/extension_connect";
import {Record} from "./gen/stub/record_connect";

const transport = createGrpcTransport({
    baseUrl: `http://${APBR_ADDR}`,
    httpVersion: '2',
});

export const extensionClient = createPromiseClient(Extension, transport);
export const recordClient = createPromiseClient(Record, transport);