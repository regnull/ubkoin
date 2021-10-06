/* eslint-disable */
import * as Long from 'long';
import { util, configure, Writer, Reader } from 'protobufjs/minimal';
export const protobufPackage = 'regnull.ubkoin.ubkoin';
export var Protocol;
(function (Protocol) {
    Protocol[Protocol["PL_UNKNOWN"] = 0] = "PL_UNKNOWN";
    Protocol[Protocol["PL_EMAIL"] = 1] = "PL_EMAIL";
    Protocol[Protocol["PL_GENERIC"] = 100] = "PL_GENERIC";
    Protocol[Protocol["UNRECOGNIZED"] = -1] = "UNRECOGNIZED";
})(Protocol || (Protocol = {}));
export function protocolFromJSON(object) {
    switch (object) {
        case 0:
        case 'PL_UNKNOWN':
            return Protocol.PL_UNKNOWN;
        case 1:
        case 'PL_EMAIL':
            return Protocol.PL_EMAIL;
        case 100:
        case 'PL_GENERIC':
            return Protocol.PL_GENERIC;
        case -1:
        case 'UNRECOGNIZED':
        default:
            return Protocol.UNRECOGNIZED;
    }
}
export function protocolToJSON(object) {
    switch (object) {
        case Protocol.PL_UNKNOWN:
            return 'PL_UNKNOWN';
        case Protocol.PL_EMAIL:
            return 'PL_EMAIL';
        case Protocol.PL_GENERIC:
            return 'PL_GENERIC';
        default:
            return 'UNKNOWN';
    }
}
const baseKey = { registrationTimestamp: 0 };
export const Key = {
    encode(message, writer = Writer.create()) {
        if (message.registrationTimestamp !== 0) {
            writer.uint32(8).int64(message.registrationTimestamp);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseKey };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.registrationTimestamp = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseKey };
        if (object.registrationTimestamp !== undefined && object.registrationTimestamp !== null) {
            message.registrationTimestamp = Number(object.registrationTimestamp);
        }
        else {
            message.registrationTimestamp = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.registrationTimestamp !== undefined && (obj.registrationTimestamp = message.registrationTimestamp);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseKey };
        if (object.registrationTimestamp !== undefined && object.registrationTimestamp !== null) {
            message.registrationTimestamp = object.registrationTimestamp;
        }
        else {
            message.registrationTimestamp = 0;
        }
        return message;
    }
};
const baseRegisterKey = {};
export const RegisterKey = {
    encode(message, writer = Writer.create()) {
        if (message.key.length !== 0) {
            writer.uint32(10).bytes(message.key);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseRegisterKey };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.key = reader.bytes();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseRegisterKey };
        if (object.key !== undefined && object.key !== null) {
            message.key = bytesFromBase64(object.key);
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.key !== undefined && (obj.key = base64FromBytes(message.key !== undefined ? message.key : new Uint8Array()));
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseRegisterKey };
        if (object.key !== undefined && object.key !== null) {
            message.key = object.key;
        }
        else {
            message.key = new Uint8Array();
        }
        return message;
    }
};
const baseEncryptionKey = { validUntil: 0 };
export const EncryptionKey = {
    encode(message, writer = Writer.create()) {
        if (message.key.length !== 0) {
            writer.uint32(10).bytes(message.key);
        }
        if (message.validUntil !== 0) {
            writer.uint32(16).int64(message.validUntil);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEncryptionKey };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.key = reader.bytes();
                    break;
                case 2:
                    message.validUntil = longToNumber(reader.int64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseEncryptionKey };
        if (object.key !== undefined && object.key !== null) {
            message.key = bytesFromBase64(object.key);
        }
        if (object.validUntil !== undefined && object.validUntil !== null) {
            message.validUntil = Number(object.validUntil);
        }
        else {
            message.validUntil = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.key !== undefined && (obj.key = base64FromBytes(message.key !== undefined ? message.key : new Uint8Array()));
        message.validUntil !== undefined && (obj.validUntil = message.validUntil);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEncryptionKey };
        if (object.key !== undefined && object.key !== null) {
            message.key = object.key;
        }
        else {
            message.key = new Uint8Array();
        }
        if (object.validUntil !== undefined && object.validUntil !== null) {
            message.validUntil = object.validUntil;
        }
        else {
            message.validUntil = 0;
        }
        return message;
    }
};
const baseEndpoint = { protocol: 0, address: '' };
export const Endpoint = {
    encode(message, writer = Writer.create()) {
        if (message.protocol !== 0) {
            writer.uint32(8).int32(message.protocol);
        }
        if (message.address !== '') {
            writer.uint32(18).string(message.address);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseEndpoint };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.protocol = reader.int32();
                    break;
                case 2:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseEndpoint };
        if (object.protocol !== undefined && object.protocol !== null) {
            message.protocol = protocolFromJSON(object.protocol);
        }
        else {
            message.protocol = 0;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = String(object.address);
        }
        else {
            message.address = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.protocol !== undefined && (obj.protocol = protocolToJSON(message.protocol));
        message.address !== undefined && (obj.address = message.address);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseEndpoint };
        if (object.protocol !== undefined && object.protocol !== null) {
            message.protocol = object.protocol;
        }
        else {
            message.protocol = 0;
        }
        if (object.address !== undefined && object.address !== null) {
            message.address = object.address;
        }
        else {
            message.address = '';
        }
        return message;
    }
};
const baseNameRecord = { name: '' };
export const NameRecord = {
    encode(message, writer = Writer.create()) {
        if (message.name !== '') {
            writer.uint32(10).string(message.name);
        }
        if (message.ownerKey.length !== 0) {
            writer.uint32(18).bytes(message.ownerKey);
        }
        for (const v of message.endpoint) {
            Endpoint.encode(v, writer.uint32(26).fork()).ldelim();
        }
        for (const v of message.encryptionKey) {
            EncryptionKey.encode(v, writer.uint32(34).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseNameRecord };
        message.endpoint = [];
        message.encryptionKey = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.ownerKey = reader.bytes();
                    break;
                case 3:
                    message.endpoint.push(Endpoint.decode(reader, reader.uint32()));
                    break;
                case 4:
                    message.encryptionKey.push(EncryptionKey.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseNameRecord };
        message.endpoint = [];
        message.encryptionKey = [];
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = '';
        }
        if (object.ownerKey !== undefined && object.ownerKey !== null) {
            message.ownerKey = bytesFromBase64(object.ownerKey);
        }
        if (object.endpoint !== undefined && object.endpoint !== null) {
            for (const e of object.endpoint) {
                message.endpoint.push(Endpoint.fromJSON(e));
            }
        }
        if (object.encryptionKey !== undefined && object.encryptionKey !== null) {
            for (const e of object.encryptionKey) {
                message.encryptionKey.push(EncryptionKey.fromJSON(e));
            }
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.name !== undefined && (obj.name = message.name);
        message.ownerKey !== undefined && (obj.ownerKey = base64FromBytes(message.ownerKey !== undefined ? message.ownerKey : new Uint8Array()));
        if (message.endpoint) {
            obj.endpoint = message.endpoint.map((e) => (e ? Endpoint.toJSON(e) : undefined));
        }
        else {
            obj.endpoint = [];
        }
        if (message.encryptionKey) {
            obj.encryptionKey = message.encryptionKey.map((e) => (e ? EncryptionKey.toJSON(e) : undefined));
        }
        else {
            obj.encryptionKey = [];
        }
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseNameRecord };
        message.endpoint = [];
        message.encryptionKey = [];
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = '';
        }
        if (object.ownerKey !== undefined && object.ownerKey !== null) {
            message.ownerKey = object.ownerKey;
        }
        else {
            message.ownerKey = new Uint8Array();
        }
        if (object.endpoint !== undefined && object.endpoint !== null) {
            for (const e of object.endpoint) {
                message.endpoint.push(Endpoint.fromPartial(e));
            }
        }
        if (object.encryptionKey !== undefined && object.encryptionKey !== null) {
            for (const e of object.encryptionKey) {
                message.encryptionKey.push(EncryptionKey.fromPartial(e));
            }
        }
        return message;
    }
};
const baseReserveName = { name: '', ownerAddress: '' };
export const ReserveName = {
    encode(message, writer = Writer.create()) {
        if (message.name !== '') {
            writer.uint32(10).string(message.name);
        }
        if (message.ownerAddress !== '') {
            writer.uint32(18).string(message.ownerAddress);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseReserveName };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                case 2:
                    message.ownerAddress = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseReserveName };
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = '';
        }
        if (object.ownerAddress !== undefined && object.ownerAddress !== null) {
            message.ownerAddress = String(object.ownerAddress);
        }
        else {
            message.ownerAddress = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.name !== undefined && (obj.name = message.name);
        message.ownerAddress !== undefined && (obj.ownerAddress = message.ownerAddress);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseReserveName };
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = '';
        }
        if (object.ownerAddress !== undefined && object.ownerAddress !== null) {
            message.ownerAddress = object.ownerAddress;
        }
        else {
            message.ownerAddress = '';
        }
        return message;
    }
};
var globalThis = (() => {
    if (typeof globalThis !== 'undefined')
        return globalThis;
    if (typeof self !== 'undefined')
        return self;
    if (typeof window !== 'undefined')
        return window;
    if (typeof global !== 'undefined')
        return global;
    throw 'Unable to locate global object';
})();
const atob = globalThis.atob || ((b64) => globalThis.Buffer.from(b64, 'base64').toString('binary'));
function bytesFromBase64(b64) {
    const bin = atob(b64);
    const arr = new Uint8Array(bin.length);
    for (let i = 0; i < bin.length; ++i) {
        arr[i] = bin.charCodeAt(i);
    }
    return arr;
}
const btoa = globalThis.btoa || ((bin) => globalThis.Buffer.from(bin, 'binary').toString('base64'));
function base64FromBytes(arr) {
    const bin = [];
    for (let i = 0; i < arr.byteLength; ++i) {
        bin.push(String.fromCharCode(arr[i]));
    }
    return btoa(bin.join(''));
}
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER');
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
