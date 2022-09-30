//                       **************************************************************
//                       *                          FUNCTION                          *
//                       **************************************************************
//                       undefined _lj_vm_ffi_call()
//       undefined         w0:1           <RETURN>
//       undefined8        Stack[-0x10]:8 local_10                                XREF[2]:     100009a00(W),
//                                                                                             100009a70(*)
//       undefined8        Stack[-0x20]:8 local_20                                XREF[1]:     1000099fc(W)
//                       _lj_vm_ffi_call                                 XREF[1]:     _lj_ccall_func:1000905e4(c)
// 1000099fc f4 4f be a9     stp        x20,x19,[sp, #local_20]!
// 100009a00 fd 7b 01 a9     stp        x29,x30,[sp, #local_10]
// 100009a04 fd 43 00 91     add        x29,sp,#0x10
// 100009a08 f3 03 00 aa     mov        x19,x0
// 100009a0c 08 08 40 b9     ldr        w8,[x0, #0x8]
// 100009a10 69 32 40 39     ldrb       w9,[x19, #0xc]
// 100009a14 6a 62 02 91     add        x10,x19,#0x98
// 100009a18 29 05 00 f1     subs       x9,x9,#0x1
// 100009a1c 6b 02 40 f9     ldr        x11,[x19]
// 100009a20 ff 63 28 cb     sub        sp,sp,x8
// 100009a24 a4 00 00 54     b.mi       LAB_100009a38
//                       LAB_100009a28                                   XREF[1]:     100009a34(j)
// 100009a28 48 79 69 f8     ldr        x8,[x10, x9, LSL #0x3]
// 100009a2c e8 7b 29 f8     str        x8,[sp, x9, LSL #0x3]
// 100009a30 29 05 00 f1     subs       x9,x9,#0x1
// 100009a34 a5 ff ff 54     b.pl       LAB_100009a28
//                       LAB_100009a38                                   XREF[1]:     100009a24(j)
// 100009a38 60 86 45 a9     ldp        x0,x1,[x19, #0x58]
// 100009a3c 60 86 41 6d     ldp        d0,d1,[x19, #0x18]
// 100009a40 62 8e 46 a9     ldp        x2,x3,[x19, #0x68]
// 100009a44 62 8e 42 6d     ldp        d2,d3,[x19, #0x28]
// 100009a48 64 96 47 a9     ldp        x4,x5,[x19, #0x78]
// 100009a4c 64 96 43 6d     ldp        d4,d5,[x19, #0x38]
// 100009a50 66 9e 48 a9     ldp        x6,x7,[x19, #0x88]
// 100009a54 66 9e 44 6d     ldp        d6,d7,[x19, #0x48]
// 100009a58 68 0a 40 f9     ldr        x8,[x19, #0x10]
// 100009a5c 60 01 3f d6     blr        x11
// 100009a60 bf 43 00 d1     sub        sp,x29,#0x10
// 100009a64 60 86 05 a9     stp        x0,x1,[x19, #0x58]
// 100009a68 60 86 01 6d     stp        d0,d1,[x19, #0x18]
// 100009a6c 62 8e 02 6d     stp        d2,d3,[x19, #0x28]
// 100009a70 fd 7b 41 a9     ldp        x29=>local_10,x30,[sp, #0x10]
// 100009a74 f4 4f c2 a8     ldp        x20,x19,[sp], #0x20
// 100009a78 c0 03 5f d6     ret

// |->vm_ffi_call:			// Call C function via FFI.
// |  // Caveat: needs special frame unwinding, see below.
// |.if FFI
// |  .type CCSTATE, CCallState, x19
// |  stp x20, CCSTATE, [sp, #-32]!
// |  stp fp, lr, [sp, #16]
// |  add fp, sp, #16
// |  mov CCSTATE, x0
// |  ldr TMP0w, CCSTATE:x0->spadj
// |   ldrb TMP1w, CCSTATE->nsp
// |    add TMP2, CCSTATE, #offsetof(CCallState, stack)
// |   subs TMP1, TMP1, #1
// |    ldr TMP3, CCSTATE->func
// |  sub sp, sp, TMP0
// |   bmi >2
// |1:  // Copy stack slots
// |  ldr TMP0, [TMP2, TMP1, lsl #3]
// |  str TMP0, [sp, TMP1, lsl #3]
// |  subs TMP1, TMP1, #1
// |  bpl <1
// |2:
// |  ldp x0, x1, CCSTATE->gpr[0]
// |   ldp d0, d1, CCSTATE->fpr[0]
// |  ldp x2, x3, CCSTATE->gpr[2]
// |   ldp d2, d3, CCSTATE->fpr[2]
// |  ldp x4, x5, CCSTATE->gpr[4]
// |   ldp d4, d5, CCSTATE->fpr[4]
// |  ldp x6, x7, CCSTATE->gpr[6]
// |   ldp d6, d7, CCSTATE->fpr[6]
// |  ldr x8, CCSTATE->retp
// |  blr TMP3
// |  sub sp, fp, #16
// |  stp x0, x1, CCSTATE->gpr[0]
// |   stp d0, d1, CCSTATE->fpr[0]
// |   stp d2, d3, CCSTATE->fpr[2]
// |  ldp fp, lr, [sp, #16]
// |  ldp x20, CCSTATE, [sp], #32
// |  ret
// |.endif
// |// Note: vm_ffi_call must be the last function in this object file!
