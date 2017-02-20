package client

import (
	_ "flag"
	"fmt"
	"os"
	//"sync"
	"testing"
)

/*
var (
    cw *ClientWrapper
    //once sync.Once
    tc *testing.T
    //done chan bool := make(chan bool)
)
*/

/*
func init() {
    flag.StringVar(&server, "server", "", "server")
    flag.StringVar(&user, "user", "", "user")
    flag.StringVar(&pass, "pass", "", "pass")
    flag.Parse()
}
*/

/*
func onceBody() {
    server := os.Getenv("zzzserver")
    user := os.Getenv("zzzuser")
    pass := os.Getenv("zzzpass")

    if server =="" || user == "" || pass == "" {
        tc.Errorf("no enough args")
    }

    cw, err := NewClientWrapper(server, user, pass)
    if err != nil {
        tc.Errorf(err.Error())
    }

    if cw == nil {
        tc.Errorf("no client")
    }
}
*/

func TestReplicationController_List(t *testing.T) {
	//tc = t
	//once.Do(onceBody)

	server := os.Getenv("zzzserver")
	user := os.Getenv("zzzuser")
	pass := os.Getenv("zzzpass")

	if server == "" || user == "" || pass == "" {
		t.Errorf("no enough args")
	}

	cw, err := NewClientWrapper(server, user, pass)
	if err != nil {
		t.Errorf(err.Error())
	}

	if rclist, err := cw.FindReplicationControllerList(); err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(rclist, "\n----")
	}
}

func TestGetRC(t *testing.T) {
	//tc = t
	//once.Do(onceBody)

	server := os.Getenv("zzzserver")
	user := os.Getenv("zzzuser")
	pass := os.Getenv("zzzpass")

	if server == "" || user == "" || pass == "" {
		t.Errorf("no enough args")
	}

	cw, err := NewClientWrapper(server, user, pass)
	if err != nil {
		t.Errorf(err.Error())
	}

	if rc, err := cw.FindReplicationController("my-nginx"); err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(rc, "\n----")
	}
}

func TestListSVC(t *testing.T) {

	server := os.Getenv("zzzserver")
	user := os.Getenv("zzzuser")
	pass := os.Getenv("zzzpass")

	if server == "" || user == "" || pass == "" {
		t.Errorf("no enough args")
	}

	cw, err := NewClientWrapper(server, user, pass)
	if err != nil {
		t.Errorf(err.Error())
	}

	if svcList, err := cw.FindServiceList(); err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(svcList, "\n----")
	}
}

func TestGetSVC(t *testing.T) {

	server := os.Getenv("zzzserver")
	user := os.Getenv("zzzuser")
	pass := os.Getenv("zzzpass")

	if server == "" || user == "" || pass == "" {
		t.Errorf("no enough args")
	}

	cw, err := NewClientWrapper(server, user, pass)
	if err != nil {
		t.Errorf(err.Error())
	}

	if svc, err := cw.FindReplicationController("example-nginx"); err != nil {
		t.Errorf(err.Error())
	} else {
		fmt.Println(svc, "\n----")
	}
}
