package structfieldname

import (
	"reflect"
	"testing"
)

func requireEqualError(t *testing.T, actualErr error, expectedErrMsg string) {
	t.Helper()

	if actualErr == nil {
		t.Fatal("An error is expected but got nil.")
	}

	if actualErr.Error() != expectedErrMsg {
		t.Fatalf("Error message not equal:\n"+
			"expected: %q\n"+
			"actual  : %q", expectedErrMsg, actualErr.Error())
	}
}

func requireNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("Received unexpected error:\n%+v", err)
	}
}

func requireEqual(t *testing.T, expected, actual any) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected, actual)
	}
}
