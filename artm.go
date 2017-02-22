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
func ArtmCreateMasterModel(config *MasterModelConfig) (int, error) {
	message, err := proto.Marshal(config)
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
func ArtmImportModel(masterModelID int, config *ImportModelArgs) error {
	message, err := proto.Marshal(config)
	if err != nil {
		return fmt.Errorf("Protobuf ImportModelArgs marshaling error: %s", err)
	}
	errorID := C.ArtmImportModel(C.int(masterModelID), C.int64_t(len(message)), C.CString(string(message)))
	if errorID < 0 {
		fmt.Printf("Load model error: %s\n", ARTM_ERRORS[-errorID])
	}
	errorStr := C.GoString(C.ArtmGetLastErrorMessage())
	if len(errorStr) > 0 {
		return fmt.Errorf("Load model error: %s\n", errorStr)
	}
	return nil
}

//ArtmDisposeMasterComponent disposes master component
func ArtmDisposeMasterComponent(masterModelID int) error {
	errorID := C.ArtmDisposeMasterComponent(C.int(masterModelID))
	if errorID < 0 {
		fmt.Errorf("Dispose model error %s\n", ARTM_ERRORS[-errorID])
	}
	return nil
}

//ArtmRequestDictionary request dictionary
func ArtmRequestDictionary(masterModelID int, conf *GetDictionaryArgs) (*DictionaryData, error) {

	message, err := proto.Marshal(conf)
	if err != nil {
		return nil, fmt.Errorf("Protobuf GetDictionaryArgs marshaling error: %s", err)
	}
	messageLength := C.ArtmRequestDictionary(C.int(masterModelID), C.int64_t(len(message)), C.CString(string(message)))
	errorStr := C.GoString(C.ArtmGetLastErrorMessage())
	if len(errorStr) > 0 {
		return nil, fmt.Errorf("Get model dictionary error: %s\n", errorStr)
	}
	if messageLength < 0 {
		return nil, fmt.Errorf("Get model dictionary error: %s\n", ARTM_ERRORS[-messageLength])
	}
	buffer := make([]byte, messageLength)
	errorID := C.ArtmCopyRequestedMessage(messageLength, C.CString(string(buffer)))
	if len(errorStr) > 0 {
		return nil, fmt.Errorf("Get model dictionary error: %s\n", errorStr)
	}
	if errorID < 0 {
		return nil, fmt.Errorf("Get dictionary data error: %s\n", ARTM_ERRORS[-errorID])
	}

	dictionaryData := &DictionaryData{}
	err = proto.Unmarshal(buffer, dictionaryData)
	if err != nil {
		return nil, fmt.Errorf("protobuf DictionaryData unmarshaling error: ", err)
	}
	return dictionaryData, nil
}
