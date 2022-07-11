/*
sudo yum -y install tmux htop irssi util-linux-user

sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
sudo chsh -s $(which zsh) $(whoami)

curl -fsSL https://tailscale.com/install.sh | sh
sudo tailscale up

sudo yum -y install autoconf automake gcc gcc-c++ make boost-devel zlib-devel ncurses-devel protobuf-devel openssl-devel
cd /usr/local/src
sudo wget http://mosh.mit.edu/mosh-1.3.2.tar.gz
sudo tar xvf mosh-1.3.2.tar.gz
cd mosh-1.3.2
sudo ./autogen.sh
sudo ./configure
sudo make
sudo make install

irssi -n wrmsr

export GOPATH="$HOME/go/1.18.3"
export PATH="$PATH:/usr/local/go/bin:$GOPATH/bin"

*/
package aws

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type EC2DescribeInstancesAPI interface {
	DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

func GetInstances(c context.Context, api EC2DescribeInstancesAPI, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return api.DescribeInstances(c, input)
}

func TestEc2Describe(t *testing.T) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstancesInput{}

	result, err := GetInstances(context.Background(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err)
		return
	}

	for _, r := range result.Reservations {
		fmt.Println("Reservation ID: " + *r.ReservationId)
		fmt.Println("Instance IDs:")
		for _, i := range r.Instances {
			fmt.Println("   " + *i.InstanceId)
		}

		fmt.Println("")
	}
}

type EC2CreateInstanceAPI interface {
	RunInstances(ctx context.Context, params *ec2.RunInstancesInput, optFns ...func(*ec2.Options)) (*ec2.RunInstancesOutput, error)
}

func MakeInstance(c context.Context, api EC2CreateInstanceAPI, input *ec2.RunInstancesInput) (*ec2.RunInstancesOutput, error) {
	return api.RunInstances(c, input)
}

/*
Amazon Linux 2 AMI (HVM) - Kernel 5.10, SSD Volume Type
ami-0d9858aa3c6322f73 (64-bit (x86))
ami-0825c6b3598e9754d (64-bit (Arm))

Ubuntu Server 22.04 LTS (HVM), SSD Volume Type
ami-085284d24fe829cd0 (64-bit (x86))
ami-05d39b0802b3a7f8e (64-bit (Arm))

Ubuntu Server 20.04 LTS (HVM), SSD Volume Type
ami-01154c8b2e9a14885 (64-bit (x86))
ami-044e9967964c7ed7e (64-bit (Arm))
*/

//func TestEc2Launch(t *testing.T) {
//	cfg, err := config.LoadDefaultConfig(context.TODO())
//	if err != nil {
//		panic("configuration error, " + err.Error())
//	}
//
//	client := ec2.NewFromConfig(cfg)
//
//	input := &ec2.RunInstancesInput{
//		ImageId:          aws.String("ami-0d9858aa3c6322f73"),
//		InstanceType:     types.InstanceTypeT2Micro,
//		MinCount:         aws.Int32(1),
//		MaxCount:         aws.Int32(1),
//		KeyName:          aws.String("spinlock-1"),
//		SecurityGroupIds: []string{"sg-53b64637"},
//		SubnetId:         aws.String("subnet-30d7b469"),
//		// vpc-23f98946
//		TagSpecifications: []types.TagSpecification{
//			{
//				ResourceType: types.ResourceTypeInstance,
//				Tags: []types.Tag{
//					{
//						Key:   aws.String("foo"),
//						Value: aws.String("bar"),
//					},
//				},
//			},
//		},
//	}
//
//	result, err := MakeInstance(context.TODO(), client, input)
//	if err != nil {
//		fmt.Println("Got an error creating an instance:")
//		fmt.Println(err)
//		return
//	}
//
//	fmt.Println("Created tagged instance with ID " + *result.Instances[0].InstanceId)
//}
