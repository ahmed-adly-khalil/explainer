package sdp

import (
	"testing"

	"github.com/pion/explainer/pkg/output"
	"github.com/stretchr/testify/require"
)

const minimalSessionDescription = `v=0
o=- 0 0 IN IP4 127.0.0.1
s=-
c=IN IP4 127.0.0.1
t=0 0`

const miminalMediaSection = `m=audio 4000 RTP/AVP 111
a=rtpmap:111 OPUS/48000/2`

// func Test_Unmarshal(t *testing.T) {
// 	s := SessionDescription{}
// 	require.NoError(t, s.Unmarshal(minimalSessionDescription))
// }
//
// Assert that we have v, o, s
func Test_GlobalValues(t *testing.T) {
	s := SessionDescription{}

	t.Run("v", func(t *testing.T) {
		// No value
		require.Equal(t, s.Unmarshal(""), output.NewMessage(errEarlyEndVersion, []output.Source{{Line: 0}}))

		// Wrong key
		require.Equal(t, s.Unmarshal("a=b"), output.NewMessage(errProtocolVersionNotFound, []output.Source{{Line: 1}}))

		// Invalid value
		require.Equal(t, s.Unmarshal("v=b"), output.NewMessage(errInvalidProtocolVersion, []output.Source{{Line: 1}}))
	})

	t.Run("o", func(t *testing.T) {
		// No value
		require.Equal(t, s.Unmarshal("v=2\r\n"), output.NewMessage(errEarlyEndOriginator, []output.Source{{Line: 1}}))

		// Wrong key
		require.Equal(t, s.Unmarshal("v=2\r\na=b"), output.NewMessage(errOriginatorNotFound, []output.Source{{Line: 2}}))
	})

	t.Run("s", func(t *testing.T) {
		// No value
		require.Equal(t, s.Unmarshal("v=2\r\no=o"), output.NewMessage(errEarlyEndSessionName, []output.Source{{Line: 2}}))

		// Wrong key
		require.Equal(t, s.Unmarshal("v=2\r\no=o\r\na=b"), output.NewMessage(errSessionNameNotFound, []output.Source{{Line: 3}}))
	})
}

func Test_LineParsing(t *testing.T) {
	s := SessionDescription{}

	require.Equal(t, s.Unmarshal("a="), output.NewMessage(errShortLine, []output.Source{{Line: 0}}))
	require.Equal(t, s.Unmarshal("a!b"), output.NewMessage(errInvalidLine, []output.Source{{Line: 0}}))
}
