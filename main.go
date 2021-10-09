package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userd struct {
	id       string `json:"id,omitempty" bson:"id,omitempty"`
	name     string `json:"name,omitempty" bson:"name,omitempty"`
	email    string `json:"email,omitempty" bson:"email,omitempty"`
	password string `json:"password,omitempty" bson:"password,omitempty"`
}

type postd struct {
	id        string `json:"id,omitempty" bson:"id,omitempty"`
	caption   string `json:"caption,omitempty" bson:"caption,omitempty"`
	img       string `json:"img,omitempty" bson:"img,omitempty"`
	timestamp string `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

var connection_string string = ""

func makeconnection() {
	client, err := mongo.NewClient(options.Client().ApplyURI(connection_string))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	fmt.Println("Connection made")
}

func adduser(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		response.Header().Add("content-type", "application/json")
		client, err := mongo.NewClient(options.Client().ApplyURI(connection_string))
		_ = err
		if err != nil {
			log.Fatal(err)
		}
		insta := client.Database("InstagramAPI")
		userscol := insta.Collection("users")
		var user userd
		json.NewDecoder(request.Body).Decode(&user)
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)
		res, err := userscol.InsertOne(ctx, user)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(response).Encode(res)
		_ = res
	} else {
		fmt.Printf(request.Method)
	}
}

func getuser(response http.ResponseWriter, request *http.Request) {
	userid := request.URL.Query().Get("id")
	var id, _ = primitive.ObjectIDFromHex(userid)
	response.Header().Add("content-type", "application/json")
	client, err := mongo.NewClient(options.Client().ApplyURI(connection_string))
	_ = err
	if err != nil {
		log.Fatal(err)
	}
	insta := client.Database("InstagramAPI")
	userscol := insta.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	userdetails, err := userscol.Find(ctx, bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(response).Encode(userdetails)
}

func addpost(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		response.Header().Add("content-type", "application/json")
		client, err := mongo.NewClient(options.Client().ApplyURI(connection_string))
		_ = err
		if err != nil {
			log.Fatal(err)
		}
		insta := client.Database("InstagramAPI")
		postcol := insta.Collection("posts")
		var post postd
		json.NewDecoder(request.Body).Decode(&post)
		ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)
		res, err := postcol.InsertOne(ctx, post)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(response).Encode(res)
		_ = res
	} else {
		fmt.Printf(request.Method)
	}
}

func listpost(response http.ResponseWriter, request *http.Request) {
	userid := request.URL.Query().Get("id")
	response.Header().Add("content-type", "application/json")
	client, err := mongo.NewClient(options.Client().ApplyURI(connection_string))
	_ = err
	if err != nil {
		log.Fatal(err)
	}
	insta := client.Database("InstagramAPI")
	postcol := insta.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	userdetails, err := postcol.Find(ctx, bson.M{"id": userid})
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(response).Encode(userdetails)
}

func getpost(response http.ResponseWriter, request *http.Request) {
	postid := request.URL.Query().Get("id")
	var id, _ = primitive.ObjectIDFromHex(postid)
	response.Header().Add("content-type", "application/json")
	client, err := mongo.NewClient(options.Client().ApplyURI(connection_string))
	_ = err
	if err != nil {
		log.Fatal(err)
	}
	insta := client.Database("InstagramAPI")
	postcol := insta.Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	postdetails, err := postcol.Find(ctx, bson.M{"id": id})
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(response).Encode(postdetails)
}

func main() {
	makeconnection()
	http.HandleFunc("/user", adduser)
	http.HandleFunc("/user/", getuser)
	http.HandleFunc("/posts", addpost)
	http.HandleFunc("/posts/", getpost)
	http.HandleFunc("/posts/users/", listpost)
	http.ListenAndServe(":5000", nil)
}
