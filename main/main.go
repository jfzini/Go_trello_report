package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	JSONPath       = "" // Path to the JSON file
	DoneColumnID   = "" // ID of the "Done" column
	ReportFileName = "report.txt"
)

type Action struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
	Date string                 `json:"date"`
}

type Card struct {
	ID     string  `json:"id"`
	Labels []Label `json:"labels"`
	Name   string  `json:"name"`
}

type Label struct {
	Color string `json:"color"`
	Name  string `json:"name"`
}

func filterActions(action Action) bool {
	endDate := time.Now().UTC()
	startDate := endDate.AddDate(0, 0, -7)

	if action.Type != "updateCard" && action.Type != "updateCheckItemStateOnCard" {
		return false
	}
	if _, exists := action.Data["checkItem"]; !exists {
		if _, exists := action.Data["listAfter"]; !exists {
			return false
		}
	}
	if checkItem, exists := action.Data["checkItem"].(map[string]interface{}); exists && checkItem["state"] != "complete" {
		return false
	}
	if listAfter, exists := action.Data["listAfter"].(map[string]interface{}); exists && listAfter["id"] != DoneColumnID {
		return false
	}
	actionDate, err := time.Parse("2006-01-02T15:04:05.000Z", action.Date)
	if err != nil {
		return false
	}
	if actionDate.Before(startDate) {
		return false
	}

	return true
}

func createReportDict(filteredActions []Action, cards []Card) map[string]map[string][]string {
	reportDict := make(map[string]map[string][]string)

	for _, action := range filteredActions {
		var labels []Label
		var cardName string
		var checkItem string

		for _, card := range cards {
			if card.ID == action.Data["card"].(map[string]interface{})["id"].(string) {
				labels = card.Labels
				cardName = card.Name
			}
		}

		label := "Outro"
		for _, l := range labels {
			if l.Color != "black_light" {
				label = l.Name
				break
			}
		}

		if checkItemData, exists := action.Data["checkItem"].(map[string]interface{}); exists {
			checkItem = checkItemData["name"].(string)
		}

		if _, exists := reportDict[label]; !exists {
			reportDict[label] = make(map[string][]string)
		}
		if _, exists := reportDict[label][cardName]; !exists {
			reportDict[label][cardName] = []string{}
		}
		if checkItem != "" {
			reportDict[label][cardName] = append(reportDict[label][cardName], checkItem)
		}
	}

	return reportDict
}

func createReportString(reportDict map[string]map[string][]string) string {
	var reportString string
	for label, cards := range reportDict {
		reportString += fmt.Sprintf("\n%s:\n", label)
		for card, checkItems := range cards {
			reportString += fmt.Sprintf("    - %s\n", card)
			for _, checkItem := range checkItems {
				reportString += fmt.Sprintf("        - %s\n", checkItem)
			}
		}
	}
	return reportString
}

func main() {
	dataBytes, err := os.ReadFile(JSONPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var data struct {
		Actions []Action `json:"actions"`
		Cards   []Card   `json:"cards"`
	}
	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	var filteredActions []Action
	for _, action := range data.Actions {
		if filterActions(action) {
			filteredActions = append(filteredActions, action)
		}
	}

	reportDict := createReportDict(filteredActions, data.Cards)
	reportString := createReportString(reportDict)

	err = os.WriteFile(ReportFileName, []byte(reportString), 0644)
	if err != nil {
		fmt.Println("Error writing report file:", err)
		return
	}

	fmt.Println("\n🚀 Arquivo criado com sucesso 🚀")
	fmt.Println("\n🚨 Lembre-se de conferir se o arquivo json está atualizado! 🚨")
	fmt.Println("\n🤔 Qualquer dúvida, leia o README 🤔")
	fmt.Println("\n😬\n🫶")
}
