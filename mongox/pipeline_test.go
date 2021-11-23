package mongox_test

import (
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/w6d-io/x/mongox"
)

var _ = Describe("Create Pipeline", func() {
	Context("", func() {
		var ()
		BeforeEach(func() {

		})
		AfterEach(func() {
		})
		It("pipeline success success", func() {
			template := `
			[
				{
					"$match": {
						"$and": [
							{
								 "priority": {
									 "$eq": {{.priority}}
								 }
							}
						]
					}
				}
			]
			`
			fields := make(map[string]interface{})
			fields["priority"] = 8

			_, err := CreatePipelineFromTemplate(
				template, fields,
			)
			Expect(err).To(Succeed())
		})
		It("pipeline error, wrong field", func() {
			template := `
			[
				{
					"$match": {
						"$and": [
							{
								 "priority": {
									 "$eq": {{.priority}}
								 }
							}
						]
					}
				}
			]
			`
			field := 6
			_, err := CreatePipelineFromTemplate(
				template, reflect.TypeOf(field),
			)
			Expect(err).NotTo(Succeed())
		})
		It("pipeline unmarshal error", func() {
			template := `
			[
				{
					"$match": {
						"$and": [
							{
								 "priority": {
									 "$eq": {{.priority}}
								 }
							}
						]
					}
				}
			]
			`
			_, err := CreatePipelineFromTemplate(
				template, nil,
			)
			Expect(err).NotTo(Succeed())
		})
	})
})
