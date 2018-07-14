package message

type FeatureOpAction int

const (
	FeatureOpAdd FeatureOpAction = iota
	FeatureOpRemove
)

type FeatureOp struct {
	OpAction FeatureOpAction
	Feature  string
}
