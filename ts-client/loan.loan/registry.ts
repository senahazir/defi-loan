import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgLiquidateLoan } from "./types/loan/tx";
import { MsgRequestLoan } from "./types/loan/tx";
import { MsgRequest } from "./types/loan/tx";
import { MsgRepayLoan } from "./types/loan/tx";
import { MsgApproveLoan } from "./types/loan/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.MsgLiquidateLoan", MsgLiquidateLoan],
    ["/loan.loan.MsgRequestLoan", MsgRequestLoan],
    ["/loan.loan.MsgRequest", MsgRequest],
    ["/loan.loan.MsgRepayLoan", MsgRepayLoan],
    ["/loan.loan.MsgApproveLoan", MsgApproveLoan],
    
];

export { msgTypes }