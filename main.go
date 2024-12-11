package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Question struct {
    ID             string `json:"id" bson:"id"`
    Category       string `json:"category" bson:"category"`
    PlaceHolderCpp string `json:"placeHolderCpp" bson:"placeHolderCpp"`
    PlaceHolderGo  string `json:"placeHolderGo" bson:"placeHolderGo"`
    SpaceComplexity string `json:"spaceComplexity" bson:"spaceComplexity"`
    TimeComplexity  string `json:"timeComplexity" bson:"timeComplexity"`
    Difficulty      string `json:"difficulty" bson:"difficulty"`
    Question        string `json:"question" bson:"question"`
    AnswerImage     string `json:"answerImage" bson:"answerImage"`
    AnswerCpp       string `json:"answerCpp" bson:"answerCpp"`
    AnswerGo        string `json:"answerGo" bson:"answerGo"`
}


type AllQuestionsData struct {
	Programming []Question 	   `json:"Programming"`
	SystemDesign []interface{} `json:"SystemDesign" bson:"SystemDesign"`
}

func getQuestionsFromDB() ([]Question, error) {
    const uri = "mongodb://mongodb:27017"

    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatalf("Failed to create MongoDB client: %v", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := client.Connect(ctx); err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }
    defer client.Disconnect(ctx)

    database := client.Database("questionDB")
    collection := database.Collection("programmingQuestions")

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        log.Fatalf("Failed to execute find query: %v", err)
    }
    defer cursor.Close(ctx)

	var questions []Question

    for cursor.Next(ctx) {
        var doc AllQuestionsData
        if err := cursor.Decode(&doc); err != nil {
            log.Printf("Failed to decode document: %v", err)
            continue
        }

		for _, question := range doc.Programming {
			questions = append(questions, question)
        }
    }

    if err := cursor.Err(); err != nil {
        log.Fatalf("Cursor encountered an error: %v", err)
    }

	return questions, nil
}

func main() {
	http.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		questions, err := getQuestionsFromDB()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to fetch questions: %v", err), http.StatusInternalServerError)
			return
		}

		response := AllQuestionsData{
			Programming: questions,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
	})

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
