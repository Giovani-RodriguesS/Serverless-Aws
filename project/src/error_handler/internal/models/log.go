package models

type Log struct {
	ID        string `dynamodbav:"ID"`
	Timestamp string `dynamodbav:"Timestamp"`
	Level     string `dynamodbav:"Level"`
	Message   string `dynamodbav:"Message"`
	Cause     string `dynamodbav:"Cause"`
}
