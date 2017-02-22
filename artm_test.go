package goartm

import "testing"

func TestCreateMainComponent(t *testing.T) {
	conf := &MasterModelConfig{}
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
	conf := &MasterModelConfig{}
	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	// load model
	model := "model"
	importModelConfig := &ImportModelArgs{}
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
	conf := &MasterModelConfig{}
	modelID, err := ArtmCreateMasterModel(conf)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	// load model
	model := "model"
	importModelConfig := &ImportModelArgs{}
	importModelConfig.FileName = &model
	err = ArtmImportModel(modelID, importModelConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
	// get dictionary
	dictionaryConfig := &GetDictionaryArgs{}
	dic, err := ArtmRequestDictionary(modelID, dictionaryConfig)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully requested dictionary of model %d", modelID)

	t.Logf("Dictionary name is [%s]", dic.Name)

	// dospose
	err = ArtmDisposeMasterComponent(modelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Successfully disposed master model with id %d", modelID)
}
