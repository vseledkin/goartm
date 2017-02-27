package goartm

import "testing"

func TestCreateMainComponent(t *testing.T) {
	conf := New(&MasterModelConfig{}).(*MasterModelConfig)
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
	conf := New(&MasterModelConfig{}).(*MasterModelConfig)

	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	// load model
	model := "model"
	importModelConfig := New(&ImportModelArgs{}).(*ImportModelArgs)
	importModelConfig.FileName = &model
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
	conf := New(&MasterModelConfig{}).(*MasterModelConfig)
	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	// load model
	model := "model"
	importModelConfig := New(&ImportModelArgs{}).(*ImportModelArgs)
	importModelConfig.FileName = &model
	err = ArtmImportModel(modelID, importModelConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// load dictionary

	importDictionaryArgsConfig := New(&ImportDictionaryArgs{}).(*ImportDictionaryArgs)
	dicFile := "dictionary.dict"
	dicName := "main_dictionary"
	importDictionaryArgsConfig.FileName = &dicFile
	importDictionaryArgsConfig.DictionaryName = &dicName
	err = ArtmImportDictionary(modelID, importDictionaryArgsConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully loaded dictionary into model %d", modelID)

	// get dictionary
	dictionaryConfig := New(&GetDictionaryArgs{}).(*GetDictionaryArgs)
	dictionaryConfig.DictionaryName = &dicName
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
	err = ArtmDisposeDictionary(modelID, dicName)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed model %d dictionary %s", modelID, dicName)
	// dospose
	err = ArtmDisposeMasterComponent(modelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed master model with id %d", modelID)
}
