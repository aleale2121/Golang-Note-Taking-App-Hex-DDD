package grpc_client

import (
	"context"
	protos "github.com/aleale2121/Golang-TODO-Hex-DDD/internal/grpc/note"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:9090")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewNoteServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	pfx := t.Format(time.RFC3339Nano)
	//Create Note
	requestCreate := protos.CreateNoteRequest{
		Note: &protos.Note{
			Id:     nil,
			Title:  "title (" + pfx + ")",
			Detail: "detail(" + pfx + ")",
		}}
	res1, err := c.CreateNote(ctx, &requestCreate)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	id := res1.Id

	requestRead := protos.GetNoteRequest{
		Id: id,
	}

	res2, err := c.GetNote(ctx, &requestRead)
	if err != nil {
		log.Fatalf("Get Note failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", res2)

	// Update
	req3 := protos.UpdateNoteRequest{
		Note: &protos.Note{
			Id:     res2.Note.Id,
			Title:  res2.Note.Title,
			Detail: res2.Note.Detail,
		},
	}
	res3, err := c.UpdateNote(ctx, &req3)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", res3)

	// Call List All Notes
	req4 := protos.ListNoteRequest{}
	res4, err := c.ListNotes(ctx, &req4)
	if err != nil {
		log.Fatalf("List All Notes failed: %v", err)
	}
	log.Printf("List All Notes result: <%+v>\n\n", res4)

	// Delete
	req5 := protos.DeleteNoteRequest{
		Id: res2.Note.Id,
	}
	res5, err := c.DeleteNote(ctx, &req5)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", res5)
}
