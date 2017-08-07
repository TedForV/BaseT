package BaseT

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestContains_NoSensitive1(t *testing.T) {
	arr := []string{"a","b","c","ab","ac","bc"}
	target := "A"
	assert.Equal(t,true,Contains(&arr,target,false))
}

func TestContains_NoSensitive2(t *testing.T) {
	arr := []string{"a","b","c","ab","ac","bc"}
	target := "Ab"
	assert.Equal(t,true,Contains(&arr,target,false))
}

func TestContains_Sensitive1(t *testing.T) {
	arr := []string{"a","b","c","ab","ac","bc"}
	target := "A"
	assert.Equal(t,false,Contains(&arr,target,true))
}

func TestContains_Sensitive2(t *testing.T) {
	arr := []string{"a","b","c","ab","ac","bc"}
	target := "Ab"
	assert.Equal(t,false,Contains(&arr,target,true))
}
