package golang

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestValidateCommand(t *testing.T) {
	tests := []struct {
		name       string
		modulePath string
		wantErr    bool
	}{
		{
			name:       "emptyModule",
			modulePath: "",
			wantErr:    true,
		},
		{
			name:       "validModule",
			modulePath: "github.com/foo/bar",
			wantErr:    false,
		},
		{
			name:       "invalidModule",
			modulePath: "invalid-module-path",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("module_path", tt.modulePath, "")

			err := ValidateCommand(cmd, []string{})
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCommand_NoModuleFlag(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	cmd := &cobra.Command{}
	ValidateCommand(cmd, []string{})
}

func Test_validateModulePath(t *testing.T) {
	type args struct {
		modulePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "validModulePath",
			args: args{
				modulePath: "github.com/foo/bar",
			},
			wantErr: false,
		},
		{
			name: "invalidModulePath",
			args: args{
				modulePath: "invalid/module/path",
			},
			wantErr: true,
		},
		{
			name: "emptyModulePath",
			args: args{
				modulePath: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateModulePath(tt.args.modulePath); (err != nil) != tt.wantErr {
				t.Errorf("validateModulePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
