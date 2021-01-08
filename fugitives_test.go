package fbi_test

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/fbi"
	"testing"
)

func TestFugitives_List(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Basic success", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fbi.Fugitives{}
			err := f.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// An example usage of listing all of the fugitives
func ExampleList() {
	f := new(fbi.Fugitives)
	f.List()
	j, err := json.MarshalIndent(&f, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}

func TestFugitives_Find(t *testing.T) {
	type args struct {
		opt *fbi.Options
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Basic Success", args{&fbi.Options{
			Title:        "HOME",
			FieldOffices: "",
			Page:         1,
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fbi.Fugitives{}
			err := f.Find(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
