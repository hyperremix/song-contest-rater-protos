// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v5.29.3
// source: stat_service.proto

/* eslint-disable */
import { BinaryReader } from "@bufbuild/protobuf/wire";
import { Empty } from "./google/protobuf/empty";
import { GlobalStatsResponse, ListUserStatsResponse, UserStatsResponse } from "./stat";

export const protobufPackage = "songcontestrater";

export interface Stat {
  ListUserStats(request: Empty): Promise<ListUserStatsResponse>;
  GetMyStats(request: Empty): Promise<UserStatsResponse>;
  GetGlobalStats(request: Empty): Promise<GlobalStatsResponse>;
}

export const StatServiceName = "songcontestrater.Stat";
export class StatClientImpl implements Stat {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || StatServiceName;
    this.rpc = rpc;
    this.ListUserStats = this.ListUserStats.bind(this);
    this.GetMyStats = this.GetMyStats.bind(this);
    this.GetGlobalStats = this.GetGlobalStats.bind(this);
  }
  ListUserStats(request: Empty): Promise<ListUserStatsResponse> {
    const data = Empty.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListUserStats", data);
    return promise.then((data) => ListUserStatsResponse.decode(new BinaryReader(data)));
  }

  GetMyStats(request: Empty): Promise<UserStatsResponse> {
    const data = Empty.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetMyStats", data);
    return promise.then((data) => UserStatsResponse.decode(new BinaryReader(data)));
  }

  GetGlobalStats(request: Empty): Promise<GlobalStatsResponse> {
    const data = Empty.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetGlobalStats", data);
    return promise.then((data) => GlobalStatsResponse.decode(new BinaryReader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
