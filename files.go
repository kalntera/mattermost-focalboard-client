// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package boards

import (
	"encoding/json"
	"io"

	model "github.com/mattermost/mattermost/server/public/model"

)

var UnsafeContentTypes = [...]string{
	"application/javascript",
	"application/ecmascript",
	"text/javascript",
	"text/ecmascript",
	"application/x-javascript",
	"text/html",
}

var MediaContentTypes = [...]string{
	"image/jpeg",
	"image/png",
	"image/bmp",
	"image/gif",
	"image/tiff",
	"video/avi",
	"video/mpeg",
	"video/mp4",
	"audio/mpeg",
	"audio/wav",
}

// FileUploadResponse is the response to a file upload
// swagger:model
type FileUploadResponse struct {
	// The FileID to retrieve the uploaded file
	// required: true
	FileID string `json:"fileId"`
}

func FileUploadResponseFromJSON(data io.Reader) (*FileUploadResponse, error) {
	var fileUploadResponse FileUploadResponse

	if err := json.NewDecoder(data).Decode(&fileUploadResponse); err != nil {
		return nil, err
	}
	return &fileUploadResponse, nil
}

func FileInfoResponseFromJSON(data io.Reader) (*model.FileInfo, error) {
	var fileInfo model.FileInfo

	if err := json.NewDecoder(data).Decode(&fileInfo); err != nil {
		return nil, err
	}
	return &fileInfo, nil
}
