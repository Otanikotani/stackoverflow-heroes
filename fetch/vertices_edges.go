package main

import (
	"strconv"
)

func questionCsv(question *Question) []string {
	return []string{
		questionVerticeID(question),
		"Question",
		question.Title,
		strconv.Itoa(question.ViewCount),
		strconv.Itoa(question.AnswerCount),
		strconv.Itoa(question.Score),
		strconv.FormatBool(question.IsAnswered),
		strconv.Itoa(question.CreationDate),
	}
}

func answerCsv(answer *Answer) []string {
	return []string{
		answerVerticeID(answer),
		"Answer",
		answer.Title,
		strconv.FormatBool(answer.IsAccepted),
		strconv.Itoa(answer.Score),
		strconv.Itoa(answer.CreationDate),
	}
}

func shallowUserCsv(shallowUser *ShallowUser) []string {
	return []string{
		shallowUserVerticeID(shallowUser),
		"Person",
		shallowUser.DisplayName,
		strconv.Itoa(shallowUser.Reputation),
	}
}

func shallowUserVerticeID(shallowUser *ShallowUser) string {
	return "u" + strconv.Itoa(shallowUser.UserID)
}

func answerVerticeID(answer *Answer) string {
	return "a" + strconv.Itoa(answer.AnswerID)
}

func questionVerticeID(question *Question) string {
	return "q" + strconv.Itoa(question.QuestionID)
}

func questionVertices(questions *[]Question) [][]string {
	var questionVertices [][]string
	questionVerticesHeader := []string{"~id", "~label", "title:String", "viewCount:Int", "answerCount:Int", "score:Int", "isAnswered:Bool", "creationDate:Int"}
	questionVertices = append(questionVertices, questionVerticesHeader)

	for _, question := range *questions {
		questionCsvRow := questionCsv(&question)
		questionVertices = append(questionVertices, questionCsvRow)
	}
	return questionVertices
}

func answerVertices(questions *[]Question) [][]string {
	var answerVertices [][]string
	answerVerticesHeader := []string{"~id", "~label", "title:String", "accepted:Bool", "score:Int", "creationDate:Int"}
	answerVertices = append(answerVertices, answerVerticesHeader)

	for _, question := range *questions {
		for _, answer := range question.Answers {
			answerCsvRow := answerCsv(&answer)
			answerVertices = append(answerVertices, answerCsvRow)
		}
	}
	return answerVertices
}

func peopleVertices(questions *[]Question) [][]string {
	var peopleVertices [][]string
	peopleVerticesHeader := []string{"~id", "~label", "title:DisplayName", "reputation:Int"}
	peopleVertices = append(peopleVertices, peopleVerticesHeader)

	peopleByIds := make(map[int]ShallowUser)

	for _, question := range *questions {
		peopleByIds[question.Owner.UserID] = question.Owner
		for _, answer := range question.Answers {
			peopleByIds[answer.Owner.UserID] = answer.Owner
		}
	}

	for _, user := range peopleByIds {
		userCsvRow := shallowUserCsv(&user)
		peopleVertices = append(peopleVertices, userCsvRow)
	}

	return peopleVertices
}

func edges(questions *[]Question) [][]string {
	var edges [][]string
	edgeHeader := []string{"~id", "~from", "~to", "~label"}
	edges = append(edges, edgeHeader)
	edgeCount := 0

	for _, question := range *questions {
		edgeCount++
		edgeID := "e" + strconv.Itoa(edgeCount)
		askedEdge := []string{edgeID, shallowUserVerticeID(&question.Owner), questionVerticeID(&question), "Asked"}
		edges = append(edges, askedEdge)

		for _, answer := range question.Answers {
			edgeCount++
			edgeID := "e" + strconv.Itoa(edgeCount)
			answerEdge := []string{edgeID, shallowUserVerticeID(&answer.Owner), answerVerticeID(&answer), "Answered"}
			edges = append(edges, answerEdge)
		}
	}

	return edges
}

func toVerticesAndEdges(questions *[]Question) ([][]string, [][]string, [][]string, [][]string) {
	return questionVertices(questions), answerVertices(questions), peopleVertices(questions), edges(questions)
}
