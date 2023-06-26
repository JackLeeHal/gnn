package main

import (
	"math"
)

func main() {
	// Initialize a new CNN with the specified architecture
	cnn := NewCNN()

	// Train the CNN with some training data
	cnn.Train()

	// Make predictions with the trained CNN
	cnn.Predict()
}

// Sigmoid activation function
func Sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// SigmoidDerivative Derivative of the sigmoid function
func SigmoidDerivative(x float64) float64 {
	return Sigmoid(x) * (1 - Sigmoid(x))
}
