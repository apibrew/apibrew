package nano

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/apibrew/apibrew/modules/nano/pkg/abs"
	"github.com/apibrew/apibrew/modules/nano/pkg/addons"
	"github.com/apibrew/apibrew/modules/nano/pkg/model"
	util2 "github.com/apibrew/apibrew/modules/nano/pkg/util"
	"github.com/apibrew/apibrew/pkg/api"
	"github.com/apibrew/apibrew/pkg/formats/unstructured"
	"github.com/apibrew/apibrew/pkg/service"
	backend_event_handler "github.com/apibrew/apibrew/pkg/service/backend-event-handler"
	"github.com/apibrew/apibrew/pkg/util"
	"github.com/clarkmcc/go-typescript"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"github.com/hashicorp/go-metrics"
	log "github.com/sirupsen/logrus"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type codeExecutorService struct {
	container           service.Container
	backendEventHandler backend_event_handler.BackendEventHandler
	codeContext         util2.Map[string, []*codeExecutionContext]
	globalObject        *globalObject
	modules             util2.Map[string, string]
	systemModules       util2.Map[string, string]
	registry            *require.Registry
	native              util2.Map[string, interface{}]
}

func (s *codeExecutorService) GetContainer() service.Container {
	return s.container
}

func (s *codeExecutorService) GetBackendEventHandler() backend_event_handler.BackendEventHandler {
	return s.backendEventHandler
}

func (s *codeExecutorService) GetGlobalObject() abs.GlobalObject {
	return s.globalObject
}

func (s *codeExecutorService) RunScript(ctx context.Context, script *model.Script) (output interface{}, err error) {
	metrics.IncrCounterWithLabels([]string{"NanoMetrics"}, float32(1), []metrics.Label{
		{Name: "type", Value: "executeScript"},
	})

	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	var source = script.Source

	decodedBytes, err := base64.StdEncoding.DecodeString(script.Source)

	if err == nil {
		source = string(decodedBytes)
	}

	if script.Language == model.ScriptLanguage_TYPESCRIPT {
		transpiled, err := typescript.TranspileCtx(ctx, strings.NewReader(source), s.typescriptOptions)

		if err != nil {
			return nil, err
		}

		source = transpiled

		source = "exports = {};" + source
	}

	log.Debug("Registering script: " + script.Id.String())

	vm := goja.New()
	vm.SetFieldNameMapper(goja.UncapFieldNameMapper())

	s.registry.Enable(vm)

	cec := &codeExecutionContext{}
	cec.id = util.RandomHex(8)
	ctx, cancel := context.WithCancel(util.WithSystemContext(ctx))
	cec.codeCtx = ctx
	cec.cancel = cancel
	cec.identifier = script.Id.String() + "-" + strconv.Itoa(int(script.Version))
	cec.scriptMode = true

	cleanUpContext := cec.WithContext(ctx)
	defer cleanUpContext()

	err = s.registerBuiltIns("["+cec.identifier+"]", vm, cec)

	if err != nil {
		return nil, err
	}

	result, err := vm.RunString(source)

	if err != nil {
		return nil, err
	}

	return result.Export(), nil
}

func (s *codeExecutorService) RunInlineScript(ctx context.Context, identifier string, source string) (result any, err error) {
	metrics.IncrCounterWithLabels([]string{"NanoMetrics"}, float32(1), []metrics.Label{
		{Name: "type", Value: "executeInlineScript"},
	})

	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	vm := goja.New()
	vm.SetFieldNameMapper(goja.UncapFieldNameMapper())

	s.registry.Enable(vm)

	cec := &codeExecutionContext{}
	cec.id = util.RandomHex(8)
	ctx, cancel := context.WithCancel(util.WithSystemContext(ctx))
	cec.codeCtx = ctx
	cec.cancel = cancel
	cec.identifier = identifier
	cec.scriptMode = true

	cleanUpContext := cec.WithContext(ctx)
	defer cleanUpContext()

	err = s.registerBuiltIns("["+cec.identifier+"]", vm, cec)

	if err != nil {
		return nil, err
	}

	var res goja.Value
	res, err = vm.RunString(source)

	if err != nil {
		return nil, err
	}

	return res.Export(), nil
}

func (s *codeExecutorService) registerCode(ctx context.Context, code *model.Code) (err error) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	var concurrencyLevel = 8
	if code.ConcurrencyLevel != nil {
		concurrencyLevel = int(*code.ConcurrencyLevel)
	}

	metrics.IncrCounterWithLabels([]string{"NanoMetrics"}, float32(concurrencyLevel), []metrics.Label{
		{Name: "type", Value: "registerCode"},
		{Name: "name", Value: code.Name},
	})

	decodedBytes, err := base64.StdEncoding.DecodeString(code.Content)

	var source = code.Content

	if err == nil {
		source = string(decodedBytes)
	}

	if code.Language == model.CodeLanguage_TYPESCRIPT {
		transpiled, err := typescript.TranspileCtx(ctx, strings.NewReader(source), s.typescriptOptions)

		if err != nil {
			return err
		}

		source = "exports = {};" + transpiled
	}

	log.Debug("Registering code: " + code.Name)

	program, err := goja.Compile(code.Name, source, false)

	if err != nil {
		return err
	}

	cleanUpList := make([]func(), 0)

	defer func() {
		for _, cleanUp := range cleanUpList {
			cleanUp()
		}
	}()

	cecList := make([]*codeExecutionContext, 0)

	var handlerMap = util2.NewConcurrentSyncMap[string, *abs.HandlerData]()

	for i := 0; i < concurrencyLevel; i++ {
		cec := &codeExecutionContext{}
		cec.id = util.RandomHex(8)
		ctx, cancel := context.WithCancel(util.WithSystemContext(context.Background()))
		cec.codeCtx = ctx
		cec.cancel = cancel

		cleanUpContext := cec.WithContext(ctx)
		cleanUpList = append(cleanUpList, cleanUpContext)

		cec.handlerMap = handlerMap
		cec.identifier = code.Id.String() + "-" + strconv.Itoa(int(code.Version))

		vm := goja.New()
		s.registry.Enable(vm)
		vm.SetFieldNameMapper(goja.UncapFieldNameMapper())
		err = s.registerBuiltIns(code.Name, vm, cec)

		if err != nil {
			return err
		}

		_, err = vm.RunProgram(program)
		if err != nil {
			return err
		}

		cec.vm = vm

		cecList = append(cecList, cec)
	}

	s.codeContext.Set(code.Name, cecList)

	return nil
}

func (s *codeExecutorService) updateCode(ctx context.Context, code *model.Code) error {
	if err := s.unRegisterCode(ctx, code); err != nil {
		return err
	}

	if err := s.registerCode(ctx, code); err != nil {
		return err
	}

	return nil
}

func (s *codeExecutorService) unRegisterCode(ctx context.Context, code *model.Code) error {
	cecs, ok := s.codeContext.Find(code.Name)

	if !ok {
		return errors.New("code context not found")
	}

	log.Info("Unregistering code: " + code.Name)

	for _, cec := range cecs {
		cec.cancel()
	}

	return nil
}

func (s *codeExecutorService) registerBuiltIns(codeName string, vm *goja.Runtime, cec *codeExecutionContext) error {
	if err := addons.Register(vm, cec, s, codeName, s.container); err != nil {
		return err
	}

	if err := vm.Set("global", s.globalObject); err != nil {
		return err
	}

	if err := s.registerTimeoutFunctions(vm, cec); err != nil {
		return err
	}

	if _, err := vm.RunScript("resource.js", GetBuiltinJs("resource.js")); err != nil {
		return err
	}

	if _, err := vm.RunScript("targetResource.js", GetBuiltinJs("targetResource.js")); err != nil {
		return err
	}

	if _, err := vm.RunScript("transactional.js", GetBuiltinJs("transactional.js")); err != nil {
		return err
	}

	if !cec.scriptMode {
		if _, err := vm.RunScript("handle.js", GetBuiltinJs("handle.js")); err != nil {
			return err
		}
	}

	for _, key := range s.native.Keys() {
		if err := vm.Set(key, s.native.Get(key)); err != nil {
			return err
		}
	}

	return nil
}

func (s *codeExecutorService) registerTimeoutFunctions(vm *goja.Runtime, cec *codeExecutionContext) error {
	if err := vm.Set("setTimeout", s.setTimeoutFn(cec)); err != nil {
		return err
	}

	if err := vm.Set("clearTimeout", s.clearTimeoutFn(cec)); err != nil {
		return err
	}

	if err := vm.Set("setInterval", s.setIntervalFn(cec)); err != nil {
		return err
	}

	if err := vm.Set("clearInterval", s.clearIntervalFn(cec)); err != nil {
		return err
	}

	if err := vm.Set("sleep", s.sleepFn(cec)); err != nil {
		return err
	}

	return nil
}

func (s *codeExecutorService) setTimeoutFn(cec *codeExecutionContext) func(fn func(), duration int64) func() {
	return func(fn func(), duration int64) func() {
		cancel := make(chan struct{})
		cancelFn := func() {
			close(cancel)
		}

		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Warn(r)
				}
			}()
			select {
			case <-time.After(time.Duration(duration) * time.Millisecond):
				fn()
			case <-cancel:
			case <-cec.codeCtx.Done():
				// Cancel the timeout
			}
		}()
		return cancelFn
	}
}

func (s *codeExecutorService) clearTimeoutFn(cec *codeExecutionContext) func(clearFn func()) {
	return func(clearFn func()) {
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				log.Warn(r)
			}
		}()
		clearFn()
	}
}

func (s *codeExecutorService) setIntervalFn(cec *codeExecutionContext) func(fn func(), duration int64) func() {
	return func(fn func(), duration int64) func() {
		cancel := make(chan struct{})
		cancelFn := func() {
			close(cancel)
		}

		go func() {
		Loop:
			for {
				defer func() {
					if r := recover(); r != nil {
						log.Warn(r)
					}
				}()
				select {
				case <-time.After(time.Duration(duration) * time.Millisecond):
					fn()
				case <-cec.codeCtx.Done():
					break Loop
				case <-cancel:
					break Loop
				}
			}
		}()
		return cancelFn
	}
}

func (s *codeExecutorService) clearIntervalFn(cec *codeExecutionContext) func(clearFn func()) {
	return func(clearFn func()) {
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
				log.Warn(r)
			}
		}()
		clearFn()
	}
}

func (s *codeExecutorService) sleepFn(cec *codeExecutionContext) func(duration int32) {
	return func(duration int32) {
		select {
		case <-time.After(time.Duration(duration) * time.Millisecond):
		case <-cec.codeCtx.Done():
		}
	}
}

func (s *codeExecutorService) typescriptOptions(config *typescript.Config) {
	config.Verbose = true
	config.CompileOptions = map[string]interface{}{
		"module": "commonJS",
		"target": "es5",
	}
}

func (s *codeExecutorService) registerModule(ctx context.Context, module *model.Module) error {
	metrics.IncrCounterWithLabels([]string{"NanoMetrics"}, float32(1), []metrics.Label{
		{Name: "type", Value: "registerModule"},
		{Name: "name", Value: module.Name},
	})

	var source = module.Source

	decodedBytes, err := base64.StdEncoding.DecodeString(module.Source)

	if err == nil {
		source = string(decodedBytes)
	}

	if module.Language == model.ModuleLanguage_TYPESCRIPT {
		transpiled, err := typescript.TranspileCtx(ctx, strings.NewReader(source), s.typescriptOptions)

		if err != nil {
			return err
		}

		source = transpiled
	}

	s.modules.Set(module.Name, source)

	log.Println("Registering module: " + module.Name)

	return nil
}

func (s *codeExecutorService) updateModule(ctx context.Context, module *model.Module) error {
	if err := s.unRegisterModule(ctx, module); err != nil {
		return err
	}

	if err := s.registerModule(ctx, module); err != nil {
		return err
	}

	return nil
}

func (s *codeExecutorService) unRegisterModule(ctx context.Context, module *model.Module) error {
	s.modules.Delete(module.Name)

	return nil
}

func (s *codeExecutorService) srcLoader(path string) ([]byte, error) {
	path = strings.ReplaceAll(path, "node_modules/", "")
	path = strings.TrimPrefix(path, "./")
	path = strings.TrimPrefix(path, "/")

	if source, ok := s.systemModules.Find(path); ok {
		return []byte(source), nil
	}

	if source, ok := s.modules.Find(path); ok {
		log.Println("Loading module: "+path, " len: ", len(source))
		return []byte(source), nil
	}

	return nil, errors.New("module not found with name: " + path)
}

func (s *codeExecutorService) init() {
	s.systemModules.Set("@apibrew/nano", GetBuiltinJs("nano.js"))
}

func (s *codeExecutorService) restartCodeContext(ctx context.Context) {
	s.registry = require.NewRegistryWithLoader(s.srcLoader)

	apiInterface := api.NewInterface(s.container)

	list, err := apiInterface.List(ctx, api.ListParams{
		Type: "nano/Code",
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, codeObj := range list.Content {
		codeRecord, err := unstructured.ToRecord(codeObj)

		if err != nil {
			log.Fatal(err)
		}

		code := model.CodeMapperInstance.FromRecord(codeRecord)

		if err := s.updateCode(ctx, code); err != nil {
			log.Fatal(err)
		}
	}
}

func (s *codeExecutorService) RegisterNative(name string, val interface{}) {
	s.native.Set(name, val)
}

func newCodeExecutorService(container service.Container, backendEventHandler backend_event_handler.BackendEventHandler) *codeExecutorService {
	ces := &codeExecutorService{
		container:           container,
		backendEventHandler: backendEventHandler,
		codeContext:         util2.NewConcurrentSyncMap[string, []*codeExecutionContext](),
		systemModules:       util2.NewConcurrentSyncMap[string, string](),
		modules:             util2.NewConcurrentSyncMap[string, string](),
		native:              util2.NewConcurrentSyncMap[string, interface{}](),
		globalObject:        newGlobalObject(),
	}

	ces.init()

	registry := require.NewRegistryWithLoader(ces.srcLoader)

	ces.registry = registry

	return ces
}
