package main

import (
	// "fmt"
    "testing"
    "reflect"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func TestGenerateTestSimpleTag(t *testing.T) {
    actual := generateTags("add", "Key=foo,Value=bar")
    var expected []*ec2.Tag
    Tag := &ec2.Tag{Key: aws.String("foo"),Value: aws.String("bar"),}
	expected = append(expected, Tag)

    if ! reflect.DeepEqual(actual, expected) {
        t.Fatal("failed test")
    }
}

func TestGenerateTestMultipleTag(t *testing.T) {
    actual := generateTags("add", "Key=foo,Value=bar Key=baz,Value=qux")
    var expected []*ec2.Tag
    Tag := &ec2.Tag{Key: aws.String("foo"),Value: aws.String("bar"),}
    expected = append(expected, Tag)
    Tag = &ec2.Tag{Key: aws.String("baz"),Value: aws.String("qux"),}
    expected = append(expected, Tag)

    if ! reflect.DeepEqual(actual, expected) {
        t.Fatal("failed test")
    }
}

func TestGenerateTestMultipleTagWithJson(t *testing.T) {
    actual := generateTags("add", "Key=foo,Value=bar Key=amirotate:default,Value={\"NoReboot\":true,\"Retention\":{\"Count\":3}}")
    var expected []*ec2.Tag
    Tag := &ec2.Tag{Key: aws.String("foo"),Value: aws.String("bar"),}
    expected = append(expected, Tag)
    Tag = &ec2.Tag{Key: aws.String("amirotate:default"),Value: aws.String("{\"NoReboot\":true,\"Retention\":{\"Count\":3}}"),}
    expected = append(expected, Tag)

    if ! reflect.DeepEqual(actual, expected) {
        t.Fatal("failed test")
    }
}

func TestGenerateTestMultipleTagNoValue(t *testing.T) {
    actual := generateTags("add", "Key=foo,Value=bar Key=baz,Value=")
    var expected []*ec2.Tag
    Tag := &ec2.Tag{Key: aws.String("foo"),Value: aws.String("bar"),}
    expected = append(expected, Tag)
    Tag = &ec2.Tag{Key: aws.String("baz"),Value: aws.String(""),}
    expected = append(expected, Tag)

    if ! reflect.DeepEqual(actual, expected) {
        t.Fatal("failed test")
    }
}