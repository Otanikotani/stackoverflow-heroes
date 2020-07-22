package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

//Questions from Stack Overflow with paging
type Questions struct {
	Items          []Question `json:"items"`
	HasMore        bool       `json:"has_more"`
	QuotaMax       int        `json:"quota_max"`
	QuotaRemaining int        `json:"quota_remaining"`
}

//Question describes a question asked in SO, the fields we are interested in only.
type Question struct {
	QuestionID       int         `json:"question_id"`
	Title            string      `json:"title"`
	IsAnswered       bool        `json:"is_answered"`
	ViewCount        int         `json:"view_count"`
	AnswerCount      int         `json:"answer_count"`
	Score            int         `json:"score"`
	LastActivityDate int         `json:"last_activity_date"`
	CreationDate     int         `json:"creation_date"`
	Link             string      `json:"link"`
	Owner            ShallowUser `json:"owner"`
	Answers          []Answer    `json:"answers"`
}

//ShallowUser brief info on a SO user
type ShallowUser struct {
	UserID      int    `json:"user_id"`
	Reputation  int    `json:"reputation"`
	DisplayName string `json:"display_name"`
	Link        string `json:"link"`
}

//Answer describes answer to an SO question
type Answer struct {
	AnswerID     int         `json:"answer_id"`
	Title        string      `json:"title"`
	CreationDate int         `json:"creation_date"`
	IsAccepted   bool        `json:"is_accepted"`
	Owner        ShallowUser `json:"owner"`
	Score        int         `json:"score"`
}

func getQuestions(accessToken string, key string) (*[]Question, error) {
	var result []Question
	page := 1
	hasMore := true
	for hasMore {
		questionsPage, err := getQuestionsPage(accessToken, key, page)
		if err != nil {
			return nil, err
		}
		log.Printf("Received page %d with %d items. Quota max: %d, remaining: %d\n",
			page, len(questionsPage.Items), questionsPage.QuotaMax, questionsPage.QuotaRemaining)
		result = append(result, questionsPage.Items...)
		page++
		hasMore = questionsPage.HasMore
		if questionsPage.QuotaRemaining <= 0 {
			fmt.Printf("Quota exceeded after retrieving %d items\n", len(result))
			break
		}
	}
	return &result, nil
}

func getQuestionsPage(accessToken string, key string, page int) (*Questions, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.stackexchange.com/2.2/questions?page=%d&pagesize=%d&order=desc&sort=activity&tagged=%v&site=%v&filter=%v&key=%v",
		page,
		100,
		"serverless",
		"stackoverflow",
		")v)boUYhKuVjzx1W___XpJX0VuTxdbR_1fbxpeOV",
		key)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer: "+accessToken)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Invalid response " + resp.Status)
	}

	var questions Questions
	err = json.NewDecoder(resp.Body).Decode(&questions)
	if err != nil {
		return nil, err
	}

	return &questions, nil
}
