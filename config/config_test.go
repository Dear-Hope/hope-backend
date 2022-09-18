package config

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoadConfig(t *testing.T) {
	Convey("Test load config", t, func() {
		Convey("from env variables", func() {
			os.Setenv("DBCONFIG_HOST", "testhost")
			os.Setenv("DBCONFIG_NAME", "testing")

			cfg, err := LoadConfig(".")
			So(err, ShouldBeNil)
			So(cfg.DBConfig.Host, ShouldEqual, "testhost")
			So(cfg.DBConfig.Name, ShouldEqual, "testing")
		})
	})
}
