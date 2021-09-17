package chords

import (
	"github.com/nate-droid/core/notes"
	"github.com/nate-droid/core/scales"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProgression(t *testing.T) {
	type args struct {
		mode      scales.ModeName
		intervals []ChordInterval
		key       notes.Note
	}
	tests := []struct {
		name string
		args args
		want []Chord
	}{
		{
			name: "C Major I - IV -V",
			args: args{
				mode:      scales.Ionian,
				intervals: []ChordInterval{I, IV, V},
				key:       notes.C,
			},
			want: []Chord{
				{Root: notes.C.Name},
				{Root: notes.F.Name},
				{Root: notes.G.Name},
			},
		},
		{
			name: "C Major ii - V - I",
			args: args{
				mode:	scales.Ionian,
				intervals: []ChordInterval{ii, V, I},
				key: notes.C,
			},
			want: []Chord{
				{Root: notes.D.Name},
				{Root: notes.G.Name},
				{Root: notes.C.Name},
			},
		},
		{
			name: "C Major I - vi - IV - V",
			args: args{
				mode: scales.Ionian,
				intervals: []ChordInterval{I, vi, IV, V},
				key: notes.C,
			},
			want: []Chord{
				{Root: notes.C.Name},
				{Root: notes.A.Name},
				{Root: notes.F.Name},
				{Root: notes.G.Name},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			progression, err := Progression(tt.args.mode, tt.args.intervals, tt.args.key)
			assert.NoError(t, err)
			assert.Equal(t, tt.want[0].Root, progression[0].Root)
			assert.Equal(t, tt.want[1].Root, progression[1].Root)
			assert.Equal(t, tt.want[2].Root, progression[2].Root)
		})
	}
}
