// +build mage

package main

import (
        "github.com/elek/magelib"
)

func Build() error {
   url := "https://github.com/chrislusf/seaweedfs/releases/download/1.77/linux_amd64.tar.gz"
   cache, err := flokkr.DownloadUrlIfRequired(url, "1.77", make([]string,0), false)
   if err != nil {
      return err
   }
   return flokkr.BuildImage("37-light", "elek/seaweedfs", ".",cache)
}
