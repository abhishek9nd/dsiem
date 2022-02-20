// Code generated by command: go run gen.go -out ../encodeblock_amd64.s -stubs ../encodeblock_amd64.go -pkg=s2. DO NOT EDIT.

// +build !appengine
// +build !noasm
// +build gc

package s2

// encodeBlockAsm encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 4294967295 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeBlockAsm(dst []byte, src []byte) int

// encodeBlockAsm4MB encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 4194304 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeBlockAsm4MB(dst []byte, src []byte) int

// encodeBlockAsm12B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 16383 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeBlockAsm12B(dst []byte, src []byte) int

// encodeBlockAsm10B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 4095 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeBlockAsm10B(dst []byte, src []byte) int

// encodeBlockAsm8B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 511 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeBlockAsm8B(dst []byte, src []byte) int

// encodeBetterBlockAsm encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 4294967295 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeBetterBlockAsm(dst []byte, src []byte) int

// encodeBetterBlockAsm4MB encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 4194304 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeBetterBlockAsm4MB(dst []byte, src []byte) int

// encodeBetterBlockAsm12B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 16383 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeBetterBlockAsm12B(dst []byte, src []byte) int

// encodeBetterBlockAsm10B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 4095 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeBetterBlockAsm10B(dst []byte, src []byte) int

// encodeBetterBlockAsm8B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 511 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeBetterBlockAsm8B(dst []byte, src []byte) int

// encodeSnappyBlockAsm encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 4294967295 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeSnappyBlockAsm(dst []byte, src []byte) int

// encodeSnappyBlockAsm64K encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 65535 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeSnappyBlockAsm64K(dst []byte, src []byte) int

// encodeSnappyBlockAsm12B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 16383 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeSnappyBlockAsm12B(dst []byte, src []byte) int

// encodeSnappyBlockAsm10B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 4095 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeSnappyBlockAsm10B(dst []byte, src []byte) int

// encodeSnappyBlockAsm8B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 511 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeSnappyBlockAsm8B(dst []byte, src []byte) int

// encodeSnappyBetterBlockAsm encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 4294967295 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeSnappyBetterBlockAsm(dst []byte, src []byte) int

// encodeSnappyBetterBlockAsm64K encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 65535 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeSnappyBetterBlockAsm64K(dst []byte, src []byte) int

// encodeSnappyBetterBlockAsm12B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 16383 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeSnappyBetterBlockAsm12B(dst []byte, src []byte) int

// encodeSnappyBetterBlockAsm10B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 4095 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeSnappyBetterBlockAsm10B(dst []byte, src []byte) int

// encodeSnappyBetterBlockAsm8B encodes a non-empty src to a guaranteed-large-enough dst.
// Maximum input 511 bytes.
// It assumes that the varint-encoded length of the decompressed bytes has already been written.
//
//go:noescape
func encodeSnappyBetterBlockAsm8B(dst []byte, src []byte) int

// emitLiteral writes a literal chunk and returns the number of bytes written.
//
// It assumes that:
//   dst is long enough to hold the encoded bytes with margin of 0 bytes
//   0 <= len(lit) && len(lit) <= math.MaxUint32
//
//go:noescape
func emitLiteral(dst []byte, lit []byte) int

// emitRepeat writes a repeat chunk and returns the number of bytes written.
// Length must be at least 4 and < 1<<32
//
//go:noescape
func emitRepeat(dst []byte, offset int, length int) int

// emitCopy writes a copy chunk and returns the number of bytes written.
//
// It assumes that:
//   dst is long enough to hold the encoded bytes
//   1 <= offset && offset <= math.MaxUint32
//   4 <= length && length <= 1 << 24
//
//go:noescape
func emitCopy(dst []byte, offset int, length int) int

// emitCopyNoRepeat writes a copy chunk and returns the number of bytes written.
//
// It assumes that:
//   dst is long enough to hold the encoded bytes
//   1 <= offset && offset <= math.MaxUint32
//   4 <= length && length <= 1 << 24
//
//go:noescape
func emitCopyNoRepeat(dst []byte, offset int, length int) int

// matchLen returns how many bytes match in a and b
//
// It assumes that:
//   len(a) <= len(b)
//
//go:noescape
func matchLen(a []byte, b []byte) int
