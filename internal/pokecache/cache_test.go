package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddAndGet(t *testing.T) {
	const interval = 10 * time.Millisecond
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "example.com/",
			val: []byte("testdata"),
		},
		{
			key: "example.com/",
			val: []byte("moretestdata"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case: %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestAddOverwrite(t *testing.T) {
	const interval = 1 * time.Second
	cases := []struct {
		key     string
		val_old []byte
		val_new []byte
	}{
		{
			key:     "google.com/hellow",
			val_old: []byte("orld"),
			val_new: []byte("hola hola!"),
		},
		{
			key:     "https://go.dev/doc/tutorial/add-a-test",
			val_old: []byte("i hope you pass them ALL!"),
			val_new: []byte("You did!"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case: %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val_old)
			cache.Add(c.key, c.val_new)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
			} else if string(val) != string(c.val_new) {
				t.Errorf("expected the value to be overwritten: expected - %v, got - %v", string(c.val_new), string(val))
			}
		})
	}
}

func TestReepLoop(t *testing.T) {
	const interval = 10 * time.Millisecond
	const baseTime = 5 * time.Millisecond
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "example.com/",
			val: []byte("testcaseuno"),
		},
		{
			key: "pipetka.ru/",
			val: []byte("ultimatepipetka"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case: %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected the key to still exist")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected key to still exist")
			}
			time.Sleep(interval + baseTime)
			_, ok = cache.Get(c.key)
			if ok {
				t.Errorf("expected the key be gone")
				return
			}
		})
	}
}

func TestGetFromEmpty(t *testing.T) {
	const interval = 10 * time.Millisecond
	const baseTime = 50 * time.Millisecond
	t.Run("test Get from new cache", func(t *testing.T) {
		cache := NewCache(interval)
		val, ok := cache.Get("missing")
		if ok {
			t.Errorf("expected to find nothing, found: %v", val)
		}
	})

	t.Run("test Get after expiration", func(t *testing.T) {
		cache := NewCache(interval)
		cache.Add("my key", []byte("my value"))
		time.Sleep(interval + baseTime)
		val, ok := cache.Get("missing")
		if ok {
			t.Errorf("expected to find nothing, found: %v", val)
		}
	})
}

// just run with 'go test -race'
func TestConcurrency(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)
	key := "pokeapi"

	go func() {
		for i := 0; i < 1000; i++ {
			cache.Add(key, []byte(fmt.Sprintf("val: %v", i)))
		}
	}()

	for i := 0; i < 1000; i++ {
		cache.Get(key)
	}
}
