package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_DefaultTheme(t *testing.T) {
	c := NewConfig(CliTestContext())

	assert.Equal(t, "default", c.DefaultTheme())
	c.options.Demo = false
	c.options.Sponsor = false
	c.options.Test = false
	c.options.DefaultTheme = "grayscale"
	assert.Equal(t, "default", c.DefaultTheme())
	c.options.Sponsor = true
	assert.Equal(t, "grayscale", c.DefaultTheme())
	c.options.Sponsor = false
	c.options.Test = true
	assert.Equal(t, "default", c.DefaultTheme())
	c.options.Sponsor = false
	c.options.Test = false
	assert.Equal(t, "default", c.DefaultTheme())
	c.options.Sponsor = true
	c.options.DefaultTheme = ""
	assert.Equal(t, "default", c.DefaultTheme())
	c.options.Sponsor = false
	assert.Equal(t, "default", c.DefaultTheme())
}

func TestConfig_DefaultLocale(t *testing.T) {
	c := NewConfig(CliTestContext())

	assert.Equal(t, "en", c.DefaultLocale())
	c.options.DefaultLocale = "de"
	assert.Equal(t, "de", c.DefaultLocale())
	c.options.DefaultLocale = ""
	assert.Equal(t, "en", c.DefaultLocale())
}

func TestConfig_WallpaperUri(t *testing.T) {
	c := NewConfig(CliTestContext())

	assert.Equal(t, "", c.WallpaperUri())
	c.options.WallpaperUri = "kashmir"
	assert.Equal(t, "/static/img/wallpaper/kashmir.jpg", c.WallpaperUri())
	c.options.WallpaperUri = "https://cdn.photoprism.app/wallpaper/welcome.jpg"
	assert.Equal(t, "https://cdn.photoprism.app/wallpaper/welcome.jpg", c.WallpaperUri())
	c.options.Test = false
	assert.Equal(t, "https://cdn.photoprism.app/wallpaper/welcome.jpg", c.WallpaperUri())
	c.options.Test = true
	assert.Equal(t, "https://cdn.photoprism.app/wallpaper/welcome.jpg", c.WallpaperUri())
	c.options.Sponsor = false
	assert.Equal(t, "", c.WallpaperUri())
	c.options.Sponsor = true
	assert.Equal(t, "https://cdn.photoprism.app/wallpaper/welcome.jpg", c.WallpaperUri())
	c.options.WallpaperUri = "kashmir"
	assert.Equal(t, "/static/img/wallpaper/kashmir.jpg", c.WallpaperUri())
	c.options.WallpaperUri = ""
	assert.Equal(t, "", c.WallpaperUri())
}
