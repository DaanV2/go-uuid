package uuid_test

import (
	"encoding/json"
	"testing"

	"github.com/DaanV2/go-uuid"
	"github.com/stretchr/testify/require"
)

func Test_MarshalText_Roundtrip(t *testing.T) {
	u := uuid.V4.New()

	text, err := u.MarshalText()
	require.NoError(t, err)
	require.Equal(t, u.String(), string(text))

	var got uuid.UUID
	require.NoError(t, got.UnmarshalText(text))
	require.Equal(t, u, got)
}

func Test_MarshalBinary_Roundtrip(t *testing.T) {
	u := uuid.V4.New()

	bin, err := u.MarshalBinary()
	require.NoError(t, err)
	require.Len(t, bin, 16)

	var got uuid.UUID
	require.NoError(t, got.UnmarshalBinary(bin))
	require.Equal(t, u, got)
}

func Test_UnmarshalBinary_InvalidLength(t *testing.T) {
	var u uuid.UUID
	require.Error(t, u.UnmarshalBinary([]byte{0x00, 0x01}))
}

func Test_MarshalJSON_Roundtrip(t *testing.T) {
	type payload struct {
		ID uuid.UUID `json:"id"`
	}

	in := payload{ID: uuid.V4.New()}
	data, err := json.Marshal(in)
	require.NoError(t, err)
	require.Equal(t, `{"id":"`+in.ID.String()+`"}`, string(data))

	var out payload
	require.NoError(t, json.Unmarshal(data, &out))
	require.Equal(t, in, out)
}

func Test_UnmarshalJSON_Null(t *testing.T) {
	u := uuid.V4.New()
	require.NoError(t, u.UnmarshalJSON([]byte("null")))
	require.Equal(t, uuid.Nil, u)
}

func Test_UnmarshalJSON_Invalid(t *testing.T) {
	var u uuid.UUID
	require.Error(t, u.UnmarshalJSON([]byte(`123`)))
	require.Error(t, u.UnmarshalJSON([]byte(`"nope"`)))
}