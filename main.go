package main

import (
	"os"
	"strings"

	"github.com/skyaxl/gk/cmd"
	"github.com/skyaxl/gk/utils"
	"github.com/skyaxl/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	gosrc := utils.GetGOPATH() + afero.FilePathSeparator + "src" + afero.FilePathSeparator
	pwd, err := os.Getwd()
	if err != nil {
		logrus.Error(err)
		return
	}
	if !strings.HasPrefix(pwd, gosrc) {
		logrus.Error("The project must be in the $GOPATH/src folder for the generator to work.")
		return
	}
	cmd.Execute()
}
