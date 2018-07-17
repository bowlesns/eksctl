package eks_test

import (
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"github.com/kubicorn/kubicorn/pkg/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	. "github.com/weaveworks/eksctl/pkg/eks"
	"github.com/weaveworks/eksctl/pkg/eks/mocks"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type MockProvider struct {
	cfn *mocks.CloudFormationAPI
	eks *mocks.EKSAPI
	ec2 *mocks.EC2API
	sts *mocks.STSAPI
}

func (m MockProvider) CloudFormation() cloudformationiface.CloudFormationAPI { return m.cfn }
func (m MockProvider) mockCloudFormation() *mocks.CloudFormationAPI {
	return m.CloudFormation().(*mocks.CloudFormationAPI)
}

func (m MockProvider) EKS() eksiface.EKSAPI   { return m.eks }
func (m MockProvider) mockEKS() *mocks.EKSAPI { return m.EKS().(*mocks.EKSAPI) }
func (m MockProvider) EC2() ec2iface.EC2API   { return m.ec2 }
func (m MockProvider) mockEC2() *mocks.EC2API { return m.EC2().(*mocks.EC2API) }
func (m MockProvider) STS() stsiface.STSAPI   { return m.sts }
func (m MockProvider) mockSTS() *mocks.STSAPI { return m.STS().(*mocks.STSAPI) }

var _ = Describe("Eks", func() {
	var (
		c *ClusterProvider
		p *MockProvider
	)

	BeforeEach(func() {

	})

	Describe("SelectAvailabilityZones", func() {
		Context("TODO", func() {
			var (
				err error
			)

			BeforeEach(func() {

				p = &MockProvider{
					cfn: &mocks.CloudFormationAPI{},
					eks: &mocks.EKSAPI{},
					ec2: &mocks.EC2API{},
					sts: &mocks.STSAPI{},
				}

				c = &ClusterProvider{
					Provider: p,
				}

				p.mockEC2().On("DescribeAvailabilityZones",
					mock.MatchedBy(func(input *ec2.DescribeAvailabilityZonesInput) bool {
						return true
					}),
				).Return(&ec2.DescribeAvailabilityZonesOutput{})
			})

			Context("and normal log level", func() {
				BeforeEach(func() {
					logger.Level = 3
				})

				JustBeforeEach(func() {
					err = c.SelectAvailabilityZones()
				})

				It("should not error", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("should have called AWS EC2 DescribeAvailabilityZones", func() {
					Expect(p.mockEC2().AssertNumberOfCalls(GinkgoT(), "DescribeAvailabilityZones", 1)).To(BeTrue())
				})
			})
		})
	})
})
