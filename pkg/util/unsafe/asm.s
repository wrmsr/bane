#include "go_asm.h"
#include "textflag.h"

TEXT ·__addints(SB),NOSPLIT,$0-24
    ADD R1, R0, 16(SB)
	MOVD	$ret+32(FP), R8
