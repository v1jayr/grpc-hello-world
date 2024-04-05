/**
 * @fileoverview gRPC-Web generated client stub for todoapp.grpc.models
 * @enhanceable
 * @public
 */
import * as grpcWeb from 'grpc-web';
import * as service_pb from './service_pb';
export declare class HelloWorldClient {
    client_: grpcWeb.AbstractClientBase;
    hostname_: string;
    credentials_: null | {
        [index: string]: string;
    };
    options_: null | {
        [index: string]: any;
    };
    constructor(hostname: string, credentials?: null | {
        [index: string]: string;
    }, options?: null | {
        [index: string]: any;
    });
    methodDescriptorGreet: any;
    greet(request: service_pb.GreetRequest, metadata?: grpcWeb.Metadata | null): Promise<service_pb.GreetResponse>;
    greet(request: service_pb.GreetRequest, metadata: grpcWeb.Metadata | null, callback: (err: grpcWeb.RpcError, response: service_pb.GreetResponse) => void): grpcWeb.ClientReadableStream<service_pb.GreetResponse>;
}
