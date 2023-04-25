package opencl

import "C"
import "fmt"

type OpenClErrorCode int

const (
	ErrSuccess OpenClErrorCode = 0

	ErrDeviceNotFound                     = -1
	ErrDeviceNotAvailable                 = -2
	ErrCompilerNotAvailable               = -3
	ErrMemObjectAllocationFailure         = -4
	ErrOutOfResources                     = -5
	ErrOutOfHostMemory                    = -6
	ErrProfilingInfoNotAvailable          = -7
	ErrMemCopyOverlap                     = -8
	ErrImageFormatMismatch                = -9
	ErrImageFormatNotSupported            = -10
	ErrBuildProgramFailure                = -11
	ErrMapFailure                         = -12
	ErrMisalignedSubBufferOffset          = -13
	ErrExecStatusErrorForEventsInWaitList = -14
	ErrCompileProgramFailure              = -15
	ErrLinkerNotAvailable                 = -16
	ErrLinkProgramFailure                 = -17
	ErrDevicePartitionFailed              = -18
	ErrKernelArgInfoNotAvailable          = -19

	ErrInvalidValue                 = -30
	ErrInvalidDevice_Type           = -31
	ErrInvalidPlatform              = -32
	ErrInvalidDevice                = -33
	ErrInvalidContext               = -34
	ErrInvalidQueueProperties       = -35
	ErrInvalidCommandQueue          = -36
	ErrInvalidHostPtr               = -37
	ErrInvalidMemObject             = -38
	ErrInvalidImageFormatDescriptor = -39
	ErrInvalidImageSize             = -40
	ErrInvalidSampler               = -41
	ErrInvalidBinary                = -42
	ErrInvalidBuildOptions          = -43
	ErrInvalidProgram               = -44
	ErrInvalidProgramExecutable     = -45
	ErrInvalidKernelName            = -46
	ErrInvalidKernelDefinition      = -47
	ErrInvalidKernel                = -48
	ErrInvalidArgIndex              = -49
	ErrInvalidArgValue              = -50
	ErrInvalidArgSize               = -51
	ErrInvalidKernelArgs            = -52
	ErrInvalidWorkDimension         = -53
	ErrInvalidWorkGroupSize         = -54
	ErrInvalidWorkItem_size         = -55
	ErrInvalidGlobalOffset          = -56
	ErrInvalidEventWaitList         = -57
	ErrInvalidEvent                 = -58
	ErrInvalidOperation             = -59
	ErrInvalidGlObject              = -60
	ErrInvalidBufferSize            = -61
	ErrInvalidMipLevel              = -62
	ErrInvalidGlobalWorkSize        = -63
	ErrInvalidProperty              = -64
	ErrInvalidImageDescriptor       = -65
	ErrInvalidCompilerOptions       = -66
	ErrInvalidLinkerOptions         = -67
	ErrInvalidDevicePartitionCount  = -68
)

type OpenClError struct {
	Code int
}

func (e OpenClError) Error() string {
	return fmt.Sprintf("opencl error: %d", e.Code)
}
