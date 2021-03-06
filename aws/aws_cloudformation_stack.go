// Code is generated. DO NOT EDIT.

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func ListCloudformationStack(client *Client) ([]Resource, error) {
	req := client.cloudformationconn.DescribeStacksRequest(&cloudformation.DescribeStacksInput{})

	var result []Resource

	p := cloudformation.NewDescribeStacksPaginator(req)
	for p.Next(context.Background()) {
		page := p.CurrentPage()

		for _, r := range page.Stacks {

			tags := map[string]string{}
			for _, t := range r.Tags {
				tags[*t.Key] = *t.Value
			}
			t := *r.CreationTime
			result = append(result, Resource{
				Type:      "aws_cloudformation_stack",
				ID:        *r.StackId,
				Region:    client.cloudformationconn.Config.Region,
				Tags:      tags,
				CreatedAt: &t,
			})
		}
	}

	if err := p.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
