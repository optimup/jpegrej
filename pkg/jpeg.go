/*
Copyright © 2020 Joel Curtis <joel@curti.se>

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

// Package pkg x
package pkg

import (
	"bytes"
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
)

var (
	jpegstart = []byte{0xff, 0xda}
	jpegend   = []byte{0xff, 0xd9}
)

// Jpegreturn test
type Jpegreturn struct {
	Start  int
	End    int
	Data   []byte
	Path   string
	seed   int64
	//rand   *rand.Rand
	amount int64
}

// Jpegload test
func Jpegload(path string) (*Jpegreturn, error) {
	if stat, err := os.Stat(path); err != nil || stat.IsDir() {
		return nil, errors.New("Could not find specified file.")
	}

	jpeg := new(Jpegreturn)
	var err error

	jpeg.Data, err = ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if i := bytes.Index(jpeg.Data, jpegstart); i > 0 {
		jpeg.Start = i
	} else {
		return nil, errors.New("Can't find start")
	}

	if i := bytes.Index(jpeg.Data, jpegend); i > 0 {
		jpeg.End = i
	} else {
		return nil, errors.New("Can't find end")
	}

	jpeg.Path = path

	return jpeg, nil
}

func (j *Jpegreturn) Load(path string) error {
	if stat, err := os.Stat(path); err != nil || stat.IsDir() {
		return errors.New("Could not find specified file.")
	}

	var err error

	j.Data, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if i := bytes.Index(j.Data, jpegstart); i > 0 {
		j.Start = i
	} else {
		return errors.New("Can't find start")
	}

	if i := bytes.Index(j.Data, jpegend); i > 0 {
		j.End = i
	} else {
		return errors.New("Can't find end")
	}

	j.Path = path
	return nil
}

// Seed saves the seed
func (j *Jpegreturn) Seed(seed int64, replace int64) {
	j.seed = seed
	rand.Seed(j.seed)
	//j.rand = rand.New(rand.NewSource(seed))
	j.amount = replace
}

// Mosh replaces bytes in range
func (j *Jpegreturn) Mosh(filepath string) error {
	buf := make([]byte, j.amount)
	nbytes, err := rand.Read(buf)
	if err != nil || nbytes != len(buf) {
		return errors.New("No bytes in range")
	}

	for _, b := range buf {
		loc := rand.Intn(j.End - j.Start)
		j.Data[j.Start:j.End][loc] = b
	}

	if err := ioutil.WriteFile(filepath, j.Data, 0660); err != nil {
		return err
	}

	return nil
}
