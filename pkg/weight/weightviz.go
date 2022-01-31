package weight

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fusion/golyz/pkg/wrap"
)

type ModuleInfo struct {
	Path string
	Name string
	Size uint64
}

var packageMatcher = regexp.MustCompile("packagefile (.*)=(.*)")

func BuildWeightData(l *log.Logger) wrap.PieData {
	modulesList := make(wrap.PieData)

	out, err := exec.Command(
		"go", "build", "-o", "goweight-bin-target", "-work", "-a").CombinedOutput()
	if err != nil {
		l.Fatal(err)
	}
	os.Remove("goweight-bin-target")
	mainPath := strings.Split(strings.TrimSuffix(string(out), "\n"), "=")[1]
	bitPaths, err := filepath.Glob(fmt.Sprintf("%s/*/importcfg", mainPath))
	if err != nil {
		l.Fatal(err)
	}
	for bIdx := range bitPaths {
		filePath := bitPaths[bIdx]
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			l.Fatal(err)
		}
		packageList := strings.Split(string(fileContent), "\n")
		for pIdx := range packageList {
			bitsModule := packageMatcher.FindStringSubmatch(packageList[pIdx])
			if bitsModule == nil {
				continue
			}
			moduleInfo := ModuleInfo{
				Path: bitsModule[2],
				Name: bitsModule[1],
			}
			modulePath := bitsModule[2]
			stat, err := os.Stat(modulePath)
			if err != nil {
				l.Fatal(err)
			}
			moduleInfo.Size = uint64(stat.Size())
			_, ok := modulesList[moduleInfo.Name]
			if !ok {
				modulesList[moduleInfo.Name] = moduleInfo.Size
			}
		}
	}

	return modulesList
}
