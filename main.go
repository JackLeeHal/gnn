package main

import (
	"fmt"
	"gnn/snn"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"log"
)

func main() {
	// Simple Neural Network
	// Form the training matrices.
	inputs, labels := snn.MakeInputsAndLabels("train.csv")

	// Define our network architecture and learning parameters.
	config := snn.NeuralNetConfig{
		InputNeurons:  4,
		OutputNeurons: 3,
		HiddenNeurons: 3,
		NumEpochs:     5000,
		LearningRate:  0.3,
	}

	// Train the neural network.
	network := snn.NewNetwork(config)
	if err := network.Train(inputs, labels); err != nil {
		log.Fatal(err)
	}

	// Form the testing matrices.
	testInputs, testLabels := snn.MakeInputsAndLabels("test.csv")

	// Make the predictions using the trained model.
	predictions, err := network.Predict(testInputs)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the accuracy of our model.
	var truePosNeg int
	numPreds, _ := predictions.Dims()
	for i := 0; i < numPreds; i++ {

		// Get the label.
		labelRow := mat.Row(nil, i, testLabels)
		var prediction int
		for idx, label := range labelRow {
			if label == 1.0 {
				prediction = idx
				break
			}
		}

		// Accumulate the true positive/negative count.
		if predictions.At(i, prediction) == floats.Max(mat.Row(nil, i, predictions)) {
			truePosNeg++
		}
	}

	// Calculate the accuracy (subset accuracy).
	accuracy := float64(truePosNeg) / float64(numPreds)

	// Output the Accuracy value to standard out.
	fmt.Printf("\nAccuracy = %0.2f\n\n", accuracy)
}
