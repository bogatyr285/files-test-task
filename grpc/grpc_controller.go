package grpc

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/bogatyr285/test_task/aws"
	pb "github.com/bogatyr285/test_task/proto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCFilesAPIController struct {
	pb.UnimplementedFilesAPIServer
	awsUploader IAWSUploader
	tmpFilesDir string
}

func NewGRPCFilesAPIController(awsUploader IAWSUploader, tmpFilesDir string) *GRPCFilesAPIController {
	return &GRPCFilesAPIController{
		awsUploader: awsUploader,
		tmpFilesDir: tmpFilesDir,
	}
}

type IAWSUploader interface {
	UploadFile(fileToUpload io.Reader, id, filename, extension string) (*aws.UploadedFileModel, error)
}

func (c *GRPCFilesAPIController) UploadFile(stream pb.FilesAPI_UploadFileServer) error {
	fileReq, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive file info")
	}

	fileInfo := fileReq.GetInfo()
	ownerID := fileInfo.OwnerId
	ownerType := fileInfo.OwnerType
	if ownerID == "" || ownerType == "" {
		return status.Errorf(codes.InvalidArgument, "owner_id and owner_type must be provided")
	}

	// open output file
	tmpFilename := fmt.Sprintf("files-api_*.%s", fileInfo.FileExtension)
	tmpFile, err := ioutil.TempFile(c.tmpFilesDir, tmpFilename)
	if err != nil {
		return status.Error(codes.Internal, "failed to create file")
	}

	log.Debug().Interface("tmpFile.Name()", tmpFile.Name()).Msg("created file for client req")

	// save file
	if saveErr := c.saveFileFromStream(tmpFile, stream); saveErr != nil {
		log.Error().Err(saveErr).Msg("saveFileFromStream")

		return status.Error(codes.Internal, "stream reading err")
	}

	fileDBm, err := c.awsUploader.UploadFile(
		tmpFile,
		"some_id",
		fileInfo.Filename,
		fileInfo.FileExtension,
	)
	if err != nil {
		return status.Error(codes.Internal, "uploading file err")
	}

	if err := stream.SendAndClose(&pb.File{
		Id:        fileDBm.ID,
		Status:    fileDBm.Status,
		OwnerId:   fileDBm.OwnerID,
		OwnerType: fileDBm.OwnerType,
		Url:       fileDBm.URL,
		ObjectKey: fileDBm.ObjectKey,
	}); err != nil {
		return status.Error(codes.Internal, "sending response err")
	}

	return nil
}

func (c *GRPCFilesAPIController) getFileSize(file *os.File) (int64, error) {
	fileStat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return fileStat.Size(), nil
}
func (c *GRPCFilesAPIController) saveFileFromStream(
	file *os.File,
	stream pb.FilesAPI_UploadFileServer,
) error {
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			return errors.Wrap(err, "stream reading err")
		}

		if _, err := file.Write(res.GetChunkData()); err != nil {
			return errors.Wrap(err, "saving file err")
		}
	}

	return nil
}

func (s *GRPCFilesAPIController) GetFiles(ctx context.Context, in *pb.GetFilesRequest) (*pb.FilesList, error) {
	return nil, errors.New("unimplemented")
}
func (s *GRPCFilesAPIController) UpdateFilesStatuses(ctx context.Context, in *pb.UpdateFilesStatusesRequest) (*pb.StatusResponse, error) {
	return nil, errors.New("unimplemented")
}
