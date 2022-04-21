package imagekitServicePkg

import "context"

type ImagekitService interface {
	UploadImage(ctx context.Context, theFile []byte, fileName string) (bool, error)
	GetImage(ctx context.Context, imageName string) (bool, error)
}
