/*
Copyright 2016 The Rook Authors. All rights reserved.

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
package ceph

import (
	"github.com/rook/rook/cmd/rook/rook"
	"github.com/rook/rook/pkg/daemon/ceph/ganesha"
	"github.com/rook/rook/pkg/daemon/ceph/mon"
	"github.com/rook/rook/pkg/util/flags"
	"github.com/spf13/cobra"
)

var (
	exportPool   string
	exportObject string
)

var ganeshaCmd = &cobra.Command{
	Use:    "ganesha",
	Short:  "Generates config and runs the nfs ganesha daemon",
	Hidden: true,
}

func init() {
	ganeshaCmd.Flags().StringVar(&exportPool, "nfs-export-pool", "", "name of the pool in which nfs exports are stored")
	ganeshaCmd.Flags().StringVar(&exportObject, "nfs-export-object", "", "name of the object in which nfs exports are stored in the pool")
	addCephFlags(ganeshaCmd)

	flags.SetFlagsFromEnv(ganeshaCmd.Flags(), rook.RookEnvVarPrefix)

	ganeshaCmd.RunE = startGanesha
}

func startGanesha(cmd *cobra.Command, args []string) error {
	required := []string{"mon-endpoints", "cluster-name", "admin-secret", "nfs-export-pool", "nfs-export-object", "public-ipv4", "private-ipv4"}
	if err := flags.VerifyRequiredFlags(ganeshaCmd, required); err != nil {
		return err
	}

	rook.SetLogLevel()

	rook.LogStartupInfo(ganeshaCmd.Flags())

	clusterInfo.Monitors = mon.ParseMonEndpoints(cfg.monEndpoints)
	config := &ganesha.Config{
		Pool:        exportPool,
		Object:      exportObject,
		ClusterInfo: &clusterInfo,
	}

	err := ganesha.Run(createContext(), config)
	if err != nil {
		rook.TerminateFatal(err)
	}

	return nil
}
