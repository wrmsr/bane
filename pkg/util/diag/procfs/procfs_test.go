package procfs

import (
	"fmt"
	"strings"
	"testing"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	stru "github.com/wrmsr/bane/pkg/util/strings"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

const procStatusContent = `
Name:	cat
Umask:	0022
State:	R (running)
Tgid:	1
Ngid:	0
Pid:	1
PPid:	0
TracerPid:	0
Uid:	0	0	0	0
Gid:	0	0	0	0
FDSize:	64
Groups:	 
NStgid:	1
NSpid:	1
NSpgid:	1
NSsid:	1
VmPeak:	    2344 kB
VmSize:	    2344 kB
VmLck:	       0 kB
VmPin:	       0 kB
VmHWM:	     780 kB
VmRSS:	     780 kB
RssAnon:	      80 kB
RssFile:	     700 kB
RssShmem:	       0 kB
VmData:	     344 kB
VmStk:	     132 kB
VmExe:	      24 kB
VmLib:	    1748 kB
VmPTE:	      44 kB
VmSwap:	       0 kB
HugetlbPages:	       0 kB
CoreDumping:	0
THP_enabled:	1
Threads:	1
SigQ:	1/31345
SigPnd:	0000000000000000
ShdPnd:	0000000000000000
SigBlk:	0000000000000000
SigIgn:	0000000000000000
SigCgt:	0000000000000000
CapInh:	0000000000000000
CapPrm:	00000000a80425fb
CapEff:	00000000a80425fb
CapBnd:	00000000a80425fb
CapAmb:	0000000000000000
NoNewPrivs:	0
Seccomp:	2
Seccomp_filters:	1
Speculation_Store_Bypass:	vulnerable
Cpus_allowed:	3f
Cpus_allowed_list:	0-5
Mems_allowed:	1
Mems_allowed_list:	0
voluntary_ctxt_switches:	13
nonvoluntary_ctxt_switches:	1
`

func TestProcStatus(t *testing.T) {
	s := procStatusContent

	lines := its.OfSlice(stru.TrimSpaceSplit(s, "\n"))
	kvs := its.FlatMap(lines, func(l string) its.Iterable[bt.Kv[string, string]] {
		k, v, ok := stru.TrimSpaceCut(l, ":")
		if !ok {
			return opt.None[bt.Kv[string, string]]()
		}
		return opt.Just(bt.KvOf(strings.TrimSpace(k), v))
	})
	m := ctr.NewOrdMap(kvs)

	fmt.Println(m)
	tu.AssertDeepEqual(t, m.Get("Threads"), "1")
}

const procNetNetstatContent = `
TcpExt: SyncookiesSent SyncookiesRecv SyncookiesFailed EmbryonicRsts PruneCalled RcvPruned OfoPruned OutOfWindowIcmps LockDroppedIcmps ArpFilter TW TWRecycled TWKilled PAWSActive PAWSEstab DelayedACKs DelayedACKLocked DelayedACKLost ListenOverflows ListenDrops TCPHPHits TCPPureAcks TCPHPAcks TCPRenoRecovery TCPSackRecovery TCPSACKReneging TCPSACKReorder TCPRenoReorder TCPTSReorder TCPFullUndo TCPPartialUndo TCPDSACKUndo TCPLossUndo TCPLostRetransmit TCPRenoFailures TCPSackFailures TCPLossFailures TCPFastRetrans TCPSlowStartRetrans TCPTimeouts TCPLossProbes TCPLossProbeRecovery TCPRenoRecoveryFail TCPSackRecoveryFail TCPRcvCollapsed TCPBacklogCoalesce TCPDSACKOldSent TCPDSACKOfoSent TCPDSACKRecv TCPDSACKOfoRecv TCPAbortOnData TCPAbortOnClose TCPAbortOnMemory TCPAbortOnTimeout TCPAbortOnLinger TCPAbortFailed TCPMemoryPressures TCPMemoryPressuresChrono TCPSACKDiscard TCPDSACKIgnoredOld TCPDSACKIgnoredNoUndo TCPSpuriousRTOs TCPMD5NotFound TCPMD5Unexpected TCPMD5Failure TCPSackShifted TCPSackMerged TCPSackShiftFallback TCPBacklogDrop PFMemallocDrop TCPMinTTLDrop TCPDeferAcceptDrop IPReversePathFilter TCPTimeWaitOverflow TCPReqQFullDoCookies TCPReqQFullDrop TCPRetransFail TCPRcvCoalesce TCPOFOQueue TCPOFODrop TCPOFOMerge TCPChallengeACK TCPSYNChallenge TCPFastOpenActive TCPFastOpenActiveFail TCPFastOpenPassive TCPFastOpenPassiveFail TCPFastOpenListenOverflow TCPFastOpenCookieReqd TCPFastOpenBlackhole TCPSpuriousRtxHostQueues BusyPollRxPackets TCPAutoCorking TCPFromZeroWindowAdv TCPToZeroWindowAdv TCPWantZeroWindowAdv TCPSynRetrans TCPOrigDataSent TCPHystartTrainDetect TCPHystartTrainCwnd TCPHystartDelayDetect TCPHystartDelayCwnd TCPACKSkippedSynRecv TCPACKSkippedPAWS TCPACKSkippedSeq TCPACKSkippedFinWait2 TCPACKSkippedTimeWait TCPACKSkippedChallenge TCPWinProbe TCPKeepAlive TCPMTUPFail TCPMTUPSuccess TCPDelivered TCPDeliveredCE TCPAckCompressed TCPZeroWindowDrop TCPRcvQDrop TCPWqueueTooBig TCPFastOpenPassiveAltKey TcpTimeoutRehash TcpDuplicateDataRehash TCPDSACKRecvSegs TCPDSACKIgnoredDubious
TcpExt: 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
IpExt: InNoRoutes InTruncatedPkts InMcastPkts OutMcastPkts InBcastPkts OutBcastPkts InOctets OutOctets InMcastOctets OutMcastOctets InBcastOctets OutBcastOctets InCsumErrors InNoECTPkts InECT1Pkts InECT0Pkts InCEPkts ReasmOverlaps
IpExt: 0 0 0 0 420 0 0 0 0 0 0 0 0 0 0 0 0 0
`

func TestProcNetNetstat(t *testing.T) {
	s := procNetNetstatContent
	lines := its.OfSlice(stru.TrimSpaceSplit(s, "\n"))

	kvs := its.FlatMap(its.ChunkShared(lines, 2), func(ls []string) its.Iterable[bt.Kv[string, ctr.OrdMap[string, string]]] {
		if len(ls) != 2 {
			return opt.None[bt.Kv[string, ctr.OrdMap[string, string]]]()
		}

		kt, kl, kok := stru.TrimSpaceCut(ls[0], ":")
		vt, vl, vok := stru.TrimSpaceCut(ls[1], ":")
		if !kok || !vok || kt != vt {
			return opt.None[bt.Kv[string, ctr.OrdMap[string, string]]]()
		}

		kvs := its.Kvs(its.Zip[string, string](
			its.OfSlice(stru.TrimSpaceSplit(kl, " ")),
			its.OfSlice(stru.TrimSpaceSplit(vl, " ")),
		))
		return opt.Just(bt.KvOf(kt, ctr.NewOrdMap(kvs)))
	})
	m := ctr.NewOrdMap(kvs)

	fmt.Println(m)
	tu.AssertDeepEqual(t, m.Get("IpExt").Get("InBcastPkts"), "420")
}
