package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {

	password := RandomString(6)
	hashPass, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashPass)

	err = CheckPassword(password, hashPass)
	require.NoError(t, err)

	fakePass := RandomString(6)
	err = CheckPassword(fakePass, hashPass)
	require.Error(t, bcrypt.ErrMismatchedHashAndPassword)
}
