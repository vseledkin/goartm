package goartm

import "testing"

func TestCreateMainComponent(t *testing.T) {
	modelID, err := ArtmCreateMasterModel()
	if err != nil {
		t.Error(err)
	}
	t.Logf("Successfully created master model with id %d", modelID)
}
