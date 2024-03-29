package validate_test

import (
	"testing"

	"github.com/Masterminds/semver"
	"github.com/ismtabo/magus/v2/context"
	"github.com/ismtabo/magus/v2/manifest"
	"github.com/ismtabo/magus/v2/validate"
)

func TestValidateVersion(t *testing.T) {
	t.Run("it should return nil if version is valid", func(t *testing.T) {
		mf := manifest.Manifest{
			Version: manifest.Version{
				Version: semver.MustParse("2"),
			},
		}
		err := validate.Version(context.New(), mf)
		if err != nil {
			t.Errorf("expected nil, got %v", err)
		}
	})

	t.Run("it should return error if version is invalid", func(t *testing.T) {
		mf := manifest.Manifest{
			Version: manifest.Version{
				Version: semver.MustParse("1.0.0"),
			},
		}
		err := validate.Version(context.New(), mf)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
