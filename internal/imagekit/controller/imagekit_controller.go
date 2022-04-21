package imagekitControllerPkg

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codedius/imagekit-go"
	"github.com/gorilla/mux"

	imagekitServicePkg "github.com/kuroyamii/go-backend-learn/internal/imagekit/service"
	"github.com/kuroyamii/go-backend-learn/pkg/entity/response"
)

type ImagekitController struct {
	router *mux.Router
	is     imagekitServicePkg.ImagekitService
}

func (ic *ImagekitController) UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")

	if err != nil {
		fmt.Println("Error Retrieving File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error converting file: ", err)
	}
	data, err := ic.is.UploadImage(r.Context(), fileBytes, handler.Filename)
	fmt.Println(data)
}

func ProvideImagekitController(router *mux.Router, is imagekitServicePkg.ImagekitService) ImagekitController {
	return ImagekitController{
		router: router,
		is:     is,
	}
}

func (ic ImagekitController) GetImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	opts := imagekit.Options{
		PublicKey:  "public_6wZ5GF54GznOEsbAMpCICw1Z3pg=",
		PrivateKey: "private_KTZgYb5hlB1Su1VADcFHIdFllYw=",
	}
	fmt.Println("Success making options")

	ik, err := imagekit.NewClient(&opts)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Success making Client")
	er := imagekit.ListAndSearchFileRequest{
		Path:  "my_image",
		Limit: 10,
	}
	fmt.Println("Success making Request")

	list, err := ik.Media.ListAndSearchFile(r.Context(), &er)
	if err != nil {
		fmt.Println("error getting files", err)
		return
	}
	fmt.Println("Success Getting Files")
	response.NewBaseResponse(200, response.RESPONSE_SUCCESS_MESSAGE, nil, list).ToJSON(w)
	fmt.Println("Success Writing JSON")

}

func (ic ImagekitController) InitEndpoints() {
	ic.router.HandleFunc("/upload", ic.UploadFile).Methods("POST")
	ic.router.HandleFunc("/download", ic.GetImage).Methods("GET")
}
