package iomanager

// I'm creating an interface so that CMDManager's methods can pass as an acceptable
// parameter to TaxIncludedPriceJob.
type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}
