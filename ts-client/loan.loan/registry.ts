import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRequestLoan } from "./types/loan/tx";
import { MsgRequest } from "./types/loan/tx";
import { MsgRepayLoan } from "./types/loan/tx";
import { MsgApproveLoan } from "./types/loan/tx";
import { MsgLiquidateLoan } from "./types/loan/tx";
import { MsgCancelLoan } from "./types/loan/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.MsgRequestLoan", MsgRequestLoan],
    ["/loan.loan.MsgRequest", MsgRequest],
    ["/loan.loan.MsgRepayLoan", MsgRepayLoan],
    ["/loan.loan.MsgApproveLoan", MsgApproveLoan],
    ["/loan.loan.MsgLiquidateLoan", MsgLiquidateLoan],
    ["/loan.loan.MsgCancelLoan", MsgCancelLoan],
    
];

export { msgTypes }