package api_v1

// generated by lightweight-swagger-codegen@eaglesakura

import (
	"encoding/json"
	"github.com/eaglesakura/swagger-go-core"
	"net/http"
)

// ビルドマシンのスペック要求
type MachineRequest struct {

	// 作成されるZone指定 デフォルトでは \"us-central1-b\" が使用される
	Zone *string `json:"zone,omitempty"`

	// リクエストされるvCPU数 デフォルトで24 vCPU使用される
	Cpu *int32 `json:"cpu,omitempty"`

	// リクエストされるRAM(GB) デフォルトで64GB使用される
	Ram *float32 `json:"ram,omitempty"`

	// リクエストされるストレージ容量(GB) デフォルトで48GB使用される
	Storage *float32 `json:"storage,omitempty"`
}

// encode to json
func (it MachineRequest) String() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}

type MachineRequestArray []MachineRequest

func (it *MachineRequest) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Zone, it.Zone == nil).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Cpu, it.Cpu == nil).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Ram, it.Ram == nil).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.Storage, it.Storage == nil).
		Valid(factory) {
		return false
	}

	return true
}

func (it *MachineRequest) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (arr *MachineRequestArray) Valid(factory swagger.ValidatorFactory) bool {
	for _, it := range *arr {
		if !factory.NewValidator(it.Zone, it.Zone == nil).
			Valid(factory) {
			return false
		}
		if !factory.NewValidator(it.Cpu, it.Cpu == nil).
			Valid(factory) {
			return false
		}
		if !factory.NewValidator(it.Ram, it.Ram == nil).
			Valid(factory) {
			return false
		}
		if !factory.NewValidator(it.Storage, it.Storage == nil).
			Valid(factory) {
			return false
		}
	}
	return true
}

func (it *MachineRequestArray) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}