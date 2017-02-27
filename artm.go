package goartm

// #cgo LDFLAGS: libartm.dylib
// #include <stdlib.h>
// #include "c_interface.h"
import "C"
import (
	"fmt"

	"unsafe"

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

func New(conf interface{}) interface{} {
	switch t := conf.(type) {
	case *MasterModelConfig:
		var MasterModelConfig_PwtName string = Default_MasterModelConfig_PwtName
		var MasterModelConfig_NwtName string = Default_MasterModelConfig_NwtName
		var MasterModelConfig_ReuseTheta bool = Default_MasterModelConfig_ReuseTheta
		var MasterModelConfig_OptForAvx bool = Default_MasterModelConfig_OptForAvx
		var MasterModelConfig_CacheTheta bool = Default_MasterModelConfig_CacheTheta

		t.NwtName = &MasterModelConfig_NwtName
		t.PwtName = &MasterModelConfig_PwtName
		t.ReuseTheta = &MasterModelConfig_ReuseTheta
		t.OptForAvx = &MasterModelConfig_OptForAvx
		t.CacheTheta = &MasterModelConfig_CacheTheta
		return t
	case *GetDictionaryArgs:
		return t
	case *ImportModelArgs:
		return t
	case *ImportDictionaryArgs:
		return t
	default:
		panic("unsupported type")
	}
}

func ArtmGetLastErrorMessage() error {
	err := C.ArtmGetLastErrorMessage()
	defer C.free(unsafe.Pointer(err))
	errorStr := C.GoString(err)
	if len(errorStr) > 0 {
		return fmt.Errorf("%s", errorStr)
	}
	return nil
}

//ArtmCreateMasterModel create master model
func ArtmCreateMasterModel(config *MasterModelConfig) (int, error) {
	message, err := proto.Marshal(config)
	if err != nil {
		return 0, fmt.Errorf("Protobuf MasterModelConfig marshaling error: %s", err)
	}
	masterID := C.int(C.ArtmCreateMasterModel(C.int64_t(len(message)), (*C.char)(unsafe.Pointer(&message[0]))))
	if masterID < 0 {
		return 0, fmt.Errorf("Create Master error: %s", ARTM_ERRORS[-masterID])
	}
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return 0, err
	}
	return int(masterID), nil
}

//ArtmImportModel loads model from file ie matrices of size |T|*|W| topics/words
func ArtmImportModel(masterModelID int, config *ImportModelArgs) error {
	message, err := proto.Marshal(config)
	if err != nil {
		return fmt.Errorf("Protobuf ImportModelArgs marshaling error: %s", err)
	}
	errorID := C.ArtmImportModel(C.int(masterModelID), C.int64_t(len(message)), (*C.char)(unsafe.Pointer(&message[0])))
	if errorID < 0 {
		fmt.Printf("Load model error: %s\n", ARTM_ERRORS[-errorID])
	}
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

//ArtmDisposeMasterComponent disposes master component
func ArtmDisposeMasterComponent(masterModelID int) error {
	errorID := C.ArtmDisposeMasterComponent(C.int(masterModelID))
	if errorID < 0 {
		fmt.Errorf("Dispose model error %s\n", ARTM_ERRORS[-errorID])
	}
	err := ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

func ArtmImportDictionary(masterModelID int, conf *ImportDictionaryArgs) error {
	message, err := proto.Marshal(conf)
	if err != nil {
		return fmt.Errorf("Protobuf ImportDictionaryArgs marshaling error: %s", err)
	}
	errorID := C.ArtmImportDictionary(C.int(masterModelID), C.int64_t(len(message)), (*C.char)(unsafe.Pointer(&message[0])))
	if errorID < 0 {
		fmt.Printf("Load dictionary error: %s\n", ARTM_ERRORS[-errorID])
	}
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

//ArtmRequestDictionary request dictionary
func ArtmRequestDictionary(masterModelID int, conf *GetDictionaryArgs) (*DictionaryData, error) {
	message, err := proto.Marshal(conf)
	if err != nil {
		return nil, fmt.Errorf("Protobuf GetDictionaryArgs marshaling error: %s", err)
	}
	p := unsafe.Pointer(&message[0])
	messageLength := C.ArtmRequestDictionary(C.int(masterModelID), C.int64_t(len(message)), (*C.char)(p))
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}

	if messageLength < 0 {
		return nil, fmt.Errorf("Get model dictionary error: %s\n", ARTM_ERRORS[-messageLength])
	}

	buffer := make([]byte, messageLength)

	errorID := C.ArtmCopyRequestedMessage(messageLength, (*C.char)(unsafe.Pointer(&buffer[0])))
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}
	if errorID < 0 {
		return nil, fmt.Errorf("Copy requested model dictionary: %s\n", ARTM_ERRORS[-errorID])
	}

	fmt.Printf("[%#v]", buffer[:10])

	dictionaryData := &DictionaryData{}

	err = proto.Unmarshal(buffer, dictionaryData)
	if err != nil {
		return nil, fmt.Errorf("Protobuf DictionaryData unmarshaling error: %s", err)
	}

	return dictionaryData, nil
}

func ArtmDisposeDictionary(masterModelID int, dictionaryName string) error {
	bytesPointer := []byte(dictionaryName)
	errorID := C.ArtmDisposeDictionary(C.int(masterModelID), (*C.char)(unsafe.Pointer(&bytesPointer[0])))
	if errorID < 0 {
		return fmt.Errorf("Dictionary dispose error: %s\n", ARTM_ERRORS[-errorID])
	}
	err := ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}
