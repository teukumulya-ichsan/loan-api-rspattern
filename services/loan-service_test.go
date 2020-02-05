package services

import (
	"github.com/stretchr/testify/assert"
	"testing"

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
			}else {
				assert.Error(t,err, "Ini harusna error")
				assert.EqualError(t, err, tt.wantErr)
			}
		})
	}
}
