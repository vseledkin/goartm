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
