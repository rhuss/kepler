package collector

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test Metric Unit", func() {
	It("Check feature values", func() {
		Expect(len(uintFeatures)).Should(BeNumerically(">", 0))
		Expect(len(podEnergyLabels)).Should(BeNumerically(">", 0))
		Expect(len(podEnergyLabels)).Should(BeNumerically(">", 0))
		fmt.Printf("%v\n%v\n%v\n", uintFeatures, podEnergyLabels, podEnergyLabels)
	})

	It("Check convert values", func() {
		v := NewPodEnergy("podA", "default")
		v.EnergyInCore = &UInt64Stat{
			Curr: 10,
			Aggr: 10,
		}
		v.CgroupFSStats = map[string]*ContainerUInt64Stat{
			CPU_USAGE_TOTAL_KEY: &ContainerUInt64Stat{
				Stat: map[string]*UInt64Stat{
					"cA": &UInt64Stat{
						Curr: SAMPLE_CURR,
						Aggr: SAMPLE_AGGR,
					},
				},
			},
		}
		collectedValues := v.ToPrometheusValues()
		Expect(len(collectedValues)).To(Equal(len(podEnergyLabels)))
		fmt.Printf("%v\n", collectedValues)
	})

})