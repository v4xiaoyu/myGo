package main

import (
	"../utils"
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	"github.com/tensorflow/tensorflow/tensorflow/go/op"
	"time"
)

func client() {
	c := make(chan string, 10)
	comp := make(chan int64)
	go utils.GetDate(c)
	go utils.GetTime(c)

	fmt.Println(<-c, <-c)

	t := time.Now()
	go utils.CompareNow(t, comp)

	b := <-comp

	if b > 0 {
		fmt.Println("after")
	} else if b < 0 {
		fmt.Println("before")
	} else {
		fmt.Println("now")
	}

	time.Sleep(3000)
	t2 := time.Now()
	go utils.CompareTime(t, t2, comp)

	a := <-comp

	if a > 0 {
		fmt.Println("t after t2")
	} else if a < 0 {
		fmt.Println("t before t2")
	} else {
		fmt.Println("t is t2")
	}
}

func learn() {
	// Load in a dataset, with headers. Header attributes will be stored.
	// Think of instances as a Data Frame structure in R or Pandas.
	// You can also create instances from scratch.
	rawData, err := base.ParseCSVToInstances("../../datasets/house-votes-84.csv", true)
	if err != nil {
		panic(err)
	}

	// Print a pleasant summary of your data.
	fmt.Println(rawData)

	//Initialises a new KNN classifier
	cls := knn.NewKnnClassifier("euclidean", "linear", 2)

	//Do a training-test split
	trainData, testData := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(trainData)

	//Calculates the Euclidean distance and returns the most popular label
	predictions, err := cls.Predict(testData)
	if err != nil {
		panic(err)
	}

	// Prints precision/recall metrics
	confusionMat, err := evaluation.GetConfusionMatrix(testData, predictions)
	if err != nil {
		panic(fmt.Sprintf("Unable to get confusion matrix: %s", err.Error()))
	}
	fmt.Println(evaluation.GetSummary(confusionMat))
}

func testTF() {
	// Let's describe what we want: create the graph

	// We want to define two placeholder to fill at runtime
	// the first placeholder A will be a [2, 2] tensor of integers
	// the second placeholder x will be a [2, 1] tensor of intergers

	// Then we want to compute Y = Ax

	// Create the first node of the graph: an empty node, the root of our graph
	root := op.NewScope()

	// define 2 subscopes of the root subscopes, called "input". In this
	// way we expect to have a input/ and a input_1/ scope under the root scope
	A := op.Placeholder(root.SubScope("input"), tf.Complex128, op.PlaceholderShape(tf.MakeShape(2, 2)))
	x := op.Placeholder(root.SubScope("input"), tf.Complex128, op.PlaceholderShape(tf.MakeShape(2, 1)))
	fmt.Println(A.Op.Name(), x.Op.Name())

	// Define the operation node that accepts A & x as inputs
	product := op.MatMul(root, A, x)

	// Every time we passed a `Scope` to an operation, we placed that
	// operation **under** that scope.
	// As you can see, we have an empty scope (created with NewScope): the empty scope
	// is the root of our graph and thus we denote it with "/".

	// Now we ask tensorflow to build the graph from our definition.
	// The concrete graph is created from the "abstract" graph we defined
	// using the combination of scope and op.

	graph, err := root.Finalize()
	if err != nil {
		// It's useless trying to handle this error in any way:
		// if we defined the graph wrongly we have to manually fix the definition.

		// It's like a SQL query: if the query is not syntactically valid
		// we have to rewrite it
		panic(err.Error())
	}

	// If here: our graph is syntatically valid.
	// We can now place it within a Session and execute it.

	var sess *tf.Session
	sess, err = tf.NewSession(graph, &tf.SessionOptions{})
	if err != nil {
		panic(err.Error())
	}

	// In order to use placeholders, we have to create the Tensors
	// containing the values to feed into the network
	var matrix, column *tf.Tensor

	// A = [ [1, 2], [-1, -2] ]
	if matrix, err = tf.NewTensor([2][2]complex128{{1, 2}, {-1, -2}}); err != nil {
		panic(err.Error())
	}
	// x = [ [10], [100] ]
	if column, err = tf.NewTensor([2][1]complex128{{10}, {100}}); err != nil {
		panic(err.Error())
	}

	var results []*tf.Tensor
	if results, err = sess.Run(map[tf.Output]*tf.Tensor{
		A: matrix,
		x: column,
	}, []tf.Output{product}, nil); err != nil {
		panic(err.Error())
	}
	for _, result := range results {
		fmt.Println(result.Value().([][]complex128))
	}
}

func main() {
	//client()
	learn()

	//testTF()
}
