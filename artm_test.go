package goartm

import (
	"os"
	"strings"
	"testing"
)

func TestCreateMainComponent(t *testing.T) {

	conf := NewMasterModelConfig()
	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	err = ArtmDisposeMasterComponent(modelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed master model with id %d", modelID)
}

func TestArtmImportModel(t *testing.T) {
	// initialize
	conf := NewMasterModelConfig()
	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	// load model

	err = ArtmImportModel(modelID, "model")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// dospose
	err = ArtmDisposeMasterComponent(modelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed master model with id %d", modelID)
}

func TestArtmRequestDictionary(t *testing.T) {
	// initialize
	conf := NewMasterModelConfig()
	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	// load model

	err = ArtmImportModel(modelID, "model")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// load dictionary

	dicName := "main_dictionary"
	err = ArtmImportDictionary(modelID, dicName, "dictionary.dict")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully loaded dictionary into model %d", modelID)

	// get dictionary
	dictionaryConfig := NewGetDictionaryArgs(dicName)
	dic, err := ArtmRequestDictionary(modelID, dictionaryConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully requested dictionary of model %d", modelID)

	t.Logf("Dictionary name is [%s]", *dic.Name)

	t.Logf("Dictionary has %d words", len(dic.Token))
	for i, token := range dic.Token {
		t.Logf("Dictionary has %d = %s", i, token)
		if i == 10 {
			break
		}
	}
	// dispose dictionary
	err = ArtmDisposeDictionary(modelID, *dictionaryConfig.DictionaryName)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed model %d dictionary %s", modelID, *dictionaryConfig.DictionaryName)
	// dospose
	err = ArtmDisposeMasterComponent(modelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed master model with id %d", modelID)
}

func TestArtmRequestScore(t *testing.T) {
	// initialize
	conf := NewMasterModelConfig()
	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	// load model

	err = ArtmImportModel(modelID, "model")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// load dictionary

	dicName := "main_dictionary"
	err = ArtmImportDictionary(modelID, dicName, "dictionary.dict")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully loaded dictionary into model %d", modelID)

	// get dictionary
	dictionaryConfig := NewGetDictionaryArgs(dicName)
	dic, err := ArtmRequestDictionary(modelID, dictionaryConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully requested dictionary of model %d", modelID)

	t.Logf("Dictionary name is [%s]", dic.GetName())

	t.Logf("Dictionary has %d words", len(dic.Token))
	for i, token := range dic.Token {
		t.Logf("Dictionary has %d = %s", i, token)
		if i == 10 {
			break
		}
	}

	// request score
	score, err := ArtmRequestScore(modelID, "TopTokensScore")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Successfully got score %s of model with id %d", score.GetName(), modelID)
	// dispose dictionary
	err = ArtmDisposeDictionary(modelID, *dictionaryConfig.DictionaryName)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed model %d dictionary %s", modelID, *dictionaryConfig.DictionaryName)
	// dospose
	err = ArtmDisposeMasterComponent(modelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed master model with id %d", modelID)
}

func TestArtmMasterComponentInfo(t *testing.T) {
	// initialize
	conf := NewMasterModelConfig()
	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	// load model

	err = ArtmImportModel(modelID, "model")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// load dictionary

	dicName := "main_dictionary"
	err = ArtmImportDictionary(modelID, dicName, "dictionary.dict")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully loaded dictionary into model %d", modelID)

	// get dictionary
	dictionaryConfig := NewGetDictionaryArgs(dicName)
	dic, err := ArtmRequestDictionary(modelID, dictionaryConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully requested dictionary of model %d", modelID)

	t.Logf("Dictionary name is [%s]", dic.GetName())

	t.Logf("Dictionary has %d words", len(dic.Token))
	for i, token := range dic.Token {
		t.Logf("Dictionary has %d = %s", i, token)
		if i == 10 {
			break
		}
	}

	info, err := ArtmRequestMasterComponentInfo(modelID, NewGetMasterComponentInfoArgs())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully got MasterComponentInfo of model %d:-> [%s]", modelID, info.String())
	t.Logf("Successfully got MasterComponentInfo of model %d:-> score [%s]", modelID, info.GetScore())
	t.Logf("Successfully got MasterComponentInfo of model %d:-> cache [%s]", modelID, info.GetCacheEntry())
	t.Logf("Successfully got MasterComponentInfo of model %d:-> dic [%s]", modelID, info.GetDictionary())
	t.Logf("Successfully got MasterComponentInfo of model %d:-> batch [%s]", modelID, info.GetBatch())
	t.Logf("Successfully got MasterComponentInfo of model %d:-> config [%s]", modelID, info.GetConfig())
	t.Logf("Successfully got MasterComponentInfo of model %d:-> model [%s]", modelID, info.GetModel())

	// dispose dictionary
	err = ArtmDisposeDictionary(modelID, *dictionaryConfig.DictionaryName)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed model %d dictionary %s", modelID, *dictionaryConfig.DictionaryName)
	// dospose
	err = ArtmDisposeMasterComponent(modelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed master model with id %d", modelID)
}

func TestArtmRequestTopicModel(t *testing.T) {
	// initialize
	conf := NewMasterModelConfig()
	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	// load model

	err = ArtmImportModel(modelID, "model")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// load dictionary

	dicName := "main_dictionary"
	err = ArtmImportDictionary(modelID, dicName, "dictionary.dict")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully loaded dictionary into model %d", modelID)

	// get dictionary
	dictionaryConfig := NewGetDictionaryArgs(dicName)
	dic, err := ArtmRequestDictionary(modelID, dictionaryConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully requested dictionary of model %d", modelID)

	t.Logf("Dictionary name is [%s]", dic.GetName())

	t.Logf("Dictionary has %d words", len(dic.Token))
	for i, token := range dic.Token {
		t.Logf("Dictionary has %d = %s", i, token)
		if i == 10 {
			break
		}
	}

	// request score
	getTopicModelArgs := NewGetTopicModelArgs()

	topicModel, err := ArtmRequestTopicModel(modelID, getTopicModelArgs)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully got [%s] model with id %d", topicModel.GetName(), modelID)
	t.Logf("Successfully got [%s] model with %d topics", topicModel.GetName(), topicModel.GetNumTopics())
	t.Logf("Successfully got [%s] model with %d topic names", topicModel.GetName(), len(topicModel.TopicName))
	for i, topicName := range topicModel.GetTopicName() {
		t.Logf("Topic %d [%s]", i, topicName)
		if i == 10 {
			break
		}
	}
	t.Logf("Successfully got [%s] model with %d classes", topicModel.GetName(), len(topicModel.ClassId))
	for i, className := range topicModel.GetClassId() {
		t.Logf("Class %d [%s]", i, className)
		if i == 10 {
			break
		}
	}

	t.Logf("Successfully got [%s] model with %d tokens", topicModel.GetName(), len(topicModel.GetToken()))
	tokens := topicModel.GetToken()
	for i, token := range tokens {
		t.Logf("Token %d [%s]\n", i, token)
		if i == 10 {
			break
		}
	}
	t.Logf("Successfully got [%s] model with %d token weights", topicModel.GetName(), len(topicModel.GetTokenWeights()))
	for i, tokenWeight := range topicModel.GetTokenWeights() {
		t.Logf("Token %d [%s] %#v", i, tokens[i], tokenWeight.GetValue())
		if i == 100 {
			break
		}
	}

	// dispose dictionary
	err = ArtmDisposeDictionary(modelID, *dictionaryConfig.DictionaryName)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed model %d dictionary %s", modelID, *dictionaryConfig.DictionaryName)
	// dospose
	err = ArtmDisposeMasterComponent(modelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed master model with id %d", modelID)
}

func TestPrintTopTopicTokens(t *testing.T) {
	// initialize
	conf := NewMasterModelConfig()
	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	// load model

	err = ArtmImportModel(modelID, "model")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// load dictionary

	dicName := "main_dictionary"
	err = ArtmImportDictionary(modelID, dicName, "dictionary.dict")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully loaded dictionary into model %d", modelID)

	// get dictionary
	dictionaryConfig := NewGetDictionaryArgs(dicName)
	dic, err := ArtmRequestDictionary(modelID, dictionaryConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully requested dictionary of model %d", modelID)

	t.Logf("Dictionary name is [%s]", dic.GetName())

	t.Logf("Dictionary has %d words", len(dic.Token))
	for i, token := range dic.Token {
		t.Logf("Dictionary has %d = %s", i, token)
		if i == 10 {
			break
		}
	}

	// request score
	getTopicModelArgs := NewGetTopicModelArgs()

	topicModel, err := ArtmRequestTopicModel(modelID, getTopicModelArgs)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully got [%s] model with id %d", topicModel.GetName(), modelID)
	t.Logf("Successfully got [%s] model with %d topics", topicModel.GetName(), topicModel.GetNumTopics())
	t.Logf("Successfully got [%s] model with %d topic names", topicModel.GetName(), len(topicModel.TopicName))
	for i, topicName := range topicModel.GetTopicName() {
		t.Logf("Topic %d [%s]", i, topicName)
		if i == 10 {
			break
		}
	}
	t.Logf("Successfully got [%s] model with %d classes", topicModel.GetName(), len(topicModel.ClassId))
	for i, className := range topicModel.GetClassId() {
		t.Logf("Class %d [%s]", i, className)
		if i == 10 {
			break
		}
	}

	tokens := topicModel.GetToken()
	t.Logf("Successfully got [%s] model with %d tokens", topicModel.GetName(), len(tokens))
	for i, token := range tokens {
		t.Logf("Token %d [%s]\n", i, token)
		if i == 10 {
			break
		}
	}
	t.Logf("Successfully got [%s] model with %d token weights", topicModel.GetName(), len(topicModel.GetTokenWeights()))
	for i, tokenWeight := range topicModel.GetTokenWeights() {
		t.Logf("Token %d [%s] %#v", i, tokens[i], tokenWeight.GetValue())
		if i == 10 {
			break
		}
	}
	teta := topicModel.GetTokenWeights()
	for i := 0; i < int(topicModel.GetNumTopics()); i++ {
		top, total := GetTopTopicTokens(i, tokens, teta, 20)
		t.Logf("%d %s (%d=%f)\n", i, topicModel.GetTopicName()[i], total, 100*float32(total)/float32(len(tokens)))
		for _, tw := range top {
			t.Logf("\tWord %d [%s] %#v\n", tw.ID, tw.Object, tw.Weight)
			if i == 10 {
				break
			}
		}
	}

	// dispose dictionary
	err = ArtmDisposeDictionary(modelID, *dictionaryConfig.DictionaryName)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed model %d dictionary %s", modelID, *dictionaryConfig.DictionaryName)
	// dospose
	err = ArtmDisposeMasterComponent(modelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed master model with id %d", modelID)
}

func TestInference(t *testing.T) {
	// initialize
	conf := NewMasterModelConfig()
	var passes int32 = 100
	conf.NumDocumentPasses = &passes

	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	// load model

	err = ArtmImportModel(modelID, "model")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// load dictionary

	dicName := "main_dictionary"
	err = ArtmImportDictionary(modelID, dicName, "dictionary.dict")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully loaded dictionary into model %d", modelID)

	// get dictionary
	dictionaryConfig := NewGetDictionaryArgs(dicName)
	dic, err := ArtmRequestDictionary(modelID, dictionaryConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully requested dictionary of model %d", modelID)

	t.Logf("Dictionary name is [%s]", dic.GetName())

	t.Logf("Dictionary has %d words", len(dic.Token))
	for i, token := range dic.Token {
		t.Logf("Dictionary has %d = %s", i, token)
		if i == 10 {
			break
		}
	}

	// request score
	getTopicModelArgs := NewGetTopicModelArgs()

	topicModel, err := ArtmRequestTopicModel(modelID, getTopicModelArgs)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully got [%s] model with id %d", topicModel.GetName(), modelID)
	t.Logf("Successfully got [%s] model with %d topics", topicModel.GetName(), topicModel.GetNumTopics())
	t.Logf("Successfully got [%s] model with %d topic names", topicModel.GetName(), len(topicModel.TopicName))
	for i, topicName := range topicModel.GetTopicName() {
		t.Logf("Topic %d [%s]", i, topicName)
		if i == 10 {
			break
		}
	}
	t.Logf("Successfully got [%s] model with %d classes", topicModel.GetName(), len(topicModel.ClassId))
	for i, className := range topicModel.GetClassId() {
		t.Logf("Class %d [%s]", i, className)
		if i == 10 {
			break
		}
	}

	documents := [][]string{
		strings.Fields("нейронные сети используются как замена программиста"),
		strings.Fields("нейронные сети мозга программиста используются как средство создания нейронных сетей для замены программиста"),
	}

	batch := NewBatchFromData(documents)

	t.Logf("Created batch with uid %s", *batch.Id)
	t.Logf("Created batch %v", *batch)
	transformMasterModelArgs := NewTransformMasterModelArgs()
	transformMasterModelArgs.Batch = []*Batch{batch}

	thetaMatrix, err := ArtmRequestTransformMasterModel(modelID, transformMasterModelArgs)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully inferred theta distribution for new documnts %d topics", thetaMatrix.GetNumTopics())

	for did, array := range thetaMatrix.GetItemWeights() {
		for tid, v := range array.Value {
			if v > 1e-3 {
				top, _ := GetTopTopicTokens(tid, topicModel.GetToken(), topicModel.GetTokenWeights(), 10)
				topicLabel := ""
				for _, tw := range top {
					topicLabel += " " + tw.Object
				}
				t.Logf("for document:%d got topic:%d with weight:%f and label:%s", did, tid, v, topicLabel)
			}
		}
	}
	// test external variant
	/*
		thetaMatrix, err = ArtmRequestTransformMasterModelExternal(modelID, transformMasterModelArgs)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Successfully inferred theta distribution for new documnts %d topics", thetaMatrix.GetNumTopics())

		for did, array := range thetaMatrix.GetItemWeights() {
			for tid, v := range array.Value {
				if v > 1e-3 {
					top, _ := GetTopTopicTokens(tid, topicModel.GetToken(), topicModel.GetTokenWeights(), 10)
					topicLabel := ""
					for _, tw := range top {
						topicLabel += " " + tw.Object
					}
					t.Logf("for document:%d got topic:%d with weight:%f and label:%s", did, tid, v, topicLabel)
				}
			}
		}
	*/
	// dispose dictionary
	err = ArtmDisposeDictionary(modelID, *dictionaryConfig.DictionaryName)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed model %d dictionary %s", modelID, *dictionaryConfig.DictionaryName)
	// dospose
	err = ArtmDisposeMasterComponent(modelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed master model with id %d", modelID)
}

func TestImportBatch(t *testing.T) {

	documents := [][]string{
		strings.Fields("нейронные сети используются как замена программиста"),
		strings.Fields("нейронные сети мозга программиста используются как средство создания нейронных сетей для замены программиста"),
	}

	batch := NewBatchFromData(documents)

	t.Logf("Created batch with uid %s", *batch.Id)
	t.Logf("Created batch %#v", *batch)

	err := os.MkdirAll("./store_tmp", 0777)
	if err != nil {
		t.Fatal(err)
	}

	// save batch to disk
	err = ArtmSaveBatch("./store_tmp", batch)
	if err != nil {
		t.Fatal(err)
	}

}
