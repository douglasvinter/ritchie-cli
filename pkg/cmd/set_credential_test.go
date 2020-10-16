/*
 * Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"errors"
	"testing"

	"github.com/ZupIT/ritchie-cli/pkg/credential"
	"github.com/ZupIT/ritchie-cli/pkg/prompt"
	"github.com/ZupIT/ritchie-cli/pkg/stream"
	sMocks "github.com/ZupIT/ritchie-cli/pkg/stream/mocks"
)

var creds = make(map[string][]credential.Field)

func Test_setCredentialCmd_runPrompt(t *testing.T) {
	type in struct {
		Setter        credential.Setter
		credFile      credential.ReaderWriterPather
		file          stream.FileReadExister
		InputText     prompt.InputText
		InputBool     prompt.InputBool
		InputList     prompt.InputList
		InputPassword prompt.InputPassword
	}
	var tests = []struct {
		name    string
		in      in
		wantErr bool
	}{
		{
			name: "success run with no data",
			in: in{
				Setter:    credSetterMock{},
				credFile:  credSettingsMock{},
				InputText: inputSecretMock{},
				InputBool: inputFalseMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return credential.AddNew, nil
					},
				},
				InputPassword: inputPasswordMock{},
			},
			wantErr: false,
		},
		{
			name: "success run with full data",
			in: in{
				Setter: credSetterMock{},
				credFile: credSettingsCustomMock{
					ReadCredentialsFieldsMock: func(path string) (credential.Fields, error) {
						cred := credential.Field{
							Name: "accesskeyid",
							Type: "plain text",
						}
						credArr := []credential.Field{}
						credArr = append(credArr, cred)
						creds["file"] = credArr
						return creds, nil
					},
				},
				file: sMocks.FileReadExisterCustomMock{
					ExistsMock: func(path string) bool {
						return true
					},
					ReadMock: func(path string) ([]byte, error) {
						return []byte("some data"), nil
					},
				},
				InputText: inputTextMock{},
				InputBool: inputFalseMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return "file", nil
					},
				},
				InputPassword: inputPasswordMock{},
			},
			wantErr: false,
		},
		{
			name: "fail text with full data and file input",
			in: in{
				Setter: credSetterMock{},
				credFile: credSettingsCustomMock{
					ReadCredentialsFieldsMock: func(path string) (credential.Fields, error) {
						cred := credential.Field{
							Name: "accesskeyid",
							Type: "plain text",
						}
						credArr := []credential.Field{}
						credArr = append(credArr, cred)
						creds["file"] = credArr
						return creds, nil
					},
				},
				file: sMocks.FileReadExisterCustomMock{
					ExistsMock: func(path string) bool {
						return true
					},
					ReadMock: func(path string) ([]byte, error) {
						return []byte("some data"), nil
					},
				},
				InputText: inputTextCustomMock{
					text: func(name string, required bool) (string, error) {
						return "", errors.New("text error")
					},
				},
				InputBool: inputFalseMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return "file", nil
					},
				},
				InputPassword: inputPasswordMock{},
			},
			wantErr: true,
		},
		{
			name: "fail to read file",
			in: in{
				Setter: credSetterMock{},
				credFile: credSettingsCustomMock{
					ReadCredentialsFieldsMock: func(path string) (credential.Fields, error) {
						cred := credential.Field{
							Name: "accesskeyid",
							Type: "plain text",
						}
						credArr := []credential.Field{}
						credArr = append(credArr, cred)
						creds["type"] = credArr
						return creds, nil
					},
				},
				file: sMocks.FileReadExisterCustomMock{
					ExistsMock: func(path string) bool {
						return true
					},
					ReadMock: func(path string) ([]byte, error) {
						return nil, errors.New("error reading file")
					},
				},
				InputText: inputTextMock{},
				InputBool: inputFalseMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return "file", nil
					},
				},
				InputPassword: inputPasswordMock{},
			},
			wantErr: true,
		},
		{
			name: "fail empty credential file",
			in: in{
				Setter: credSetterMock{},
				credFile: credSettingsCustomMock{
					ReadCredentialsFieldsMock: func(path string) (credential.Fields, error) {
						cred := credential.Field{
							Name: "accesskeyid",
							Type: "plain text",
						}
						credArr := []credential.Field{}
						credArr = append(credArr, cred)
						creds["type"] = credArr
						return creds, nil
					},
				},
				file: sMocks.FileReadExisterCustomMock{
					ExistsMock: func(path string) bool {
						return true
					},
					ReadMock: func(path string) ([]byte, error) {
						return []byte(""), nil
					},
				},
				InputText: inputTextMock{},
				InputBool: inputFalseMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return "file", nil
					},
				},
				InputPassword: inputPasswordMock{},
			},
			wantErr: true,
		},
		{
			name: "fail no file to read",
			in: in{
				Setter: credSetterMock{},
				credFile: credSettingsCustomMock{
					ReadCredentialsFieldsMock: func(path string) (credential.Fields, error) {
						cred := credential.Field{
							Name: "accesskeyid",
							Type: "plain text",
						}
						credArr := []credential.Field{}
						credArr = append(credArr, cred)
						creds["type"] = credArr
						return creds, nil
					},
				},
				file: sMocks.FileReadExisterCustomMock{
					ExistsMock: func(path string) bool {
						return false
					},
				},
				InputText: inputTextMock{},
				InputBool: inputFalseMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return "file", nil
					},
				},
				InputPassword: inputPasswordMock{},
			},
			wantErr: true,
		},
		{
			name: "fail cannot find any credential in path ",
			in: in{
				Setter: credSetterMock{},
				credFile: credSettingsCustomMock{
					ReadCredentialsFieldsMock: func(path string) (credential.Fields, error) {
						cred := credential.Field{
							Name: "accesskeyid",
							Type: "plain text",
						}
						credArr := []credential.Field{}
						credArr = append(credArr, cred)
						creds["type"] = credArr
						return creds, nil
					},
				},
				file: sMocks.FileReadExisterCustomMock{
					ExistsMock: func(path string) bool {
						return false
					},
					ReadMock: func(path string) ([]byte, error) {
						return []byte("some data"), nil
					},
				},
				InputText: inputTextCustomMock{
					text: func(name string, required bool) (string, error) {
						return "", errors.New("text error")
					},
				},
				InputBool: inputFalseMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return "type", nil
					},
				},
				InputPassword: inputPasswordMock{},
			},
			wantErr: true,
		},
		{
			name: "fail when password return err",
			in: in{
				Setter: credSetterMock{},
				credFile: credSettingsCustomMock{
					ReadCredentialsFieldsMock: func(path string) (credential.Fields, error) {
						cred := credential.Field{
							Name: "accesskeyid",
							Type: "secret",
						}
						credArr := []credential.Field{}
						credArr = append(credArr, cred)
						creds["type"] = credArr
						return creds, nil
					},
				},
				file: sMocks.FileReadExisterCustomMock{},
				InputText: inputTextCustomMock{
					text: func(name string, required bool) (string, error) {
						return "./path/to/my/credentialFile", nil
					},
				},
				InputBool: inputFalseMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return "type", nil
					},
				},
				InputPassword: inputPasswordErrorMock{},
			},
			wantErr: true,
		},
		{
			name: "fail when write credential fields return err",
			in: in{
				Setter: credSetterMock{},
				credFile: credSettingsCustomMock{
					ReadCredentialsFieldsMock: func(path string) (credential.Fields, error) {
						return credential.Fields{}, errors.New("error reading credentials")
					},
				},
				file: sMocks.FileReadExisterCustomMock{},
				InputText: inputTextCustomMock{
					text: func(name string, required bool) (string, error) {
						return "./path/to/my/credentialFile", nil
					},
				},
				InputBool: inputFalseMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return "type", nil
					},
				},
				InputPassword: inputPasswordErrorMock{},
			},
			wantErr: true,
		},
		{
			name: "fail when list return err",
			in: in{
				Setter:        credSetterMock{},
				credFile:      credSettingsMock{},
				InputText:     inputSecretMock{},
				InputBool:     inputFalseMock{},
				InputList:     inputListErrorMock{},
				InputPassword: inputPasswordErrorMock{},
			},
			wantErr: true,
		},
		{
			name: "fail when text return err",
			in: in{
				Setter:    credSetterMock{},
				credFile:  credSettingsMock{},
				InputText: inputTextErrorMock{},
				InputBool: inputFalseMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return credential.AddNew, nil
					},
				},
				InputPassword: inputPasswordMock{},
			},
			wantErr: true,
		},
		{
			name: "fail when text bool err",
			in: in{
				Setter:    credSetterMock{},
				credFile:  credSettingsMock{},
				InputText: inputTextMock{},
				InputBool: inputBoolErrorMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return credential.AddNew, nil
					},
				},
				InputPassword: inputPasswordMock{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewSetCredentialCmd(
				tt.in.Setter,
				tt.in.credFile,
				tt.in.file,
				tt.in.InputText,
				tt.in.InputBool,
				tt.in.InputList,
				tt.in.InputPassword,
			)
			cmd.PersistentFlags().Bool("stdin", false, "input by stdin")
			if err := cmd.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("set credential command error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setCredentialCmd_promptResolver(t *testing.T) {
	type fields struct {
		Setter          credential.Setter
		Settings        credential.Settings
		SingleSettings  credential.SingleSettings
		edition         api.Edition
		InputText       prompt.InputText
		InputBool       prompt.InputBool
		InputList       prompt.InputList
		InputPassword   prompt.InputPassword
		InputMultiline  prompt.InputMultiline
		FileReadExister stream.FileReadExister
	}
	tests := []struct {
		name    string
		fields  fields
		want    credential.Detail
		wantErr bool
	}{
		{
			name: "reach default",
			fields: fields{

			},
			want:    credential.Detail{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setCredentialCmd{
				Setter:          tt.fields.Setter,
				Settings:        tt.fields.Settings,
				SingleSettings:  tt.fields.SingleSettings,
				edition:         tt.fields.edition,
				InputText:       tt.fields.InputText,
				InputBool:       tt.fields.InputBool,
				InputList:       tt.fields.InputList,
				InputPassword:   tt.fields.InputPassword,
				InputMultiline:  tt.fields.InputMultiline,
				FileReadExister: tt.fields.FileReadExister,
			}
			got, err := s.promptResolver()
			if (err != nil) != tt.wantErr {
				t.Errorf("promptResolver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("promptResolver() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setCredentialCmd_singlePrompt(t *testing.T) {
	type fields struct {
		Setter          credential.Setter
		Settings        credential.Settings
		SingleSettings  credential.SingleSettings
		edition         api.Edition
		InputText       prompt.InputText
		InputBool       prompt.InputBool
		InputList       prompt.InputList
		InputPassword   prompt.InputPassword
		InputMultiline  prompt.InputMultiline
		FileReadExister stream.FileReadExister
	}
	tests := []struct {
		name    string
		fields  fields
		want    credential.Detail
		wantErr bool
	}{
		{
			name: "error on write default credentials",
			fields: fields{
				SingleSettings: singleCredSettingsCustomMock{
					writeDefaultCredentials: func(path string) error {
						return errors.New("some error")
					},
				},
			},
			want:    credential.Detail{},
			wantErr: true,
		},
		{
			name: "error on read credentials",
			fields: fields{
				SingleSettings: singleCredSettingsCustomMock{
					writeDefaultCredentials: func(path string) error {
						return nil
					},
					readCredentials: func(path string) (credential.Fields, error) {
						return nil, errors.New("some error")
					},
				},
			},
			want:    credential.Detail{},
			wantErr: true,
		},
		{
			name: "error on provider choose",
			fields: fields{
				SingleSettings: singleCredSettingsMock{},
				InputList:      inputListErrorMock{},
			},
			want:    credential.Detail{},
			wantErr: true,
		},
		{
			name: "error on input text",
			fields: fields{
				SingleSettings: singleCredSettingsMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return credsingle.AddNew, nil
					},
				},
				InputText: inputTextCustomMock{
					text: func(name string, required bool) (string, error) {
						return "", errors.New("some error")
					},
				},
			},
			want:    credential.Detail{},
			wantErr: true,
		},
		{
			name: "error on input bool",
			fields: fields{
				SingleSettings: singleCredSettingsMock{},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return credsingle.AddNew, nil
					},
				},
				InputText: inputTextMock{},
				InputBool: inputBoolCustomMock{
					bool: func(name string, items []string) (bool, error) {
						return false, errors.New("some error")
					},
				},
			},
			want:    credential.Detail{},
			wantErr: true,
		},
		{
			name: "error on write credentials",
			fields: fields{
				SingleSettings: singleCredSettingsCustomMock{
					writeDefaultCredentials: func(path string) error {
						return nil
					},
					readCredentials: func(path string) (credential.Fields, error) {
						return credential.Fields{}, nil
					},
					writeCredentials: func(fields credential.Fields, path string) error {
						return errors.New("some error")
					},
				},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return credsingle.AddNew, nil
					},
				},
				InputText: inputTextMock{},
				InputBool: inputFalseMock{},
			},
			want:    credential.Detail{},
			wantErr: true,
		},
		{
			name: "success with empty inputs",
			fields: fields{
				SingleSettings: singleCredSettingsCustomMock{
					writeDefaultCredentials: func(path string) error {
						return nil
					},
					readCredentials: func(path string) (credential.Fields, error) {
						return credential.Fields{}, nil
					},
					writeCredentials: func(fields credential.Fields, path string) error {
						return nil
					},
				},
				InputList: inputListCustomMock{
					list: func(name string, items []string) (string, error) {
						return credsingle.AddNew, nil
					},
				},
				InputText: inputTextMock{},
				InputBool: inputFalseMock{},
			},
			want: credential.Detail{
				Username: "",
				Credential: credential.Credential{
					"mocked text": "mocked text",
				},
				Service: "mocked text",
				Type:    "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setCredentialCmd{
				Setter:          tt.fields.Setter,
				Settings:        tt.fields.Settings,
				SingleSettings:  tt.fields.SingleSettings,
				edition:         tt.fields.edition,
				InputText:       tt.fields.InputText,
				InputBool:       tt.fields.InputBool,
				InputList:       tt.fields.InputList,
				InputPassword:   tt.fields.InputPassword,
				InputMultiline:  tt.fields.InputMultiline,
				FileReadExister: tt.fields.FileReadExister,
			}
			got, err := s.singlePrompt()
			if (err != nil) != tt.wantErr {
				t.Errorf("singlePrompt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singlePrompt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setCredentialCmd_inputFile(t *testing.T) {
	type fields struct {
		Setter          credential.Setter
		Settings        credential.Settings
		SingleSettings  credential.SingleSettings
		edition         api.Edition
		InputText       prompt.InputText
		InputBool       prompt.InputBool
		InputList       prompt.InputList
		InputPassword   prompt.InputPassword
		InputMultiline  prompt.InputMultiline
		FileReadExister stream.FileReadExister
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "run with success",
			fields: fields{
				InputText: inputTextMock{},
				FileReadExister: FileManagerMock{},
			},
			want:    "Some response",
			wantErr: false,
		},
		{
			name:    "error on input text",
			fields:  fields{
				InputText: inputTextCustomMock{
					text: func(name string, required bool) (string, error) {
						return "", errors.New("some error")
					},
				},
			},
			want:    "",
			wantErr: true,
		},
		{
			name:    "error on read file",
			fields:  fields{
				InputText: inputTextMock{},
				FileReadExister: FileManagerCustomMock{
					read: func(path string) ([]byte, error) {
						return nil, errors.New("some error")
					},
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setCredentialCmd{
				Setter:          tt.fields.Setter,
				Settings:        tt.fields.Settings,
				SingleSettings:  tt.fields.SingleSettings,
				edition:         tt.fields.edition,
				InputText:       tt.fields.InputText,
				InputBool:       tt.fields.InputBool,
				InputList:       tt.fields.InputList,
				InputPassword:   tt.fields.InputPassword,
				InputMultiline:  tt.fields.InputMultiline,
				FileReadExister: tt.fields.FileReadExister,
			}
			got, err := s.inputFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("inputFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("inputFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setCredentialCmd_stdinResolver(t *testing.T) {
	type fields struct {
		Setter          credential.Setter
		Settings        credential.Settings
		SingleSettings  credential.SingleSettings
		edition         api.Edition
		InputText       prompt.InputText
		InputBool       prompt.InputBool
		InputList       prompt.InputList
		InputPassword   prompt.InputPassword
		InputMultiline  prompt.InputMultiline
		FileReadExister stream.FileReadExister
	}
	tests := []struct {
		name    string
		fields  fields
		want    credential.Detail
		wantErr bool
	}{
		{
			name:    "run with empty edition",
			fields:  fields{
				edition: "",
			},
			want:    credential.Detail{},
			wantErr: true,
		},
		{
			name:    "error on stdin inputs",
			fields:  fields{
				edition: api.Team,
			},
			want:    credential.Detail{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := setCredentialCmd{
				Setter:          tt.fields.Setter,
				Settings:        tt.fields.Settings,
				SingleSettings:  tt.fields.SingleSettings,
				edition:         tt.fields.edition,
				InputText:       tt.fields.InputText,
				InputBool:       tt.fields.InputBool,
				InputList:       tt.fields.InputList,
				InputPassword:   tt.fields.InputPassword,
				InputMultiline:  tt.fields.InputMultiline,
				FileReadExister: tt.fields.FileReadExister,
			}
			got, err := s.stdinResolver()
			if (err != nil) != tt.wantErr {
				t.Errorf("stdinResolver() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stdinResolver() got = %v, want %v", got, tt.want)
			}
		})
	}
}