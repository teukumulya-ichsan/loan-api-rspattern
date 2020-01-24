package repositories

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/teukumulya-ichsan/go-loan/src/loan/models"
)

func TestNewLoanRepositoryPg(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *loanRepositoryPg
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoanRepositoryPg(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoanRepositoryPg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loanRepositoryPg_FindAll(t *testing.T) {
	tests := []struct {
		name    string
		r       *loanRepositoryPg
		want    models.Loans
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("loanRepositoryPg.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loanRepositoryPg.FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loanRepositoryPg_FindByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		r       *loanRepositoryPg
		args    args
		want    *models.Loan
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.FindByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("loanRepositoryPg.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loanRepositoryPg.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loanRepositoryPg_CreateLoan(t *testing.T) {
	type args struct {
		ctx context.Context
		l   *models.Loan
	}
	tests := []struct {
		name    string
		r       *loanRepositoryPg
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.CreateLoan(tt.args.ctx, tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("loanRepositoryPg.CreateLoan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("loanRepositoryPg.CreateLoan() = %v, want %v", got, tt.want)
			}
		})
	}
}
