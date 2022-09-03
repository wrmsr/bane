package numpy

import "testing"

/*
https://numpy.org/devdocs/reference/generated/numpy.lib.format.html

====

Format Version 1.0
The first 6 bytes are a magic string: exactly \x93NUMPY.

The next 1 byte is an unsigned byte: the major version number of the file format, e.g. \x01.

The next 1 byte is an unsigned byte: the minor version number of the file format, e.g. \x00. Note: the version of the
file format is not tied to the version of the numpy package.

The next 2 bytes form a little-endian unsigned short int: the length of the header data HEADER_LEN.

The next HEADER_LEN bytes form the header data describing the array’s format. It is an ASCII string which contains a
Python literal expression of a dictionary. It is terminated by a newline (\n) and padded with spaces (\x20) to make the
total of len(magic string) + 2 + len(length) + HEADER_LEN be evenly divisible by 64 for alignment purposes.

The dictionary contains three keys:

“descr”dtype.descr
An object that can be passed as an argument to the numpy.dtype constructor to create the array’s dtype.

“fortran_order”bool
Whether the array data is Fortran-contiguous or not. Since Fortran-contiguous arrays are a common form of
non-C-contiguity, we allow them to be written directly to disk for efficiency.

“shape”tuple of int
The shape of the array.

For repeatability and readability, the dictionary keys are sorted in alphabetic order. This is for convenience only. A
writer SHOULD implement this if possible. A reader MUST NOT depend on this.

Following the header comes the array data. If the dtype contains Python objects (i.e. dtype.hasobject is True), then the
data is a Python pickle of the array. Otherwise the data is the contiguous (either C- or Fortran-, depending on
fortran_order) bytes of the array. Consumers can figure out the number of bytes by multiplying the number of elements
given by the shape (noting that shape=() means there is 1 element) by dtype.itemsize.

====

Format Version 2.0
The version 1.0 format only allowed the array header to have a total size of 65535 bytes. This can be exceeded by
structured arrays with a large number of columns. The version 2.0 format extends the header size to 4 GiB. numpy.save
will automatically save in 2.0 format if the data requires it, else it will always use the more compatible 1.0 format.

The description of the fourth element of the header therefore has become: “The next 4 bytes form a little-endian
unsigned int: the length of the header data HEADER_LEN.”

====

Format Version 3.0
This version replaces the ASCII string (which in practice was latin1) with a utf8-encoded string, so supports structured
types with any unicode field names.
*/

func TestNumpy(t *testing.T) {

}
