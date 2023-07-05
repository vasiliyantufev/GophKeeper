package handlers

import (
	"io"

	grpc "github.com/vasiliyantufev/gophkeeper/internal/server/proto"
	"github.com/vasiliyantufev/gophkeeper/internal/server/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handler) Upload(stream grpc.Gophkeeper_UploadServer) error {
	name := "some-unique-name.png"
	/// := "some-unique-name.png"
	file := storage.NewFile(name)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if err := h.storage.Store(file); err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			return stream.SendAndClose(&grpc.UploadResponse{Name: name})
		}
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err := file.Write(req.GetChunk()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
	return nil
}
