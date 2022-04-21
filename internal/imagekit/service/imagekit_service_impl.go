package imagekitServicePkg

import (
	"context"
	"fmt"
	"log"

	"github.com/codedius/imagekit-go"
)

type imagekitServiceImpl struct {
	imageKitClient *imagekit.Client
}

func ProvideImagekitService(imagekit *imagekit.Client) imagekitServiceImpl {
	return imagekitServiceImpl{imageKitClient: imagekit}
}

func CreateNewClient(publicKey string, privateKey string) (*imagekit.Client, error) {
	opts := imagekit.Options{
		PublicKey:  "public_6wZ5GF54GznOEsbAMpCICw1Z3pg=",
		PrivateKey: "private_KTZgYb5hlB1Su1VADcFHIdFllYw=",
	}

	imageKit, err := imagekit.NewClient(&opts)
	if err != nil {
		log.Println("ERROR (ImageKit): Error while creating imagekit client options")
		return imageKit, err
	}
	return imageKit, nil
}

func (is *imagekitServiceImpl) UploadImage(ctx context.Context, theFile []byte, fileName string) (bool, error) {
	ur := imagekit.UploadRequest{
		File:              theFile,
		FileName:          fileName,
		UseUniqueFileName: false,
		Tags:              []string{},
		Folder:            "/",
		IsPrivateFile:     false,
		CustomCoordinates: "",
		ResponseFields:    nil,
	}

	upr, err := is.imageKitClient.Upload.ServerUpload(ctx, &ur)
	if err != nil {
		fmt.Println(upr)
	}
	return false, nil
}

func (is *imagekitServiceImpl) GetImage(ctx context.Context, imageName string) (bool, error) {

	return false, nil
}
