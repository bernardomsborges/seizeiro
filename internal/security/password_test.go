package security

import "testing"

func TestPasswordHashing(t *testing.T) {
	t.Parallel()

	pwd := "zyyzx"
	pwHash, err := HashPassword(pwd)
	if err != nil {
		t.Fatal(err)
	}

	ok, err := VerifyPassword(pwHash, pwd)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("expected hash and password to match")
	}

	ok, err = VerifyPassword(pwHash, "foobar")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("expected hash and invalid password to not match")
	}
}
