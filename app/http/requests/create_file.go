package requests

import (
	"mime/multipart"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type CreateFile struct {
	UploadID    string                `form:"upload_id"`
	ChunkIdx    int                   `form:"chunk_index"`
	File        *multipart.FileHeader `form:"chunk"`
	Filename    string                `form:"filename"`
	TotalChunks int                   `form:"total_chunks"`
}

func (r *CreateFile) Authorize(ctx http.Context) error {
	return nil
}

func (r *CreateFile) Rules(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateFile) Messages(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateFile) Attributes(ctx http.Context) map[string]string {
	return map[string]string{}
}

func (r *CreateFile) PrepareForValidation(ctx http.Context, data validation.Data) error {
	return nil
}
