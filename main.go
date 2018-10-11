package main

import (
        "fmt"
        "io/ioutil"
        "log"
        "os
        "github.com/aws/aws-sdk-go-v2/aws"
        "github.com/aws/aws-sdk-go-v2/aws/endpoints"
        "github.com/aws/aws-sdk-go-v2/aws/external"
        "github.com/aws/aws-sdk-go-v2/service/ec2"
        "github.com/spf13/viper"
        _ "https://github.com/asmodeus70/"
)

viper.AddRemoteProvider("git", "https://github.com/asmodeus70/","viper/blob/master/modules/config.yml")
viper.SetConfigType("yaml")
err := viper.ReadRemoteConfig()

func main() {

svc := ec2.New(session.New())
input := &ec2.RunInstancesInput{
    BlockDeviceMappings: []*ec2.BlockDeviceMapping{
        {
            DeviceName: GetString("device"),
            Ebs: &ec2.EbsBlockDevice{
                VolumeSize: GetInt("volsize"),
            },
        },
    },

    ImageId:      GetString("AMI"),
    InstanceType: GetString("instance"),
    KeyName:      GetString("keys"),
    MaxCount:     aws.Int64(1),
    MinCount:     aws.Int64(1),
    SecurityGroupIds: []*string{
        GetString("sg"),
	},

    SubnetId: GetString("subnet"),
    TagSpecifications: []*ec2.TagSpecification{
        {
            ResourceType: aws.String("instance"),
            Tags: []*ec2.Tag{
                {
                    Key:   GetString("tags.key"),
                    Value: GetString("tags.value"),
                },
            },
        },
    },
}

result, err := svc.RunInstances(input)
if err != nil {
    if aerr, ok := err.(awserr.Error); ok {
        switch aerr.Code() {
        default:
            fmt.Println(aerr.Error())
        }
    } else {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
    }
    return
}

fmt.Println(result)

}