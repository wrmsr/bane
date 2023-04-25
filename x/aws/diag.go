package aws

import (
	"context"
	"reflect"

	httpu "github.com/wrmsr/bane/core/http"
	rfl "github.com/wrmsr/bane/core/reflect"
)

//

type LocalMetadataNetworkInterface struct {
	DeviceNumber string `ec2:"device-number"`
	InterfaceId  string `ec2:"interface-id"`

	// ipv4-associations/

	LocalHostname string `ec2:"local-hostname"`
	LocalIpv4s    string `ec2:"local-ipv4s"`

	Mac     string `ec2:"mac"`
	OwnerId string `ec2:"owner-id"`

	PublicHostname string `ec2:"public-hostname"`
	PublicIpv4s    string `ec2:"public-ipv4s"`

	SecurityGroupIds string `ec2:"security-group-ids"`
	SecurityGroups   string `ec2:"security-groups"`

	SubnetId            string `ec2:"subnet-id"`
	SubnetIpv4CidrBlock string `ec2:"subnet-ipv4-cidr-block"`

	VpcId            string `ec2:"vpc-id"`
	VpcIpv4CidrBlock string `ec2:"vpc-ipv4-cidr-block"`
}

//

type LocalMetadata struct {
	AmiId           string `ec2:"ami-id"`
	AmiLaunchIndex  string `ec2:"ami-launch-index"`
	AmiManifestPath string `ec2:"ami-manifest-path"`

	// block-device-mapping/

	// events/

	Hostname string `ec2:"hostname"`

	// identity-credentials/

	InstanceAction    string `ec2:"instance-action"`
	InstanceId        string `ec2:"instance-id"`
	InstanceLifeCycle string `ec2:"instance-life-cycle"`
	InstanceType      string `ec2:"instance-type"`

	LocalHostname string `ec2:"local-hostname"`
	LocalIpv4     string `ec2:"local-ipv4"`
	Mac           string `ec2:"mac"`

	// metrics/
	// network/
	// placement/

	Profile        string `ec2:"profile"`
	PublicHostname string `ec2:"public-hostname"`
	PublicIpv4     string `ec2:"public-ipv4"`

	// public-keys/

	ReservationId  string `ec2:"reservation-id"`
	SecurityGroups string `ec2:"security-groups"`
}

//

func fillLocalMetadataStruct(ctx context.Context, s any, url string) error {
	v := reflect.ValueOf(s)
	for i, f := range rfl.Fields(v.Type()) {
		p := f.Tag.Get("ec2")
		if p == "" {
			continue
		}
		b, err := httpu.FetchBytes(ctx, url+p)
		if err != nil {
			return err
		}
		v.Field(i).Set(reflect.ValueOf(string(b)))
	}
	return nil
}

const LocalMetadataUrl = "http://169.254.169.254/latest/meta-data/"

func GetLocalMetadata(ctx context.Context) (LocalMetadata, error) {
	var m LocalMetadata
	if err := fillLocalMetadataStruct(ctx, &m, LocalMetadataUrl); err != nil {
		return m, err
	}
	return m, nil
}
