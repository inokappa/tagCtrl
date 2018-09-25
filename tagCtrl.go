package main

import (
    "os"
    "fmt"
    "flag"
    "strings"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/ec2"

    "github.com/olekukonko/tablewriter"
)

const AppVersion = "0.0.1"

var (
    argProfile = flag.String("profile", "", "Profile 名を指定.")
    argRegion = flag.String("region", "ap-northeast-1", "Region 名を指定.")
    argEndpoint = flag.String("endpoint", "", "AWS API のエンドポイントを指定.")
    argInstances = flag.String("instances", "", "Instance ID 又は Instance Tag 名を指定.")
    argTags = flag.String("tags", "", "Tag Key(Key=) 及び Tag Value(Value=) を指定.")
    argAdd = flag.Bool("add", false, "タグをインスタンスに付与.")
    argDel = flag.Bool("del", false, "タグをインスタンスから削除.")
    argList = flag.Bool("list", false, "インスタンスのタグ一覧を取得.")
    argVersion = flag.Bool("version", false, "バージョンを出力.")
    // argJson = flag.Bool("json", false, "JSON 形式で出力する")
)

func awsEc2Client(profile string, region string) *ec2.EC2 {
    var config aws.Config
    if profile != "" {
        creds := credentials.NewSharedCredentials("", profile)
        config = aws.Config{Region: aws.String(region), Credentials: creds, Endpoint: aws.String(*argEndpoint)}
    } else {
        config = aws.Config{Region: aws.String(region), Endpoint: aws.String(*argEndpoint)}
    }
    sess := session.New(&config)
    ec2Client := ec2.New(sess)
    return ec2Client
}

func createTag(ec2Client *ec2.EC2, instances string, tags []*ec2.Tag) {
    splitedInstances := strings.Split(instances, ",")
    var instanceIds []*string
    for _, i := range splitedInstances {
        instanceIds = append(instanceIds, aws.String(i))
    }
    input := &ec2.CreateTagsInput{
        Resources: instanceIds,
        Tags: tags,
    }
    _, err := ec2Client.CreateTags(input)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}

func deleteTag(ec2Client *ec2.EC2, instances string, tags []*ec2.Tag) {
    splitedInstances := strings.Split(instances, ",")
    var instanceIds []*string
    for _, i := range splitedInstances {
        instanceIds = append(instanceIds, aws.String(i))
    }
    input := &ec2.DeleteTagsInput{
        Resources: instanceIds,
        Tags: tags,
    }
    _, err := ec2Client.DeleteTags(input)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}

func listTag(ec2Client *ec2.EC2, instances string) {
    splitedInstances := strings.Split(instances, ",")
    var instanceIds []*string
    for _, i := range splitedInstances {
        instanceIds = append(instanceIds, aws.String(i))
    }
    input := &ec2.DescribeTagsInput{
        Filters: []*ec2.Filter{
            {
                Name: aws.String("resource-id"),
                Values: instanceIds,
            },
        },
    }
    result, err := ec2Client.DescribeTags(input)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    instance_tags := [][]string{}
    for _, t := range result.Tags {
        instance_tag := []string{
            *t.ResourceId,
            *t.Key,
            *t.Value,
        }
        instance_tags = append(instance_tags, instance_tag)
    }
    outputTbl(instance_tags)
}

func outputTbl(data [][]string) {
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"InstanceId", "Key", "Value"})
    table.SetAutoMergeCells(true)
    table.SetRowLine(true)

    for _, value := range data {
        table.Append(value)
    }
    table.Render()
}

func generateTags(action string, tags string) []*ec2.Tag {
    var allTags []*ec2.Tag
    splitedTags := strings.Split(tags, " ")
    for _, tag := range splitedTags {
        var Tag *ec2.Tag
        var tagKey string
        var tagValue string
        splitedTag := strings.Split(tag, ",Value=")
        splitedT := strings.Split(splitedTag[0], "=")
        if splitedT[0] == "Key" {
            tagKey = splitedT[1]
        } else {
            fmt.Println("Tags Parse error.")
            os.Exit(1)
        }

        if action == "add" || (action == "del" && len(splitedTag) > 1) {
            tagValue = splitedTag[1]
            Tag = &ec2.Tag{
                Key: aws.String(tagKey),
                Value: aws.String(tagValue),
            }
        } else if action == "del" && len(splitedTag) == 1 {
            Tag = &ec2.Tag{
                Key: aws.String(tagKey),
            }
        }
        allTags = append(allTags, Tag)
    }
    return allTags
}

func main() {
    flag.Parse()

    if *argVersion {
        fmt.Println(AppVersion)
        os.Exit(0)
    }

    if *argInstances == "" {
        fmt.Println("Please Set `-instances`.")
        os.Exit(1)
    }

    ec2Client := awsEc2Client(*argProfile, *argRegion)

    if *argAdd {
        allTags := generateTags("add", *argTags)
        createTag(ec2Client, *argInstances, allTags)
    } else if *argDel {
        allTags := generateTags("del", *argTags)
        deleteTag(ec2Client, *argInstances, allTags)
    } else if *argList {
        listTag(ec2Client, *argInstances)
    } else {
        listTag(ec2Client, *argInstances)
    }
}
