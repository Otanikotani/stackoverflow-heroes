package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var payload = []Question{{
	IsAnswered:       false,
	ViewCount:        123,
	AnswerCount:      1,
	Score:            22,
	LastActivityDate: 1594888088,
	CreationDate:     1594888088,
	QuestionID:       62930607,
	Link:             "",
	Title:            "How to allow a Lambda Function",
	Owner: ShallowUser{
		Reputation:  3230,
		UserID:      37867,
		DisplayName: "zitterion",
		Link:        "https://stackoverflow.com/users/37867/zitterion",
	},
	Answers: []Answer{{
		AnswerID:     62931544,
		CreationDate: 1594891505,
		IsAccepted:   false,
		Owner: ShallowUser{
			Reputation:  111,
			UserID:      5382111,
			DisplayName: "neyhar",
			Link:        "https://stackoverflow.com/users/5382111/neyhar",
		},
		Score: 0,
		Title: "This is how",
	},
	}},
	{
		IsAnswered:       true,
		ViewCount:        22,
		AnswerCount:      2,
		Score:            202,
		LastActivityDate: 1594764629,
		CreationDate:     1593627036,
		QuestionID:       62682939,
		Link:             "",
		Title:            "",
		Owner: ShallowUser{
			Reputation:  943,
			UserID:      1464514,
			DisplayName: "magne",
			Link:        "https://stackoverflow.com/users/1464514/magne",
		},
		Answers: []Answer{{
			AnswerID:     62683507,
			CreationDate: 1593629259,
			IsAccepted:   false,
			Owner: ShallowUser{
				Reputation:  517,
				UserID:      9175871,
				DisplayName: "saxa",
				Link:        "https://stackoverflow.com/users/9175871/saxa",
			},
			Score: 0,
			Title: "Serverless solution for long running query on AWS",
		},
			{
				AnswerID:     62883758,
				CreationDate: 1594672012,
				IsAccepted:   true,
				Owner: ShallowUser{
					Reputation:  1153,
					UserID:      4903466,
					DisplayName: "inge arda",
					Link:        "https://stackoverflow.com/users/4903467/inge",
				},
				Score: 6,
				Title: "How to deal with large dependencies in AWS Lambda?",
			},
		},
	}}

func TestQuestionConversion(t *testing.T) {
	fmt.Println("Test Question to csv vertices conversion")
	result := questionVertices(&payload)
	if len(result) != 3 {
		t.Errorf("Expected %d rows, but got %d", 3, len(result))
	}
	assert.Equal(t,
		[]string{"~id", "~label", "title:String", "viewCount:Int", "answerCount:Int", "score:Int", "isAnswered:Bool", "creationDate:Int"},
		result[0],
		"Invalid header")
	assert.Contains(t,
		result,
		[]string{"q62682939", "Question", "", "22", "2", "202", "true", "1593627036"},
		"Can't find expected question record")
}

func TestAnswerConversion(t *testing.T) {
	fmt.Println("Test Answer to csv vertices conversion")
	result := answerVertices(&payload)
	if len(result) != 4 {
		t.Errorf("Expected %d rows, but got %d", 4, len(result))
	}
	assert.Equal(t, []string{"~id", "~label", "title:String", "accepted:Bool", "score:Int", "creationDate:Int"},
		result[0],
		"Invalid header")
	assert.Contains(t,
		result,
		[]string{"a62683507", "Answer", "Serverless solution for long running query on AWS", "false", "0", "1593629259"},
		"Can't find expected answer record")
}

func TestPeopleConversion(t *testing.T) {
	fmt.Println("Test People to csv vertices conversion")
	result := peopleVertices(&payload)
	if len(result) != 6 {
		t.Errorf("Expected %d rows, but got %d", 6, len(result))
	}
	assert.Equal(t, []string{"~id", "~label", "title:DisplayName", "reputation:Int"},
		result[0],
		"Invalid header")
	assert.Contains(t,
		result,
		[]string{"u5382111", "Person", "neyhar", "111"},
		"Invalid second row")
}

func TestEdgeConversion(t *testing.T) {
	fmt.Println("Test csv edge conversion")
	result := edges(&payload)
	if len(result) != 6 {
		t.Errorf("Expected %d rows, but got %d", 6, len(result))
	}
	assert.Equal(t, []string{"~id", "~from", "~to", "~label"},
		result[0],
		"Invalid header")
	assert.Contains(t,
		result,
		[]string{"e3", "u1464514", "q62682939", "Asked"},
		"Can't find expected 'asked' edge")
	assert.Contains(t,
		result,
		[]string{"e5", "u4903466", "a62883758", "Answered"},
		"Can't find expected 'answered' edge")
}
