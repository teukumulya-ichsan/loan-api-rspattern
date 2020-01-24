package controllers

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/teukumulya-ichsan/go-loan/config"
)

func TestNewLoanController(t *testing.T) {
	type args struct {
		db *config.DB
	}
	tests := []struct {
		name string
		args args
		want *Loan
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoanController(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoanController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoan_FindAll(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		p    *Loan
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.FindAll(tt.args.w, tt.args.r)
		})
	}
}

func TestLoan_FindByID(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		p    *Loan
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.FindByID(tt.args.w, tt.args.r)
		})
	}
}

func TestLoan_CreateLoan(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		p    *Loan
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.CreateLoan(tt.args.w, tt.args.r)
		})
	}
}

func Test_respondwithJSON(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		code    int
		payload interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondwithJSON(tt.args.w, tt.args.code, tt.args.payload)
		})
	}
}

func Test_respondWithError(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		code int
		msg  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondWithError(tt.args.w, tt.args.code, tt.args.msg)
		})
	}
}
