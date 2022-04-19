/*
Copyright © 2020 Romber Li <romber2001@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/pingcap/errors"
	"github.com/romberli/das/config"
	"github.com/romberli/das/global"
	"github.com/romberli/das/pkg/message"
	msgrouter "github.com/romberli/das/pkg/message/router"
	"github.com/romberli/das/router"
	"github.com/romberli/das/server"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start command",
	Long:  `start the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err           error
			pidFileExists bool
			isRunning     bool
		)

		// init config
		err = initConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf("%+v", message.NewMessage(message.ErrInitConfig, err)))
			os.Exit(constant.DefaultAbnormalExitCode)
		}

		// check pid file
		serverPidFile = viper.GetString(config.ServerPidFileKey)
		pidFileExists, err = linux.PathExists(serverPidFile)
		if err != nil {
			log.Errorf("%+v", message.NewMessage(message.ErrCheckServerPid, err))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
		if pidFileExists {
			isRunning, err = linux.IsRunningWithPidFile(serverPidFile)
			if err != nil {
				log.Errorf("%+v", message.NewMessage(message.ErrCheckServerRunningStatus, err))
				os.Exit(constant.DefaultAbnormalExitCode)
			}
			if isRunning {
				log.Errorf("%+v", message.NewMessage(message.ErrServerIsRunning, serverPidFile))
				os.Exit(constant.DefaultAbnormalExitCode)
			}
		}

		// check if runs in daemon mode
		daemon := viper.GetBool(config.DaemonKey)
		if daemon {
			// set daemon to false
			args = os.Args[1:]
			for i, arg := range os.Args[1:] {
				if config.TrimSpaceOfArg(arg) == config.DaemonArgTrue {
					args[i] = config.DaemonArgFalse
				}
			}

			// start server with new process
			startCommand := exec.Command(os.Args[constant.ZeroInt], args...)
			err = startCommand.Start()
			if err != nil {
				log.Errorf("%+v", message.NewMessage(message.ErrStartAsForeground, errors.Trace(err)))
				os.Exit(constant.DefaultAbnormalExitCode)
			}

			time.Sleep(time.Second)
			os.Exit(constant.DefaultNormalExitCode)
		} else {
			// get pid
			serverPid = os.Getpid()

			// save pid
			err = linux.SavePid(serverPid, serverPidFile, constant.DefaultFileMode)
			if err != nil {
				log.Errorf("%+v", message.NewMessage(message.ErrSavePidToFile, err))
				os.Exit(constant.DefaultAbnormalExitCode)
			}

			// init connection pool
			err = global.InitDASMySQLPool()
			if err != nil {
				log.Errorf("%+v", message.NewMessage(message.ErrInitConnectionPool, err))
				os.Exit(constant.DefaultAbnormalExitCode)
			}

			// init token auth
			ta := router.NewTokenAuthWithGlobal()
			tokens, err := ta.GetTokens()
			if err != nil {
				log.Errorf("%+v", message.NewMessage(msgrouter.ErrRouterGetHandlerFunc, err))
				os.Exit(constant.DefaultAbnormalExitCode)
			}
			// init router
			r := router.NewGinRouter()
			r.Use(ta.GetHandlerFunc(tokens))
			// init server
			s := server.NewServer(
				viper.GetString(config.ServerAddrKey),
				serverPidFile,
				viper.GetInt(config.ServerReadTimeoutKey),
				viper.GetInt(config.ServerWriteTimeoutKey),
				r,
			)
			// start server
			go s.Run()

			log.CloneStdoutLogger().Info(message.NewMessage(message.InfoServerStart, s.Addr(), serverPid, serverPidFile).Error())
			log.Info(message.NewMessage(message.InfoServerStart, s.Addr(), serverPid, serverPidFile).Error())

			// handle signal
			linux.HandleSignalsWithPidFile(serverPidFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
