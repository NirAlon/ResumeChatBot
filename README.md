# your Resume Chatbot API

A lightweight Go backend that answers questions about my skills, experience, and hobbies using the OpenAI API.
You can integrate this API with any frontend (React, etc.) to provide a conversational Q&A bot about a resume.

---

## Features

- Simple REST API for Q&A about a resume
- Powered by OpenAI’s GPT-3.5 Turbo for accurate, concise answers
- Reads the resume from a local file (not tracked in git for privacy)
- Easy to deploy, cross-platform, and Docker-friendly



## Requirements

- Go 1.20 or later (get it from [https://go.dev/dl/](https://go.dev/dl/))
- An OpenAI API key ([get one here](https://platform.openai.com/api-keys))
- Your `resume.txt` in the root of the project (plain text or markdown)



## Environment Variables

- `OPENAI_API_KEY` – Your OpenAI API key (required)



## Setup & Usage

1. **Clone the repository:**

    ```bash
    git clone https://github.com/yourusername/resume-chatbot-go.git
    cd resume-chatbot-go
    ```

2. **Add your resume:**

    - Copy your resume as `resume.txt` in the project root.
    - **DO NOT COMMIT your real resume!**  
      `resume.txt` is in `.gitignore` by default.

3. **Set your OpenAI API key:**

    ```bash
    export OPENAI_API_KEY=sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
    ```

4. **Build and run:**

    ```bash
    go build -o your-bot
    ./your-bot
    ```

    Or, just:

    ```bash
    go run main.go
    ```

5. **API Endpoint:**

    - The API will be available at:  
      `http://localhost:8080/chat`

    - Example request:

      ```bash
      curl -X POST http://localhost:8080/chat \
        -H "Content-Type: application/json" \
        -d '{"message": "What is your’s experience with Go?"}'
      ```

---

## Docker (optional)

You can build and run with Docker (if you prefer):

```bash
docker build -t your-bot .
docker run -e OPENAI_API_KEY=sk-xxxxxx -v $(pwd)/resume.txt:/app/resume.txt -p 8080:8080 your-bot
```


---

...

**Questions or Issues?**  
Open an issue or reach out to Me!

---

Cheers! 🎉


