package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	resp "github.com/teukumulya-ichsan/loan-api-rspattern/response"

	"github.com/teukumulya-ichsan/loan-api-rspattern/models"
)

func Test_service_Validate(t *testing.T) {
	type args struct {
		loan *models.Loan
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		// TODO 1: Add test cases.
		{
			name: "IsValidLengthKtp() when KTP length is 16",
			args: args{
				&models.Loan{
					Name:      "Ichsan",
					DateLoan:  "2020-01-01",
					Gender:    "Male",
					Ktp:       "1208040025100111",
					BirthDate: "1995-06-02",
					Amount:    2000000,
					Period:    6,
				},
			},
			wantErr: "",
		},
		{
			name: "IsValidLengthKtp() when KTP length is not 16",
			args: args{
				&models.Loan{
					Name:      "Ichsan",
					DateLoan:  "2020-01-01",
					Gender:    "Male",
					Ktp:       "1208040025100111a",
					BirthDate: "1995-06-02",
					Amount:    2000000,
					Period:    6,
				},
			},
			wantErr: "the Length KTP it must 16 character",
		},
		{
			name: "KtpIsNumber()",
			args: args{
				&models.Loan{
					Name:      "Ichsan",
					DateLoan:  "2020-01-01",
					Gender:    "Male",
					Ktp:       "120804002510111a",
					BirthDate: "1995-06-02",
					Amount:    2000000,
					Period:    6,
				},
			},
			wantErr: "KTP must the number",
		},
		{
			name: "IsNotEmpty()",
			args: args{
				&models.Loan{
					Name:      "",
					DateLoan:  "2020-01-01",
					Gender:    "Male",
					Ktp:       "120804002510111a",
					BirthDate: "1995-06-02",
					Amount:    2000000,
					Period:    6,
				},
			},
			wantErr: "The Name is cannot be null.",
		},
		{
			name: "verifyProvince()",
			args: args{
				&models.Loan{
					Name:      "Ichsan",
					DateLoan:  "2020-01-01",
					Gender:    "Male",
					Ktp:       "1008040025100111",
					BirthDate: "1995-06-02",
					Amount:    2000000,
					Period:    6,
				},
			},
			wantErr: "Provinsi dari KTP tidak di izinkan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// TODO 2. Setup testing service.
			testService := NewLoanServices()

			// TODO 3. Do Testing
			err := testService.Validate(tt.args.loan)

			// TODO 4. Data Asserting
			if tt.wantErr == "" {
				assert.Nil(t, err)
			} else {
				assert.Error(t, err, "Ini harusna error")
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetSummary(t *testing.T) {
	type args struct {
		loan []models.Loan
	}
	tests := []struct {
		name    string
		args    args
		want    []resp.Info
		wantErr bool
	}{
		// TODO 1: Add test cases.
		{
			name: "Test GetSummary()",
			args: args{
				loan: []models.Loan{
					{
						ID: 1,
						Name: "Ichsan",
						DateLoan: "2020-01-29T00:00:00+07:00",
						Gender: "Male",
						Ktp: "3522582509010004",
						BirthDate: "1995-06-02T00:00:00Z",
						Amount: 20000000,
						Period: 12,
					},
					{
						ID: 2,
						Name: "Mulia Ichsan",
						DateLoan: "2020-01-29T00:00:00+07:00",
						Gender: "Male",
						Ktp: "1208040025100111",
						BirthDate: "1995-06-02T00:00:00Z",
						Amount: 2000000,
						Period: 6,
					},
				},
			},
			want:   []resp.Info{
				{
					Count:2,
					Sum:22000000,
					Average:11000000,
				},
			} ,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO 2. Setup testing service.
			testService := NewLoanServices()

			// TODO 3. Do Testing
			err, _ := testService.GetSummary(tt.args.loan)

			// TODO 4. Data Asserting
			assert.Equal(t,tt.want,err)

		})
	}
}
