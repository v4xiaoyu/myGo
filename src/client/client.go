package main

import (
	"../utils"
	"bufio"
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
	"net"
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

func main() {
	client()
	learn()

	conn, err := net.Dial("tcp", ":18888")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(conn, "hello server\n")
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", data)
}
