package goartm

import (
	"fmt"
	"math/rand"
	"unsafe"
	"time"

	"github.com/golang/protobuf/proto"
)

// #cgo darwin LDFLAGS: -L. -lstdc++ -lm -lboost_system -lboost_timer -lboost_thread-mt -lboost_iostreams -lboost_filesystem -lartm-static -lgflags -lglog -linternals_proto -lmessages_proto -lprotobuf -lprotobuf-lite -lprotoc
// #cgo linux LDFLAGS: -L. -lstdc++ -lm -lboost_system -lboost_thread -lboost_iostreams -lboost_filesystem -lartm-static -lgoogle-glog -lgflags -linternals_proto -lmessages_proto -lprotobuf -lprotobuf-lite -lprotoc
// #include <stdlib.h>
// #include "c_interface.h"
import "C"

var TRUE = true
var FALSE = false

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

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func NewGetThetaMatrixArgs() *GetThetaMatrixArgs {
	eps := Default_GetThetaMatrixArgs_Eps
	layout := Default_GetThetaMatrixArgs_MatrixLayout
	return &GetThetaMatrixArgs{Eps: &eps, MatrixLayout: &layout}
}

func NewTransformMasterModelArgs() *TransformMasterModelArgs {
	matrixType := Default_TransformMasterModelArgs_ThetaMatrixType
	return &TransformMasterModelArgs{ThetaMatrixType: &matrixType}
}

func NewBatch() *Batch {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	guid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return &Batch{Id: &guid}
}

func NewItem() *Item {
	return &Item{}
}

func NewGetMasterComponentInfoArgs() *GetMasterComponentInfoArgs {
	return &GetMasterComponentInfoArgs{}
}

func NewMasterModelConfig() *MasterModelConfig {
	c := &MasterModelConfig{}
	var MasterModelConfig_PwtName string = Default_MasterModelConfig_PwtName
	var MasterModelConfig_NwtName string = Default_MasterModelConfig_NwtName
	var MasterModelConfig_ReuseTheta bool = Default_MasterModelConfig_ReuseTheta
	var MasterModelConfig_OptForAvx bool = Default_MasterModelConfig_OptForAvx
	var MasterModelConfig_CacheTheta bool = Default_MasterModelConfig_CacheTheta

	c.NwtName = &MasterModelConfig_NwtName
	c.PwtName = &MasterModelConfig_PwtName
	c.ReuseTheta = &MasterModelConfig_ReuseTheta
	c.OptForAvx = &MasterModelConfig_OptForAvx
	c.CacheTheta = &MasterModelConfig_CacheTheta
	return c
}

func NewGatherDictionaryArgs(name, dataPath, vocabFilePath string) *GatherDictionaryArgs {
	gda := new(GatherDictionaryArgs)
	gda.DictionaryTargetName = &name
	gda.DataPath = &dataPath
	gda.VocabFilePath = &vocabFilePath
	symmetricCoocValues := gda.GetSymmetricCoocValues()
	gda.SymmetricCoocValues = &symmetricCoocValues
	return gda
}

func NewFilterDictionaryArgs(name, targetName string, minCount, maxDfRate float32, max int64) *FilterDictionaryArgs {
	fda := new(FilterDictionaryArgs)
	fda.DictionaryName = &name
	fda.DictionaryTargetName = &targetName
	fda.MinDf = &minCount
	if max != 0 {
		fda.MaxDictionarySize = &max
	}
	if maxDfRate != 0 {
		fda.MaxDfRate = &maxDfRate
	}
	return fda
}
func ArtmGetLastErrorMessage() error {
	err := C.ArtmGetLastErrorMessage()

	errorStr := C.GoString(err)
	if len(errorStr) > 0 {
		//C.free(err) ???? should we dispose error string allocated within library
		return fmt.Errorf("%s", errorStr)
	}
	return nil
}

func artmCopyRequestedMessage(length C.int64_t, messagePointer proto.Message) error {
	// allocate memory for message being filled
	buffer := make([]byte, length)
	// fill memory with message data

	errorID := C.ArtmCopyRequestedMessage(length, (*C.char)(unsafe.Pointer(&buffer[0])))
	// check errors
	if errorID < 0 {
		return fmt.Errorf("Copy requested data error: %s\n", ARTM_ERRORS[-errorID])
	}

	if err := ArtmGetLastErrorMessage(); err != nil {
		return err
	}

	if err := proto.Unmarshal(buffer, messagePointer); err != nil {
		return fmt.Errorf("Protobuf message unmarshaling error: %s", err)
	}
	return nil
}

//ArtmRequestScore create master model
func ArtmRequestScore(masterModelID int, scoreName string) (*ScoreData, error) {
	config := new(GetScoreValueArgs)
	config.ScoreName = &scoreName
	message, err := proto.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("Protobuf GetScoreValueArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	messageLength := C.ArtmRequestScore(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}
	if messageLength < 0 {
		return nil, fmt.Errorf("Get requested data error: %s\n", ARTM_ERRORS[-messageLength])
	}

	scoreData := &ScoreData{}
	err = artmCopyRequestedMessage(messageLength, scoreData)
	if err != nil {
		return nil, err
	}

	return scoreData, nil
}

//ArtmRequestMasterComponentInfo create master model
func ArtmRequestMasterComponentInfo(masterModelID int, config *GetMasterComponentInfoArgs) (*MasterComponentInfo, error) {
	message, err := proto.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("Protobuf GetMasterComponentInfoArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	messageLength := C.ArtmRequestMasterComponentInfo(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}
	if messageLength < 0 {
		return nil, fmt.Errorf("Get requested data error: %s\n", ARTM_ERRORS[-messageLength])
	}

	masterComponentInfo := &MasterComponentInfo{}
	err = artmCopyRequestedMessage(messageLength, masterComponentInfo)
	if err != nil {
		return nil, err
	}

	return masterComponentInfo, nil
}

//ArtmRequestTopicModel
func ArtmRequestTopicModel(masterModelID int, topicNames []string, sparse bool) (*TopicModel, error) {
	eps := float32(1e-8)
	ml := Default_GetTopicModelArgs_MatrixLayout
	config := &GetTopicModelArgs{Eps: &eps, MatrixLayout: &ml}
	config.TopicName = topicNames
	config.UseSparseFormat = &sparse
	message, err := proto.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("Protobuf GetTopicModelArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	messageLength := C.ArtmRequestTopicModel(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}
	if messageLength < 0 {
		return nil, fmt.Errorf("Get requested data error: %s\n", ARTM_ERRORS[-messageLength])
	}

	topicModel := &TopicModel{}
	err = artmCopyRequestedMessage(messageLength, topicModel)
	if err != nil {
		return nil, err
	}

	return topicModel, nil
}

//ArtmRequestTopicModelExternal
func ArtmRequestTopicModelExternal(masterModelID int, config *GetTopicModelArgs) (*TopicModel, error) {
	message, err := proto.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("Protobuf GetTopicModelArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	messageLength := C.ArtmRequestTopicModelExternal(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}
	if messageLength < 0 {
		return nil, fmt.Errorf("Get requested data error: %s\n", ARTM_ERRORS[-messageLength])
	}

	topicModel := &TopicModel{}
	err = artmCopyRequestedMessage(messageLength, topicModel)
	if err != nil {
		return nil, err
	}

	return topicModel, nil
}

//ArtmCreateMasterModel create master model
func ArtmCreateMasterModel(config *MasterModelConfig) (int, error) {
	message, err := proto.Marshal(config)
	if err != nil {
		return 0, fmt.Errorf("Protobuf MasterModelConfig marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	masterID := C.int(C.ArtmCreateMasterModel(C.int64_t(len(message)), p_message))
	C.free(unsafe.Pointer(p_message))
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
func ArtmImportModel(masterModelID int, fileName string) error {
	config := &ImportModelArgs{FileName: &fileName}
	message, err := proto.Marshal(config)
	if err != nil {
		return fmt.Errorf("Protobuf ImportModelArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	errorID := C.ArtmImportModel(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
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

func ArtmDisposeBatch(masterModelID int, batchName string) error {
	p_batchName := C.CString(batchName)
	errorID := C.ArtmDisposeBatch(C.int(masterModelID), p_batchName)
	C.free(unsafe.Pointer(p_batchName))
	if errorID < 0 {
		fmt.Errorf("Dispose batch error %s\n", ARTM_ERRORS[-errorID])
	}
	err := ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

func ArtmImportDictionary(masterModelID int, dictionaryName, dictionaryFile string) error {
	conf := new(ImportDictionaryArgs)
	conf.DictionaryName = &dictionaryName
	conf.FileName = &dictionaryFile

	message, err := proto.Marshal(conf)
	if err != nil {
		return fmt.Errorf("Protobuf ImportDictionaryArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	errorID := C.ArtmImportDictionary(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
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
func ArtmRequestDictionary(masterModelID int, dicName string) (*DictionaryData, error) {
	conf := &GetDictionaryArgs{DictionaryName: &dicName}
	message, err := proto.Marshal(conf)
	if err != nil {
		return nil, fmt.Errorf("Protobuf GetDictionaryArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	messageLength := C.ArtmRequestDictionary(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}

	dictionaryData := &DictionaryData{}
	if err = artmCopyRequestedMessage(messageLength, dictionaryData); err != nil {
		return nil, err
	}
	return dictionaryData, nil
}

//ArtmRequestLoadBatch
func ArtmRequestLoadBatch(fileName string) (*Batch, error) {
	p_fileName := C.CString(fileName)
	messageLength := C.ArtmRequestLoadBatch(p_fileName)
	C.free(unsafe.Pointer(p_fileName))
	err := ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}
	batch := new(Batch)
	if err = artmCopyRequestedMessage(messageLength, batch); err != nil {
		return nil, err
	}
	return batch, nil
}

func ArtmDisposeDictionary(masterModelID int, dictionaryName string) error {
	p_dictionaryName := C.CString(dictionaryName)
	errorID := C.ArtmDisposeDictionary(C.int(masterModelID), p_dictionaryName)
	C.free(unsafe.Pointer(p_dictionaryName))
	if errorID < 0 {
		return fmt.Errorf("Dictionary dispose error: %s\n", ARTM_ERRORS[-errorID])
	}
	err := ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

//ArtmRequestTransformMasterModel apply model to new data
func ArtmRequestTransformMasterModel(masterModelID int, conf *TransformMasterModelArgs) (*ThetaMatrix, error) {
	message, err := proto.Marshal(conf)
	if err != nil {
		return nil, fmt.Errorf("Protobuf TransformMasterModelArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	messageLength := C.ArtmRequestTransformMasterModel(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}

	thetaMatrix := &ThetaMatrix{}
	if err = artmCopyRequestedMessage(messageLength, thetaMatrix); err != nil {
		return nil, err
	}
	return thetaMatrix, nil
}

//ArtmRequestTransformMasterModelExternal apply model to new data
func ArtmRequestTransformMasterModelExternal(masterModelID int, conf *TransformMasterModelArgs) (*ThetaMatrix, error) {
	message, err := proto.Marshal(conf)
	if err != nil {
		return nil, fmt.Errorf("Protobuf TransformMasterModelArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	messageLength := C.ArtmRequestTransformMasterModelExternal(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}

	thetaMatrix := &ThetaMatrix{}
	if err = artmCopyRequestedMessage(messageLength, thetaMatrix); err != nil {
		return nil, err
	}
	return thetaMatrix, nil
}

//ArtmRequestThetaMatrix
func ArtmRequestThetaMatrix(masterModelID int, conf *GetThetaMatrixArgs) (*ThetaMatrix, error) {
	message, err := proto.Marshal(conf)
	if err != nil {
		return nil, fmt.Errorf("Protobuf GetThetaMatrixArgs marshaling error: %s", err)
	}

	p_message := C.CString(string(message))
	messageLength := C.ArtmRequestThetaMatrix(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}
	if messageLength == 0 {
		return nil, fmt.Errorf("ThetaMatrix is empty")
	}
	thetaMatrix := &ThetaMatrix{}
	if err = artmCopyRequestedMessage(messageLength, thetaMatrix); err != nil {
		return nil, err
	}
	return thetaMatrix, nil
}

//ArtmRequestThetaMatrixExternal
func ArtmRequestThetaMatrixExternal(masterModelID int, conf *GetThetaMatrixArgs) (*ThetaMatrix, error) {
	message, err := proto.Marshal(conf)
	if err != nil {
		return nil, fmt.Errorf("Protobuf GetThetaMatrixArgs marshaling error: %s", err)
	}

	p_message := C.CString(string(message))
	messageLength := C.ArtmRequestThetaMatrixExternal(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return nil, err
	}

	thetaMatrix := &ThetaMatrix{}
	if err = artmCopyRequestedMessage(messageLength, thetaMatrix); err != nil {
		return nil, err
	}
	return thetaMatrix, nil
}

//ArtmSaveBatch save batch to disk
func ArtmSaveBatch(disk_path string, batch *Batch) error {
	message, err := proto.Marshal(batch)
	if err != nil {
		return fmt.Errorf("Protobuf Batch marshaling error: %s", err)
	}

	p_message := C.CString(string(message))
	p_diskPath := C.CString(disk_path)
	C.ArtmSaveBatch(p_diskPath, C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_diskPath))
	C.free(unsafe.Pointer(p_message))

	err = ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

func ArtmGatherDictionary(masterModelID int, conf *GatherDictionaryArgs) error {
	message, err := proto.Marshal(conf)
	if err != nil {
		return fmt.Errorf("Protobuf GatherDictionaryArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	errorID := C.ArtmGatherDictionary(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))

	if errorID < 0 {
		return fmt.Errorf("ArtmGatherDictionary error: %s\n", ARTM_ERRORS[-errorID])
	}
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

func ArtmFilterDictionary(masterModelID int, conf *FilterDictionaryArgs) error {
	message, err := proto.Marshal(conf)
	if err != nil {
		return fmt.Errorf("Protobuf FilterDictionaryArgs marshaling error: %s", err)
	}

	p_message := C.CString(string(message))
	errorID := C.ArtmFilterDictionary(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	if errorID < 0 {
		return fmt.Errorf("ArtmFilterDictionary error: %s\n", ARTM_ERRORS[-errorID])
	}
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

func ArtmExportDictionary(masterModelID int, dictionaryName, fileName string) error {
	conf := new(ExportDictionaryArgs)
	conf.DictionaryName = &dictionaryName
	conf.FileName = &fileName
	message, err := proto.Marshal(conf)
	if err != nil {
		return fmt.Errorf("Protobuf ExportDictionaryArgs marshaling error: %s", err)
	}

	p_message := C.CString(string(message))
	errorID := C.ArtmExportDictionary(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))

	if errorID < 0 {
		return fmt.Errorf("ArtmExportDictionary error: %s\n", ARTM_ERRORS[-errorID])
	}
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

func ArtmFitOfflineMasterModel(masterModelID int, batchFolder string, numCollectionPasses int32) error {
	conf := new(FitOfflineMasterModelArgs)
	conf.BatchFolder = &batchFolder
	conf.NumCollectionPasses = &numCollectionPasses

	message, err := proto.Marshal(conf)
	if err != nil {
		return fmt.Errorf("Protobuf FitOfflineMasterModelArgs marshaling error: %s", err)
	}

	p_message := C.CString(string(message))
	errorID := C.ArtmFitOfflineMasterModel(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	if errorID < 0 {
		return fmt.Errorf("ArtmFitOfflineMasterModel error: %s\n", ARTM_ERRORS[-errorID])
	}
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

func ArtmFitOnlineMasterModel(masterModelID int, batchFileName []string) error {
	conf := new(FitOnlineMasterModelArgs)
	conf.BatchFilename = batchFileName

	message, err := proto.Marshal(conf)
	if err != nil {
		return fmt.Errorf("Protobuf FitOnlineMasterModel marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	errorID := C.ArtmFitOfflineMasterModel(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	if errorID < 0 {
		return fmt.Errorf("ArtmFitOnlineMasterModel error: %s\n", ARTM_ERRORS[-errorID])
	}
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

func ArtmExportModel(masterModelID int, fileName, modelName string) error {
	conf := new(ExportModelArgs)
	conf.FileName = &fileName
	conf.ModelName = &modelName
	message, err := proto.Marshal(conf)
	if err != nil {
		return fmt.Errorf("Protobuf ExportModelArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	errorID := C.ArtmExportModel(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	if errorID < 0 {
		return fmt.Errorf("ArtmExportModel error: %s\n", ARTM_ERRORS[-errorID])
	}
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

//ArtmInitializeModel wrapper
func ArtmInitializeModel(masterModelID int, modelName, dictionaryName string, topics ...[]string) error {
	conf := new(InitializeModelArgs)
	conf.ModelName = &modelName
	conf.DictionaryName = &dictionaryName
	var allTopics []string
	for _, t := range topics {
		allTopics = append(allTopics, t...)
	}
	conf.TopicName = allTopics
	rnd := rand.Int31()
	conf.Seed = &rnd
	message, err := proto.Marshal(conf)
	if err != nil {
		return fmt.Errorf("Protobuf InitializeModelArgs marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	errorID := C.ArtmInitializeModel(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	if errorID < 0 {
		return fmt.Errorf("ArtmInitializeModel error: %s\n", ARTM_ERRORS[-errorID])
	}
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}

//ArtmReconfigureMasterModel wrapper
func ArtmReconfigureMasterModel(masterModelID int, conf *MasterModelConfig) error {

	message, err := proto.Marshal(conf)
	if err != nil {
		return fmt.Errorf("Protobuf MasterModelConfig marshaling error: %s", err)
	}
	p_message := C.CString(string(message))
	errorID := C.ArtmReconfigureMasterModel(C.int(masterModelID), C.int64_t(len(message)), p_message)
	C.free(unsafe.Pointer(p_message))
	if errorID < 0 {
		return fmt.Errorf("ArtmReconfigureMasterModel error: %s\n", ARTM_ERRORS[-errorID])
	}
	err = ArtmGetLastErrorMessage()
	if err != nil {
		return err
	}
	return nil
}
