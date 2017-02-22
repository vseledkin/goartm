package goartm

// #cgo LDFLAGS: libartm.dylib
// #include "c_interface.h"
import "C"
import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

var ARTM_ERRORS = []string{
	"ARTM_SUCCESS",
	"ARTM_STILL_WORKING",
	"ARTM_INTERNAL_ERROR",
	"ARTM_ARGUMENT_OUT_OF_RANGE",
	"ARTM_INVALID_MASTER_ID",
	"ARTM_CORRUPTED_MESSAGE",
	"ARTM_INVALID_OPERATION",
	"ARTM_DISK_READ_ERROR",
	"ARTM_DISK_WRITE_ERROR",
}

//ArtmCreateMasterModel create master model
func ArtmCreateMasterModel() (int, error) {
	masterConfig := &MasterModelConfig{}
	message, err := proto.Marshal(masterConfig)
	if err != nil {
		return 0, fmt.Errorf("Protobuf MasterModelConfig marshaling error: %s", err)
	}
	masterID := C.int(C.ArtmCreateMasterModel(C.int64_t(len(message)), C.CString(string(message))))
	if masterID < 0 {
		return 0, fmt.Errorf("Create Master error: %s", ARTM_ERRORS[-masterID])
	}
	return int(masterID), nil
}

//ArtmImportModel loads model from file ie matrices of size |T|*|W| topics/words
func ArtmImportModel(masterModelID int, modelFile string) error {
	importModelConfig := &ImportModelArgs{}
	importModelConfig.FileName = &modelFile
	message, err := proto.Marshal(importModelConfig)
	if err != nil {
		return fmt.Errorf("Protobuf ImportModelArgs marshaling error: %s", err)
	}
	errorID := C.ArtmImportModel(C.int(masterModelID), C.int64_t(len(message)), C.CString(string(message)))
	if errorID < 0 {
		fmt.Printf("Load model error %s\n", ARTM_ERRORS[-errorID])
	}
	return nil
}

/*
var MODEL_FILE string

func main() {
	flag.StringVar(&MODEL_FILE, "model", "", "artm model file")
	flag.Parse()

	if len(MODEL_FILE) == 0 {
		flag.Usage()
		return
	}
	masterConfig := &MasterModelConfig{}
	data, err := proto.Marshal(masterConfig)
	if err != nil {
		log.Fatal("protobuf marshaling error: ", err)
	}

	/*
			        newTest := &example.Test{}
		        err = proto.Unmarshal(data, newTest)
		        if err != nil {
		            log.Fatal("unmarshaling error: ", err)
		        }


	// create master object
	masterID := C.int(C.ArtmCreateMasterModel(C.int64_t(len(data)), C.CString(string(data))))
	if masterID < 0 {
		fmt.Printf("Master creation error %s ok\n", ARTM_ERRORS[-masterID])
	} else {
		fmt.Printf("Master created with id %d ok\n", masterID)
	}
	// load model

	importModelConfig := &ImportModelArgs{}
	importModelConfig.FileName = &MODEL_FILE
	data, err = proto.Marshal(importModelConfig)
	if err != nil {
		log.Fatal("protobuf marshaling error: ", err)
	}
	errorID := C.ArtmImportModel(masterID, C.int64_t(len(data)), C.CString(string(data)))
	fmt.Printf("Master load message %s\n", ARTM_ERRORS[-errorID])

	// destroy master component
	errorID = C.ArtmDisposeMasterComponent(masterID)
	fmt.Printf("Master destroy message %s\n", ARTM_ERRORS[-errorID])
}
*/
