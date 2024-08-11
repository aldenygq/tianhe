package pkg 

import (
  cs20151215  "github.com/alibabacloud-go/cs-20151215/v5/client"
  openapi  "github.com/alibabacloud-go/darabonba-openapi/v2/client"
  util  "github.com/alibabacloud-go/tea-utils/v2/service"
  "github.com/alibabacloud-go/tea/tea"
)

func NewClient(ak,sk string) (result *cs20151215.Client, err error) {
  // 工程代码泄露可能会导致 AccessKey 泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考。
  // 建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html。
  config := &openapi.Config{
    // 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_ID。
    AccessKeyId: tea.String(ak),
    // 必填，请确保代码运行环境设置了环境变量 ALIBABA_CLOUD_ACCESS_KEY_SECRET。
    AccessKeySecret: tea.String(sk),
  }
  // Endpoint 请参考 https://api.aliyun.com/product/CS
  config.Endpoint = tea.String("cs.cn-qingdao.aliyuncs.com")
  result = &cs20151215.Client{}
  result, err = cs20151215.NewClient(config)
  if err != nil {
    return nil,err 
  }
  return result, nil 
}
func CloseClient(client *cs20151215.Client) {
    client = nil 
}

func NodeGroupListByAliyun(clusterid,ak,sk string) (*cs20151215.DescribeClusterNodePoolsResponse,error) {
  client,err := NewClient(ak,sk)
  if err != nil {
    return nil,err 
  }
  defer CloseClient(client)
  describeClusterNodePoolsRequest := &cs20151215.DescribeClusterNodePoolsRequest{}
  runtime := &util.RuntimeOptions{}
  headers := make(map[string]*string)
  resp, err := client.DescribeClusterNodePoolsWithOptions(tea.String(clusterid), describeClusterNodePoolsRequest, headers, runtime)
  if err != nil {
    return nil,err
  }

  return resp,err
}
func NodeListByNodeGroup(clusterid,nodegroupid,ak,sk string) (*cs20151215.DescribeClusterNodesResponse,error) {
  client,err := NewClient(ak,sk)
  if err != nil {
    return nil,err 
  }
  defer CloseClient(client)
  describeClusterNodesRequest := &cs20151215.DescribeClusterNodesRequest{
    NodepoolId: tea.String(nodegroupid),
  }
  runtime := &util.RuntimeOptions{}
  headers := make(map[string]*string)
  // 复制代码运行请自行打印 API 的返回值
  resp, err := client.DescribeClusterNodesWithOptions(tea.String(clusterid), describeClusterNodesRequest, headers, runtime)
  if err != nil {
    return nil,err
  }
  return resp,nil 
}