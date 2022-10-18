#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

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

// (lldb) reg read
// General Purpose Registers:
//         x0 = 0x000000018e8f0ee8  libsystem_c.dylib`printf
//         x1 = 0x0000000000000000
//         x2 = 0xffffffffffffffff
//         x3 = 0x0000000000000000
//         x4 = 0x00000001001d2ae1  main`runtime.arm64HasATOMICS
//         x5 = 0x0000000040000000
//         x6 = 0x0000000100000000  main`_mh_execute_header
//         x7 = 0x0000000000000001
//         x8 = 0x000001400013dd08
//         x9 = 0x000001400013e000
//        x10 = 0x0000000248dfd488
//        x11 = 0x00000001002bc058
//        x12 = 0x0000000000000000
//        x13 = 0x00000001003bc000
//        x14 = 0x0000000000100000
//        x15 = 0x00000000000003fd
//        x16 = 0x000001400013c3a0
//        x17 = 0x00000001e8b23ed0  (void *)0x000000018ea14608: os_unfair_recursive_lock_unlock
//        x18 = 0x0000000000000000
//        x19 = 0x0000000000000008
//        x20 = 0x000001400013d930
//        x21 = 0x0000000000000000
//        x22 = 0x0000000100062134  main`runtime.goexit.abi0 + 4
//        x23 = 0x000001400013de30
//        x24 = 0xffffffffffffffff
//        x25 = 0x00000001000bece8  main`go.func.*
//        x26 = 0x000001400013dde0
//        x27 = 0x0000000040000000
//        x28 = 0x00000140000021a0
//         fp = 0x000001400013df38
//         lr = 0x00000001000b2b74  main`main.main + 36
//         sp = 0x000001400013df40
//         pc = 0x00000001000b2abc  main`github.com/wrmsr/bane/exp/util/lj/ffi.Ffi_call.abi0 + 12
//       cpsr = 0x60001000

TEXT ·Ffi_call(SB),NOSPLIT,$0
    BRK $0

    CALL ·runtime_procPin(SB)

    MOVD cs+0(FP), R9
    LDP CCallState_Gpr(R9), (R0, R1)
    LDP 0(R9), (R2, R3)
    LDP 0(R9), (R4, R5)
    LDP 0(R9), (R6, R7)

    MOVD cs+0(FP), R9
    FLDPD CCallState_Fpr(R9), (F0, F1)
    FLDPD 0(R9), (F2, F3)
    FLDPD 0(R9), (F4, F5)
    FLDPD 0(R9), (F6, F7)

    MOVD cs+0(FP), R9
    MOVD CCallState_Fn(R9), R9
    CALL R9

    CALL ·runtime_procUnpin(SB)

	RET
