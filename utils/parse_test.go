package utils_test

import (
	"aoc/utils"
	"reflect"
	"testing"
)

func ParseTest(t *testing.T) {
	t.Parallel()

	type args struct {
		fileContent string
		separator   string
	}

	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "space-separated",
			args: args{
				fileContent: "Game 1\nGame 2",
				separator:   " ",
			},
			want: [][]string{{"Game", "1"}, {"Game", "2"}},
		},
		{
			name: "space-separated lines",
			args: args{
				fileContent: "Game 1\nGame 2",
				separator:   " ",
			},
			want: [][]string{{"Game", "1"}, {"Game", "2"}},
		},
		{
			name: "space-separated",
			args: args{
				fileContent: "apple banana\ncarrot potato",
				separator:   " ",
			},
			want: [][]string{{"apple", "banana"}, {"carrot", "potato"}},
		},
		{
			name: "comma-separated, no skip empty lines",
			args: args{
				fileContent: " apple,banana \n carrot,potato ",
				separator:   ",",
			},
			want: [][]string{{"apple", "banana"}, {"carrot", "potato"}},
		},
		{
			name: "tab-separated lines",
			args: args{
				fileContent: "row1\tcol1\tcol2\nrow2\tcol3\tcol4",
				separator:   "\t",
			},
			want: [][]string{{"row1", "col1", "col2"}, {"row2", "col3", "col4"}},
		},
		{
			name: "comma-separated lines",
			args: args{
				fileContent: "apple, orange , banana\n, , \n ,kiwi , ",
				separator:   ",",
			},
			want: [][]string{{"apple", "orange", "banana"}, {"", "", ""}, {"", "kiwi", ""}},
		},
		{
			name: "semicolon-separated lines",
			args: args{
				fileContent: "word1;word2;word3\nword4;word5;word6",
				separator:   ";",
			},
			want: [][]string{{"word1", "word2", "word3"}, {"word4", "word5", "word6"}},
		},
		{
			name: "complex mixed separators",
			args: args{
				fileContent: "val1 ,  val2 ,val3\nval4; val5;val6",
				separator:   ",",
			},
			want: [][]string{{"val1", "val2", "val3"}, {"val4; val5;val6"}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := utils.Parse(tt.args.fileContent, tt.args.separator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got  %v want %v", got, tt.want)
			}
		})
	}
}
