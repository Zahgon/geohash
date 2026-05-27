package geohash

// invalid is a placeholder for invalid character decodings.
const invalid = 0xff

// encoding encapsulates an encoding defined by a given base32 alphabet.
type encoding struct {
	encode string
	decode [256]byte
}

// newEncoding constructs a new encoding defined by the given alphabet,
// which must be a 32-byte string.
func newEncoding(encoder string) *encoding { _ = "STUB: not implemented"; return nil }

// ValidByte reports whether b is part of the encoding.
func (e *encoding) ValidByte(b byte) bool { _ = "STUB: not implemented"; return false }

// Decode string into bits of a 64-bit word. The string s may be at most 12
// characters.
func (e *encoding) Decode(s string) uint64 { _ = "STUB: not implemented"; return 0 }

// Encode bits of 64-bit word into a string.
func (e *encoding) Encode(x uint64) string { _ = "STUB: not implemented"; return "" }

// Base32Encoding with the Geohash alphabet.
var base32encoding = newEncoding("0123456789bcdefghjkmnpqrstuvwxyz")
