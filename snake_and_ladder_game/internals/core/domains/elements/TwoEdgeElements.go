package elements

//Elements : common interface for all the elements
type TwoEdgeElement interface {
	IsValid() error
	GetStart() int
	GetEnd() int
}
