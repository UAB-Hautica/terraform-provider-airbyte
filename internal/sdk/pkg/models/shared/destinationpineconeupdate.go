// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type DestinationPineconeUpdateSchemasEmbeddingEmbedding5Mode string

const (
	DestinationPineconeUpdateSchemasEmbeddingEmbedding5ModeOpenaiCompatible DestinationPineconeUpdateSchemasEmbeddingEmbedding5Mode = "openai_compatible"
)

func (e DestinationPineconeUpdateSchemasEmbeddingEmbedding5Mode) ToPointer() *DestinationPineconeUpdateSchemasEmbeddingEmbedding5Mode {
	return &e
}

func (e *DestinationPineconeUpdateSchemasEmbeddingEmbedding5Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai_compatible":
		*e = DestinationPineconeUpdateSchemasEmbeddingEmbedding5Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateSchemasEmbeddingEmbedding5Mode: %v", v)
	}
}

// DestinationPineconeUpdateOpenAICompatible - Use a service that's compatible with the OpenAI API to embed text.
type DestinationPineconeUpdateOpenAICompatible struct {
	APIKey *string `default:"" json:"api_key"`
	// The base URL for your OpenAI-compatible service
	BaseURL string `json:"base_url"`
	// The number of dimensions the embedding model is generating
	Dimensions int64                                                    `json:"dimensions"`
	mode       *DestinationPineconeUpdateSchemasEmbeddingEmbedding5Mode `const:"openai_compatible" json:"mode"`
	// The name of the model to use for embedding
	ModelName *string `default:"text-embedding-ada-002" json:"model_name"`
}

func (d DestinationPineconeUpdateOpenAICompatible) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeUpdateOpenAICompatible) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeUpdateOpenAICompatible) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *DestinationPineconeUpdateOpenAICompatible) GetBaseURL() string {
	if o == nil {
		return ""
	}
	return o.BaseURL
}

func (o *DestinationPineconeUpdateOpenAICompatible) GetDimensions() int64 {
	if o == nil {
		return 0
	}
	return o.Dimensions
}

func (o *DestinationPineconeUpdateOpenAICompatible) GetMode() *DestinationPineconeUpdateSchemasEmbeddingEmbedding5Mode {
	return DestinationPineconeUpdateSchemasEmbeddingEmbedding5ModeOpenaiCompatible.ToPointer()
}

func (o *DestinationPineconeUpdateOpenAICompatible) GetModelName() *string {
	if o == nil {
		return nil
	}
	return o.ModelName
}

type DestinationPineconeUpdateSchemasEmbeddingEmbeddingMode string

const (
	DestinationPineconeUpdateSchemasEmbeddingEmbeddingModeAzureOpenai DestinationPineconeUpdateSchemasEmbeddingEmbeddingMode = "azure_openai"
)

func (e DestinationPineconeUpdateSchemasEmbeddingEmbeddingMode) ToPointer() *DestinationPineconeUpdateSchemasEmbeddingEmbeddingMode {
	return &e
}

func (e *DestinationPineconeUpdateSchemasEmbeddingEmbeddingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "azure_openai":
		*e = DestinationPineconeUpdateSchemasEmbeddingEmbeddingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateSchemasEmbeddingEmbeddingMode: %v", v)
	}
}

// DestinationPineconeUpdateAzureOpenAI - Use the Azure-hosted OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type DestinationPineconeUpdateAzureOpenAI struct {
	// The base URL for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	APIBase string `json:"api_base"`
	// The deployment for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	Deployment string                                                  `json:"deployment"`
	mode       *DestinationPineconeUpdateSchemasEmbeddingEmbeddingMode `const:"azure_openai" json:"mode"`
	// The API key for your Azure OpenAI resource.  You can find this in the Azure portal under your Azure OpenAI resource
	OpenaiKey string `json:"openai_key"`
}

func (d DestinationPineconeUpdateAzureOpenAI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeUpdateAzureOpenAI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeUpdateAzureOpenAI) GetAPIBase() string {
	if o == nil {
		return ""
	}
	return o.APIBase
}

func (o *DestinationPineconeUpdateAzureOpenAI) GetDeployment() string {
	if o == nil {
		return ""
	}
	return o.Deployment
}

func (o *DestinationPineconeUpdateAzureOpenAI) GetMode() *DestinationPineconeUpdateSchemasEmbeddingEmbeddingMode {
	return DestinationPineconeUpdateSchemasEmbeddingEmbeddingModeAzureOpenai.ToPointer()
}

func (o *DestinationPineconeUpdateAzureOpenAI) GetOpenaiKey() string {
	if o == nil {
		return ""
	}
	return o.OpenaiKey
}

type DestinationPineconeUpdateSchemasEmbeddingMode string

const (
	DestinationPineconeUpdateSchemasEmbeddingModeFake DestinationPineconeUpdateSchemasEmbeddingMode = "fake"
)

func (e DestinationPineconeUpdateSchemasEmbeddingMode) ToPointer() *DestinationPineconeUpdateSchemasEmbeddingMode {
	return &e
}

func (e *DestinationPineconeUpdateSchemasEmbeddingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fake":
		*e = DestinationPineconeUpdateSchemasEmbeddingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateSchemasEmbeddingMode: %v", v)
	}
}

// DestinationPineconeUpdateFake - Use a fake embedding made out of random vectors with 1536 embedding dimensions. This is useful for testing the data pipeline without incurring any costs.
type DestinationPineconeUpdateFake struct {
	mode *DestinationPineconeUpdateSchemasEmbeddingMode `const:"fake" json:"mode"`
}

func (d DestinationPineconeUpdateFake) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeUpdateFake) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeUpdateFake) GetMode() *DestinationPineconeUpdateSchemasEmbeddingMode {
	return DestinationPineconeUpdateSchemasEmbeddingModeFake.ToPointer()
}

type DestinationPineconeUpdateSchemasMode string

const (
	DestinationPineconeUpdateSchemasModeCohere DestinationPineconeUpdateSchemasMode = "cohere"
)

func (e DestinationPineconeUpdateSchemasMode) ToPointer() *DestinationPineconeUpdateSchemasMode {
	return &e
}

func (e *DestinationPineconeUpdateSchemasMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cohere":
		*e = DestinationPineconeUpdateSchemasMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateSchemasMode: %v", v)
	}
}

// DestinationPineconeUpdateCohere - Use the Cohere API to embed text.
type DestinationPineconeUpdateCohere struct {
	CohereKey string                                `json:"cohere_key"`
	mode      *DestinationPineconeUpdateSchemasMode `const:"cohere" json:"mode"`
}

func (d DestinationPineconeUpdateCohere) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeUpdateCohere) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeUpdateCohere) GetCohereKey() string {
	if o == nil {
		return ""
	}
	return o.CohereKey
}

func (o *DestinationPineconeUpdateCohere) GetMode() *DestinationPineconeUpdateSchemasMode {
	return DestinationPineconeUpdateSchemasModeCohere.ToPointer()
}

type DestinationPineconeUpdateMode string

const (
	DestinationPineconeUpdateModeOpenai DestinationPineconeUpdateMode = "openai"
)

func (e DestinationPineconeUpdateMode) ToPointer() *DestinationPineconeUpdateMode {
	return &e
}

func (e *DestinationPineconeUpdateMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "openai":
		*e = DestinationPineconeUpdateMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateMode: %v", v)
	}
}

// DestinationPineconeUpdateOpenAI - Use the OpenAI API to embed text. This option is using the text-embedding-ada-002 model with 1536 embedding dimensions.
type DestinationPineconeUpdateOpenAI struct {
	mode      *DestinationPineconeUpdateMode `const:"openai" json:"mode"`
	OpenaiKey string                         `json:"openai_key"`
}

func (d DestinationPineconeUpdateOpenAI) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeUpdateOpenAI) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeUpdateOpenAI) GetMode() *DestinationPineconeUpdateMode {
	return DestinationPineconeUpdateModeOpenai.ToPointer()
}

func (o *DestinationPineconeUpdateOpenAI) GetOpenaiKey() string {
	if o == nil {
		return ""
	}
	return o.OpenaiKey
}

type DestinationPineconeUpdateEmbeddingType string

const (
	DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateOpenAI           DestinationPineconeUpdateEmbeddingType = "destination-pinecone-update_OpenAI"
	DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateCohere           DestinationPineconeUpdateEmbeddingType = "destination-pinecone-update_Cohere"
	DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateFake             DestinationPineconeUpdateEmbeddingType = "destination-pinecone-update_Fake"
	DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateAzureOpenAI      DestinationPineconeUpdateEmbeddingType = "destination-pinecone-update_Azure OpenAI"
	DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateOpenAICompatible DestinationPineconeUpdateEmbeddingType = "destination-pinecone-update_OpenAI-compatible"
)

type DestinationPineconeUpdateEmbedding struct {
	DestinationPineconeUpdateOpenAI           *DestinationPineconeUpdateOpenAI
	DestinationPineconeUpdateCohere           *DestinationPineconeUpdateCohere
	DestinationPineconeUpdateFake             *DestinationPineconeUpdateFake
	DestinationPineconeUpdateAzureOpenAI      *DestinationPineconeUpdateAzureOpenAI
	DestinationPineconeUpdateOpenAICompatible *DestinationPineconeUpdateOpenAICompatible

	Type DestinationPineconeUpdateEmbeddingType
}

func CreateDestinationPineconeUpdateEmbeddingDestinationPineconeUpdateOpenAI(destinationPineconeUpdateOpenAI DestinationPineconeUpdateOpenAI) DestinationPineconeUpdateEmbedding {
	typ := DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateOpenAI

	return DestinationPineconeUpdateEmbedding{
		DestinationPineconeUpdateOpenAI: &destinationPineconeUpdateOpenAI,
		Type:                            typ,
	}
}

func CreateDestinationPineconeUpdateEmbeddingDestinationPineconeUpdateCohere(destinationPineconeUpdateCohere DestinationPineconeUpdateCohere) DestinationPineconeUpdateEmbedding {
	typ := DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateCohere

	return DestinationPineconeUpdateEmbedding{
		DestinationPineconeUpdateCohere: &destinationPineconeUpdateCohere,
		Type:                            typ,
	}
}

func CreateDestinationPineconeUpdateEmbeddingDestinationPineconeUpdateFake(destinationPineconeUpdateFake DestinationPineconeUpdateFake) DestinationPineconeUpdateEmbedding {
	typ := DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateFake

	return DestinationPineconeUpdateEmbedding{
		DestinationPineconeUpdateFake: &destinationPineconeUpdateFake,
		Type:                          typ,
	}
}

func CreateDestinationPineconeUpdateEmbeddingDestinationPineconeUpdateAzureOpenAI(destinationPineconeUpdateAzureOpenAI DestinationPineconeUpdateAzureOpenAI) DestinationPineconeUpdateEmbedding {
	typ := DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateAzureOpenAI

	return DestinationPineconeUpdateEmbedding{
		DestinationPineconeUpdateAzureOpenAI: &destinationPineconeUpdateAzureOpenAI,
		Type:                                 typ,
	}
}

func CreateDestinationPineconeUpdateEmbeddingDestinationPineconeUpdateOpenAICompatible(destinationPineconeUpdateOpenAICompatible DestinationPineconeUpdateOpenAICompatible) DestinationPineconeUpdateEmbedding {
	typ := DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateOpenAICompatible

	return DestinationPineconeUpdateEmbedding{
		DestinationPineconeUpdateOpenAICompatible: &destinationPineconeUpdateOpenAICompatible,
		Type: typ,
	}
}

func (u *DestinationPineconeUpdateEmbedding) UnmarshalJSON(data []byte) error {

	destinationPineconeUpdateFake := new(DestinationPineconeUpdateFake)
	if err := utils.UnmarshalJSON(data, &destinationPineconeUpdateFake, "", true, true); err == nil {
		u.DestinationPineconeUpdateFake = destinationPineconeUpdateFake
		u.Type = DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateFake
		return nil
	}

	destinationPineconeUpdateOpenAI := new(DestinationPineconeUpdateOpenAI)
	if err := utils.UnmarshalJSON(data, &destinationPineconeUpdateOpenAI, "", true, true); err == nil {
		u.DestinationPineconeUpdateOpenAI = destinationPineconeUpdateOpenAI
		u.Type = DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateOpenAI
		return nil
	}

	destinationPineconeUpdateCohere := new(DestinationPineconeUpdateCohere)
	if err := utils.UnmarshalJSON(data, &destinationPineconeUpdateCohere, "", true, true); err == nil {
		u.DestinationPineconeUpdateCohere = destinationPineconeUpdateCohere
		u.Type = DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateCohere
		return nil
	}

	destinationPineconeUpdateAzureOpenAI := new(DestinationPineconeUpdateAzureOpenAI)
	if err := utils.UnmarshalJSON(data, &destinationPineconeUpdateAzureOpenAI, "", true, true); err == nil {
		u.DestinationPineconeUpdateAzureOpenAI = destinationPineconeUpdateAzureOpenAI
		u.Type = DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateAzureOpenAI
		return nil
	}

	destinationPineconeUpdateOpenAICompatible := new(DestinationPineconeUpdateOpenAICompatible)
	if err := utils.UnmarshalJSON(data, &destinationPineconeUpdateOpenAICompatible, "", true, true); err == nil {
		u.DestinationPineconeUpdateOpenAICompatible = destinationPineconeUpdateOpenAICompatible
		u.Type = DestinationPineconeUpdateEmbeddingTypeDestinationPineconeUpdateOpenAICompatible
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPineconeUpdateEmbedding) MarshalJSON() ([]byte, error) {
	if u.DestinationPineconeUpdateOpenAI != nil {
		return utils.MarshalJSON(u.DestinationPineconeUpdateOpenAI, "", true)
	}

	if u.DestinationPineconeUpdateCohere != nil {
		return utils.MarshalJSON(u.DestinationPineconeUpdateCohere, "", true)
	}

	if u.DestinationPineconeUpdateFake != nil {
		return utils.MarshalJSON(u.DestinationPineconeUpdateFake, "", true)
	}

	if u.DestinationPineconeUpdateAzureOpenAI != nil {
		return utils.MarshalJSON(u.DestinationPineconeUpdateAzureOpenAI, "", true)
	}

	if u.DestinationPineconeUpdateOpenAICompatible != nil {
		return utils.MarshalJSON(u.DestinationPineconeUpdateOpenAICompatible, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// DestinationPineconeUpdateIndexing - Pinecone is a popular vector store that can be used to store and retrieve embeddings.
type DestinationPineconeUpdateIndexing struct {
	// Pinecone index in your project to load data into
	Index string `json:"index"`
	// Pinecone Cloud environment to use
	PineconeEnvironment string `json:"pinecone_environment"`
	// The Pinecone API key to use matching the environment (copy from Pinecone console)
	PineconeKey string `json:"pinecone_key"`
}

func (o *DestinationPineconeUpdateIndexing) GetIndex() string {
	if o == nil {
		return ""
	}
	return o.Index
}

func (o *DestinationPineconeUpdateIndexing) GetPineconeEnvironment() string {
	if o == nil {
		return ""
	}
	return o.PineconeEnvironment
}

func (o *DestinationPineconeUpdateIndexing) GetPineconeKey() string {
	if o == nil {
		return ""
	}
	return o.PineconeKey
}

type DestinationPineconeUpdateFieldNameMappingConfigModel struct {
	// The field name in the source
	FromField string `json:"from_field"`
	// The field name to use in the destination
	ToField string `json:"to_field"`
}

func (o *DestinationPineconeUpdateFieldNameMappingConfigModel) GetFromField() string {
	if o == nil {
		return ""
	}
	return o.FromField
}

func (o *DestinationPineconeUpdateFieldNameMappingConfigModel) GetToField() string {
	if o == nil {
		return ""
	}
	return o.ToField
}

// DestinationPineconeUpdateLanguage - Split code in suitable places based on the programming language
type DestinationPineconeUpdateLanguage string

const (
	DestinationPineconeUpdateLanguageCpp      DestinationPineconeUpdateLanguage = "cpp"
	DestinationPineconeUpdateLanguageGo       DestinationPineconeUpdateLanguage = "go"
	DestinationPineconeUpdateLanguageJava     DestinationPineconeUpdateLanguage = "java"
	DestinationPineconeUpdateLanguageJs       DestinationPineconeUpdateLanguage = "js"
	DestinationPineconeUpdateLanguagePhp      DestinationPineconeUpdateLanguage = "php"
	DestinationPineconeUpdateLanguageProto    DestinationPineconeUpdateLanguage = "proto"
	DestinationPineconeUpdateLanguagePython   DestinationPineconeUpdateLanguage = "python"
	DestinationPineconeUpdateLanguageRst      DestinationPineconeUpdateLanguage = "rst"
	DestinationPineconeUpdateLanguageRuby     DestinationPineconeUpdateLanguage = "ruby"
	DestinationPineconeUpdateLanguageRust     DestinationPineconeUpdateLanguage = "rust"
	DestinationPineconeUpdateLanguageScala    DestinationPineconeUpdateLanguage = "scala"
	DestinationPineconeUpdateLanguageSwift    DestinationPineconeUpdateLanguage = "swift"
	DestinationPineconeUpdateLanguageMarkdown DestinationPineconeUpdateLanguage = "markdown"
	DestinationPineconeUpdateLanguageLatex    DestinationPineconeUpdateLanguage = "latex"
	DestinationPineconeUpdateLanguageHTML     DestinationPineconeUpdateLanguage = "html"
	DestinationPineconeUpdateLanguageSol      DestinationPineconeUpdateLanguage = "sol"
)

func (e DestinationPineconeUpdateLanguage) ToPointer() *DestinationPineconeUpdateLanguage {
	return &e
}

func (e *DestinationPineconeUpdateLanguage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "cpp":
		fallthrough
	case "go":
		fallthrough
	case "java":
		fallthrough
	case "js":
		fallthrough
	case "php":
		fallthrough
	case "proto":
		fallthrough
	case "python":
		fallthrough
	case "rst":
		fallthrough
	case "ruby":
		fallthrough
	case "rust":
		fallthrough
	case "scala":
		fallthrough
	case "swift":
		fallthrough
	case "markdown":
		fallthrough
	case "latex":
		fallthrough
	case "html":
		fallthrough
	case "sol":
		*e = DestinationPineconeUpdateLanguage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateLanguage: %v", v)
	}
}

type DestinationPineconeUpdateSchemasProcessingTextSplitterTextSplitterMode string

const (
	DestinationPineconeUpdateSchemasProcessingTextSplitterTextSplitterModeCode DestinationPineconeUpdateSchemasProcessingTextSplitterTextSplitterMode = "code"
)

func (e DestinationPineconeUpdateSchemasProcessingTextSplitterTextSplitterMode) ToPointer() *DestinationPineconeUpdateSchemasProcessingTextSplitterTextSplitterMode {
	return &e
}

func (e *DestinationPineconeUpdateSchemasProcessingTextSplitterTextSplitterMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "code":
		*e = DestinationPineconeUpdateSchemasProcessingTextSplitterTextSplitterMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateSchemasProcessingTextSplitterTextSplitterMode: %v", v)
	}
}

// DestinationPineconeUpdateByProgrammingLanguage - Split the text by suitable delimiters based on the programming language. This is useful for splitting code into chunks.
type DestinationPineconeUpdateByProgrammingLanguage struct {
	// Split code in suitable places based on the programming language
	Language DestinationPineconeUpdateLanguage                                       `json:"language"`
	mode     *DestinationPineconeUpdateSchemasProcessingTextSplitterTextSplitterMode `const:"code" json:"mode"`
}

func (d DestinationPineconeUpdateByProgrammingLanguage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeUpdateByProgrammingLanguage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeUpdateByProgrammingLanguage) GetLanguage() DestinationPineconeUpdateLanguage {
	if o == nil {
		return DestinationPineconeUpdateLanguage("")
	}
	return o.Language
}

func (o *DestinationPineconeUpdateByProgrammingLanguage) GetMode() *DestinationPineconeUpdateSchemasProcessingTextSplitterTextSplitterMode {
	return DestinationPineconeUpdateSchemasProcessingTextSplitterTextSplitterModeCode.ToPointer()
}

type DestinationPineconeUpdateSchemasProcessingTextSplitterMode string

const (
	DestinationPineconeUpdateSchemasProcessingTextSplitterModeMarkdown DestinationPineconeUpdateSchemasProcessingTextSplitterMode = "markdown"
)

func (e DestinationPineconeUpdateSchemasProcessingTextSplitterMode) ToPointer() *DestinationPineconeUpdateSchemasProcessingTextSplitterMode {
	return &e
}

func (e *DestinationPineconeUpdateSchemasProcessingTextSplitterMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "markdown":
		*e = DestinationPineconeUpdateSchemasProcessingTextSplitterMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateSchemasProcessingTextSplitterMode: %v", v)
	}
}

// DestinationPineconeUpdateByMarkdownHeader - Split the text by Markdown headers down to the specified header level. If the chunk size fits multiple sections, they will be combined into a single chunk.
type DestinationPineconeUpdateByMarkdownHeader struct {
	mode *DestinationPineconeUpdateSchemasProcessingTextSplitterMode `const:"markdown" json:"mode"`
	// Level of markdown headers to split text fields by. Headings down to the specified level will be used as split points
	SplitLevel *int64 `default:"1" json:"split_level"`
}

func (d DestinationPineconeUpdateByMarkdownHeader) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeUpdateByMarkdownHeader) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeUpdateByMarkdownHeader) GetMode() *DestinationPineconeUpdateSchemasProcessingTextSplitterMode {
	return DestinationPineconeUpdateSchemasProcessingTextSplitterModeMarkdown.ToPointer()
}

func (o *DestinationPineconeUpdateByMarkdownHeader) GetSplitLevel() *int64 {
	if o == nil {
		return nil
	}
	return o.SplitLevel
}

type DestinationPineconeUpdateSchemasProcessingMode string

const (
	DestinationPineconeUpdateSchemasProcessingModeSeparator DestinationPineconeUpdateSchemasProcessingMode = "separator"
)

func (e DestinationPineconeUpdateSchemasProcessingMode) ToPointer() *DestinationPineconeUpdateSchemasProcessingMode {
	return &e
}

func (e *DestinationPineconeUpdateSchemasProcessingMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "separator":
		*e = DestinationPineconeUpdateSchemasProcessingMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationPineconeUpdateSchemasProcessingMode: %v", v)
	}
}

// DestinationPineconeUpdateBySeparator - Split the text by the list of separators until the chunk size is reached, using the earlier mentioned separators where possible. This is useful for splitting text fields by paragraphs, sentences, words, etc.
type DestinationPineconeUpdateBySeparator struct {
	// Whether to keep the separator in the resulting chunks
	KeepSeparator *bool                                           `default:"false" json:"keep_separator"`
	mode          *DestinationPineconeUpdateSchemasProcessingMode `const:"separator" json:"mode"`
	// List of separator strings to split text fields by. The separator itself needs to be wrapped in double quotes, e.g. to split by the dot character, use ".". To split by a newline, use "\n".
	Separators []string `json:"separators,omitempty"`
}

func (d DestinationPineconeUpdateBySeparator) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeUpdateBySeparator) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeUpdateBySeparator) GetKeepSeparator() *bool {
	if o == nil {
		return nil
	}
	return o.KeepSeparator
}

func (o *DestinationPineconeUpdateBySeparator) GetMode() *DestinationPineconeUpdateSchemasProcessingMode {
	return DestinationPineconeUpdateSchemasProcessingModeSeparator.ToPointer()
}

func (o *DestinationPineconeUpdateBySeparator) GetSeparators() []string {
	if o == nil {
		return nil
	}
	return o.Separators
}

type DestinationPineconeUpdateTextSplitterType string

const (
	DestinationPineconeUpdateTextSplitterTypeDestinationPineconeUpdateBySeparator           DestinationPineconeUpdateTextSplitterType = "destination-pinecone-update_By Separator"
	DestinationPineconeUpdateTextSplitterTypeDestinationPineconeUpdateByMarkdownHeader      DestinationPineconeUpdateTextSplitterType = "destination-pinecone-update_By Markdown header"
	DestinationPineconeUpdateTextSplitterTypeDestinationPineconeUpdateByProgrammingLanguage DestinationPineconeUpdateTextSplitterType = "destination-pinecone-update_By Programming Language"
)

type DestinationPineconeUpdateTextSplitter struct {
	DestinationPineconeUpdateBySeparator           *DestinationPineconeUpdateBySeparator
	DestinationPineconeUpdateByMarkdownHeader      *DestinationPineconeUpdateByMarkdownHeader
	DestinationPineconeUpdateByProgrammingLanguage *DestinationPineconeUpdateByProgrammingLanguage

	Type DestinationPineconeUpdateTextSplitterType
}

func CreateDestinationPineconeUpdateTextSplitterDestinationPineconeUpdateBySeparator(destinationPineconeUpdateBySeparator DestinationPineconeUpdateBySeparator) DestinationPineconeUpdateTextSplitter {
	typ := DestinationPineconeUpdateTextSplitterTypeDestinationPineconeUpdateBySeparator

	return DestinationPineconeUpdateTextSplitter{
		DestinationPineconeUpdateBySeparator: &destinationPineconeUpdateBySeparator,
		Type:                                 typ,
	}
}

func CreateDestinationPineconeUpdateTextSplitterDestinationPineconeUpdateByMarkdownHeader(destinationPineconeUpdateByMarkdownHeader DestinationPineconeUpdateByMarkdownHeader) DestinationPineconeUpdateTextSplitter {
	typ := DestinationPineconeUpdateTextSplitterTypeDestinationPineconeUpdateByMarkdownHeader

	return DestinationPineconeUpdateTextSplitter{
		DestinationPineconeUpdateByMarkdownHeader: &destinationPineconeUpdateByMarkdownHeader,
		Type: typ,
	}
}

func CreateDestinationPineconeUpdateTextSplitterDestinationPineconeUpdateByProgrammingLanguage(destinationPineconeUpdateByProgrammingLanguage DestinationPineconeUpdateByProgrammingLanguage) DestinationPineconeUpdateTextSplitter {
	typ := DestinationPineconeUpdateTextSplitterTypeDestinationPineconeUpdateByProgrammingLanguage

	return DestinationPineconeUpdateTextSplitter{
		DestinationPineconeUpdateByProgrammingLanguage: &destinationPineconeUpdateByProgrammingLanguage,
		Type: typ,
	}
}

func (u *DestinationPineconeUpdateTextSplitter) UnmarshalJSON(data []byte) error {

	destinationPineconeUpdateByMarkdownHeader := new(DestinationPineconeUpdateByMarkdownHeader)
	if err := utils.UnmarshalJSON(data, &destinationPineconeUpdateByMarkdownHeader, "", true, true); err == nil {
		u.DestinationPineconeUpdateByMarkdownHeader = destinationPineconeUpdateByMarkdownHeader
		u.Type = DestinationPineconeUpdateTextSplitterTypeDestinationPineconeUpdateByMarkdownHeader
		return nil
	}

	destinationPineconeUpdateByProgrammingLanguage := new(DestinationPineconeUpdateByProgrammingLanguage)
	if err := utils.UnmarshalJSON(data, &destinationPineconeUpdateByProgrammingLanguage, "", true, true); err == nil {
		u.DestinationPineconeUpdateByProgrammingLanguage = destinationPineconeUpdateByProgrammingLanguage
		u.Type = DestinationPineconeUpdateTextSplitterTypeDestinationPineconeUpdateByProgrammingLanguage
		return nil
	}

	destinationPineconeUpdateBySeparator := new(DestinationPineconeUpdateBySeparator)
	if err := utils.UnmarshalJSON(data, &destinationPineconeUpdateBySeparator, "", true, true); err == nil {
		u.DestinationPineconeUpdateBySeparator = destinationPineconeUpdateBySeparator
		u.Type = DestinationPineconeUpdateTextSplitterTypeDestinationPineconeUpdateBySeparator
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationPineconeUpdateTextSplitter) MarshalJSON() ([]byte, error) {
	if u.DestinationPineconeUpdateBySeparator != nil {
		return utils.MarshalJSON(u.DestinationPineconeUpdateBySeparator, "", true)
	}

	if u.DestinationPineconeUpdateByMarkdownHeader != nil {
		return utils.MarshalJSON(u.DestinationPineconeUpdateByMarkdownHeader, "", true)
	}

	if u.DestinationPineconeUpdateByProgrammingLanguage != nil {
		return utils.MarshalJSON(u.DestinationPineconeUpdateByProgrammingLanguage, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationPineconeUpdateProcessingConfigModel struct {
	// Size of overlap between chunks in tokens to store in vector store to better capture relevant context
	ChunkOverlap *int64 `default:"0" json:"chunk_overlap"`
	// Size of chunks in tokens to store in vector store (make sure it is not too big for the context if your LLM)
	ChunkSize int64 `json:"chunk_size"`
	// List of fields to rename. Not applicable for nested fields, but can be used to rename fields already flattened via dot notation.
	FieldNameMappings []DestinationPineconeUpdateFieldNameMappingConfigModel `json:"field_name_mappings,omitempty"`
	// List of fields in the record that should be stored as metadata. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered metadata fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array. When specifying nested paths, all matching values are flattened into an array set to a field named by the path.
	MetadataFields []string `json:"metadata_fields,omitempty"`
	// List of fields in the record that should be used to calculate the embedding. The field list is applied to all streams in the same way and non-existing fields are ignored. If none are defined, all fields are considered text fields. When specifying text fields, you can access nested fields in the record by using dot notation, e.g. `user.name` will access the `name` field in the `user` object. It's also possible to use wildcards to access all fields in an object, e.g. `users.*.name` will access all `names` fields in all entries of the `users` array.
	TextFields []string `json:"text_fields,omitempty"`
	// Split text fields into chunks based on the specified method.
	TextSplitter *DestinationPineconeUpdateTextSplitter `json:"text_splitter,omitempty"`
}

func (d DestinationPineconeUpdateProcessingConfigModel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationPineconeUpdateProcessingConfigModel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationPineconeUpdateProcessingConfigModel) GetChunkOverlap() *int64 {
	if o == nil {
		return nil
	}
	return o.ChunkOverlap
}

func (o *DestinationPineconeUpdateProcessingConfigModel) GetChunkSize() int64 {
	if o == nil {
		return 0
	}
	return o.ChunkSize
}

func (o *DestinationPineconeUpdateProcessingConfigModel) GetFieldNameMappings() []DestinationPineconeUpdateFieldNameMappingConfigModel {
	if o == nil {
		return nil
	}
	return o.FieldNameMappings
}

func (o *DestinationPineconeUpdateProcessingConfigModel) GetMetadataFields() []string {
	if o == nil {
		return nil
	}
	return o.MetadataFields
}

func (o *DestinationPineconeUpdateProcessingConfigModel) GetTextFields() []string {
	if o == nil {
		return nil
	}
	return o.TextFields
}

func (o *DestinationPineconeUpdateProcessingConfigModel) GetTextSplitter() *DestinationPineconeUpdateTextSplitter {
	if o == nil {
		return nil
	}
	return o.TextSplitter
}

type DestinationPineconeUpdate struct {
	// Embedding configuration
	Embedding DestinationPineconeUpdateEmbedding `json:"embedding"`
	// Pinecone is a popular vector store that can be used to store and retrieve embeddings.
	Indexing   DestinationPineconeUpdateIndexing              `json:"indexing"`
	Processing DestinationPineconeUpdateProcessingConfigModel `json:"processing"`
}

func (o *DestinationPineconeUpdate) GetEmbedding() DestinationPineconeUpdateEmbedding {
	if o == nil {
		return DestinationPineconeUpdateEmbedding{}
	}
	return o.Embedding
}

func (o *DestinationPineconeUpdate) GetIndexing() DestinationPineconeUpdateIndexing {
	if o == nil {
		return DestinationPineconeUpdateIndexing{}
	}
	return o.Indexing
}

func (o *DestinationPineconeUpdate) GetProcessing() DestinationPineconeUpdateProcessingConfigModel {
	if o == nil {
		return DestinationPineconeUpdateProcessingConfigModel{}
	}
	return o.Processing
}
