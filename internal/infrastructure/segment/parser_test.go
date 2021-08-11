package segment_test

import (
	"github.com/cipma/smsbin/infrastructure/segment"
	"github.com/stretchr/testify/assert"
	"testing"
)

type sgmt struct {
	text string
	size int
}

var testData = []struct {
	text     string
	size     int
	segments []sgmt
}{
	{
		"",
		0,
		[]sgmt{},
	},
	{
		"hello world!",
		12,
		[]sgmt{
			{"hello world!", 12},
		},
	},
	{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. In lobortis pellentesque ultricies. Vivamus vitae arcu turpis. Morbi malesuada ex vel posuere accumsan.",
		160,
		[]sgmt{
			{
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit. In lobortis pellentesque ultricies. Vivamus vitae arcu turpis. Morbi malesuada ex vel posuere accumsan.",
				160,
			},
		},
	},
	{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. In lobortis pellentesque ultricies. Vivamus vitae arcu turpis. Morbi malesuada ex vel posuere accumsan. Aenean consectetur efficitur ultrices. Nullam sollicitudin, mauris vel rutrum sodales, turpis est ultricies nisl, facilisis feugiat lectus sem in turpis. Nunc pharetra mauris ut sem tempor viverra. Nam diam lectus, efficitur eu vehicula quis, rhoncus eu tellus. Sed malesuada odio ligula, sed varius metus dignissim ac. Quisque molestie imperdiet cursus. Ut quam diam, tempus at velit a, molestie blandit nulla. Duis commodo iaculis elit nec consequat.",
		613,
		[]sgmt{
			{
				"Lorem ipsum dolor sit amet, consectetur adipiscing elit. In lobortis pellentesque ultricies. Vivamus vitae arcu turpis. Morbi malesuada ex vel posuere ac",
				153,
			},
			{
				"cumsan. Aenean consectetur efficitur ultrices. Nullam sollicitudin, mauris vel rutrum sodales, turpis est ultricies nisl, facilisis feugiat lectus sem in",
				153,
			},
			{
				" turpis. Nunc pharetra mauris ut sem tempor viverra. Nam diam lectus, efficitur eu vehicula quis, rhoncus eu tellus. Sed malesuada odio ligula, sed variu",
				153,
			},
			{
				"s metus dignissim ac. Quisque molestie imperdiet cursus. Ut quam diam, tempus at velit a, molestie blandit nulla. Duis commodo iaculis elit nec consequat",
				153,
			},
			{
				".",
				1,
			},
		},
	},
	{
		"^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€",
		162,
		[]sgmt{
			{"^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\[~]|€^{}\\", 152},
			{"[~]|€", 10},
		},
	},
	{
		"This is and example µessage with some ©haracte® unavailable in the GSM-7 encoding",
		81,
		[]sgmt{
			{"This is and example µessage with some ©haracte® unavailable in the ", 67},
			{"GSM-7 encoding", 14},
		},
	},
}

func TestSegmentCount(t *testing.T) {
	t.Parallel()

	for _, td := range testData {
		td := td
		t.Run("", func(t *testing.T) {
			parser := segment.Parser{}

			metadata := parser.ExtractSegments(td.text)

			assert.Equal(t, td.size, metadata.TotalSize)
			assert.Equal(t, len(td.segments), len(metadata.Segments), "Segment number mismatch")

			for i, s := range metadata.Segments {
				assert.Equal(t, td.segments[i].text, s.Data)
				assert.Equal(t, td.segments[i].size, s.Size)
			}
		})
	}
}
