package opencl

import (
	"testing"
)

func TestOpenCl(t *testing.T) {
	GetPlatforms()

	/*
		kernel = clCreateKernel(program, "myGEMM1", &err)
		err = clSetKernelArg(kernel, 0, sizeof(int), (void*)&M);
		err = clSetKernelArg(kernel, 1, sizeof(int), (void*)&N);
		err = clSetKernelArg(kernel, 2, sizeof(int), (void*)&K);
		err = clSetKernelArg(kernel, 3, sizeof(cl_mem), (void*)&A);
		err = clSetKernelArg(kernel, 4, sizeof(cl_mem), (void*)&B);
		err = clSetKernelArg(kernel, 5, sizeof(cl_mem), (void*)&C);
		const int TS = 32;
		const size_t local[2] = { TS, TS };
		const size_t global[2] = { M, N };
		err = clEnqueueNDRangeKernel(queue, kernel, 2, NULL,
		                             global, local, 0, NULL, &event);
		err = clWaitForEvents(1, &event);
	*/
}
