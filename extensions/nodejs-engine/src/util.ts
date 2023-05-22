import * as jspb from "google-protobuf";
import * as grpc from "@grpc/grpc-js";

type GrpcApiCall<Request extends jspb.Message, Response extends jspb.Message> = (request: Request, callback: (error: grpc.ServiceError | null, response: Response) => void) => void;

export function toPromise<Request extends jspb.Message, Response extends jspb.Message>(call: GrpcApiCall<Request, Response>): (request: Request) => Promise<Response> {
    return (request) => new Promise((resolve, reject) => {
        call(request, (error, response) => {
            if (error) {
                reject(error);
            } else {
                resolve(response);
            }
        });
    })
}