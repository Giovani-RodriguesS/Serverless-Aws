package models

type Log struct {
	ID        string `dynamodbav:"id"`
	Timestamp string `dynamodbav:"name"`
	Level     string `dynamodbav:"level"`
	Message   string `dynamodbav:"message"`
	Cause     string `dynamodbav:"cause"`
}
