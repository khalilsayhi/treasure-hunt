package tests

import (
	"github.com/khalilsayhi/treasure-hunt/domain/enum"
	"github.com/khalilsayhi/treasure-hunt/utils"
	"testing"
)

func TestIsRotation(t *testing.T) {
	tests := []struct {
		name string
		move enum.PlayerMove
		want bool
	}{
		{"Rotation Droite", enum.Droite, true},
		{"Rotation Gauche", enum.Gauche, true},
		{"Avancer", enum.Avancer, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.IsRotation(tt.move); got != tt.want {
				t.Errorf("IsRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNextCoordinates(t *testing.T) {
	type args struct {
		x           int
		y           int
		orientation enum.PlayerOrientation
	}
	tests := []struct {
		name string
		move enum.PlayerMove
		args args
		outX int
		outY int
	}{
		{"Nord", enum.Avancer, args{1, 1, enum.Nord}, 1, 0},
		{"Est", enum.Avancer, args{1, 1, enum.Est}, 2, 1},
		{"Sud", enum.Avancer, args{1, 1, enum.Sud}, 1, 2},
		{"Ouest", enum.Avancer, args{1, 1, enum.Ouest}, 0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x, y := utils.GetNextCoordinates(tt.args.x, tt.args.y, tt.args.orientation)
			if x != tt.outX || y != tt.outY {
				t.Errorf("GetNextCoordinates() = (%d-%d), want (%d-%d)", x, y, tt.outX, tt.outY)
			}
		})
	}
}

func TestGetNextOrientation(t *testing.T) {
	type args struct {
		currentOrientation enum.PlayerOrientation
		rotation           enum.PlayerMove
	}
	tests := []struct {
		name string
		args args
		want enum.PlayerOrientation
	}{
		{"Nord Droite", args{enum.Nord, enum.Droite}, enum.Est},
		{"Nord Gauche", args{enum.Nord, enum.Gauche}, enum.Ouest},
		{"Est Droite", args{enum.Est, enum.Droite}, enum.Sud},
		{"Est Gauche", args{enum.Est, enum.Gauche}, enum.Nord},
		{"Sud Droite", args{enum.Sud, enum.Droite}, enum.Ouest},
		{"Sud Gauche", args{enum.Sud, enum.Gauche}, enum.Est},
		{"Ouest Droite", args{enum.Ouest, enum.Droite}, enum.Nord},
		{"Ouest Gauche", args{enum.Ouest, enum.Gauche}, enum.Sud},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.GetNextOrientation(tt.args.currentOrientation, tt.args.rotation); got != tt.want {
				t.Errorf("GetNextOrientation() = %v, want %v", got, tt.want)
			}
		})
	}
}
