package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"google.golang.org/grpc"
	"log"

	"msisdn/pkg/pb"
	"msisdn/pkg/producer"
	"net/http"
)

type JsonMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func getMsisdn(
	render *render.Render,
	producer producer.MsisdnProducer,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		vars := mux.Vars(r)

		msisdnInput, _ := vars["input"]

		resp, err := producer.GetMsisdnDetails(ctx, msisdnInput)
		if err != nil {

			err = render.JSON(w, http.StatusOK, JsonMessage{
				Message: err.Error(),
				Code:    http.StatusUnauthorized,
			})

			if err!=nil{
				fmt.Println("failed to render json response")
			}

			return
		}

		err = render.JSON(w, http.StatusOK, resp)
		if err!=nil{
			fmt.Println("failed to render json response")
		}
		return

	})
}

func main() {

	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMsisdnServiceClient(conn)

	msisdnProducer := producer.NewMsisdnProducer(c)

	r := render.New()
	router := mux.NewRouter()

	router.Methods(http.MethodGet).Path("/msisdn/{input:[0-9]+}").Handler(getMsisdn(r, msisdnProducer))

	http.ListenAndServe(":1212", router)

}

