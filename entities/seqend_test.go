package entities

import (
	"github.com/stretchr/testify/assert"
	"github.com/whutwxn/dxf-go/core"
	"strings"
	"testing"
)

func TestSeqEnd(t *testing.T) {
	expected := SeqEnd{
		BaseEntity: BaseEntity{
			On:      true,
			Visible: true,
		},
	}

	next := core.Tagger(strings.NewReader("  0\nSEQEND"))
	seqend, err := NewSeqEnd(core.TagSlice(core.AllTags(next)))

	assert.Nil(t, err)
	assert.True(t, expected.Equals(seqend))
	assert.False(t, SeqEnd{}.Equals(core.NewStringValue("SEQOTHER")))

	assert.True(t, SeqEnd{}.IsSeqEnd())
	assert.False(t, SeqEnd{}.HasNestedEntities())
}
