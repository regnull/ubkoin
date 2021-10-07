/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'
import { NameRecord } from '../ubkoin/ubkoin'
import { PageRequest, PageResponse } from '../cosmos/base/query/v1beta1/pagination'

export const protobufPackage = 'regnull.ubkoin.ubkoin'

export interface QueryGetNameRequest {
  name: string
}

export interface QueryGetNameResponse {
  record: NameRecord | undefined
}

export interface QueryAllNameRequest {
  pagination: PageRequest | undefined
}

export interface QueryAllNameResponse {
  record: NameRecord[]
  pagination: PageResponse | undefined
}

const baseQueryGetNameRequest: object = { name: '' }

export const QueryGetNameRequest = {
  encode(message: QueryGetNameRequest, writer: Writer = Writer.create()): Writer {
    if (message.name !== '') {
      writer.uint32(10).string(message.name)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNameRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNameRequest } as QueryGetNameRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetNameRequest {
    const message = { ...baseQueryGetNameRequest } as QueryGetNameRequest
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name)
    } else {
      message.name = ''
    }
    return message
  },

  toJSON(message: QueryGetNameRequest): unknown {
    const obj: any = {}
    message.name !== undefined && (obj.name = message.name)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetNameRequest>): QueryGetNameRequest {
    const message = { ...baseQueryGetNameRequest } as QueryGetNameRequest
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name
    } else {
      message.name = ''
    }
    return message
  }
}

const baseQueryGetNameResponse: object = {}

export const QueryGetNameResponse = {
  encode(message: QueryGetNameResponse, writer: Writer = Writer.create()): Writer {
    if (message.record !== undefined) {
      NameRecord.encode(message.record, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryGetNameResponse } as QueryGetNameResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.record = NameRecord.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryGetNameResponse {
    const message = { ...baseQueryGetNameResponse } as QueryGetNameResponse
    if (object.record !== undefined && object.record !== null) {
      message.record = NameRecord.fromJSON(object.record)
    } else {
      message.record = undefined
    }
    return message
  },

  toJSON(message: QueryGetNameResponse): unknown {
    const obj: any = {}
    message.record !== undefined && (obj.record = message.record ? NameRecord.toJSON(message.record) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryGetNameResponse>): QueryGetNameResponse {
    const message = { ...baseQueryGetNameResponse } as QueryGetNameResponse
    if (object.record !== undefined && object.record !== null) {
      message.record = NameRecord.fromPartial(object.record)
    } else {
      message.record = undefined
    }
    return message
  }
}

const baseQueryAllNameRequest: object = {}

export const QueryAllNameRequest = {
  encode(message: QueryAllNameRequest, writer: Writer = Writer.create()): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNameRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllNameRequest } as QueryAllNameRequest
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllNameRequest {
    const message = { ...baseQueryAllNameRequest } as QueryAllNameRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllNameRequest): unknown {
    const obj: any = {}
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllNameRequest>): QueryAllNameRequest {
    const message = { ...baseQueryAllNameRequest } as QueryAllNameRequest
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

const baseQueryAllNameResponse: object = {}

export const QueryAllNameResponse = {
  encode(message: QueryAllNameResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.record) {
      NameRecord.encode(v!, writer.uint32(10).fork()).ldelim()
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim()
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNameResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseQueryAllNameResponse } as QueryAllNameResponse
    message.record = []
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.record.push(NameRecord.decode(reader, reader.uint32()))
          break
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32())
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): QueryAllNameResponse {
    const message = { ...baseQueryAllNameResponse } as QueryAllNameResponse
    message.record = []
    if (object.record !== undefined && object.record !== null) {
      for (const e of object.record) {
        message.record.push(NameRecord.fromJSON(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  },

  toJSON(message: QueryAllNameResponse): unknown {
    const obj: any = {}
    if (message.record) {
      obj.record = message.record.map((e) => (e ? NameRecord.toJSON(e) : undefined))
    } else {
      obj.record = []
    }
    message.pagination !== undefined && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined)
    return obj
  },

  fromPartial(object: DeepPartial<QueryAllNameResponse>): QueryAllNameResponse {
    const message = { ...baseQueryAllNameResponse } as QueryAllNameResponse
    message.record = []
    if (object.record !== undefined && object.record !== null) {
      for (const e of object.record) {
        message.record.push(NameRecord.fromPartial(e))
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination)
    } else {
      message.pagination = undefined
    }
    return message
  }
}

/** Query defines the gRPC querier service. */
export interface Query {
  /** this line is used by starport scaffolding # 2 */
  Name(request: QueryGetNameRequest): Promise<QueryGetNameResponse>
  NameAll(request: QueryAllNameRequest): Promise<QueryAllNameResponse>
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Name(request: QueryGetNameRequest): Promise<QueryGetNameResponse> {
    const data = QueryGetNameRequest.encode(request).finish()
    const promise = this.rpc.request('regnull.ubkoin.ubkoin.Query', 'Name', data)
    return promise.then((data) => QueryGetNameResponse.decode(new Reader(data)))
  }

  NameAll(request: QueryAllNameRequest): Promise<QueryAllNameResponse> {
    const data = QueryAllNameRequest.encode(request).finish()
    const promise = this.rpc.request('regnull.ubkoin.ubkoin.Query', 'NameAll', data)
    return promise.then((data) => QueryAllNameResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>
