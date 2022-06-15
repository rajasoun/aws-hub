package spike

import (
	"os"
	"reflect"
	"testing"
)

func Test_perimeter_Perimeter(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
		{
			name: "perimeter",
			fields: fields{
				width:  12,
				height: 12,
			},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &perimeter{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := p.Perimeter(); got != tt.want {
				t.Errorf("perimeter.Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newaudio(t *testing.T) {
	type args struct {
		lang string
	}
	tests := []struct {
		name string
		args args
		want Audio
	}{
		// TODO: Add test cases.
		{
			name: "Check for English",
			args: args{
				lang: "Hello",
			},
			want: English{},
		},
		{
			name: "check for spanish",
			args: args{
				lang: "Hola",
			},
			want: Spanish{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newaudio(tt.args.lang); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newaudio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayerDetails(t *testing.T) {
	tests := []struct {
		name string
		want *player
	}{
		// TODO: Add test cases.
		{
			name: "define the player name ",
			want: &player{
				name: "ajit kumar",
				age:  25,
				perimeter: perimeter{
					width:  10,
					height: 20,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlayerDetails(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlayerDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_player_playingGames(t *testing.T) {
	type fields struct {
		name      string
		age       int64
		langu     string
		perimeter perimeter
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "checking the details ",
			fields: fields{
				name:  "ajit kumar",
				age:   25,
				langu: "English",
				perimeter: perimeter{
					width:  20,
					height: 20,
				},
			},
			want: "Hola",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &player{
				name:      tt.fields.name,
				age:       tt.fields.age,
				langu:     tt.fields.langu,
				perimeter: tt.fields.perimeter,
			}
			if got := p.playingGames(); got != tt.want {
				t.Errorf("player.playingGames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreatingFile(t *testing.T) {
	t.Run("testing the file path", func(t *testing.T) {
		//assert := assert.New(t)
		t.Parallel()
		name := CreatingFile().Name()
		//path, _ := os.Getwd()
		fileInfo, _ := os.Stat(name)

		got := name
		want := fileInfo.Name()
		if got != want {
			t.Errorf("got %q want %q", got, want)

		}
		//log.Println(baseP)
		//got := CreatingFile()

	})

}
