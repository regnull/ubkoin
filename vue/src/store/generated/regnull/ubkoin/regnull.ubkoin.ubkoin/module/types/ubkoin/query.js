/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal';
import { NameRecord } from '../ubkoin/ubkoin';
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination';
export const protobufPackage = 'regnull.ubkoin.ubkoin';
const baseQueryGetNameRequest = { name: '' };
export const QueryGetNameRequest = {
    encode(message, writer = Writer.create()) {
        if (message.name !== '') {
            writer.uint32(10).string(message.name);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetNameRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.name = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetNameRequest };
        if (object.name !== undefined && object.name !== null) {
            message.name = String(object.name);
        }
        else {
            message.name = '';
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.name !== undefined && (obj.name = message.name);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetNameRequest };
        if (object.name !== undefined && object.name !== null) {
            message.name = object.name;
        }
        else {
            message.name = '';
        }
        return message;
    }
};
const baseQueryGetNameResponse = {};
export const QueryGetNameResponse = {
    encode(message, writer = Writer.create()) {
        if (message.record !== undefined) {
            NameRecord.encode(message.record, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryGetNameResponse };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.record = NameRecord.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryGetNameResponse };
        if (object.record !== undefined && object.record !== null) {
            message.record = NameRecord.fromJSON(object.record);
        }
        else {
            message.record = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.record !== undefined && (obj.record = message.record ? NameRecord.toJSON(message.record) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryGetNameResponse };
        if (object.record !== undefined && object.record !== null) {
            message.record = NameRecord.fromPartial(object.record);
        }
        else {
            message.record = undefined;
        }
        return message;
    }
};
const baseQueryAllNameRequest = {};
export const QueryAllNameRequest = {
    encode(message, writer = Writer.create()) {
        if (message.pagination !== undefined) {
            PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllNameRequest };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.pagination = PageRequest.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllNameRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllNameRequest };
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageRequest.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
const baseQueryAllNameResponse = {};
export const QueryAllNameResponse = {
    encode(message, writer = Writer.create()) {
        for (const v of message.record) {
            NameRecord.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.pagination !== undefined) {
            PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseQueryAllNameResponse };
        message.record = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.record.push(NameRecord.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.pagination = PageResponse.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseQueryAllNameResponse };
        message.record = [];
        if (object.record !== undefined && object.record !== null) {
            for (const e of object.record) {
                message.record.push(NameRecord.fromJSON(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromJSON(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.record) {
            obj.record = message.record.map((e) => (e ? NameRecord.toJSON(e) : undefined));
        }
        else {
            obj.record = [];
        }
        message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseQueryAllNameResponse };
        message.record = [];
        if (object.record !== undefined && object.record !== null) {
            for (const e of object.record) {
                message.record.push(NameRecord.fromPartial(e));
            }
        }
        if (object.pagination !== undefined && object.pagination !== null) {
            message.pagination = PageResponse.fromPartial(object.pagination);
        }
        else {
            message.pagination = undefined;
        }
        return message;
    }
};
export class QueryClientImpl {
    constructor(rpc) {
        this.rpc = rpc;
    }
    Name(request) {
        const data = QueryGetNameRequest.encode(request).finish();
        const promise = this.rpc.request('regnull.ubkoin.ubkoin.Query', 'Name', data);
        return promise.then((data) => QueryGetNameResponse.decode(new Reader(data)));
    }
    NameAll(request) {
        const data = QueryAllNameRequest.encode(request).finish();
        const promise = this.rpc.request('regnull.ubkoin.ubkoin.Query', 'NameAll', data);
        return promise.then((data) => QueryAllNameResponse.decode(new Reader(data)));
    }
}
