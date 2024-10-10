#!/usr/bin/env python3

import sys
import os
import openai
from dotenv import load_dotenv

# Load environment variables from .env file
load_dotenv()

# Set up OpenAI API key
openai.api_key = os.getenv("OPENAI_API_KEY")


def read_file(file_path):
    """Read the content of a file."""
    with open(file_path, "r") as file:
        return file.read()


def generate_resume(resume_md, job_description):
    """Generate a resume using OpenAI's GPT model."""
    prompt = f"""
    I have a resume formatted in Markdown and a job description. \
    Please adapt my resume to better align with the job requirements while \
    maintaining a professional tone tailored to this job.

    ### Here is my resume in Markdown:
        {resume_md}

    ### Here is the job description:
        {job_description}

   Please modify the resume to:
   - Use keywords and phrases from the job description.
   - Adjust the bullet points under each role to emphasize relevant skills and achievements.
   - Make sure my experiences are presented in a way that matches the required qualifications.
   - Maintain clarity, conciseness, and professionalism throughout.

   Return the updated resume in Markdown format.
    """
    response = openai.chat.completions.create(
        model="gpt-4o-mini",
        messages=[
            {"role": "system", "content": "You are a professional resume writer."},
            {"role": "user", "content": prompt},
        ],
        max_tokens=1000,
    )

    return response.choices[0].message.content


def main():
    if len(sys.argv) != 4:
        print(
            "Usage: python openai_resume.py <resume in markdown> <job_description_file> <generated resume>"
        )
        sys.exit(1)

    job_description_file = sys.argv[1]
    resume_md = sys.argv[2]
    generated_resume = sys.argv[3]

    # Read job description and current resume from files
    job_description = read_file(job_description_file)
    resume_md = read_file(resume_md)

    # Generate resume
    resume = generate_resume(resume_md, job_description)

    # Print the generated resume
    print("\nGenerated Resume:\n")
    print(resume)

    # Save the resume to a file in markdown format
    with open(generated_resume, "w") as file:
        file.write(resume)
    print("\nResume saved to " + generated_resume)


if __name__ == "__main__":
    main()
