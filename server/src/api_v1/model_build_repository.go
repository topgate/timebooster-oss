package api_v1

// generated by lightweight-swagger-codegen@eaglesakura

import (
	"encoding/json"
	"github.com/eaglesakura/swagger-go-core"
	"net/http"
)

// ビルド対象のソースリポジトリを示す
type BuildRepository struct {

	// gitリポジトリのpathを示す。 git cloneとして有効な値が設定される。
	Git *string `json:"git,omitempty"`

	// gitリポジトリのcheckout対象を示す。 git checkout として有効な値が設定される。 checkoutできれば良いので、branchやtagも可能。
	GitRevision *string `json:"gitRevision,omitempty"`
}

// encode to json
func (it BuildRepository) String() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}

type BuildRepositoryArray []BuildRepository

func (it *BuildRepository) Valid(factory swagger.ValidatorFactory) bool {
	if !factory.NewValidator(it.Git, it.Git == nil).
		Valid(factory) {
		return false
	}
	if !factory.NewValidator(it.GitRevision, it.GitRevision == nil).
		Valid(factory) {
		return false
	}

	return true
}

func (it *BuildRepository) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

func (arr *BuildRepositoryArray) Valid(factory swagger.ValidatorFactory) bool {
	for _, it := range *arr {
		if !factory.NewValidator(it.Git, it.Git == nil).
			Valid(factory) {
			return false
		}
		if !factory.NewValidator(it.GitRevision, it.GitRevision == nil).
			Valid(factory) {
			return false
		}
	}
	return true
}

func (it *BuildRepositoryArray) WriteResponse(writer http.ResponseWriter, producer swagger.Producer) {
	writer.WriteHeader(200)
	if err := producer.Produce(writer, it); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
