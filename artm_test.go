package goartm

import "testing"

func TestCreateMainComponent(t *testing.T) {
	modelID, err := ArtmCreateMasterModel()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
}

func TestArtmImportModel(t *testing.T) {
	modelID, err := ArtmCreateMasterModel()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
	err = ArtmImportModel(modelID, "model")
	if err != nil {
		t.Error(err)
	}
	t.Logf("Successfully imported model into %d", modelID)
}
