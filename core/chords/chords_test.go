package chords

import (
	"github.com/nate-droid/core/notes"
	"github.com/nate-droid/core/scales"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChordTones(t *testing.T) {
	type args struct {
		note      notes.Note
		chordType ChordType
	}
	tests := []struct {
		name string
		args args
		want []notes.Note
	}{
		{
			name: "C Major Chord",
			args: args{
				note:      notes.C,
				chordType: MajorChord,
			},
			want: []notes.Note{notes.C, notes.E, notes.G},
		},
		{
			name: "G Minor Chord",
			args: args{
				note:      notes.G,
				chordType: MinorChord,
			},
			want: []notes.Note{notes.G, notes.ASharp, notes.D},
		},
		{
			name: "F Major Chord",
			args: args{
				note: notes.F,
				chordType: MajorChord,
			},
			want: []notes.Note{notes.F, notes.A, notes.C},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := GetChordTones(tt.args.note, tt.args.chordType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChordTones() = %v, want %v", got, tt.want)
			}
		})
	}
}
// todo index F = 5
// interval P5 7
// index is 12

func TestGetChordTypeForScalePosition(t *testing.T) {
	type args struct {
		scalePosition int
		scale         scales.Scale
	}
	tests := []struct {
		name    string
		args    args
		want    []ChordType
		wantErr bool
	}{
		{
			name: "Test Ionian mode",
			args: args{
				scalePosition: 0,
				scale:         scales.MajorScale,
			},
			want:    []ChordType{MajorChord, MajorSeventh, MajorSixth}, // TODO add all of the chords
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChordQualitiesForScalePosition(tt.args.scalePosition, tt.args.scale)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChordQualitiesForScalePosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// assert.EqualValues(t, got, tt.want)
			assert.True(t, SameChordTypeSlice(got, tt.want))
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("GetChordQualitiesForScalePosition() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestNewChord(t *testing.T) {
	type args struct {
		root notes.Name
		name ChordType
	}
	tests := []struct {
		name string
		args args
		want Chord
	}{
		{
			name: "Test creating C Maj",
			args: args{
				root: notes.C.Name,
				name: MajorChord,
			},
			want: Chord{
				Name:      MajorChord,
				Root:      notes.C.Name,
				Intervals: []Interval{P1, major3, P5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChord(tt.args.root, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		slice []int
		val   int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "Test if integers are present (PASS)",
			args: args{
				slice: []int{1, 2, 3, 4},
				val:   2,
			},
			want:  1,
			want1: true,
		},
		{
			name: "Test integer not present (FAIL)",
			args: args{
				slice: []int{1, 2, 3, 4},
				val:   7,
			},
			want:  -1,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Find(tt.args.slice, tt.args.val)
			if got != tt.want {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSameChordTypeSlice(t *testing.T) {
	type args struct {
		x []ChordType
		y []ChordType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test if same ChordTypes are present",
			args: args{
				x: []ChordType{MajorSixth, MajorSeventh},
				y: []ChordType{MajorSixth, MajorSeventh},
			},
			want: true,
		},
		{
			name: "Test if same ChordTypes are not present",
			args: args{
				x: []ChordType{MajorSixth},
				y: []ChordType{MajorSeventh},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SameChordTypeSlice(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("SameChordTypeSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
