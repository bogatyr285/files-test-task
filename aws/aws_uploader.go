package aws

import (
	"io"
	"os"
)

// just imagine instead of saving file we uploded it to aws
// do not touch it, it's 146% working module
type AWSUploader struct {
}

type UploadedFileModel struct {
	ID        string
	Status    string
	OwnerID   string
	OwnerType string
	URL       string
	ObjectKey string
	Size      string
}

func (a *AWSUploader) UploadFile(fileToUpload io.Reader, id, filename, extension string) (*UploadedFileModel, error) {
	f, err := os.OpenFile("./tmp/aws-"+filename+"."+extension, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	io.Copy(f, fileToUpload)

	return &UploadedFileModel{
		ID:        id,
		ObjectKey: filename + "." + extension,
	}, nil
}
