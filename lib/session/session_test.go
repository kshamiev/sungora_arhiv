package session

import (
	"testing"
	"time"
)

func TestSession(t *testing.T) {
	storage := NewSession(time.Second * 5)
	ses := storage.GetSession("storage")

	// SET
	ses.Set("key 1", "value 1")
	ses.Set("key 2", "value 2")
	ses.Set("key 3", "value 3")
	time.Sleep(time.Second * 3)

	// CHECK VALUE
	if _, ok := ses.Get("key 1").(string); !ok {
		t.Fatal("error get value from 'key 1'")
	}

	if _, ok := ses.Get("key 2").(string); !ok {
		t.Fatal("error get value from 'key 2'")
	}

	if _, ok := ses.Get("key 3").(string); !ok {
		t.Fatal("error get value from 'key 3'")
	}

	// CHECK DELETE
	ses.Del("key 2")

	if _, ok := ses.Get("key 2").(string); ok {
		t.Fatal("error delete value from 'key 2'")
	}
}
