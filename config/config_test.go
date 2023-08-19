package config

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	Convey("Test get config", t, func() {
		Convey("empty", func() {
			cfg := Get()

			So(cfg.Database, ShouldBeZeroValue)
		})
	})
}

func TestLoad(t *testing.T) {
	Convey("Test load config", t, func() {
		Convey("success", func() {
			err := Load(WithConfigFolder("../config/"))
			cfg := Get()

			So(err, ShouldBeNil)
			So(cfg, ShouldNotBeEmpty)
		})
	})
}
