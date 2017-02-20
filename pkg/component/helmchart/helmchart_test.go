package helmchart

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"os"
	"testing"

	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/proto/hapi/chart"
	"k8s.io/helm/pkg/timeconv"
)

var (
	fake_name      string = "hadoop"
	fake_namespace string = "default"

	chrt     *chart.Chart  = nil
	chrtVals *chart.Config = nil

	fake_chart_dir []string = []string{
		"/work/src/github.com/tangfeixiong/charts/experimental/external-service-auto-discovery/hadoop",
		"/work/src/github.com/tangfeixiong/charts/incubator/etcd",
		"/work/src/github.com/tangfeixiong/charts/stable/mariadb",
	}
	fake_chart_index int = 0
)

func TestMain(m *testing.M) {
	flag.IntVar(&fake_chart_index, "index", 0, "chart index")
	flag.Parse()

	var err error
	chrt, err = chartutil.Load(fake_chart_dir[fake_chart_index])
	if err != nil {
		fmt.Printf("Prepare chart failed: %+v\n", err)
		os.Exit(1)
	}
	chrtVals = chrt.Values

	ret := m.Run()
	os.Exit(ret)
}

func TestChart_abc(t *testing.T) {
	ts := timeconv.Now()
	options := chartutil.ReleaseOptions{Name: fake_name, Time: ts, Namespace: fake_namespace}
	valuesToRender, err := chartutil.ToRenderValues(chrt, chrtVals, options)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("#Result\n---\nRelease: %+v\nChart: %+v\nValues: %+v\n", valuesToRender["Release"], valuesToRender["Chart"], valuesToRender["Values"])
	for k, v := range valuesToRender["Files"].(chartutil.Files) {
		t.Logf("File: %+v, data: %+v\n", k, string(v))
	}

	buf := &bytes.Buffer{}
	for i := range chrt.Templates {
		if tpl := chrt.Templates[i]; tpl != nil {
			t.Logf("Template: %+v, data: %+v\n", tpl.Name, string(tpl.Data))
			engine, err := template.New(tpl.Name).Parse(string(tpl.Data))
			if err != nil {
				t.Fatal(err)
			} else {
				err = engine.Execute(buf, valuesToRender)
				if err != nil {
					t.Fatal(err)
				} else {
					t.Logf("Render: %+v\n", buf.String())
					buf.Reset()
				}
			}
		}
	}
}
