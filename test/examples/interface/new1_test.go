package spike

import (
	"log"
	"os"
	"reflect"
	"testing"
)

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
			if got := NewAudio(tt.args.lang); !reflect.DeepEqual(got, tt.want) {
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
			if got := p.PlayingGames(); got != tt.want {
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
		log.Println(name)
		//path, _ := os.Getwd()
		fileInfo, _ := os.Stat(name)
		log.Println(fileInfo)

		got := name
		want := fileInfo.Name()
		if got != want {
			t.Errorf("got %q want %q", got, want)

		}
		//log.Println(baseP)
		//got := CreatingFile()

	})

}

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
		{
			name: "testing perimeter",
			fields: fields{
				width:  2.0,
				height: 3.1,
			},
			want: 10.2,
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

/*
func TestCreatingFile(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want *os.File
	}{
		// TODO: Add test cases.
		{
			name: "generating the test",
			args: args{
				val: "hi name is ajit ",
			},
			want: os.Stat(CreatingFile().Name()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreatingFile(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatingFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
//Without mocking

/*
func TestNewUser(t *testing.T) {
	type args struct {
		u User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "check avilable user ",
			args: args{
				User{
					Name:     "ajit kumar",
					Email:    "ajithkumar.sinha@srsconsultinginc.com",
					UserName: "ajit",
				},
			},
			wantErr: UserAvilable("ajithkumar.sinha@srsconsultinginc.com"),
		},
		// TODO: Add test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewUser(tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
*/

// with mocking
// we do mocking when we have to test the with the external dependency .

var userExistsMock func(email string) bool

type preCheckMock struct{}

func (u preCheckMock) userExists(email string) bool {
	return userExistsMock(email)
}
func TestNewUser(t *testing.T) {
	t.Parallel()
	user := User{
		Name:     "ajit kumar",
		Email:    "ajithkumar.sinha@srsconsultinginc.com",
		UserName: "ajit",
	}
	regCond = preCheckMock{}
	userExistsMock = func(email string) bool {
		return false
	}
	err1 := NewUser(user)
	if err1 != nil {
		t.Fatal(err1)
	}

	userExistsMock = func(email string) bool {
		return true
	}
	err2 := NewUser(user)
	if err2 == nil {
		t.Errorf("throw an error got nil")
	}

}

func TestEmployeeSalary(t *testing.T) {
	tests := []struct {
		name string
		//want  *os.File
		want1 bool
	}{
		// TODO: Add test cases.
		{
			name:  "testing the file size",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1 := EmployeeSalary()

			if got1 != tt.want1 {
				t.Errorf("EmployeeSalary() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
