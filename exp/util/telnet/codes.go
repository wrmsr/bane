// https://datatracker.ietf.org/doc/html/rfc854
package telnet

const (
	// No Operation
	NULL = 0

	// Moves the printer to the next print line, keeping the same horizontal position.
	LF = 10
	// Moves the printer to the left margin of the current line.
	CR = 13

	// Produces an audible or visible signal (which does NOT move the print head).
	BEL = 7
	// Moves the print head one character position towards the left margin.
	BS = 8

	// Moves the printer to the next horizontal tab stop. It remains unspecified how either party determines or
	// establishes where such tab
	HT = 9
	// Moves the printer to the next vertical tab stop. It remains unspecified how either party determines or where such
	// tab are located.
	VT = 11

	// Moves the printer to the top the next page, keeping same horizontal position.
	FF = 12

	// End of subnegotiation parameters.
	SE = 240
	// No operation.
	NOP = 241
	// The data stream portion of a Synch. This should always be accompanied by a TCP Urgent notification.
	DM = 242
	// NVT character BRK.
	BRK = 243
	// The function IP.
	IP = 244
	// The function AO.
	AO = 245
	// The function AYT.
	AYT = 246
	// The function EC.
	EC = 247
	// The function EL.
	EL = 248
	// The GA signal.
	GA = 249
	// Indicates that what follows is subnegotiation of the indicated option.
	SB = 250
	// Indicates the desire to begin performing, or confirmation that you are now performing, the indicated option.
	WILL = 251
	// Indicates the refusal to perform, or continue performing, the indicated option.
	WONT = 252
	// Indicates the request that the other party perform, or confirmation that you are expecting the other party to
	// perform, the indicated option.
	DO = 253
	// Indicates the demand that the other party stop performing, or confirmation that you are no longer expecting the
	// other party to perform, the indicated option.
	DONT = 254
	// Data Byte 255.
	IAC = 255
)
