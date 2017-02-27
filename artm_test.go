package goartm

import "testing"

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

	importModelConfig := NewImportModelArgs("model")
	err = ArtmImportModel(modelID, importModelConfig)
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

	importModelConfig := NewImportModelArgs("model")
	err = ArtmImportModel(modelID, importModelConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// load dictionary

	importDictionaryArgsConfig := NewImportDictionaryArgs("main_dictionary", "dictionary.dict")
	err = ArtmImportDictionary(modelID, importDictionaryArgsConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully loaded dictionary into model %d", modelID)

	// get dictionary
	dictionaryConfig := NewGetDictionaryArgs(*importDictionaryArgsConfig.DictionaryName)
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
	importModelConfig := NewImportModelArgs("model")
	err = ArtmImportModel(modelID, importModelConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// load dictionary

	importDictionaryArgsConfig := NewImportDictionaryArgs("main_dictionary", "dictionary.dict")
	err = ArtmImportDictionary(modelID, importDictionaryArgsConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully loaded dictionary into model %d", modelID)

	// get dictionary
	dictionaryConfig := NewGetDictionaryArgs(*importDictionaryArgsConfig.DictionaryName)
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
	scoreConfig := NewGetScoreValueArgs("TopTokensScore")

	score, err := ArtmRequestScore(modelID, scoreConfig)
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
	importModelConfig := NewImportModelArgs("model")
	err = ArtmImportModel(modelID, importModelConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// load dictionary

	importDictionaryArgsConfig := NewImportDictionaryArgs("main_dictionary", "dictionary.dict")
	err = ArtmImportDictionary(modelID, importDictionaryArgsConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully loaded dictionary into model %d", modelID)

	// get dictionary
	dictionaryConfig := NewGetDictionaryArgs(*importDictionaryArgsConfig.DictionaryName)
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
	importModelConfig := NewImportModelArgs("model")
	err = ArtmImportModel(modelID, importModelConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// load dictionary

	importDictionaryArgsConfig := NewImportDictionaryArgs("main_dictionary", "dictionary.dict")
	err = ArtmImportDictionary(modelID, importDictionaryArgsConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully loaded dictionary into model %d", modelID)

	// get dictionary
	dictionaryConfig := NewGetDictionaryArgs(*importDictionaryArgsConfig.DictionaryName)
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
	for i, token := range topicModel.GetToken() {
		t.Logf("Token %d [%s]", i, token)
		if i == 10 {
			break
		}
	}
	t.Logf("Successfully got [%s] model with %d token weights", topicModel.GetName(), len(topicModel.GetTokenWeights()))
	for i, tokenWeight := range topicModel.GetTokenWeights() {
		t.Logf("Token %d [%s]", i, tokenWeight.GetValue())
		if i == 10 {
			break
		}
	}

	topicModel.GetToken()

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
