package play

import (
	"errors"
	"net/url"
	"reflect"
	"strings"

	"github.com/containers/podman/v2/pkg/bindings/util"
	jsoniter "github.com/json-iterator/go"
)

/*
This file is generated automatically by go generate.  Do not edit.
*/

// Changed
func (o *KubeOptions) Changed(fieldName string) bool {
	r := reflect.ValueOf(o)
	value := reflect.Indirect(r).FieldByName(fieldName)
	return !value.IsNil()
}

// ToParams
func (o *KubeOptions) ToParams() (url.Values, error) {
	params := url.Values{}
	if o == nil {
		return params, nil
	}
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	s := reflect.ValueOf(o)
	if reflect.Ptr == s.Kind() {
		s = s.Elem()
	}
	sType := s.Type()
	for i := 0; i < s.NumField(); i++ {
		fieldName := sType.Field(i).Name
		if !o.Changed(fieldName) {
			continue
		}
		fieldName = strings.ToLower(fieldName)
		f := s.Field(i)
		if reflect.Ptr == f.Kind() {
			f = f.Elem()
		}
		switch {
		case util.IsSimpleType(f):
			params.Set(fieldName, util.SimpleTypeToParam(f))
		case f.Kind() == reflect.Slice:
			for i := 0; i < f.Len(); i++ {
				elem := f.Index(i)
				if util.IsSimpleType(elem) {
					params.Add(fieldName, util.SimpleTypeToParam(elem))
				} else {
					return nil, errors.New("slices must contain only simple types")
				}
			}
		case f.Kind() == reflect.Map:
			lowerCaseKeys := make(map[string][]string)
			iter := f.MapRange()
			for iter.Next() {
				lowerCaseKeys[iter.Key().Interface().(string)] = iter.Value().Interface().([]string)

			}
			s, err := json.MarshalToString(lowerCaseKeys)
			if err != nil {
				return nil, err
			}

			params.Set(fieldName, s)
		}

	}
	return params, nil
}

// WithAuthfile
func (o *KubeOptions) WithAuthfile(value string) *KubeOptions {
	v := &value
	o.Authfile = v
	return o
}

// GetAuthfile
func (o *KubeOptions) GetAuthfile() string {
	var authfile string
	if o.Authfile == nil {
		return authfile
	}
	return *o.Authfile
}

// WithCertDir
func (o *KubeOptions) WithCertDir(value string) *KubeOptions {
	v := &value
	o.CertDir = v
	return o
}

// GetCertDir
func (o *KubeOptions) GetCertDir() string {
	var certDir string
	if o.CertDir == nil {
		return certDir
	}
	return *o.CertDir
}

// WithUsername
func (o *KubeOptions) WithUsername(value string) *KubeOptions {
	v := &value
	o.Username = v
	return o
}

// GetUsername
func (o *KubeOptions) GetUsername() string {
	var username string
	if o.Username == nil {
		return username
	}
	return *o.Username
}

// WithPassword
func (o *KubeOptions) WithPassword(value string) *KubeOptions {
	v := &value
	o.Password = v
	return o
}

// GetPassword
func (o *KubeOptions) GetPassword() string {
	var password string
	if o.Password == nil {
		return password
	}
	return *o.Password
}

// WithNetwork
func (o *KubeOptions) WithNetwork(value string) *KubeOptions {
	v := &value
	o.Network = v
	return o
}

// GetNetwork
func (o *KubeOptions) GetNetwork() string {
	var network string
	if o.Network == nil {
		return network
	}
	return *o.Network
}

// WithQuiet
func (o *KubeOptions) WithQuiet(value bool) *KubeOptions {
	v := &value
	o.Quiet = v
	return o
}

// GetQuiet
func (o *KubeOptions) GetQuiet() bool {
	var quiet bool
	if o.Quiet == nil {
		return quiet
	}
	return *o.Quiet
}

// WithSignaturePolicy
func (o *KubeOptions) WithSignaturePolicy(value string) *KubeOptions {
	v := &value
	o.SignaturePolicy = v
	return o
}

// GetSignaturePolicy
func (o *KubeOptions) GetSignaturePolicy() string {
	var signaturePolicy string
	if o.SignaturePolicy == nil {
		return signaturePolicy
	}
	return *o.SignaturePolicy
}

// WithSkipTLSVerify
func (o *KubeOptions) WithSkipTLSVerify(value bool) *KubeOptions {
	v := &value
	o.SkipTLSVerify = v
	return o
}

// GetSkipTLSVerify
func (o *KubeOptions) GetSkipTLSVerify() bool {
	var skipTLSVerify bool
	if o.SkipTLSVerify == nil {
		return skipTLSVerify
	}
	return *o.SkipTLSVerify
}

// WithSeccompProfileRoot
func (o *KubeOptions) WithSeccompProfileRoot(value string) *KubeOptions {
	v := &value
	o.SeccompProfileRoot = v
	return o
}

// GetSeccompProfileRoot
func (o *KubeOptions) GetSeccompProfileRoot() string {
	var seccompProfileRoot string
	if o.SeccompProfileRoot == nil {
		return seccompProfileRoot
	}
	return *o.SeccompProfileRoot
}

// WithConfigMaps
func (o *KubeOptions) WithConfigMaps(value []string) *KubeOptions {
	v := &value
	o.ConfigMaps = v
	return o
}

// GetConfigMaps
func (o *KubeOptions) GetConfigMaps() []string {
	var configMaps []string
	if o.ConfigMaps == nil {
		return configMaps
	}
	return *o.ConfigMaps
}

// WithLogDriver
func (o *KubeOptions) WithLogDriver(value string) *KubeOptions {
	v := &value
	o.LogDriver = v
	return o
}

// GetLogDriver
func (o *KubeOptions) GetLogDriver() string {
	var logDriver string
	if o.LogDriver == nil {
		return logDriver
	}
	return *o.LogDriver
}

// WithStart
func (o *KubeOptions) WithStart(value bool) *KubeOptions {
	v := &value
	o.Start = v
	return o
}

// GetStart
func (o *KubeOptions) GetStart() bool {
	var start bool
	if o.Start == nil {
		return start
	}
	return *o.Start
}
