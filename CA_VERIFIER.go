package crypto

// Verifier interface
type Verifier interface {
	Verify(args, responseData []byte, readset, writeset [][]byte, signature, enclavePk []byte) (bool, error)
}