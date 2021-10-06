import { Writer, Reader } from 'protobufjs/minimal';
export declare const protobufPackage = "regnull.ubkoin.ubkoin";
export declare enum Protocol {
    PL_UNKNOWN = 0,
    PL_EMAIL = 1,
    PL_GENERIC = 100,
    UNRECOGNIZED = -1
}
export declare function protocolFromJSON(object: any): Protocol;
export declare function protocolToJSON(object: Protocol): string;
export interface Key {
    registrationTimestamp: number;
}
export interface RegisterKey {
    key: Uint8Array;
}
export interface EncryptionKey {
    key: Uint8Array;
    validUntil: number;
}
export interface Endpoint {
    protocol: Protocol;
    address: string;
}
export interface NameRecord {
    name: string;
    ownerKey: Uint8Array;
    endpoint: Endpoint[];
    encryptionKey: EncryptionKey[];
}
export interface ReserveName {
    name: string;
    ownerAddress: string;
}
export declare const Key: {
    encode(message: Key, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Key;
    fromJSON(object: any): Key;
    toJSON(message: Key): unknown;
    fromPartial(object: DeepPartial<Key>): Key;
};
export declare const RegisterKey: {
    encode(message: RegisterKey, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): RegisterKey;
    fromJSON(object: any): RegisterKey;
    toJSON(message: RegisterKey): unknown;
    fromPartial(object: DeepPartial<RegisterKey>): RegisterKey;
};
export declare const EncryptionKey: {
    encode(message: EncryptionKey, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): EncryptionKey;
    fromJSON(object: any): EncryptionKey;
    toJSON(message: EncryptionKey): unknown;
    fromPartial(object: DeepPartial<EncryptionKey>): EncryptionKey;
};
export declare const Endpoint: {
    encode(message: Endpoint, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Endpoint;
    fromJSON(object: any): Endpoint;
    toJSON(message: Endpoint): unknown;
    fromPartial(object: DeepPartial<Endpoint>): Endpoint;
};
export declare const NameRecord: {
    encode(message: NameRecord, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): NameRecord;
    fromJSON(object: any): NameRecord;
    toJSON(message: NameRecord): unknown;
    fromPartial(object: DeepPartial<NameRecord>): NameRecord;
};
export declare const ReserveName: {
    encode(message: ReserveName, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): ReserveName;
    fromJSON(object: any): ReserveName;
    toJSON(message: ReserveName): unknown;
    fromPartial(object: DeepPartial<ReserveName>): ReserveName;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
