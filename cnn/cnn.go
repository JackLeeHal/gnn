package main

// CNN represents a Convolutional Neural Network
type CNN struct {
	ConvLayer           *ConvLayer
	PoolLayer           *PoolLayer
	FullyConnectedLayer *FullyConnectedLayer
}

// NewCNN creates a new Convolutional Neural Network
func NewCNN() *CNN {
	return &CNN{
		ConvLayer:           NewConvLayer(),
		PoolLayer:           NewPoolLayer(),
		FullyConnectedLayer: NewFullyConnectedLayer(),
	}
}

// Train trains the CNN with the given training data
func (cnn *CNN) Train() {
	// TODO: Implement training logic
}

// Predict makes predictions with the trained CNN
func (cnn *CNN) Predict() {
	// TODO: Implement prediction logic
}
