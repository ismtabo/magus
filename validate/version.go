package validate

import (
	"github.com/Masterminds/semver"
	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/manifest"
)

var (
	// constraint is the constraint of the manifest version.
	constraint = mustCompileConstraint("<=2.0.0")
)

// Version validates the version of the manifest.
func Version(ctx context.Context, m manifest.Manifest) error {
	if ok := constraint.Check(m.Version.Version); !ok {
		return WrapInvalidVersionErrorf("invalid version %s", m.Version.Version)
	}
	return nil
}

func mustCompileConstraint(s string) *semver.Constraints {
	c, err := semver.NewConstraint(s)
	if err != nil {
		panic(err)
	}
	return c
}
