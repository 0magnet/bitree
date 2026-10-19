package bitree

import "testing"

// TestAboveRootAndSpineLabels covers Node.Above (top level and nested), a root
// with a left label and columns, and the labeled spine line under the root.
func TestAboveRootAndSpineLabels(t *testing.T) {
	root := &Node{Label: "HV", Left: ann("hypervisor"), Cols: []string{"1.2.3.4", "v1"}, Right: []*Node{
		{Label: "PK1", Above: "5.6.7.8  v2", Left: ann("prod1"), Cols: []string{"dmsg server"}, Right: []*Node{
			{Label: "services", Cols: []string{"name", "listens"}, Right: []*Node{leaf("SVCKEY", "tpd", ":9094")}},
			{Label: "sub", Above: "above nested", Right: []*Node{leaf("x")}},
		}},
		{Label: "PK2", Above: "offline", Left: ann("dmsg-1")},
	}}
	want := "hypervisor ──HV                  1.2.3.4      v1" + "\n" +
		"    linode   │   visor key       roles" + "\n" +
		"             │   5.6.7.8  v2" + "\n" +
		"     prod1 ──┼── PK1             dmsg server" + "\n" +
		"             │   ├── services    name         listens" + "\n" +
		"             │   │   └── SVCKEY  tpd          :9094" + "\n" +
		"             │   │   above nested" + "\n" +
		"             │   └── sub" + "\n" +
		"             │       └── x" + "\n" +
		"             │   offline" + "\n" +
		"    dmsg-1 ──┴── PK2"
	check(t, "above", Render(root, Options{SpineLeft: "linode", SpineLabel: "visor key", SpineCols: []string{"roles"}}), want)
}
