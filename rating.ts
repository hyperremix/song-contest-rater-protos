// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v5.29.3
// source: rating.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
import { Timestamp } from "./google/protobuf/timestamp";
import { UserResponse } from "./user";

export const protobufPackage = "songcontestrater";

export interface ListUserRatingsRequest {
  user_id: string;
}

export interface GetRatingRequest {
  id: string;
}

export interface CreateRatingRequest {
  contest_id: string;
  act_id: string;
  song: number;
  singing: number;
  show: number;
  looks: number;
  clothes: number;
}

export interface UpdateRatingRequest {
  id: string;
  song: number;
  singing: number;
  show: number;
  looks: number;
  clothes: number;
}

export interface DeleteRatingRequest {
  id: string;
}

export interface RatingResponse {
  id: string;
  contest_id: string;
  act_id: string;
  song: number;
  singing: number;
  show: number;
  looks: number;
  clothes: number;
  total: number;
  user: UserResponse | undefined;
  created_at: Timestamp | undefined;
  updated_at: Timestamp | undefined;
}

export interface ListRatingsResponse {
  ratings: RatingResponse[];
}

function createBaseListUserRatingsRequest(): ListUserRatingsRequest {
  return { user_id: "" };
}

export const ListUserRatingsRequest: MessageFns<ListUserRatingsRequest> = {
  encode(message: ListUserRatingsRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.user_id !== "") {
      writer.uint32(10).string(message.user_id);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListUserRatingsRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListUserRatingsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.user_id = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListUserRatingsRequest {
    return { user_id: isSet(object.user_id) ? globalThis.String(object.user_id) : "" };
  },

  toJSON(message: ListUserRatingsRequest): unknown {
    const obj: any = {};
    if (message.user_id !== "") {
      obj.user_id = message.user_id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListUserRatingsRequest>, I>>(base?: I): ListUserRatingsRequest {
    return ListUserRatingsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListUserRatingsRequest>, I>>(object: I): ListUserRatingsRequest {
    const message = createBaseListUserRatingsRequest();
    message.user_id = object.user_id ?? "";
    return message;
  },
};

function createBaseGetRatingRequest(): GetRatingRequest {
  return { id: "" };
}

export const GetRatingRequest: MessageFns<GetRatingRequest> = {
  encode(message: GetRatingRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): GetRatingRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGetRatingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): GetRatingRequest {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
  },

  toJSON(message: GetRatingRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<GetRatingRequest>, I>>(base?: I): GetRatingRequest {
    return GetRatingRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<GetRatingRequest>, I>>(object: I): GetRatingRequest {
    const message = createBaseGetRatingRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseCreateRatingRequest(): CreateRatingRequest {
  return { contest_id: "", act_id: "", song: 0, singing: 0, show: 0, looks: 0, clothes: 0 };
}

export const CreateRatingRequest: MessageFns<CreateRatingRequest> = {
  encode(message: CreateRatingRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.contest_id !== "") {
      writer.uint32(10).string(message.contest_id);
    }
    if (message.act_id !== "") {
      writer.uint32(18).string(message.act_id);
    }
    if (message.song !== 0) {
      writer.uint32(24).int32(message.song);
    }
    if (message.singing !== 0) {
      writer.uint32(32).int32(message.singing);
    }
    if (message.show !== 0) {
      writer.uint32(40).int32(message.show);
    }
    if (message.looks !== 0) {
      writer.uint32(48).int32(message.looks);
    }
    if (message.clothes !== 0) {
      writer.uint32(56).int32(message.clothes);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): CreateRatingRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateRatingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.contest_id = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.act_id = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.song = reader.int32();
          continue;
        }
        case 4: {
          if (tag !== 32) {
            break;
          }

          message.singing = reader.int32();
          continue;
        }
        case 5: {
          if (tag !== 40) {
            break;
          }

          message.show = reader.int32();
          continue;
        }
        case 6: {
          if (tag !== 48) {
            break;
          }

          message.looks = reader.int32();
          continue;
        }
        case 7: {
          if (tag !== 56) {
            break;
          }

          message.clothes = reader.int32();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateRatingRequest {
    return {
      contest_id: isSet(object.contest_id) ? globalThis.String(object.contest_id) : "",
      act_id: isSet(object.act_id) ? globalThis.String(object.act_id) : "",
      song: isSet(object.song) ? globalThis.Number(object.song) : 0,
      singing: isSet(object.singing) ? globalThis.Number(object.singing) : 0,
      show: isSet(object.show) ? globalThis.Number(object.show) : 0,
      looks: isSet(object.looks) ? globalThis.Number(object.looks) : 0,
      clothes: isSet(object.clothes) ? globalThis.Number(object.clothes) : 0,
    };
  },

  toJSON(message: CreateRatingRequest): unknown {
    const obj: any = {};
    if (message.contest_id !== "") {
      obj.contest_id = message.contest_id;
    }
    if (message.act_id !== "") {
      obj.act_id = message.act_id;
    }
    if (message.song !== 0) {
      obj.song = Math.round(message.song);
    }
    if (message.singing !== 0) {
      obj.singing = Math.round(message.singing);
    }
    if (message.show !== 0) {
      obj.show = Math.round(message.show);
    }
    if (message.looks !== 0) {
      obj.looks = Math.round(message.looks);
    }
    if (message.clothes !== 0) {
      obj.clothes = Math.round(message.clothes);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateRatingRequest>, I>>(base?: I): CreateRatingRequest {
    return CreateRatingRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<CreateRatingRequest>, I>>(object: I): CreateRatingRequest {
    const message = createBaseCreateRatingRequest();
    message.contest_id = object.contest_id ?? "";
    message.act_id = object.act_id ?? "";
    message.song = object.song ?? 0;
    message.singing = object.singing ?? 0;
    message.show = object.show ?? 0;
    message.looks = object.looks ?? 0;
    message.clothes = object.clothes ?? 0;
    return message;
  },
};

function createBaseUpdateRatingRequest(): UpdateRatingRequest {
  return { id: "", song: 0, singing: 0, show: 0, looks: 0, clothes: 0 };
}

export const UpdateRatingRequest: MessageFns<UpdateRatingRequest> = {
  encode(message: UpdateRatingRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.song !== 0) {
      writer.uint32(16).int32(message.song);
    }
    if (message.singing !== 0) {
      writer.uint32(24).int32(message.singing);
    }
    if (message.show !== 0) {
      writer.uint32(32).int32(message.show);
    }
    if (message.looks !== 0) {
      writer.uint32(40).int32(message.looks);
    }
    if (message.clothes !== 0) {
      writer.uint32(48).int32(message.clothes);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): UpdateRatingRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUpdateRatingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.song = reader.int32();
          continue;
        }
        case 3: {
          if (tag !== 24) {
            break;
          }

          message.singing = reader.int32();
          continue;
        }
        case 4: {
          if (tag !== 32) {
            break;
          }

          message.show = reader.int32();
          continue;
        }
        case 5: {
          if (tag !== 40) {
            break;
          }

          message.looks = reader.int32();
          continue;
        }
        case 6: {
          if (tag !== 48) {
            break;
          }

          message.clothes = reader.int32();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UpdateRatingRequest {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      song: isSet(object.song) ? globalThis.Number(object.song) : 0,
      singing: isSet(object.singing) ? globalThis.Number(object.singing) : 0,
      show: isSet(object.show) ? globalThis.Number(object.show) : 0,
      looks: isSet(object.looks) ? globalThis.Number(object.looks) : 0,
      clothes: isSet(object.clothes) ? globalThis.Number(object.clothes) : 0,
    };
  },

  toJSON(message: UpdateRatingRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.song !== 0) {
      obj.song = Math.round(message.song);
    }
    if (message.singing !== 0) {
      obj.singing = Math.round(message.singing);
    }
    if (message.show !== 0) {
      obj.show = Math.round(message.show);
    }
    if (message.looks !== 0) {
      obj.looks = Math.round(message.looks);
    }
    if (message.clothes !== 0) {
      obj.clothes = Math.round(message.clothes);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UpdateRatingRequest>, I>>(base?: I): UpdateRatingRequest {
    return UpdateRatingRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UpdateRatingRequest>, I>>(object: I): UpdateRatingRequest {
    const message = createBaseUpdateRatingRequest();
    message.id = object.id ?? "";
    message.song = object.song ?? 0;
    message.singing = object.singing ?? 0;
    message.show = object.show ?? 0;
    message.looks = object.looks ?? 0;
    message.clothes = object.clothes ?? 0;
    return message;
  },
};

function createBaseDeleteRatingRequest(): DeleteRatingRequest {
  return { id: "" };
}

export const DeleteRatingRequest: MessageFns<DeleteRatingRequest> = {
  encode(message: DeleteRatingRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DeleteRatingRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteRatingRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteRatingRequest {
    return { id: isSet(object.id) ? globalThis.String(object.id) : "" };
  },

  toJSON(message: DeleteRatingRequest): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteRatingRequest>, I>>(base?: I): DeleteRatingRequest {
    return DeleteRatingRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteRatingRequest>, I>>(object: I): DeleteRatingRequest {
    const message = createBaseDeleteRatingRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseRatingResponse(): RatingResponse {
  return {
    id: "",
    contest_id: "",
    act_id: "",
    song: 0,
    singing: 0,
    show: 0,
    looks: 0,
    clothes: 0,
    total: 0,
    user: undefined,
    created_at: undefined,
    updated_at: undefined,
  };
}

export const RatingResponse: MessageFns<RatingResponse> = {
  encode(message: RatingResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.contest_id !== "") {
      writer.uint32(18).string(message.contest_id);
    }
    if (message.act_id !== "") {
      writer.uint32(26).string(message.act_id);
    }
    if (message.song !== 0) {
      writer.uint32(32).int32(message.song);
    }
    if (message.singing !== 0) {
      writer.uint32(40).int32(message.singing);
    }
    if (message.show !== 0) {
      writer.uint32(48).int32(message.show);
    }
    if (message.looks !== 0) {
      writer.uint32(56).int32(message.looks);
    }
    if (message.clothes !== 0) {
      writer.uint32(64).int32(message.clothes);
    }
    if (message.total !== 0) {
      writer.uint32(72).int32(message.total);
    }
    if (message.user !== undefined) {
      UserResponse.encode(message.user, writer.uint32(82).fork()).join();
    }
    if (message.created_at !== undefined) {
      Timestamp.encode(message.created_at, writer.uint32(90).fork()).join();
    }
    if (message.updated_at !== undefined) {
      Timestamp.encode(message.updated_at, writer.uint32(98).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): RatingResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRatingResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.contest_id = reader.string();
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          message.act_id = reader.string();
          continue;
        }
        case 4: {
          if (tag !== 32) {
            break;
          }

          message.song = reader.int32();
          continue;
        }
        case 5: {
          if (tag !== 40) {
            break;
          }

          message.singing = reader.int32();
          continue;
        }
        case 6: {
          if (tag !== 48) {
            break;
          }

          message.show = reader.int32();
          continue;
        }
        case 7: {
          if (tag !== 56) {
            break;
          }

          message.looks = reader.int32();
          continue;
        }
        case 8: {
          if (tag !== 64) {
            break;
          }

          message.clothes = reader.int32();
          continue;
        }
        case 9: {
          if (tag !== 72) {
            break;
          }

          message.total = reader.int32();
          continue;
        }
        case 10: {
          if (tag !== 82) {
            break;
          }

          message.user = UserResponse.decode(reader, reader.uint32());
          continue;
        }
        case 11: {
          if (tag !== 90) {
            break;
          }

          message.created_at = Timestamp.decode(reader, reader.uint32());
          continue;
        }
        case 12: {
          if (tag !== 98) {
            break;
          }

          message.updated_at = Timestamp.decode(reader, reader.uint32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RatingResponse {
    return {
      id: isSet(object.id) ? globalThis.String(object.id) : "",
      contest_id: isSet(object.contest_id) ? globalThis.String(object.contest_id) : "",
      act_id: isSet(object.act_id) ? globalThis.String(object.act_id) : "",
      song: isSet(object.song) ? globalThis.Number(object.song) : 0,
      singing: isSet(object.singing) ? globalThis.Number(object.singing) : 0,
      show: isSet(object.show) ? globalThis.Number(object.show) : 0,
      looks: isSet(object.looks) ? globalThis.Number(object.looks) : 0,
      clothes: isSet(object.clothes) ? globalThis.Number(object.clothes) : 0,
      total: isSet(object.total) ? globalThis.Number(object.total) : 0,
      user: isSet(object.user) ? UserResponse.fromJSON(object.user) : undefined,
      created_at: isSet(object.created_at) ? fromJsonTimestamp(object.created_at) : undefined,
      updated_at: isSet(object.updated_at) ? fromJsonTimestamp(object.updated_at) : undefined,
    };
  },

  toJSON(message: RatingResponse): unknown {
    const obj: any = {};
    if (message.id !== "") {
      obj.id = message.id;
    }
    if (message.contest_id !== "") {
      obj.contest_id = message.contest_id;
    }
    if (message.act_id !== "") {
      obj.act_id = message.act_id;
    }
    if (message.song !== 0) {
      obj.song = Math.round(message.song);
    }
    if (message.singing !== 0) {
      obj.singing = Math.round(message.singing);
    }
    if (message.show !== 0) {
      obj.show = Math.round(message.show);
    }
    if (message.looks !== 0) {
      obj.looks = Math.round(message.looks);
    }
    if (message.clothes !== 0) {
      obj.clothes = Math.round(message.clothes);
    }
    if (message.total !== 0) {
      obj.total = Math.round(message.total);
    }
    if (message.user !== undefined) {
      obj.user = UserResponse.toJSON(message.user);
    }
    if (message.created_at !== undefined) {
      obj.created_at = fromTimestamp(message.created_at).toISOString();
    }
    if (message.updated_at !== undefined) {
      obj.updated_at = fromTimestamp(message.updated_at).toISOString();
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RatingResponse>, I>>(base?: I): RatingResponse {
    return RatingResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RatingResponse>, I>>(object: I): RatingResponse {
    const message = createBaseRatingResponse();
    message.id = object.id ?? "";
    message.contest_id = object.contest_id ?? "";
    message.act_id = object.act_id ?? "";
    message.song = object.song ?? 0;
    message.singing = object.singing ?? 0;
    message.show = object.show ?? 0;
    message.looks = object.looks ?? 0;
    message.clothes = object.clothes ?? 0;
    message.total = object.total ?? 0;
    message.user = (object.user !== undefined && object.user !== null)
      ? UserResponse.fromPartial(object.user)
      : undefined;
    message.created_at = (object.created_at !== undefined && object.created_at !== null)
      ? Timestamp.fromPartial(object.created_at)
      : undefined;
    message.updated_at = (object.updated_at !== undefined && object.updated_at !== null)
      ? Timestamp.fromPartial(object.updated_at)
      : undefined;
    return message;
  },
};

function createBaseListRatingsResponse(): ListRatingsResponse {
  return { ratings: [] };
}

export const ListRatingsResponse: MessageFns<ListRatingsResponse> = {
  encode(message: ListRatingsResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    for (const v of message.ratings) {
      RatingResponse.encode(v!, writer.uint32(10).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): ListRatingsResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseListRatingsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.ratings.push(RatingResponse.decode(reader, reader.uint32()));
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): ListRatingsResponse {
    return {
      ratings: globalThis.Array.isArray(object?.ratings)
        ? object.ratings.map((e: any) => RatingResponse.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ListRatingsResponse): unknown {
    const obj: any = {};
    if (message.ratings?.length) {
      obj.ratings = message.ratings.map((e) => RatingResponse.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<ListRatingsResponse>, I>>(base?: I): ListRatingsResponse {
    return ListRatingsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<ListRatingsResponse>, I>>(object: I): ListRatingsResponse {
    const message = createBaseListRatingsResponse();
    message.ratings = object.ratings?.map((e) => RatingResponse.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function toTimestamp(date: Date): Timestamp {
  const seconds = Math.trunc(date.getTime() / 1_000);
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = (t.seconds || 0) * 1_000;
  millis += (t.nanos || 0) / 1_000_000;
  return new globalThis.Date(millis);
}

function fromJsonTimestamp(o: any): Timestamp {
  if (o instanceof globalThis.Date) {
    return toTimestamp(o);
  } else if (typeof o === "string") {
    return toTimestamp(new globalThis.Date(o));
  } else {
    return Timestamp.fromJSON(o);
  }
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
