// Copyright (c) 2016 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package loadtestconfig

import (
	"reflect"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type LoadTestConfig struct {
	UserCreationConfiguration    UserCreationConfiguration
	TeamCreationConfiguration    TeamCreationConfiguration
	ChannelCreationConfiguration ChannelCreationConfiguration
	ConnectionConfiguration      ConnectionConfiguration
	UserEntitiesConfiguration    UserEntitiesConfiguration
}

func (config *LoadTestConfig) setDefaultsIfRequired() {
	config.UserCreationConfiguration.SetDefaultsIfRequired()
	config.TeamCreationConfiguration.SetDefaultsIfRequired()
	config.ChannelCreationConfiguration.SetDefaultsIfRequired()
	config.UserEntitiesConfiguration.SetDefaultsIfRequired()
}

func SetupConfig() error {
	viper.SetConfigName("loadtestconfig")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() *LoadTestConfig {
	var config LoadTestConfig
	UnmarshalConfigStruct(&config)

	config.setDefaultsIfRequired()

	return &config
}

func UnmarshalConfigStruct(configStruct interface{}) error {
	return viper.Unmarshal(configStruct)
}

func UnmarshalConfigSubStruct(configStruct interface{}) error {
	return viper.Sub(reflect.ValueOf(configStruct).Elem().Type().Name()).Unmarshal(configStruct)
}

func SetIntFlag(flags *pflag.FlagSet, full, short, helpText, configFileSetting string, defaultValue int) {
	flags.IntP(full, short, defaultValue, helpText)
	viper.SetDefault(configFileSetting, defaultValue)
	viper.BindPFlag(configFileSetting, flags.Lookup(full))
}

func SetBoolFlag(flags *pflag.FlagSet, full, short, helpText, configFileSetting string, defaultValue bool) {
	flags.BoolP(full, short, defaultValue, helpText)
	viper.SetDefault(configFileSetting, defaultValue)
	viper.BindPFlag(configFileSetting, flags.Lookup(full))
}
