#include "funcdata.h"
#include "textflag.h"

TEXT ·__addints(SB),NOSPLIT|NOFRAME,$0-24
    MOVD x(FP), R0
    MOVD y+8(FP), R1
	ADD R1, R0, R0
	MOVD R0, ret+16(FP)
	RET (R30)
