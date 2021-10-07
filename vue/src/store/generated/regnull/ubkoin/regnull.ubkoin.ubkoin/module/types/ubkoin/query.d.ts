import { Reader, Writer } from 'protobufjs/minimal';
import { NameRecord } from '../ubkoin/ubkoin';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export declare const protobufPackage = "regnull.ubkoin.ubkoin";
export interface QueryGetNameRequest {
    name: string;
}
export interface QueryGetNameResponse {
    record: NameRecord | undefined;
}
export interface QueryAllNameRequest {
    pagination: PageRequest | undefined;
}
export interface QueryAllNameResponse {
    record: NameRecord[];
    pagination: PageResponse | undefined;
}
export declare const QueryGetNameRequest: {
    encode(message: QueryGetNameRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNameRequest;
    fromJSON(object: any): QueryGetNameRequest;
    toJSON(message: QueryGetNameRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetNameRequest>): QueryGetNameRequest;
};
export declare const QueryGetNameResponse: {
    encode(message: QueryGetNameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetNameResponse;
    fromJSON(object: any): QueryGetNameResponse;
    toJSON(message: QueryGetNameResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetNameResponse>): QueryGetNameResponse;
};
export declare const QueryAllNameRequest: {
    encode(message: QueryAllNameRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllNameRequest;
    fromJSON(object: any): QueryAllNameRequest;
    toJSON(message: QueryAllNameRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllNameRequest>): QueryAllNameRequest;
};
export declare const QueryAllNameResponse: {
    encode(message: QueryAllNameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllNameResponse;
    fromJSON(object: any): QueryAllNameResponse;
    toJSON(message: QueryAllNameResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllNameResponse>): QueryAllNameResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** this line is used by starport scaffolding # 2 */
    Name(request: QueryGetNameRequest): Promise<QueryGetNameResponse>;
    NameAll(request: QueryAllNameRequest): Promise<QueryAllNameResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Name(request: QueryGetNameRequest): Promise<QueryGetNameResponse>;
    NameAll(request: QueryAllNameRequest): Promise<QueryAllNameResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
