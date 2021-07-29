package espay

type ISignature interface {
	CreateSignature(signatureKey string) string
}