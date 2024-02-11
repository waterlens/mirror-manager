package rank

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/schollz/progressbar/v3"
)

func downloadAndMeasureSpeed(url string) (float64, error) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	n, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return 0, err
	}
	duration := time.Since(start)
	speed := float64(n) / duration.Seconds()
	return speed, nil
}

func Rank(mirrors []string) (string, error) {
	if len(mirrors) == 0 {
		return "", errors.New("no mirrors")
	}
	type result struct {
		mirror string
		speed  float64
		err    error
	}

	var ch []chan result
	var mlen = len(mirrors) + 1
	for i := 0; i < mlen; i++ {
		ch = append(ch, make(chan result))
	}

	for i, mirror := range mirrors {
		go func(mirror string, idx int) {
			speed, err := downloadAndMeasureSpeed(mirror)
			if err != nil {
				time.Sleep(10 * time.Second)
			}
			ch[idx] <- result{mirror, speed, err}
		}(mirror, i)
	}

	bar := progressbar.Default(int64(mlen) - 1)

	go func() {
		for i := 0; i < mlen-1; i++ {
			bar.Add(1)
			time.Sleep(1 * time.Second)
		}
		ch[mlen-1] <- result{"", 0, errors.New("fetching all mirrors timed out")}
	}()

	bar.Finish()

	cases := make([]reflect.SelectCase, mlen)
	for i := 0; i < mlen; i++ {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch[i])}
	}
	_, value, ok := reflect.Select(cases)
	if !ok {
		return "", errors.New("reflect.Select failed")
	}
	r := value.Interface().(result)
	if r.err != nil {
		return "", r.err
	}
	return r.mirror, nil
}
