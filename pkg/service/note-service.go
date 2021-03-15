package service

import (
	"context"
	entity "github.com/aleale2121/Golang-TODO-Hex-DDD/internal/constant/model"
	protos "github.com/aleale2121/Golang-TODO-Hex-DDD/internal/grpc/note"
	noteServ "github.com/aleale2121/Golang-TODO-Hex-DDD/internal/module/user"
)

type noteServiceServer struct {
	service noteServ.UseCase
}

func NewNoteServer(service noteServ.UseCase) protos.NoteServiceServer {
	return &noteServiceServer{service: service}
}
func (n noteServiceServer) CreateNote(ctx context.Context, request *protos.CreateNoteRequest) (*protos.CreateNoteResponse, error) {
	noteCreated, err := n.service.AddNote(*ConvertProtoNoteToNote(request.Note))
	if len(err) > 0 {
		return nil, err[0]
	}
	return &protos.CreateNoteResponse{Id: uint32(noteCreated.ID)}, nil
}

func (n noteServiceServer) GetNote(ctx context.Context, request *protos.GetNoteRequest) (*protos.GetNoteResponse, error) {
	id := request.Id
	note, errs := n.service.FindNoteByID(uint(id))
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return &protos.GetNoteResponse{Note: ConvertToProtoNote(note)}, nil
}

func (n noteServiceServer) ListNotes(ctx context.Context, request *protos.ListNoteRequest) (*protos.ListNoteResponse, error) {
	notes, errs := n.service.GetAllNotes()
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return &protos.ListNoteResponse{
		Notes: ConvertToProtoNotes(notes),
	}, nil
}

func (n noteServiceServer) UpdateNote(ctx context.Context, request *protos.UpdateNoteRequest) (*protos.UpdateNoteResponse, error) {
	note := ConvertProtoNoteToNote(request.Note)
	_, errs := n.service.UpdateNote(*note)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return &protos.UpdateNoteResponse{}, nil
}

func (n noteServiceServer) DeleteNote(ctx context.Context, request *protos.DeleteNoteRequest) (*protos.DeleteNoteResponse, error) {
	_, errs := n.service.DeleteNote(uint(request.Id))
	if len(errs) > 0 {
		return nil, errs[0]
	}
	return &protos.DeleteNoteResponse{}, nil
}

func ConvertToProtoNotes(notes []entity.Note) []*protos.Note {
	var protoNotes []*protos.Note
	for i := 0; i < len(notes); i++ {
		pNote := protos.Note{
			Id:     uint32(notes[i].ID),
			Title:  notes[i].Title,
			Detail: notes[i].Detail,
		}
		protoNotes = append(protoNotes, &pNote)
	}
	return protoNotes
}
func ConvertProtoNoteToNote(note *protos.Note) *entity.Note {
	return &entity.Note{
		ID:     uint(note.Id),
		Title:  note.Title,
		Detail: note.Detail,
	}
}
func ConvertToProtoNote(note *entity.Note) *protos.Note {
	return &protos.Note{
		Id:     uint32(note.ID),
		Title:  note.Title,
		Detail: note.Detail,
	}
}
