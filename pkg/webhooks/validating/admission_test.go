package validating_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//	k8sfield "k8s.io/apimachinery/pkg/util/validation/field"

	//	templatev1 "github.com/openshift/api/template/v1"

	k6tv1 "kubevirt.io/client-go/api/v1"

	"github.com/fromanirh/kubevirt-template-validator/pkg/validation"
	"github.com/fromanirh/kubevirt-template-validator/pkg/webhooks/validating"
)

var _ = Describe("Admission", func() {
	Context("Without some data", func() {
		It("should admit without template", func() {
			newVM := k6tv1.VirtualMachine{}
			oldVM := k6tv1.VirtualMachine{}
			var rules []validation.Rule

			causes := validating.ValidateVMTemplate(rules, &newVM, &oldVM)

			Expect(len(causes)).To(Equal(0))
		})
	})
})
