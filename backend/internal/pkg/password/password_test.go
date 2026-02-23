package password

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	if hash == "password" {
		t.Error("expected hash to be different from password")
	}
}

func TestHashPassword_Error(t *testing.T) {
	// Invalid cost triggers bcrypt error (cost must be MinCost..MaxCost).
	old := cost
	cost = 99
	t.Cleanup(func() { cost = old })

	_, err := HashPassword("password")
	if err == nil {
		t.Error("expected error for invalid cost")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	err = ComparePassword("password", hash)
	if err != nil {
		t.Errorf("error comparing password: %v", err)
	}
}

func TestComparePassword_Error(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}

	err = ComparePassword("wrongpassword", hash)
	if err == nil {
		t.Error("expected error for wrong password")
	}
}
