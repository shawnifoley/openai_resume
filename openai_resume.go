package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
)

func readFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file %s: %w", filePath, err)
	}
	return string(content), nil
}

func generateResume(resumeMD, jobDescription string) (string, error) {
	prompt := fmt.Sprintf(`
    I have a resume formatted in Markdown and a job description.
    Please adapt my resume to better align with the job requirements while
    maintaining a professional tone tailored to this job.

    ### Here is my resume in Markdown:
        %s

    ### Here is the job description:
        %s

   Please modify the resume to:
   - Use keywords and phrases from the job description.
   - Adjust the bullet points under each role to emphasize relevant skills and achievements.
   - Make sure my experiences are presented in a way that matches the required qualifications.
   - Maintain clarity, conciseness, and professionalism throughout.

   Return the updated resume in Markdown format.
    `, resumeMD, jobDescription)

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: "gpt-4o-mini",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a professional resume writer.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: 1000,
		},
	)

	if err != nil {
		return "", fmt.Errorf("error generating resume: %w", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from openai")
	}

	return resp.Choices[0].Message.Content, nil
}

func main() {
	godotenv.Load()

	if len(os.Args) != 4 {
		fmt.Println("Usage: openai_resume <job_description_file> <resume_md_file>  <generated_resume_file>")
		os.Exit(1)
	}

	jobDescriptionFile := os.Args[1]
	resumeMDFile := os.Args[2]
	generatedResumeFile := os.Args[3]

	jobDescription, err := readFile(jobDescriptionFile)
	if err != nil {
		log.Fatalf("Error reading job description file: %v", err)
	}

	resumeMD, err := readFile(resumeMDFile)
	if err != nil {
		log.Fatalf("Error reading resume file: %v", err)
	}

	resume, err := generateResume(resumeMD, jobDescription)
	if err != nil {
		log.Fatalf("Error generating resume: %v", err)
	}
	fmt.Println("\nGenerated Resume:\n")
	fmt.Println(resume)

	// save resume
	err = ioutil.WriteFile(generatedResumeFile, []byte(resume), 0644)
	if err != nil {
		log.Fatalf("Error saving generated resume: %v", err)
	}

	fmt.Println("\nResume saved to", generatedResumeFile)

}
