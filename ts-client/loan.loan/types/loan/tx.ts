/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "loan.loan";

export interface MsgRequest {
  creator: string;
  loan: string;
}

export interface MsgRequestResponse {}

export interface MsgRequestLoan {
  creator: string;
  amount: string;
  fee: string;
  collateral: string;
  deadline: string;
}

export interface MsgRequestLoanResponse {}

const baseMsgRequest: object = { creator: "", loan: "" };

export const MsgRequest = {
  encode(message: MsgRequest, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.loan !== "") {
      writer.uint32(18).string(message.loan);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRequest } as MsgRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.loan = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRequest {
    const message = { ...baseMsgRequest } as MsgRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.loan !== undefined && object.loan !== null) {
      message.loan = String(object.loan);
    } else {
      message.loan = "";
    }
    return message;
  },

  toJSON(message: MsgRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.loan !== undefined && (obj.loan = message.loan);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRequest>): MsgRequest {
    const message = { ...baseMsgRequest } as MsgRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.loan !== undefined && object.loan !== null) {
      message.loan = object.loan;
    } else {
      message.loan = "";
    }
    return message;
  },
};

const baseMsgRequestResponse: object = {};

export const MsgRequestResponse = {
  encode(_: MsgRequestResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRequestResponse } as MsgRequestResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRequestResponse {
    const message = { ...baseMsgRequestResponse } as MsgRequestResponse;
    return message;
  },

  toJSON(_: MsgRequestResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgRequestResponse>): MsgRequestResponse {
    const message = { ...baseMsgRequestResponse } as MsgRequestResponse;
    return message;
  },
};

const baseMsgRequestLoan: object = {
  creator: "",
  amount: "",
  fee: "",
  collateral: "",
  deadline: "",
};

export const MsgRequestLoan = {
  encode(message: MsgRequestLoan, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.fee !== "") {
      writer.uint32(26).string(message.fee);
    }
    if (message.collateral !== "") {
      writer.uint32(34).string(message.collateral);
    }
    if (message.deadline !== "") {
      writer.uint32(42).string(message.deadline);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestLoan {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRequestLoan } as MsgRequestLoan;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        case 3:
          message.fee = reader.string();
          break;
        case 4:
          message.collateral = reader.string();
          break;
        case 5:
          message.deadline = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRequestLoan {
    const message = { ...baseMsgRequestLoan } as MsgRequestLoan;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = String(object.collateral);
    } else {
      message.collateral = "";
    }
    if (object.deadline !== undefined && object.deadline !== null) {
      message.deadline = String(object.deadline);
    } else {
      message.deadline = "";
    }
    return message;
  },

  toJSON(message: MsgRequestLoan): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    message.fee !== undefined && (obj.fee = message.fee);
    message.collateral !== undefined && (obj.collateral = message.collateral);
    message.deadline !== undefined && (obj.deadline = message.deadline);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRequestLoan>): MsgRequestLoan {
    const message = { ...baseMsgRequestLoan } as MsgRequestLoan;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = object.collateral;
    } else {
      message.collateral = "";
    }
    if (object.deadline !== undefined && object.deadline !== null) {
      message.deadline = object.deadline;
    } else {
      message.deadline = "";
    }
    return message;
  },
};

const baseMsgRequestLoanResponse: object = {};

export const MsgRequestLoanResponse = {
  encode(_: MsgRequestLoanResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestLoanResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRequestLoanResponse } as MsgRequestLoanResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRequestLoanResponse {
    const message = { ...baseMsgRequestLoanResponse } as MsgRequestLoanResponse;
    return message;
  },

  toJSON(_: MsgRequestLoanResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgRequestLoanResponse>): MsgRequestLoanResponse {
    const message = { ...baseMsgRequestLoanResponse } as MsgRequestLoanResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Request(request: MsgRequest): Promise<MsgRequestResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RequestLoan(request: MsgRequestLoan): Promise<MsgRequestLoanResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Request(request: MsgRequest): Promise<MsgRequestResponse> {
    const data = MsgRequest.encode(request).finish();
    const promise = this.rpc.request("loan.loan.Msg", "Request", data);
    return promise.then((data) => MsgRequestResponse.decode(new Reader(data)));
  }

  RequestLoan(request: MsgRequestLoan): Promise<MsgRequestLoanResponse> {
    const data = MsgRequestLoan.encode(request).finish();
    const promise = this.rpc.request("loan.loan.Msg", "RequestLoan", data);
    return promise.then((data) =>
      MsgRequestLoanResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
