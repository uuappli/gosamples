package library

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicMagager failed")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicMageger failed, not empty")
	}

	m0 := &MusicEntry{"1", "My Heart Will Go On", "Celin Dion", "http://asfaf.com", "MP3"}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager Add failed")

	}

	m := mm.Find(m0.Name)
	if m.Artist != m0.Artist || m.Id != m0.Id || m.Name != m0.Name || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager Find failed, Found item mismatch")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager Get Failed", err)
	}

	m = mm.Remove(m0.Name)

	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager Remove Failed")
	}
}
