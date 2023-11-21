// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

// DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod - Connect through a jump server tunnel host using username and password authentication
type DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod string

const (
	DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod = "SSH_PASSWORD_AUTH"
)

func (e DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod) ToPointer() *DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod {
	return &e
}

func (e *DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_PASSWORD_AUTH":
		*e = DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod: %v", v)
	}
}

// DestinationRedshiftUpdatePasswordAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationRedshiftUpdatePasswordAuthentication struct {
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and password authentication
	tunnelMethod DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod `const:"SSH_PASSWORD_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host
	TunnelUser string `json:"tunnel_user"`
	// OS-level password for logging into the jump server host
	TunnelUserPassword string `json:"tunnel_user_password"`
}

func (d DestinationRedshiftUpdatePasswordAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationRedshiftUpdatePasswordAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationRedshiftUpdatePasswordAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationRedshiftUpdatePasswordAuthentication) GetTunnelMethod() DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethod {
	return DestinationRedshiftUpdateSchemasTunnelMethodTunnelMethodSSHPasswordAuth
}

func (o *DestinationRedshiftUpdatePasswordAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationRedshiftUpdatePasswordAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

func (o *DestinationRedshiftUpdatePasswordAuthentication) GetTunnelUserPassword() string {
	if o == nil {
		return ""
	}
	return o.TunnelUserPassword
}

// DestinationRedshiftUpdateSchemasTunnelMethod - Connect through a jump server tunnel host using username and ssh key
type DestinationRedshiftUpdateSchemasTunnelMethod string

const (
	DestinationRedshiftUpdateSchemasTunnelMethodSSHKeyAuth DestinationRedshiftUpdateSchemasTunnelMethod = "SSH_KEY_AUTH"
)

func (e DestinationRedshiftUpdateSchemasTunnelMethod) ToPointer() *DestinationRedshiftUpdateSchemasTunnelMethod {
	return &e
}

func (e *DestinationRedshiftUpdateSchemasTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SSH_KEY_AUTH":
		*e = DestinationRedshiftUpdateSchemasTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateSchemasTunnelMethod: %v", v)
	}
}

// DestinationRedshiftUpdateSSHKeyAuthentication - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationRedshiftUpdateSSHKeyAuthentication struct {
	// OS-level user account ssh key credentials in RSA PEM format ( created with ssh-keygen -t rsa -m PEM -f myuser_rsa )
	SSHKey string `json:"ssh_key"`
	// Hostname of the jump server host that allows inbound ssh tunnel.
	TunnelHost string `json:"tunnel_host"`
	// Connect through a jump server tunnel host using username and ssh key
	tunnelMethod DestinationRedshiftUpdateSchemasTunnelMethod `const:"SSH_KEY_AUTH" json:"tunnel_method"`
	// Port on the proxy/jump server that accepts inbound ssh connections.
	TunnelPort *int64 `default:"22" json:"tunnel_port"`
	// OS-level username for logging into the jump server host.
	TunnelUser string `json:"tunnel_user"`
}

func (d DestinationRedshiftUpdateSSHKeyAuthentication) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationRedshiftUpdateSSHKeyAuthentication) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationRedshiftUpdateSSHKeyAuthentication) GetSSHKey() string {
	if o == nil {
		return ""
	}
	return o.SSHKey
}

func (o *DestinationRedshiftUpdateSSHKeyAuthentication) GetTunnelHost() string {
	if o == nil {
		return ""
	}
	return o.TunnelHost
}

func (o *DestinationRedshiftUpdateSSHKeyAuthentication) GetTunnelMethod() DestinationRedshiftUpdateSchemasTunnelMethod {
	return DestinationRedshiftUpdateSchemasTunnelMethodSSHKeyAuth
}

func (o *DestinationRedshiftUpdateSSHKeyAuthentication) GetTunnelPort() *int64 {
	if o == nil {
		return nil
	}
	return o.TunnelPort
}

func (o *DestinationRedshiftUpdateSSHKeyAuthentication) GetTunnelUser() string {
	if o == nil {
		return ""
	}
	return o.TunnelUser
}

// DestinationRedshiftUpdateTunnelMethod - No ssh tunnel needed to connect to database
type DestinationRedshiftUpdateTunnelMethod string

const (
	DestinationRedshiftUpdateTunnelMethodNoTunnel DestinationRedshiftUpdateTunnelMethod = "NO_TUNNEL"
)

func (e DestinationRedshiftUpdateTunnelMethod) ToPointer() *DestinationRedshiftUpdateTunnelMethod {
	return &e
}

func (e *DestinationRedshiftUpdateTunnelMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NO_TUNNEL":
		*e = DestinationRedshiftUpdateTunnelMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateTunnelMethod: %v", v)
	}
}

// DestinationRedshiftUpdateNoTunnel - Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
type DestinationRedshiftUpdateNoTunnel struct {
	// No ssh tunnel needed to connect to database
	tunnelMethod DestinationRedshiftUpdateTunnelMethod `const:"NO_TUNNEL" json:"tunnel_method"`
}

func (d DestinationRedshiftUpdateNoTunnel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationRedshiftUpdateNoTunnel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationRedshiftUpdateNoTunnel) GetTunnelMethod() DestinationRedshiftUpdateTunnelMethod {
	return DestinationRedshiftUpdateTunnelMethodNoTunnel
}

type DestinationRedshiftUpdateSSHTunnelMethodType string

const (
	DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateNoTunnel               DestinationRedshiftUpdateSSHTunnelMethodType = "destination-redshift-update_No Tunnel"
	DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateSSHKeyAuthentication   DestinationRedshiftUpdateSSHTunnelMethodType = "destination-redshift-update_SSH Key Authentication"
	DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdatePasswordAuthentication DestinationRedshiftUpdateSSHTunnelMethodType = "destination-redshift-update_Password Authentication"
)

type DestinationRedshiftUpdateSSHTunnelMethod struct {
	DestinationRedshiftUpdateNoTunnel               *DestinationRedshiftUpdateNoTunnel
	DestinationRedshiftUpdateSSHKeyAuthentication   *DestinationRedshiftUpdateSSHKeyAuthentication
	DestinationRedshiftUpdatePasswordAuthentication *DestinationRedshiftUpdatePasswordAuthentication

	Type DestinationRedshiftUpdateSSHTunnelMethodType
}

func CreateDestinationRedshiftUpdateSSHTunnelMethodDestinationRedshiftUpdateNoTunnel(destinationRedshiftUpdateNoTunnel DestinationRedshiftUpdateNoTunnel) DestinationRedshiftUpdateSSHTunnelMethod {
	typ := DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateNoTunnel

	return DestinationRedshiftUpdateSSHTunnelMethod{
		DestinationRedshiftUpdateNoTunnel: &destinationRedshiftUpdateNoTunnel,
		Type:                              typ,
	}
}

func CreateDestinationRedshiftUpdateSSHTunnelMethodDestinationRedshiftUpdateSSHKeyAuthentication(destinationRedshiftUpdateSSHKeyAuthentication DestinationRedshiftUpdateSSHKeyAuthentication) DestinationRedshiftUpdateSSHTunnelMethod {
	typ := DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateSSHKeyAuthentication

	return DestinationRedshiftUpdateSSHTunnelMethod{
		DestinationRedshiftUpdateSSHKeyAuthentication: &destinationRedshiftUpdateSSHKeyAuthentication,
		Type: typ,
	}
}

func CreateDestinationRedshiftUpdateSSHTunnelMethodDestinationRedshiftUpdatePasswordAuthentication(destinationRedshiftUpdatePasswordAuthentication DestinationRedshiftUpdatePasswordAuthentication) DestinationRedshiftUpdateSSHTunnelMethod {
	typ := DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdatePasswordAuthentication

	return DestinationRedshiftUpdateSSHTunnelMethod{
		DestinationRedshiftUpdatePasswordAuthentication: &destinationRedshiftUpdatePasswordAuthentication,
		Type: typ,
	}
}

func (u *DestinationRedshiftUpdateSSHTunnelMethod) UnmarshalJSON(data []byte) error {

	destinationRedshiftUpdateNoTunnel := new(DestinationRedshiftUpdateNoTunnel)
	if err := utils.UnmarshalJSON(data, &destinationRedshiftUpdateNoTunnel, "", true, true); err == nil {
		u.DestinationRedshiftUpdateNoTunnel = destinationRedshiftUpdateNoTunnel
		u.Type = DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateNoTunnel
		return nil
	}

	destinationRedshiftUpdateSSHKeyAuthentication := new(DestinationRedshiftUpdateSSHKeyAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationRedshiftUpdateSSHKeyAuthentication, "", true, true); err == nil {
		u.DestinationRedshiftUpdateSSHKeyAuthentication = destinationRedshiftUpdateSSHKeyAuthentication
		u.Type = DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdateSSHKeyAuthentication
		return nil
	}

	destinationRedshiftUpdatePasswordAuthentication := new(DestinationRedshiftUpdatePasswordAuthentication)
	if err := utils.UnmarshalJSON(data, &destinationRedshiftUpdatePasswordAuthentication, "", true, true); err == nil {
		u.DestinationRedshiftUpdatePasswordAuthentication = destinationRedshiftUpdatePasswordAuthentication
		u.Type = DestinationRedshiftUpdateSSHTunnelMethodTypeDestinationRedshiftUpdatePasswordAuthentication
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationRedshiftUpdateSSHTunnelMethod) MarshalJSON() ([]byte, error) {
	if u.DestinationRedshiftUpdateNoTunnel != nil {
		return utils.MarshalJSON(u.DestinationRedshiftUpdateNoTunnel, "", true)
	}

	if u.DestinationRedshiftUpdateSSHKeyAuthentication != nil {
		return utils.MarshalJSON(u.DestinationRedshiftUpdateSSHKeyAuthentication, "", true)
	}

	if u.DestinationRedshiftUpdatePasswordAuthentication != nil {
		return utils.MarshalJSON(u.DestinationRedshiftUpdatePasswordAuthentication, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationRedshiftUpdateSchemasMethod string

const (
	DestinationRedshiftUpdateSchemasMethodStandard DestinationRedshiftUpdateSchemasMethod = "Standard"
)

func (e DestinationRedshiftUpdateSchemasMethod) ToPointer() *DestinationRedshiftUpdateSchemasMethod {
	return &e
}

func (e *DestinationRedshiftUpdateSchemasMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Standard":
		*e = DestinationRedshiftUpdateSchemasMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateSchemasMethod: %v", v)
	}
}

// Standard - <i>(not recommended)</i> Direct loading using SQL INSERT statements. This method is extremely inefficient and provided only for quick testing. In all other cases, you should use S3 uploading.
type Standard struct {
	method DestinationRedshiftUpdateSchemasMethod `const:"Standard" json:"method"`
}

func (s Standard) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *Standard) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Standard) GetMethod() DestinationRedshiftUpdateSchemasMethod {
	return DestinationRedshiftUpdateSchemasMethodStandard
}

type DestinationRedshiftUpdateEncryptionType string

const (
	DestinationRedshiftUpdateEncryptionTypeAesCbcEnvelope DestinationRedshiftUpdateEncryptionType = "aes_cbc_envelope"
)

func (e DestinationRedshiftUpdateEncryptionType) ToPointer() *DestinationRedshiftUpdateEncryptionType {
	return &e
}

func (e *DestinationRedshiftUpdateEncryptionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aes_cbc_envelope":
		*e = DestinationRedshiftUpdateEncryptionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateEncryptionType: %v", v)
	}
}

// AESCBCEnvelopeEncryption - Staging data will be encrypted using AES-CBC envelope encryption.
type AESCBCEnvelopeEncryption struct {
	encryptionType *DestinationRedshiftUpdateEncryptionType `const:"aes_cbc_envelope" json:"encryption_type"`
	// The key, base64-encoded. Must be either 128, 192, or 256 bits. Leave blank to have Airbyte generate an ephemeral key for each sync.
	KeyEncryptingKey *string `json:"key_encrypting_key,omitempty"`
}

func (a AESCBCEnvelopeEncryption) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AESCBCEnvelopeEncryption) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *AESCBCEnvelopeEncryption) GetEncryptionType() *DestinationRedshiftUpdateEncryptionType {
	return DestinationRedshiftUpdateEncryptionTypeAesCbcEnvelope.ToPointer()
}

func (o *AESCBCEnvelopeEncryption) GetKeyEncryptingKey() *string {
	if o == nil {
		return nil
	}
	return o.KeyEncryptingKey
}

type EncryptionType string

const (
	EncryptionTypeNone EncryptionType = "none"
)

func (e EncryptionType) ToPointer() *EncryptionType {
	return &e
}

func (e *EncryptionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "none":
		*e = EncryptionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EncryptionType: %v", v)
	}
}

// NoEncryption - Staging data will be stored in plaintext.
type NoEncryption struct {
	encryptionType *EncryptionType `const:"none" json:"encryption_type"`
}

func (n NoEncryption) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *NoEncryption) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *NoEncryption) GetEncryptionType() *EncryptionType {
	return EncryptionTypeNone.ToPointer()
}

type DestinationRedshiftUpdateEncryptionUnionType string

const (
	DestinationRedshiftUpdateEncryptionUnionTypeNoEncryption             DestinationRedshiftUpdateEncryptionUnionType = "No encryption"
	DestinationRedshiftUpdateEncryptionUnionTypeAESCBCEnvelopeEncryption DestinationRedshiftUpdateEncryptionUnionType = "AES-CBC envelope encryption"
)

type DestinationRedshiftUpdateEncryption struct {
	NoEncryption             *NoEncryption
	AESCBCEnvelopeEncryption *AESCBCEnvelopeEncryption

	Type DestinationRedshiftUpdateEncryptionUnionType
}

func CreateDestinationRedshiftUpdateEncryptionNoEncryption(noEncryption NoEncryption) DestinationRedshiftUpdateEncryption {
	typ := DestinationRedshiftUpdateEncryptionUnionTypeNoEncryption

	return DestinationRedshiftUpdateEncryption{
		NoEncryption: &noEncryption,
		Type:         typ,
	}
}

func CreateDestinationRedshiftUpdateEncryptionAESCBCEnvelopeEncryption(aesCBCEnvelopeEncryption AESCBCEnvelopeEncryption) DestinationRedshiftUpdateEncryption {
	typ := DestinationRedshiftUpdateEncryptionUnionTypeAESCBCEnvelopeEncryption

	return DestinationRedshiftUpdateEncryption{
		AESCBCEnvelopeEncryption: &aesCBCEnvelopeEncryption,
		Type:                     typ,
	}
}

func (u *DestinationRedshiftUpdateEncryption) UnmarshalJSON(data []byte) error {

	noEncryption := new(NoEncryption)
	if err := utils.UnmarshalJSON(data, &noEncryption, "", true, true); err == nil {
		u.NoEncryption = noEncryption
		u.Type = DestinationRedshiftUpdateEncryptionUnionTypeNoEncryption
		return nil
	}

	aesCBCEnvelopeEncryption := new(AESCBCEnvelopeEncryption)
	if err := utils.UnmarshalJSON(data, &aesCBCEnvelopeEncryption, "", true, true); err == nil {
		u.AESCBCEnvelopeEncryption = aesCBCEnvelopeEncryption
		u.Type = DestinationRedshiftUpdateEncryptionUnionTypeAESCBCEnvelopeEncryption
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u DestinationRedshiftUpdateEncryption) MarshalJSON() ([]byte, error) {
	if u.NoEncryption != nil {
		return utils.MarshalJSON(u.NoEncryption, "", true)
	}

	if u.AESCBCEnvelopeEncryption != nil {
		return utils.MarshalJSON(u.AESCBCEnvelopeEncryption, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationRedshiftUpdateMethod string

const (
	DestinationRedshiftUpdateMethodS3Staging DestinationRedshiftUpdateMethod = "S3 Staging"
)

func (e DestinationRedshiftUpdateMethod) ToPointer() *DestinationRedshiftUpdateMethod {
	return &e
}

func (e *DestinationRedshiftUpdateMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3 Staging":
		*e = DestinationRedshiftUpdateMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateMethod: %v", v)
	}
}

// DestinationRedshiftUpdateS3BucketRegion - The region of the S3 staging bucket.
type DestinationRedshiftUpdateS3BucketRegion string

const (
	DestinationRedshiftUpdateS3BucketRegionUnknown      DestinationRedshiftUpdateS3BucketRegion = ""
	DestinationRedshiftUpdateS3BucketRegionUsEast1      DestinationRedshiftUpdateS3BucketRegion = "us-east-1"
	DestinationRedshiftUpdateS3BucketRegionUsEast2      DestinationRedshiftUpdateS3BucketRegion = "us-east-2"
	DestinationRedshiftUpdateS3BucketRegionUsWest1      DestinationRedshiftUpdateS3BucketRegion = "us-west-1"
	DestinationRedshiftUpdateS3BucketRegionUsWest2      DestinationRedshiftUpdateS3BucketRegion = "us-west-2"
	DestinationRedshiftUpdateS3BucketRegionAfSouth1     DestinationRedshiftUpdateS3BucketRegion = "af-south-1"
	DestinationRedshiftUpdateS3BucketRegionApEast1      DestinationRedshiftUpdateS3BucketRegion = "ap-east-1"
	DestinationRedshiftUpdateS3BucketRegionApSouth1     DestinationRedshiftUpdateS3BucketRegion = "ap-south-1"
	DestinationRedshiftUpdateS3BucketRegionApNortheast1 DestinationRedshiftUpdateS3BucketRegion = "ap-northeast-1"
	DestinationRedshiftUpdateS3BucketRegionApNortheast2 DestinationRedshiftUpdateS3BucketRegion = "ap-northeast-2"
	DestinationRedshiftUpdateS3BucketRegionApNortheast3 DestinationRedshiftUpdateS3BucketRegion = "ap-northeast-3"
	DestinationRedshiftUpdateS3BucketRegionApSoutheast1 DestinationRedshiftUpdateS3BucketRegion = "ap-southeast-1"
	DestinationRedshiftUpdateS3BucketRegionApSoutheast2 DestinationRedshiftUpdateS3BucketRegion = "ap-southeast-2"
	DestinationRedshiftUpdateS3BucketRegionCaCentral1   DestinationRedshiftUpdateS3BucketRegion = "ca-central-1"
	DestinationRedshiftUpdateS3BucketRegionCnNorth1     DestinationRedshiftUpdateS3BucketRegion = "cn-north-1"
	DestinationRedshiftUpdateS3BucketRegionCnNorthwest1 DestinationRedshiftUpdateS3BucketRegion = "cn-northwest-1"
	DestinationRedshiftUpdateS3BucketRegionEuCentral1   DestinationRedshiftUpdateS3BucketRegion = "eu-central-1"
	DestinationRedshiftUpdateS3BucketRegionEuNorth1     DestinationRedshiftUpdateS3BucketRegion = "eu-north-1"
	DestinationRedshiftUpdateS3BucketRegionEuSouth1     DestinationRedshiftUpdateS3BucketRegion = "eu-south-1"
	DestinationRedshiftUpdateS3BucketRegionEuWest1      DestinationRedshiftUpdateS3BucketRegion = "eu-west-1"
	DestinationRedshiftUpdateS3BucketRegionEuWest2      DestinationRedshiftUpdateS3BucketRegion = "eu-west-2"
	DestinationRedshiftUpdateS3BucketRegionEuWest3      DestinationRedshiftUpdateS3BucketRegion = "eu-west-3"
	DestinationRedshiftUpdateS3BucketRegionSaEast1      DestinationRedshiftUpdateS3BucketRegion = "sa-east-1"
	DestinationRedshiftUpdateS3BucketRegionMeSouth1     DestinationRedshiftUpdateS3BucketRegion = "me-south-1"
)

func (e DestinationRedshiftUpdateS3BucketRegion) ToPointer() *DestinationRedshiftUpdateS3BucketRegion {
	return &e
}

func (e *DestinationRedshiftUpdateS3BucketRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "us-east-1":
		fallthrough
	case "us-east-2":
		fallthrough
	case "us-west-1":
		fallthrough
	case "us-west-2":
		fallthrough
	case "af-south-1":
		fallthrough
	case "ap-east-1":
		fallthrough
	case "ap-south-1":
		fallthrough
	case "ap-northeast-1":
		fallthrough
	case "ap-northeast-2":
		fallthrough
	case "ap-northeast-3":
		fallthrough
	case "ap-southeast-1":
		fallthrough
	case "ap-southeast-2":
		fallthrough
	case "ca-central-1":
		fallthrough
	case "cn-north-1":
		fallthrough
	case "cn-northwest-1":
		fallthrough
	case "eu-central-1":
		fallthrough
	case "eu-north-1":
		fallthrough
	case "eu-south-1":
		fallthrough
	case "eu-west-1":
		fallthrough
	case "eu-west-2":
		fallthrough
	case "eu-west-3":
		fallthrough
	case "sa-east-1":
		fallthrough
	case "me-south-1":
		*e = DestinationRedshiftUpdateS3BucketRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRedshiftUpdateS3BucketRegion: %v", v)
	}
}

// S3Staging - <i>(recommended)</i> Uploads data to S3 and then uses a COPY to insert the data into Redshift. COPY is recommended for production workloads for better speed and scalability. See <a href="https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-bucket.html">AWS docs</a> for more details.
type S3Staging struct {
	// This ID grants access to the above S3 staging bucket. Airbyte requires Read and Write permissions to the given bucket. See <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">AWS docs</a> on how to generate an access key ID and secret access key.
	AccessKeyID string `json:"access_key_id"`
	// How to encrypt the staging data
	Encryption *DestinationRedshiftUpdateEncryption `json:"encryption,omitempty"`
	// Number of file buffers allocated for writing data. Increasing this number is beneficial for connections using Change Data Capture (CDC) and up to the number of streams within a connection. Increasing the number of file buffers past the maximum number of streams has deteriorating effects
	FileBufferCount *int64 `default:"10" json:"file_buffer_count"`
	// The pattern allows you to set the file-name format for the S3 staging file(s)
	FileNamePattern *string                         `json:"file_name_pattern,omitempty"`
	method          DestinationRedshiftUpdateMethod `const:"S3 Staging" json:"method"`
	// Whether to delete the staging files from S3 after completing the sync. See <a href="https://docs.airbyte.com/integrations/destinations/redshift/#:~:text=the%20root%20directory.-,Purge%20Staging%20Data,-Whether%20to%20delete"> docs</a> for details.
	PurgeStagingData *bool `default:"true" json:"purge_staging_data"`
	// The name of the staging S3 bucket.
	S3BucketName string `json:"s3_bucket_name"`
	// The directory under the S3 bucket where data will be written. If not provided, then defaults to the root directory. See <a href="https://docs.aws.amazon.com/prescriptive-guidance/latest/defining-bucket-names-data-lakes/faq.html#:~:text=be%20globally%20unique.-,For%20S3%20bucket%20paths,-%2C%20you%20can%20use">path's name recommendations</a> for more details.
	S3BucketPath *string `json:"s3_bucket_path,omitempty"`
	// The region of the S3 staging bucket.
	S3BucketRegion *DestinationRedshiftUpdateS3BucketRegion `default:"" json:"s3_bucket_region"`
	// The corresponding secret to the above access key id. See <a href="https://docs.aws.amazon.com/general/latest/gr/aws-sec-cred-types.html#access-keys-and-secret-access-keys">AWS docs</a> on how to generate an access key ID and secret access key.
	SecretAccessKey string `json:"secret_access_key"`
}

func (s S3Staging) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *S3Staging) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *S3Staging) GetAccessKeyID() string {
	if o == nil {
		return ""
	}
	return o.AccessKeyID
}

func (o *S3Staging) GetEncryption() *DestinationRedshiftUpdateEncryption {
	if o == nil {
		return nil
	}
	return o.Encryption
}

func (o *S3Staging) GetFileBufferCount() *int64 {
	if o == nil {
		return nil
	}
	return o.FileBufferCount
}

func (o *S3Staging) GetFileNamePattern() *string {
	if o == nil {
		return nil
	}
	return o.FileNamePattern
}

func (o *S3Staging) GetMethod() DestinationRedshiftUpdateMethod {
	return DestinationRedshiftUpdateMethodS3Staging
}

func (o *S3Staging) GetPurgeStagingData() *bool {
	if o == nil {
		return nil
	}
	return o.PurgeStagingData
}

func (o *S3Staging) GetS3BucketName() string {
	if o == nil {
		return ""
	}
	return o.S3BucketName
}

func (o *S3Staging) GetS3BucketPath() *string {
	if o == nil {
		return nil
	}
	return o.S3BucketPath
}

func (o *S3Staging) GetS3BucketRegion() *DestinationRedshiftUpdateS3BucketRegion {
	if o == nil {
		return nil
	}
	return o.S3BucketRegion
}

func (o *S3Staging) GetSecretAccessKey() string {
	if o == nil {
		return ""
	}
	return o.SecretAccessKey
}

type UploadingMethodType string

const (
	UploadingMethodTypeS3Staging UploadingMethodType = "S3 Staging"
	UploadingMethodTypeStandard  UploadingMethodType = "Standard"
)

type UploadingMethod struct {
	S3Staging *S3Staging
	Standard  *Standard

	Type UploadingMethodType
}

func CreateUploadingMethodS3Staging(s3Staging S3Staging) UploadingMethod {
	typ := UploadingMethodTypeS3Staging

	return UploadingMethod{
		S3Staging: &s3Staging,
		Type:      typ,
	}
}

func CreateUploadingMethodStandard(standard Standard) UploadingMethod {
	typ := UploadingMethodTypeStandard

	return UploadingMethod{
		Standard: &standard,
		Type:     typ,
	}
}

func (u *UploadingMethod) UnmarshalJSON(data []byte) error {

	standard := new(Standard)
	if err := utils.UnmarshalJSON(data, &standard, "", true, true); err == nil {
		u.Standard = standard
		u.Type = UploadingMethodTypeStandard
		return nil
	}

	s3Staging := new(S3Staging)
	if err := utils.UnmarshalJSON(data, &s3Staging, "", true, true); err == nil {
		u.S3Staging = s3Staging
		u.Type = UploadingMethodTypeS3Staging
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UploadingMethod) MarshalJSON() ([]byte, error) {
	if u.S3Staging != nil {
		return utils.MarshalJSON(u.S3Staging, "", true)
	}

	if u.Standard != nil {
		return utils.MarshalJSON(u.Standard, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type DestinationRedshiftUpdate struct {
	// Name of the database.
	Database string `json:"database"`
	// Host Endpoint of the Redshift Cluster (must include the cluster-id, region and end with .redshift.amazonaws.com)
	Host string `json:"host"`
	// Additional properties to pass to the JDBC URL string when connecting to the database formatted as 'key=value' pairs separated by the symbol '&'. (example: key1=value1&key2=value2&key3=value3).
	JdbcURLParams *string `json:"jdbc_url_params,omitempty"`
	// Password associated with the username.
	Password string `json:"password"`
	// Port of the database.
	Port *int64 `default:"5439" json:"port"`
	// The default schema tables are written to if the source does not specify a namespace. Unless specifically configured, the usual value for this field is "public".
	Schema *string `default:"public" json:"schema"`
	// Whether to initiate an SSH tunnel before connecting to the database, and if so, which kind of authentication to use.
	TunnelMethod *DestinationRedshiftUpdateSSHTunnelMethod `json:"tunnel_method,omitempty"`
	// The way data will be uploaded to Redshift.
	UploadingMethod *UploadingMethod `json:"uploading_method,omitempty"`
	// Username to use to access the database.
	Username string `json:"username"`
}

func (d DestinationRedshiftUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationRedshiftUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationRedshiftUpdate) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *DestinationRedshiftUpdate) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *DestinationRedshiftUpdate) GetJdbcURLParams() *string {
	if o == nil {
		return nil
	}
	return o.JdbcURLParams
}

func (o *DestinationRedshiftUpdate) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *DestinationRedshiftUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *DestinationRedshiftUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DestinationRedshiftUpdate) GetTunnelMethod() *DestinationRedshiftUpdateSSHTunnelMethod {
	if o == nil {
		return nil
	}
	return o.TunnelMethod
}

func (o *DestinationRedshiftUpdate) GetUploadingMethod() *UploadingMethod {
	if o == nil {
		return nil
	}
	return o.UploadingMethod
}

func (o *DestinationRedshiftUpdate) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
