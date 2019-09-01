package params

import (
	"context"
	"encoding/json"
	pb "gogui/injector"
	"log"
	"strings"
)

type Modules struct {
	mods map[string]Accessor
}

func NewModules() Modules {
	return Modules{mods: make(map[string]Accessor)}
}

func (m *Modules) AddModule(name string, acc Accessor) {
	m.mods[name] = acc
}

func (m *Modules) Get(paramName string) (string, error) {
	mod, param := m.extractAccessor(paramName)
	return mod.Get(param)
}

func (m *Modules) MultiGet(ctx context.Context, modToPaths map[string][]string) <-chan []*pb.PathAndValue {
	ctx, cancel := context.WithCancel(context.Background())
	newValsStream := make(chan []*pb.PathAndValue)
	for mod, paths := range modToPaths {
		c, _ := m.mods[mod].MultiGet(ctx, paths)
		go func() {
			for {
				select {
				case <-ctx.Done():
					cancel()
					return
				case vals := <-c:
					newValsStream <- vals
				}
			}
		}()
	}
	return newValsStream
}

func (m *Modules) Print(p string) {
	req := make(map[string]map[string]string)
	json.Unmarshal([]byte(p), &req)
	// for k, v := range p {
	// 	req[k] = v.(map[string]interface{})
	// }
	log.Println(req)
}

func (m *Modules) MultiSet(params map[string]map[string]string) error {
	for mod, paramValMap := range params {
		err := m.mods[mod].MultiSet(paramValMap)
		if err != nil {
			log.Printf("Couldnt set values for %s", mod)
		}
	}
	return nil
}

func (m *Modules) GetAllParamInfo() map[string][]ParamInfo {
	params := make(map[string][]ParamInfo)
	for mod, acc := range m.mods {
		pInfos, err := acc.GetAllParamInfo()
		if err != nil {
			log.Printf("Couldnt get parameter infos of %s", mod)
			break
		}
		params[mod] = pInfos
	}
	return params
}

func (m *Modules) GetAllPossibleEnums() (map[string]map[string][]string, error) {
	enums := make(map[string]map[string][]string)
	for mod, acc := range m.mods {
		e, err := acc.GetAllPossibleEnums()
		if err != nil {
			log.Printf("Couldnt fetch enums values of %s", mod)
			break
		}
		enums[mod] = e
	}
	return enums, nil
}

func (m *Modules) extractAccessor(path string) (Accessor, string) {
	idx := strings.Index(path, ".")
	name, param := path[:idx], path[idx+1:]
	return m.mods[name], param
}
