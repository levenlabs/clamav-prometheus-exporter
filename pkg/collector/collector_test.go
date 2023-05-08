package collector

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStat(t *testing.T) {
	matches := [][]string{{"998", "098"}, {"300", "sup"}, {"456", "123"}}

	{ // In Bounds
		stat := GetStat(matches, 2)
		assert.Equal(t, 123., stat)
	}

	{ // Out of Bounds
		stat := GetStat(matches, 4)
		assert.True(t, math.IsNaN(stat))
	}

	{ // Not a Float
		stat := GetStat(matches, 1)
		assert.True(t, math.IsNaN(stat))
	}
}

func TestGetVersionInfo(t *testing.T) {
	var float_64 float64

	{ // Test Full Version String
		version_info := GetVersionInfo("ClamAV 1.0.1/2.5/Mon Aug 13 08:23:14 2012")

		assert.Equal(t, true, version_info.versions_parsed)
		assert.Equal(t, "1.0.1", version_info.clamav_version)
		assert.Equal(t, "2.5", version_info.database_version)
		assert.IsType(t, float_64, version_info.database_age)
		assert.NotEqual(t, float64(0), version_info.database_age)
		assert.False(t, math.IsNaN(version_info.database_age))
	}

	{ // Test Limited Version String
		version_info := GetVersionInfo("ClamAV 1.0.1")

		assert.Equal(t, true, version_info.versions_parsed)
		assert.Equal(t, "1.0.1", version_info.clamav_version)
		assert.Equal(t, "", version_info.database_version)
		assert.True(t, math.IsNaN(version_info.database_age))
	}

	{ // Test Invalid Version numbers
		version_info := GetVersionInfo("ClamAV numbers")

		assert.Equal(t, true, version_info.versions_parsed)
		assert.Equal(t, "", version_info.clamav_version)
		assert.Equal(t, "", version_info.database_version)
		assert.True(t, math.IsNaN(version_info.database_age))

	}

	{ // Test Invalid
		version_info := GetVersionInfo("adfsdf 1238674")

		assert.Equal(t, false, version_info.versions_parsed)
		assert.Equal(t, "", version_info.clamav_version)
		assert.Equal(t, "", version_info.database_version)
		assert.True(t, math.IsNaN(version_info.database_age))
	}
}
