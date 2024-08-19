package generator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortLinkGenerator(t *testing.T) {
	link1 := "https://rutube.ru/video/c4e6290ace7933b84b73810771a6f31e/"
	shortLink1, _ := GenerateShortLink(link1)

	link2 := "https://habr.com/ru/companies/otus/articles/739468/"
	shortLink2, _ := GenerateShortLink(link2)

	link3 := "https://habr.com/ru/articles/837090/"
	shortLink3, _ := GenerateShortLink(link3)

	assert.Equal(t, shortLink1, "XkXzo63P")
	assert.Equal(t, shortLink2, "TjBqM2CA")
	assert.Equal(t, shortLink3, "bWMcZjo5")
}
