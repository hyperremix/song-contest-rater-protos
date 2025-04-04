// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.7.0
//   protoc               v5.29.3
// source: rating_service.proto

/* eslint-disable */
import { BinaryReader } from "@bufbuild/protobuf/wire";
import { Observable } from "rxjs";
import { map } from "rxjs/operators";
import { Empty } from "./google/protobuf/empty";
import {
  CreateRatingRequest,
  DeleteRatingRequest,
  GetRatingRequest,
  ListRatingsResponse,
  ListUserRatingsRequest,
  RatingResponse,
  UpdateRatingRequest,
} from "./rating";

export const protobufPackage = "songcontestrater";

export interface Rating {
  ListRatings(request: Empty): Promise<ListRatingsResponse>;
  ListUserRatings(request: ListUserRatingsRequest): Promise<ListRatingsResponse>;
  GetRating(request: GetRatingRequest): Promise<RatingResponse>;
  CreateRating(request: CreateRatingRequest): Promise<RatingResponse>;
  UpdateRating(request: UpdateRatingRequest): Promise<RatingResponse>;
  DeleteRating(request: DeleteRatingRequest): Promise<RatingResponse>;
  StreamRatings(request: Empty): Observable<RatingResponse>;
}

export const RatingServiceName = "songcontestrater.Rating";
export class RatingClientImpl implements Rating {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || RatingServiceName;
    this.rpc = rpc;
    this.ListRatings = this.ListRatings.bind(this);
    this.ListUserRatings = this.ListUserRatings.bind(this);
    this.GetRating = this.GetRating.bind(this);
    this.CreateRating = this.CreateRating.bind(this);
    this.UpdateRating = this.UpdateRating.bind(this);
    this.DeleteRating = this.DeleteRating.bind(this);
    this.StreamRatings = this.StreamRatings.bind(this);
  }
  ListRatings(request: Empty): Promise<ListRatingsResponse> {
    const data = Empty.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListRatings", data);
    return promise.then((data) => ListRatingsResponse.decode(new BinaryReader(data)));
  }

  ListUserRatings(request: ListUserRatingsRequest): Promise<ListRatingsResponse> {
    const data = ListUserRatingsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ListUserRatings", data);
    return promise.then((data) => ListRatingsResponse.decode(new BinaryReader(data)));
  }

  GetRating(request: GetRatingRequest): Promise<RatingResponse> {
    const data = GetRatingRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "GetRating", data);
    return promise.then((data) => RatingResponse.decode(new BinaryReader(data)));
  }

  CreateRating(request: CreateRatingRequest): Promise<RatingResponse> {
    const data = CreateRatingRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateRating", data);
    return promise.then((data) => RatingResponse.decode(new BinaryReader(data)));
  }

  UpdateRating(request: UpdateRatingRequest): Promise<RatingResponse> {
    const data = UpdateRatingRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateRating", data);
    return promise.then((data) => RatingResponse.decode(new BinaryReader(data)));
  }

  DeleteRating(request: DeleteRatingRequest): Promise<RatingResponse> {
    const data = DeleteRatingRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteRating", data);
    return promise.then((data) => RatingResponse.decode(new BinaryReader(data)));
  }

  StreamRatings(request: Empty): Observable<RatingResponse> {
    const data = Empty.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(this.service, "StreamRatings", data);
    return result.pipe(map((data) => RatingResponse.decode(new BinaryReader(data))));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
  clientStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Promise<Uint8Array>;
  serverStreamingRequest(service: string, method: string, data: Uint8Array): Observable<Uint8Array>;
  bidirectionalStreamingRequest(service: string, method: string, data: Observable<Uint8Array>): Observable<Uint8Array>;
}
