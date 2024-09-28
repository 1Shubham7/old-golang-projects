package main

import (
	"bytes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
)

var _ = Describe("Car service test", func() {
	Context("When successfully creating a new car", func() {
		It("should return the correct response", func() {
			requestBody := `{"id": "701","title": "GM","color": "Transparent"}`
			request, _ := http.NewRequest("POST", "http://localhost:8080/cars", bytes.NewBuffer([]byte(requestBody)))
			request.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			response, err := client.Do(request)
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusCreated))
		})
	})

	Context("When creating an invalid car", func() {
		It("should return an error in the response", func() {
			requestBody := ""
			request, _ := http.NewRequest("POST", "http://localhost:8080/cars", bytes.NewBuffer([]byte(requestBody)))
			request.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			response, err := client.Do(request)
			Expect(err).NotTo(HaveOccurred())
			Expect(response.StatusCode).To(Equal(http.StatusBadRequest))

			responseBody, _ := io.ReadAll(response.Body)
			Expect(string(responseBody)).To(ContainSubstring("Failed"))
		})
	})
})