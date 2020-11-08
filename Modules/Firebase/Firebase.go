package Firebase

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
	database "firebase.google.com/go/db"
)

var ctx = context.Background()
var client *database.Client

func ConnectFirebase() {

	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASECREDENTIONAL")))

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}

	firebaseClient, err := app.DatabaseWithURL(ctx, "https://xbooks-hackathon.firebaseio.com/")
	if err != nil {
		log.Fatal(err)
	}

	client = firebaseClient
}

func WriteData(path string, data interface{}) error {
	err := client.NewRef(path).Set(ctx, data)
	return err

}

func PushData(path string, data interface{}) error {
	_, err := client.NewRef(path).Push(ctx, data)
	return err
}

func GetFilteredData(path string, child string, equal string) interface{} {
	var data interface{}
	err := client.NewRef(path).OrderByChild(child).EqualTo(equal).Get(ctx, &data)
	if err != nil {
		fmt.Println(err)
	}
	itemsMap := data.(map[string]interface{})

	var responseData interface{}
	for _, v := range itemsMap {
		responseData = v
		break
	}
	return responseData
}

func UpdateFilteredData(path string, child string, equal string, updatedData interface{}) error {
	var data interface{}
	err := client.NewRef(path).OrderByChild(child).EqualTo(equal).Get(ctx, &data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	itemsMap := data.(map[string]interface{})

	var dataParentName string
	for i, _ := range itemsMap {
		dataParentName = i
		break
	}

	if dataParentName == "" {
		fmt.Println("Durum boş")
		dataParentName = uuid.New().String()
	}

	newData := map[string]interface{}{
		dataParentName: updatedData,
	}

	fmt.Println(newData)
	err = client.NewRef(path).Update(ctx, newData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Delete(path string, child string, equal string) error {
	var data interface{}
	err := client.NewRef(path).OrderByChild(child).EqualTo(equal).Get(ctx, &data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	itemsMap := data.(map[string]interface{})

	var dataParentName string
	for i, _ := range itemsMap {
		dataParentName = i
		break
	}

	err = client.NewRef(path + "/" + dataParentName).Delete(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func ReadData(path string) interface{} {
	var data interface{}
	if err := client.NewRef(path).Get(ctx, &data); err != nil {
		log.Fatal(err)
	}
	return data
}
